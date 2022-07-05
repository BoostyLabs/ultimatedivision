// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package nftsigner

import (
	"context"
	"math/big"
	"time"

	"github.com/BoostyLabs/evmsignature"
	"github.com/BoostyLabs/thelooper"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/zeebo/errs"

	"ultimatedivision/cards/waitlist"
	"ultimatedivision/internal/signature"
)

// ChoreError represents nft signer chore error type.
var ChoreError = errs.Class("nft signer chore error")

// ChoreConfig is the global configuration for nftsigner.
type ChoreConfig struct {
	RenewalInterval          time.Duration           `json:"renewalInterval"`
	PrivateKey               evmsignature.PrivateKey `json:"privateKey"`
	NFTCreateContractAddress common.Address          `json:"nftCreateContractAddress"`
	BuyMethodName            string                  `json:"buyMethodName"`
}

// Chore requests for unsigned nft tokens and sign all of them.
//
// architecture: Chore
type Chore struct {
	loop   *thelooper.Loop
	config ChoreConfig
	nfts   *waitlist.Service
}

// NewChore instantiates Chore.
func NewChore(config ChoreConfig, db waitlist.DB) *Chore {
	return &Chore{
		loop:   thelooper.NewLoop(config.RenewalInterval),
		config: config,
		nfts:   waitlist.NewService(waitlist.Config{}, db, nil, nil, nil, nil),
	}
}

// Run starts the chore for signing unsigned nft token from ultimatedivision.
func (chore *Chore) Run(ctx context.Context) (err error) {
	privateKeyECDSA, err := crypto.HexToECDSA(string(chore.config.PrivateKey))
	if err != nil {
		return ChoreError.Wrap(err)
	}

	return chore.loop.Run(ctx, func(ctx context.Context) error {
		unsignedNFTs, err := chore.nfts.ListWithoutPassword(ctx)
		if err != nil {
			return ChoreError.Wrap(err)
		}

		for _, token := range unsignedNFTs {
			var generatedSignature evmsignature.Signature
			if token.Value.Cmp(big.NewInt(0)) <= 0 {
				generatedSignature, err = evmsignature.GenerateSignatureWithValue(evmsignature.Address(token.Wallet.String()),
					evmsignature.Address(chore.config.NFTCreateContractAddress.String()), token.TokenID, privateKeyECDSA)
				if err != nil {
					return ChoreError.Wrap(err)
				}
			} else {
				nftStoreSignature := signature.NFTStoreSignature{
					MethodName:      chore.config.BuyMethodName,
					WalletAddress:   token.Wallet,
					ContractAddress: chore.config.NFTCreateContractAddress,
					TokenID:         token.TokenID,
					Value:           &token.Value,
					PrivateKey:      privateKeyECDSA,
				}

				generatedSignature, err = signature.GenerateNFTStoreSignature(nftStoreSignature)
				if err != nil {
					return ChoreError.Wrap(err)
				}
			}

			if err = chore.nfts.Update(ctx, token.TokenID, generatedSignature); err != nil {
				return ChoreError.Wrap(err)
			}
		}

		return ChoreError.Wrap(err)
	})
}
