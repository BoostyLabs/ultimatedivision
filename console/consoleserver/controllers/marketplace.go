// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/zeebo/errs"

	"ultimatedivision/internal/logger"
	"ultimatedivision/marketplaces"
)

var (
	// ErrMarketplaces is an internal error type for marketplaces controller.
	ErrMarketplaces = errs.Class("marketplaces controller error")
)

// Marketplaces is a mvc controller that handles all marketplaces related views.
type Marketplaces struct {
	log logger.Logger

	marketplaces *marketplaces.Service
}

// NewMarketplaces is a constructor for marketplaces controller.
func NewMarketplaces(log logger.Logger, marketplaces *marketplaces.Service) *Marketplaces {
	marketplacesController := &Marketplaces{
		log:          log,
		marketplaces: marketplaces,
	}

	return marketplacesController
}

// ListActive is an endpoint that returns active lots list.
func (controller *Marketplaces) ListActive(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	w.Header().Set("Content-Type", "application/json")

	lotsListActive, err := controller.marketplaces.ListActive(ctx)
	if err != nil {
		controller.log.Error("could not get active lots list", ErrMarketplaces.Wrap(err))
		controller.serveError(w, http.StatusInternalServerError, ErrMarketplaces.Wrap(err))
		return
	}

	if err = json.NewEncoder(w).Encode(lotsListActive); err != nil {
		controller.log.Error("failed to write json response", ErrMarketplaces.Wrap(err))
		return
	}
}

// Get is an endpoint that returns lot by id.
func (controller *Marketplaces) Get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")

	if vars["id"] == "" {
		controller.serveError(w, http.StatusBadRequest, ErrMarketplaces.Wrap(fmt.Errorf("id parameter is empty")))
		return
	}

	id, err := uuid.Parse(vars["id"])
	if err != nil {
		controller.serveError(w, http.StatusBadRequest, ErrMarketplaces.Wrap(err))
		return
	}

	lot, err := controller.marketplaces.Get(ctx, id)
	if err != nil {
		controller.log.Error("could not get active lot", ErrMarketplaces.Wrap(err))
		controller.serveError(w, http.StatusInternalServerError, ErrMarketplaces.Wrap(err))
		return
	}

	if err = json.NewEncoder(w).Encode(lot); err != nil {
		controller.log.Error("failed to write json response", ErrMarketplaces.Wrap(err))
		return
	}
}

// serveError replies to the request with specific code and error message.
func (controller *Marketplaces) serveError(w http.ResponseWriter, status int, err error) {
	w.WriteHeader(status)

	var response struct {
		Error string `json:"error"`
	}

	response.Error = err.Error()

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		controller.log.Error("failed to write json error response", ErrMarketplaces.Wrap(err))
	}
}
