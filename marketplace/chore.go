// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package marketplace

import (
	"context"
	"encoding/hex"
	"fmt"

	"github.com/BoostyLabs/thelooper"
	casper_ed25519 "github.com/casper-ecosystem/casper-golang-sdk/keypair/ed25519"
	"github.com/zeebo/errs"

	"ultimatedivision/cards"
	"ultimatedivision/internal/contract/casper"
	"ultimatedivision/internal/logger"
	contract "ultimatedivision/pkg/contractcasper"
)

var (
	// ChoreError represents lot chore error type.
	ChoreError = errs.Class("expiration lot chore error")
)

// Chore requests access token for contis api calls, re-requests it after token's expiration time.
//
// architecture: Chore
type Chore struct {
	log         logger.Logger
	Loop        *thelooper.Loop
	marketplace *Service
}

// NewChore instantiates Chore.
func NewChore(log logger.Logger, config Config, marketplace *Service) *Chore {
	return &Chore{
		log:         log,
		Loop:        thelooper.NewLoop(config.LotRenewalInterval),
		marketplace: marketplace,
	}
}

// Run starts the chore for re-check the expiration time of the lot.
func (chore *Chore) Run(ctx context.Context) (err error) {
	return chore.Loop.Run(ctx, func(ctx context.Context) error {
		lots, err := chore.marketplace.ListExpiredLot(ctx)
		if err != nil {
			chore.log.Error("could not get list of the expired lot", ChoreError.Wrap(err))
		}

		// TODO: the transaction may be required for all operations.
		for _, lot := range lots {
			tokenID, err := chore.marketplace.GetNFTTokenIDbyCardID(ctx, lot.CardID)
			if err != nil {
				chore.log.Error(fmt.Sprintf("could not get nft token id by card id equal %v from db", lot.CardID), ChoreError.Wrap(err))
			}

			privateAccountKey := chore.marketplace.config.ContractOwnerPrivateKey
			privateAccountKeyBytes, err := hex.DecodeString(privateAccountKey)
			if err != nil {
				chore.log.Error("could not decode privateAccountKey in bytes", ChoreError.Wrap(err))
			}

			publicAccountKey := chore.marketplace.config.ContractOwnerPublicKey
			publicAccountKeyBytes, err := hex.DecodeString(publicAccountKey)
			if err != nil {
				chore.log.Error("could not decode publicAccountKey in bytes", ChoreError.Wrap(err))
			}

			pair := casper_ed25519.ParseKeyPair(publicAccountKeyBytes, privateAccountKeyBytes)

			casperClient := contract.New(chore.marketplace.config.RPCNodeAddress)
			transfer := casper.NewTransfer(casperClient, func(b []byte) ([]byte, error) {
				casperSignature := pair.Sign(b)
				return casperSignature.SignatureData, nil
			})

			_, err = transfer.FinalListing(ctx, casper.FinalListingRequest{
				PublicKey:          pair.PublicKey(),
				ChainName:          "casper-test",
				StandardPayment:    10000000000,
				MarketContractHash: chore.marketplace.config.MarketContractAddress,
				NFTContractHash:    fmt.Sprintf("%s%s", chore.marketplace.config.NFTContractPrefix, chore.marketplace.config.NFTContractAddress),
				TokenID:            tokenID.String(),
			})
			if err != nil {
				chore.log.Error("could not connect with smart contract final listing method", ChoreError.Wrap(err))
			}

			if lot.CurrentPrice.BitLen() != 0 {
				// TODO: unhold old user's money.

				winLot := WinLot{
					CardID:    lot.CardID,
					Type:      lot.Type,
					UserID:    lot.UserID,
					ShopperID: lot.ShopperID,
					Status:    StatusSold,
					Amount:    lot.MaxPrice,
				}

				err := chore.marketplace.WinLot(ctx, winLot)
				if err != nil {
					chore.log.Error("could not changes owner of the item in marketplace", ChoreError.Wrap(err))
				}
				continue
			}

			err = chore.marketplace.UpdateStatusLot(ctx, lot.CardID, StatusExpired)
			if err != nil {
				chore.log.Error("could not update status of the lot to the StatusExpired", ChoreError.Wrap(err))
			}

			if lot.Type == TypeCard {
				if err := chore.marketplace.cards.UpdateStatus(ctx, lot.CardID, cards.StatusActive); err != nil {
					chore.log.Error("could not update status of the lot to the StatusActive", ChoreError.Wrap(err))
				}
			}
		}
		return ChoreError.Wrap(err)
	})
}

// Close closes the chore for re-check the expiration time of the lot.
func (chore *Chore) Close() {
	chore.Loop.Close()
}
