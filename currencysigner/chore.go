// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package currencysigner

import (
	"context"
	"time"

	"github.com/BoostyLabs/evmsignature"
	"github.com/BoostyLabs/thelooper"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/zeebo/errs"

	"ultimatedivision/internal/logger"
	"ultimatedivision/udts/currencywaitlist"
	"ultimatedivision/users"
)

// ChoreError represents nft signer chore error type.
var ChoreError = errs.Class("nft signer chore error")

// ChoreConfig is the global configuration for currencysigner.
type ChoreConfig struct {
	RenewalInterval    time.Duration           `json:"renewalInterval"`
	PrivateKey         evmsignature.PrivateKey `json:"privateKey"`
	UDTContractAddress common.Address          `json:"udtContractAddress"`

	NFTCreateContractAddress   common.Address `json:"nftCreateContractAddress"`
	VelasSmartContractAddress  common.Address `json:"velasSmartContractAddress"`
	CasperSmartContractAddress string         `json:"casperSmartContractAddress"`
	CasperTokenContract        string         `json:"casperTokenContract"`
}

// Chore requests for unsigned nft tokens and sign all of them .
//
// architecture: Chore
type Chore struct {
	log              logger.Logger
	currencywaitlist *currencywaitlist.Service
	Loop             *thelooper.Loop
	config           ChoreConfig
}

// NewChore instantiates Chore.
func NewChore(log logger.Logger, config ChoreConfig, db currencywaitlist.DB) *Chore {
	return &Chore{
		log:              log,
		Loop:             thelooper.NewLoop(config.RenewalInterval),
		currencywaitlist: currencywaitlist.NewService(currencywaitlist.Config{}, db, nil, nil),
		config:           config,
	}
}

// Run starts the chore for signing unsigned item of currency waitlist from ultimatedivision.
func (chore *Chore) Run(ctx context.Context) (err error) {
	return chore.Loop.Run(ctx, func(ctx context.Context) error {
		unsignedItems, err := chore.currencywaitlist.ListWithoutSignature(ctx)
		if err != nil {
			return ChoreError.Wrap(err)
		}

		privateKeyECDSA, err := crypto.HexToECDSA(string(chore.config.PrivateKey))
		if err != nil {
			return ChoreError.Wrap(err)
		}

		for _, item := range unsignedItems {

			var (
				signature           evmsignature.Signature
				smartContract       evmsignature.Address
				casperTokenContract string
				casperContract      string
				casperWallet        string
			)

			switch item.WalletType {
			case string(users.WalletTypeETH):
				smartContract = evmsignature.Address(chore.config.UDTContractAddress.String())
			case string(users.WalletTypeVelas):
				smartContract = evmsignature.Address(chore.config.VelasSmartContractAddress.String())
			case string(users.WalletTypeCasper):
				casperContract = "0x" + chore.config.CasperSmartContractAddress
				casperTokenContract = chore.config.CasperTokenContract
				casperWallet = "0x" + item.WalletAddress.String()
			}

			if casperContract != "" {
				signature, err = evmsignature.GenerateSignatureWithValueAndNonce(evmsignature.Address(casperWallet),
					evmsignature.Address(casperTokenContract), &item.Value, item.Nonce, privateKeyECDSA)
				if err != nil {
					return ChoreError.Wrap(err)
				}
			} else {
				signature, err = evmsignature.GenerateSignatureWithValueAndNonce(evmsignature.Address(item.WalletAddress.String()),
					smartContract, &item.Value, item.Nonce, privateKeyECDSA)
				if err != nil {
					return ChoreError.Wrap(err)
				}
			}

			if err != nil {
				return ChoreError.Wrap(err)
			}

			err = chore.currencywaitlist.UpdateSignature(ctx, signature, item.WalletAddress, item.Nonce)
			if err != nil {
				return ChoreError.Wrap(err)
			}
		}

		return ChoreError.Wrap(err)
	})
}

// Close closes the chore for signing unsigned item of currency waitlist from ultimatedivision.
func (chore *Chore) Close() {
	chore.Loop.Close()
}
