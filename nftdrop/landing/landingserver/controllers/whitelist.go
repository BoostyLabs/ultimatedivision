// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/zeebo/errs"

	"ultimatedivision/internal/logger"
	"ultimatedivision/nftdrop/whitelist"
)

var (
	// ErrWhitelist is an internal error type for whitelist controller.
	ErrWhitelist = errs.Class("whitelist controller error")
)

// Whitelist is a mvc controller that handles all whitelist related views.
type Whitelist struct {
	log logger.Logger

	whitelist *whitelist.Service
}

// NewWhitelist is a constructor for whitelist controller.
func NewWhitelist(log logger.Logger, whitelist *whitelist.Service) *Whitelist {
	whitelistController := &Whitelist{
		log:       log,
		whitelist: whitelist,
	}

	return whitelistController
}

// Create is an endpoint that allows to view details of whitelist.
func (controller *Whitelist) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	w.Header().Set("Content-Type", "application/json")

	var err error
	var request whitelist.Request
	if err = json.NewDecoder(r.Body).Decode(&request); err != nil {
		controller.serveError(w, http.StatusBadRequest, ErrWhitelist.Wrap(err))
		return
	}

	if isValid := request.ValidateAddress(); isValid != true {
		controller.serveError(w, http.StatusBadRequest, ErrWhitelist.New("the address is not valid"))
		return
	}

	if err = controller.whitelist.Create(ctx, request); err != nil {
		controller.log.Error("could not create whitelist", ErrWhitelist.Wrap(err))
		controller.serveError(w, http.StatusInternalServerError, ErrWhitelist.Wrap(err))
		return
	}
}

// List is an endpoint that allows to view all whitelist.
func (controller *Whitelist) List(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	w.Header().Set("Content-Type", "application/json")

	whitelistRecords, err := controller.whitelist.List(ctx)
	if err != nil {
		controller.log.Error("could not list whitelist", ErrWhitelist.Wrap(err))
		controller.serveError(w, http.StatusInternalServerError, ErrWhitelist.Wrap(err))
		return
	}

	if err = json.NewEncoder(w).Encode(whitelistRecords); err != nil {
		controller.log.Error("failed to write json response", ErrWhitelist.Wrap(err))
		return
	}
}

// serveError replies to the request with specific code and error message.
func (controller *Whitelist) serveError(w http.ResponseWriter, status int, err error) {
	w.WriteHeader(status)
	var response struct {
		Error string `json:"error"`
	}
	response.Error = err.Error()

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		controller.log.Error("failed to write json error response", ErrWhitelist.Wrap(err))
	}
}
