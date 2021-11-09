// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/zeebo/errs"

	"ultimatedivision/cards/nfts/nftwaitlist"
	"ultimatedivision/internal/logger"
	"ultimatedivision/pkg/auth"
	"ultimatedivision/users"
)

// ErrNFTWaitList is an internal error type for NFTWaitList controller.
var ErrNFTWaitList = errs.Class("NFTWaitList controller error")

// NFTWaitList is a mvc controller that handles all NFTWaitList related views.
type NFTWaitList struct {
	log         logger.Logger
	nftWaitList *nftwaitlist.Service
}

// NewNFTWaitList is a constructor for NFTWaitList controller.
func NewNFTWaitList(log logger.Logger, nftWaitList *nftwaitlist.Service) *NFTWaitList {
	nftWaitListController := &NFTWaitList{
		log:         log,
		nftWaitList: nftWaitList,
	}

	return nftWaitListController
}

// Create is an endpoint that creates nft token.
func (controller *NFTWaitList) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	w.Header().Set("Content-Type", "application/json")

	claims, err := auth.GetClaims(ctx)
	if err != nil {
		controller.serveError(w, http.StatusUnauthorized, ErrNFTWaitList.Wrap(err))
		return
	}

	var createNFT nftwaitlist.CreateNFT

	if err = json.NewDecoder(r.Body).Decode(&createNFT); err != nil {
		controller.serveError(w, http.StatusBadRequest, ErrNFTWaitList.Wrap(err))
		return
	}
	createNFT.UserID = claims.UserID

	err = controller.nftWaitList.Create(ctx, createNFT)
	if err != nil {
		controller.log.Error("could not create nft token", ErrNFTWaitList.Wrap(err))

		if users.ErrNoUser.Has(err) {
			controller.serveError(w, http.StatusNotFound, ErrNFTWaitList.Wrap(err))
			return
		}

		controller.serveError(w, http.StatusInternalServerError, ErrNFTWaitList.Wrap(err))
		return
	}
}

// serveError replies to the request with specific code and error message.
func (controller *NFTWaitList) serveError(w http.ResponseWriter, status int, err error) {
	w.WriteHeader(status)
	var response struct {
		Error string `json:"error"`
	}
	response.Error = err.Error()

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		controller.log.Error("failed to write json error response", ErrCards.Wrap(err))
	}
}
