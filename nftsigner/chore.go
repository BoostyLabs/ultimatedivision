// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package nftsigner

import (
	"context"
	"crypto/ecdsa"
	"math/big"
	"time"

	"github.com/BoostyLabs/evmsignature"
	"github.com/BoostyLabs/thelooper"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/zeebo/errs"

	"ultimatedivision/cards/waitlist"
	"ultimatedivision/users"
)

// ChoreError represents nft signer chore error type.
var ChoreError = errs.Class("nft signer chore error")

// ChoreConfig is the global configuration for nftsigner.
type ChoreConfig struct {
	RenewalInterval           time.Duration           `json:"renewalInterval"`
	PrivateKey                evmsignature.PrivateKey `json:"privateKey"`
	SmartContractAddress      evmsignature.Address    `json:"smartContractAddress"`
	PrivateKeyVelas           evmsignature.PrivateKey `json:"privateKeyVelas"`
	SmartContractAddressVelas evmsignature.Address    `json:"smartContractAddressVelas"`
}

// Chore requests for unsigned nft tokens and sign all of them .
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
	return chore.loop.Run(ctx, func(ctx context.Context) error {
		unsignedNFTs, err := chore.nfts.ListWithoutPassword(ctx)
		if err != nil {
			return ChoreError.Wrap(err)
		}

		privateKeyECDSA, err := crypto.HexToECDSA(string(chore.config.PrivateKey))
		if err != nil {
			return ChoreError.Wrap(err)
		}

		privateKeyECDSAVelas, err := crypto.HexToECDSA(string(chore.config.PrivateKeyVelas))
		if err != nil {
			return ChoreError.Wrap(err)
		}

		for _, token := range unsignedNFTs {
			var signature evmsignature.Signature
			var smartContract evmsignature.Address
			var privatKey *ecdsa.PrivateKey
			switch token.WalletType {
			case users.Wallet:
				smartContract = chore.config.SmartContractAddress
				privatKey = privateKeyECDSA
			case users.Velas:
				smartContract = chore.config.SmartContractAddressVelas
				privatKey = privateKeyECDSAVelas
			}
			if token.Value.Cmp(big.NewInt(0)) <= 0 {
				signature, err = evmsignature.GenerateSignatureWithValue(token.Wallet, smartContract, token.TokenID, privatKey)
				if err != nil {
					return ChoreError.Wrap(err)
				}
			} else {
				signature, err = evmsignature.GenerateSignatureWithValueAndNonce(token.Wallet, smartContract, &token.Value, token.TokenID, privatKey)
				if err != nil {
					return ChoreError.Wrap(err)
				}
			}

			if err = chore.nfts.Update(ctx, token.TokenID, signature); err != nil {
				return ChoreError.Wrap(err)
			}
		}

		return ChoreError.Wrap(err)
	})
}
