// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package nftwaitlist

import (
	"context"

	"github.com/zeebo/errs"

	"ultimatedivision/cards"
	"ultimatedivision/cards/avatars"
	"ultimatedivision/cards/nfts"
	"ultimatedivision/users"
)

// ErrNFTWaitlist indicated that there was an error in service.
var ErrNFTWaitlist = errs.Class("NFTWaitlist service error")

// Service is handling NFTWaitlist related logic.
//
// architecture: Service
type Service struct {
	nftWaitList DB
	cards       *cards.Service
	avatars     *avatars.Service
	users       *users.Service
	nfts        *nfts.Service
}

// NewService is a constructor for NFTWaitlist service.
func NewService(nftWaitList DB, cards *cards.Service, avatars *avatars.Service, users *users.Service, nfts *nfts.Service) *Service {
	return &Service{
		nftWaitList: nftWaitList,
		cards:       cards,
		avatars:     avatars,
		users:       users,
		nfts:        nfts,
	}
}

// Create creates nft for wait list.
func (service *Service) Create(ctx context.Context, createNFT CreateNFT) error {
	card, err := service.cards.Get(ctx, createNFT.CardID)
	if err != nil {
		return ErrNFTWaitlist.Wrap(err)
	}

	if card.UserID != createNFT.UserID {
		return ErrNFTWaitlist.New("it isn't user`s card")
	}

	avatar, err := service.avatars.Get(ctx, createNFT.CardID)
	if err != nil {
		return ErrNFTWaitlist.Wrap(err)
	}

	// TODO: save avatar

	_, err = service.nfts.Generate(ctx, card, avatar.OriginalURL)
	if err != nil {
		return ErrNFTWaitlist.Wrap(err)
	}

	// TODO: save nft
	// TODO: add transaction

	if err = service.users.UpdateWalletAddress(ctx, createNFT.WalletAddress, createNFT.UserID); err != nil {
		return ErrNFTWaitlist.Wrap(err)
	}

	return service.nftWaitList.Create(ctx, createNFT.CardID, createNFT.WalletAddress)
}

// List returns all nft for wait list.
func (service *Service) List(ctx context.Context) ([]NFTWaitList, error) {
	allNFT, err := service.nftWaitList.List(ctx)
	return allNFT, ErrNFTWaitlist.Wrap(err)
}

// Get returns nft for wait list by token id.
func (service *Service) Get(ctx context.Context, tokenID int) (NFTWaitList, error) {
	nft, err := service.nftWaitList.Get(ctx, tokenID)
	return nft, ErrNFTWaitlist.Wrap(err)
}

// GetLastTokenID returns id of latest nft for wait list.
func (service *Service) GetLastTokenID(ctx context.Context) (int, error) {
	lastID, err := service.nftWaitList.GetLast(ctx)
	return lastID, ErrNFTWaitlist.Wrap(err)
}

// ListWithoutPassword returns nft for wait list without password.
func (service *Service) ListWithoutPassword(ctx context.Context) ([]NFTWaitList, error) {
	nftWithoutPassword, err := service.nftWaitList.ListWithoutPassword(ctx)
	return nftWithoutPassword, ErrNFTWaitlist.Wrap(err)
}

// Delete deletes nft for wait list.
func (service *Service) Delete(ctx context.Context, tokenIDs []int) error {
	return ErrNFTWaitlist.Wrap(service.nftWaitList.Delete(ctx, tokenIDs))
}
