// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package waitlist

import (
	"context"
	"strconv"

	"github.com/google/uuid"
	"github.com/zeebo/errs"

	"ultimatedivision/cards"
	"ultimatedivision/cards/nfts"
	"ultimatedivision/internal/logger"
	"ultimatedivision/pkg/cryptoutils"
	"ultimatedivision/pkg/jsonrpc"
	"ultimatedivision/pkg/sync"
	"ultimatedivision/users"
)

var (
	// ChoreError represents waitlist chore error type.
	ChoreError = errs.Class("expiration waitlist chore error")
)

// Chore requests access token for contis api calls, re-requests it after token's expiration time.
//
// architecture: Chore
type Chore struct {
	config   Config
	log      logger.Logger
	Loop     *sync.Cycle
	waitList *Service
	nfts     *nfts.Service
	users    *users.Service
	cards    *cards.Service
}

// NewChore instantiates Chore.
func NewChore(config Config, log logger.Logger, waitList *Service, nfts *nfts.Service, users *users.Service, cards *cards.Service) *Chore {
	return &Chore{
		config:   config,
		log:      log,
		Loop:     sync.NewCycle(config.WaitListRenewalInterval),
		waitList: waitList,
		nfts:     nfts,
		users:    users,
		cards:    cards,
	}
}

// RunCheckMintEvent runs a task to check the nft assignment.
func (chore *Chore) RunCheckMintEvent(ctx context.Context) (err error) {
	filter := []*jsonrpc.CreateFilter{
		{
			ToBlock: cryptoutils.BlockTagLatest,
			Address: chore.config.NFTContract.Address,
			Topics:  []cryptoutils.Hex{chore.config.NFTContract.AddressEvent},
		},
	}

	transaction := jsonrpc.NewTransaction(jsonrpc.MethodEthNewFilter, filter, cryptoutils.ChainIDRinkeby)
	body, err := jsonrpc.Send(chore.config.AddressNodeServer, transaction)
	if err != nil {
		return ChoreError.Wrap(err)
	}

	addressOfFilter, err := jsonrpc.GetAddressOfFilter(body)
	if err != nil {
		return ChoreError.Wrap(err)
	}

	return chore.Loop.Run(ctx, func(ctx context.Context) error {

		transaction := jsonrpc.NewTransaction(jsonrpc.MethodEthGetFilterChanges, []cryptoutils.Address{addressOfFilter}, cryptoutils.ChainIDRinkeby)
		body, err := jsonrpc.Send(chore.config.AddressNodeServer, transaction)
		if err != nil {
			return ChoreError.Wrap(err)
		}

		events, err := jsonrpc.ListEvents(body)
		if err != nil {
			return ChoreError.Wrap(err)
		}

		for _, event := range events {
			fromStr := string(event.Topics[1])
			from, _ := strconv.ParseInt(fromStr[cryptoutils.LengthHexPrefix:], 16, 64)

			toStr := string(event.Topics[2])
			toAddress := cryptoutils.CreateValidAddress(cryptoutils.Hex(toStr))

			tokenIDStr := string(event.Topics[3])
			tokenID, err := strconv.ParseInt(tokenIDStr[cryptoutils.LengthHexPrefix:], 16, 64)
			if err != nil {
				return ChoreError.Wrap(err)
			}

			if from == 0 {
				nftWaitList, err := chore.waitList.GetByTokenID(ctx, tokenID)
				if err != nil {
					return ChoreError.Wrap(err)
				}

				nft := nfts.NFT{
					CardID:        nftWaitList.CardID,
					Chain:         cryptoutils.ChainPolygon,
					TokenID:       tokenID,
					WalletAddress: toAddress,
				}
				if err = chore.nfts.Create(ctx, nft); err != nil {
					return ChoreError.Wrap(err)
				}
				continue
			}

			nft, err := chore.nfts.Get(ctx, tokenID, cryptoutils.ChainPolygon)
			if err != nil {
				return ChoreError.Wrap(err)
			}

			user, err := chore.users.GetByWalletAddress(ctx, toAddress)
			if err != nil {
				if err = chore.nfts.Delete(ctx, nft.CardID); err != nil {
					return ChoreError.Wrap(err)
				}

				if err = chore.cards.UpdateUserID(ctx, nft.CardID, uuid.Nil); err != nil {
					return ChoreError.Wrap(err)
				}
				continue
			}

			if err = chore.nfts.Update(ctx, nft); err != nil {
				return ChoreError.Wrap(err)
			}

			if err = chore.cards.UpdateUserID(ctx, nft.CardID, user.ID); err != nil {
				return ChoreError.Wrap(err)
			}
		}

		return ChoreError.Wrap(err)
	})
}
