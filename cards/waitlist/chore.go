// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package waitlist

import (
	"context"
	"reflect"

	"github.com/BoostyLabs/evmsignature"
	"github.com/ethereum/go-ethereum/common"
	"github.com/google/uuid"
	"github.com/make-software/casper-go-sdk/sse"
	"github.com/zeebo/errs"

	// TODO: remove naming after a complete transition to a new sdk.
	new_casper_types "github.com/make-software/casper-go-sdk/types"
	new_casper_key "github.com/make-software/casper-go-sdk/types/key"

	"ultimatedivision/cards"
	"ultimatedivision/cards/nfts"
	"ultimatedivision/internal/logger"
	"ultimatedivision/pkg/eventparsing"
	"ultimatedivision/users"
)

var (
	// ChoreError represents waitlist chore error type.
	ChoreError = errs.Class("expiration waitlist chore error")
)

// Chore requests access token for contis api calls, re-requests it after token's expiration time.
//
// architecture: Chore.
type Chore struct {
	log      logger.Logger
	config   Config
	waitList *Service
	nfts     *nfts.Service
	users    *users.Service
	cards    *cards.Service
	events   *sse.Client
}

// NewChore instantiates Chore.
func NewChore(log logger.Logger, config Config, waitList *Service, nfts *nfts.Service, users *users.Service, cards *cards.Service) *Chore {
	eventsClient := sse.NewClient(config.EventNodeAddress)

	return &Chore{
		log:      log,
		config:   config,
		waitList: waitList,
		nfts:     nfts,
		users:    users,
		cards:    cards,
		events:   eventsClient,
	}
}

// RunCasperCheckMintEvent runs a task to check and create the casper nft assignment.
func (chore *Chore) RunCasperCheckMintEvent(ctx context.Context) error {
	chore.events.RegisterHandler(
		sse.DeployProcessedEventType,
		func(ctx context.Context, rawEvent sse.RawEvent) error {
			select {
			case <-ctx.Done():
				return nil
			default:
			}

			deploy, err := rawEvent.ParseAsDeployProcessedEvent()
			if err != nil {
				chore.log.Error("could not parse as deploy processed event", ChoreError.Wrap(err))
				return ChoreError.Wrap(err)
			}

			if deploy.DeployProcessed == (sse.DeployProcessedPayload{}) ||
				deploy.DeployProcessed.ExecutionResult == (new_casper_types.ExecutionResultStatus{}) ||
				deploy.DeployProcessed.ExecutionResult.Success == (&new_casper_types.ExecutionResultStatusData{}) ||
				reflect.DeepEqual(deploy.DeployProcessed.ExecutionResult.Success.Effect, new_casper_types.Effect{}) ||
				deploy.DeployProcessed.ExecutionResult.Success.Effect.Transforms == nil {
				return nil
			}

			transforms := deploy.DeployProcessed.ExecutionResult.Success.Effect.Transforms
			if len(transforms) == 0 {
				return nil
			}

			for _, transform := range transforms {
				select {
				case <-ctx.Done():
					return nil
				default:
				}

				if transform.Key == (new_casper_key.Key{}) || !transform.Transform.IsWriteCLValue() {
					continue
				}

				argument, err := transform.Transform.ParseAsWriteCLValue()
				if err != nil {
					chore.log.Error("could not parse transform as write CLValue", ChoreError.Wrap(err))
					return ChoreError.Wrap(err)
				}

				argumentBytes, err := argument.Bytes()
				if err != nil {
					chore.log.Error("could not get bytes from argument", ChoreError.Wrap(err))
					return ChoreError.Wrap(err)
				}

				if len(argumentBytes.String()) != 170 {
					continue
				}

				eventData := eventparsing.EventData{
					Bytes: argumentBytes.String(),
				}
				tokenID, err := eventData.GetTokenID(eventData)
				if err != nil {
					return ChoreError.Wrap(err)
				}

				if tokenID == uuid.Nil {
					continue
				}

				nftWaitList, err := chore.waitList.GetByTokenID(ctx, tokenID)
				if err != nil {
					chore.log.Error("could not get node events", ChoreError.Wrap(err))
				}

				toAddress := common.HexToAddress(nftWaitList.CasperWalletHash)
				nft := nfts.NFT{
					CardID:        nftWaitList.CardID,
					TokenID:       tokenID,
					Chain:         evmsignature.ChainEthereum,
					WalletAddress: toAddress,
				}

				if err = chore.nfts.Create(ctx, nft); err != nil {
					chore.log.Error("could not create nft", ChoreError.Wrap(err))
				}

				user, err := chore.users.GetByCasperHash(ctx, nftWaitList.CasperWalletHash)
				if err != nil {
					if err = chore.nfts.Delete(ctx, nft.CardID); err != nil {
						chore.log.Error("could not delete nft events", ChoreError.Wrap(err))
					}
					chore.log.Error("could get user by casper hash", ChoreError.Wrap(err))
				}

				if err = chore.nfts.Update(ctx, nft); err != nil {
					chore.log.Error("could not update nft", ChoreError.Wrap(err))
				}

				if err = chore.cards.UpdateUserID(ctx, nft.CardID, user.ID); err != nil {
					chore.log.Error("could not update user ID by card id", ChoreError.Wrap(err))
				}

				if err = chore.cards.UpdateMintedStatus(ctx, nft.CardID, cards.Minted); err != nil {
					chore.log.Error("could not update minted status to 1", ChoreError.Wrap(err))
				}

				return nil
			}

			return nil
		})

	lastEventID := 0
	err := chore.events.Start(ctx, lastEventID)
	if err != nil {
		chore.log.Error("events reading error: ", ChoreError.Wrap(err))
	}

	return nil
}
