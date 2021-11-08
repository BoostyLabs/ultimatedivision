// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package nftsigner

import (
	"context"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/zeebo/errs"

	"ultimatedivision/cards/nfts"
	"ultimatedivision/internal/logger"
	"ultimatedivision/pkg/cryptoutils"
	"ultimatedivision/pkg/sync"
)

// ChoreError represents nft signer chore error type.
var ChoreError = errs.Class("nft signer chore error")

// Chore requests access token for contis api calls, re-requests it after token's expiration time.
//
// architecture: Chore
type Chore struct {
	log  logger.Logger
	nfts *nfts.Service
	loop *sync.Cycle
}

// NewChore instantiates Chore.
func NewChore(log logger.Logger, config Config, db DB) *Chore {
	return &Chore{
		log:  log,
		loop: sync.NewCycle(config.RenewalInterval),
		nfts: nfts.NewService(nfts.Config{}, nil, nil, nil, db.NFTs()),
	}
}

// Run starts the chore for signing unsigned nft token from ultimatedivision.
func (chore *Chore) Run(ctx context.Context, smartContractAddress cryptoutils.Address, privateKey cryptoutils.PrivateKey) (err error) {
	return chore.loop.Run(ctx, func(ctx context.Context) error {
		unsignedNFTs, err := chore.nfts.ListWithoutPassword(ctx)
		if err != nil {
			return ChoreError.Wrap(err)
		}

		privateKeyECDSA, err := crypto.HexToECDSA(string(privateKey))
		if err != nil {
			return ChoreError.Wrap(err)
		}

		for _, token := range unsignedNFTs {
			signature, err := cryptoutils.GenerateSignatureWithToken(token.Wallet, smartContractAddress, token.TokenID, privateKeyECDSA)
			if err != nil {
				return ChoreError.Wrap(err)
			}

			err = chore.nfts.Update(ctx, token.TokenID, signature)
			if err != nil {
				return ChoreError.Wrap(err)
			}
		}

		return ChoreError.Wrap(err)
	})
}
