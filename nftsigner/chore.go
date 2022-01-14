// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package nftsigner

import (
	"context"
	"math/big"
	"time"

	"github.com/BoostyLabs/evmsignature"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/zeebo/errs"

	"ultimatedivision/cards/waitlist"
	"ultimatedivision/internal/logger"
	"ultimatedivision/pkg/sync"
)

// ChoreError represents nft signer chore error type.
var ChoreError = errs.Class("nft signer chore error")

// ChoreConfig is the global configuration for nftsigner.
type ChoreConfig struct {
	RenewalInterval      time.Duration           `json:"renewalInterval"`
	PrivateKey           evmsignature.PrivateKey `json:"privateKey"`
	SmartContractAddress evmsignature.Address    `json:"smartContractAddress"`
}

// Chore requests for unsigned nft tokens and sign all of them .
//
// architecture: Chore
type Chore struct {
	log    logger.Logger
	nfts   *waitlist.Service
	loop   *sync.Cycle
	config ChoreConfig
}

// NewChore instantiates Chore.
func NewChore(log logger.Logger, config ChoreConfig, db waitlist.DB) *Chore {
	return &Chore{
		log:    log,
		loop:   sync.NewCycle(config.RenewalInterval),
		nfts:   waitlist.NewService(waitlist.Config{}, db, nil, nil, nil, nil),
		config: config,
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

		for _, token := range unsignedNFTs {
			var signature evmsignature.Signature
			if token.Value.Cmp(big.NewInt(0)) <= 0 {
				signature, err = evmsignature.GenerateSignatureWithValue(token.Wallet, chore.config.SmartContractAddress, token.TokenID, privateKeyECDSA)
				if err != nil {
					return ChoreError.Wrap(err)
				}
			} else {
				signature, err = evmsignature.GenerateSignatureWithValueAndNonce(token.Wallet, chore.config.SmartContractAddress, &token.Value, token.TokenID, privateKeyECDSA)
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
