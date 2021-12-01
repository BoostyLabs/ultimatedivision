// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package waitlist

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/zeebo/errs"

	"ultimatedivision/cards"
	"ultimatedivision/cards/avatars"
	"ultimatedivision/cards/nfts"
	"ultimatedivision/pkg/cryptoutils"
	"ultimatedivision/users"
)

// ErrWaitlist indicated that there was an error in service.
var ErrWaitlist = errs.Class("waitlist service error")

// Service is handling waitList related logic.
//
// architecture: Service
type Service struct {
	config   Config
	waitList DB
	cards    *cards.Service
	avatars  *avatars.Service
	users    *users.Service
	nfts     *nfts.Service
}

// NewService is a constructor for waitlist service.
func NewService(config Config, waitList DB, cards *cards.Service, avatars *avatars.Service, users *users.Service, nfts *nfts.Service) *Service {
	return &Service{
		config:   config,
		waitList: waitList,
		cards:    cards,
		avatars:  avatars,
		users:    users,
		nfts:     nfts,
	}
}

// Create creates nft for wait list.
func (service *Service) Create(ctx context.Context, createNFT CreateNFT) (Transaction, error) {
	var transaction Transaction
	card, err := service.cards.Get(ctx, createNFT.CardID)
	if err != nil {
		return transaction, ErrWaitlist.Wrap(err)
	}

	if card.UserID != createNFT.UserID {
		return transaction, ErrWaitlist.New("it isn't user`s card")
	}

	if item, err := service.GetByCardID(ctx, createNFT.CardID); item.Password != "" && err == nil {
		transaction = Transaction{
			Password: item.Password,
			Contract: service.config.Contract,
			TokenID:  item.TokenID,
		}
		return transaction, nil
	}

	avatar, err := service.avatars.Get(ctx, createNFT.CardID)
	if err != nil {
		return transaction, ErrWaitlist.Wrap(err)
	}

	// TODO: save avatar in file storage

	service.nfts.Generate(ctx, card, avatar.OriginalURL)

	// TODO: save metadata in file storage
	// TODO: add transaction

	if err = service.users.UpdateWalletAddress(ctx, createNFT.WalletAddress, createNFT.UserID); err != nil {
		return transaction, ErrWaitlist.Wrap(err)
	}

	if err = service.waitList.Create(ctx, createNFT.CardID, createNFT.WalletAddress); err != nil {
		return transaction, ErrWaitlist.Wrap(err)
	}

	for range time.NewTicker(time.Millisecond * service.config.WaitListCheckSignature).C {
		if item, err := service.GetByCardID(ctx, createNFT.CardID); item.Password != "" && err == nil {
			transaction = Transaction{
				Password: item.Password,
				Contract: service.config.Contract,
				TokenID:  item.TokenID,
			}
			break
		}
	}

	return transaction, err
}

// List returns all nft for wait list.
func (service *Service) List(ctx context.Context) ([]Item, error) {
	allNFT, err := service.waitList.List(ctx)
	return allNFT, ErrWaitlist.Wrap(err)
}

// GetByTokenID returns nft for wait list by token id.
func (service *Service) GetByTokenID(ctx context.Context, tokenID uint64) (Item, error) {
	nft, err := service.waitList.GetByTokenID(ctx, tokenID)
	return nft, ErrWaitlist.Wrap(err)
}

// GetByCardID returns nft for wait list by card id.
func (service *Service) GetByCardID(ctx context.Context, cardID uuid.UUID) (Item, error) {
	nft, err := service.waitList.GetByCardID(ctx, cardID)
	return nft, ErrWaitlist.Wrap(err)
}

// GetLastTokenID returns id of latest nft for wait list.
func (service *Service) GetLastTokenID(ctx context.Context) (uint64, error) {
	lastID, err := service.waitList.GetLast(ctx)
	return lastID, ErrWaitlist.Wrap(err)
}

// ListWithoutPassword returns nft for wait list without password.
func (service *Service) ListWithoutPassword(ctx context.Context) ([]Item, error) {
	nftWithoutPassword, err := service.waitList.ListWithoutPassword(ctx)
	return nftWithoutPassword, ErrWaitlist.Wrap(err)
}

// Delete deletes nft for wait list.
func (service *Service) Delete(ctx context.Context, tokenIDs []uint64) error {
	return ErrWaitlist.Wrap(service.waitList.Delete(ctx, tokenIDs))
}

// Update updates signature to nft token.
func (service *Service) Update(ctx context.Context, tokenID uint64, password cryptoutils.Signature) error {
	return ErrWaitlist.Wrap(service.waitList.Update(ctx, tokenID, password))
}
