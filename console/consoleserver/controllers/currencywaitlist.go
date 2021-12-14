// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/zeebo/errs"

	"ultimatedivision/internal/logger"
	"ultimatedivision/pkg/auth"
	"ultimatedivision/pkg/cryptoutils"
	"ultimatedivision/udts"
	"ultimatedivision/udts/currencywaitlist"
	"ultimatedivision/users"
)

// ErrCurrencyWaitList is an internal error type for currency waitlist controller.
var ErrCurrencyWaitList = errs.Class("currency waitlist controller error")

// CurrencyWaitList is a mvc controller that handles all currency waitlist related views.
type CurrencyWaitList struct {
	log              logger.Logger
	currencywaitlist *currencywaitlist.Service
}

// NewCurrencyWaitList is a constructor for CurrencyWaitList controller.
func NewCurrencyWaitList(log logger.Logger, currencywaitlist *currencywaitlist.Service) *CurrencyWaitList {
	return &CurrencyWaitList{
		log:              log,
		currencywaitlist: currencywaitlist,
	}
}

// CreateRequest entity describes values which need for creating item of currency waitlist.
type CreateRequest struct {
	WalletAddress cryptoutils.Address `json:"walletAddress"`
}

// Create is an endpoint that creates item of currency waitlist.
func (controller *CurrencyWaitList) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	w.Header().Set("Content-Type", "application/json")

	_, err := auth.GetClaims(ctx)
	if err != nil {
		controller.serveError(w, http.StatusUnauthorized, ErrCurrencyWaitList.Wrap(err))
		return
	}

	createRequest := CreateRequest{}
	if err = json.NewDecoder(r.Body).Decode(&createRequest); err != nil {
		controller.serveError(w, http.StatusBadRequest, ErrCurrencyWaitList.Wrap(err))
		return
	}

	transaction, err := controller.currencywaitlist.Create(ctx, createRequest.WalletAddress)
	if err != nil {
		controller.log.Error("could not create nft token", ErrCurrencyWaitList.Wrap(err))

		if users.ErrNoUser.Has(err) {
			controller.serveError(w, http.StatusNotFound, ErrCurrencyWaitList.Wrap(err))
			return
		}

		if udts.ErrNoUDT.Has(err) {
			controller.serveError(w, http.StatusNotFound, ErrCurrencyWaitList.Wrap(err))
			return
		}

		controller.serveError(w, http.StatusInternalServerError, ErrCurrencyWaitList.Wrap(err))
		return
	}

	if err = json.NewEncoder(w).Encode(transaction); err != nil {
		controller.log.Error("failed to write json response", ErrCurrencyWaitList.Wrap(err))
		return
	}
}

// serveError replies to the request with specific code and error message.
func (controller *CurrencyWaitList) serveError(w http.ResponseWriter, status int, err error) {
	w.WriteHeader(status)
	var response struct {
		Error string `json:"error"`
	}
	response.Error = err.Error()

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		controller.log.Error("failed to write json error response", ErrCurrencyWaitList.Wrap(err))
	}
}
