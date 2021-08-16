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
	"ultimatedivision/marketplace"
)

var (
	// ErrMarketplaces is an internal error type for marketplaces controller.
	ErrMarketplaces = errs.Class("marketplaces controller error")
)

// Marketplaces is a mvc controller that handles all marketplaces related views.
type Marketplaces struct {
	log logger.Logger

	marketplaces *marketplace.Service
}

// NewMarketplaces is a constructor for marketplaces controller.
func NewMarketplaces(log logger.Logger, marketplaces *marketplace.Service) *Marketplaces {
	marketplacesController := &Marketplaces{
		log:          log,
		marketplaces: marketplaces,
	}

	return marketplacesController
}

// ListActiveLots is an endpoint that returns active lots list.
func (controller *Marketplaces) ListActiveLots(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	w.Header().Set("Content-Type", "application/json")

	listActiveLots, err := controller.marketplaces.ListActiveLots(ctx)
	if err != nil {
		controller.log.Error("could not get active lots list", ErrMarketplaces.Wrap(err))
		controller.serveError(w, http.StatusInternalServerError, ErrMarketplaces.Wrap(err))
		return
	}

	if err = json.NewEncoder(w).Encode(listActiveLots); err != nil {
		controller.log.Error("failed to write json response", ErrMarketplaces.Wrap(err))
		return
	}
}

// GetLotByID is an endpoint that returns lot by id.
func (controller *Marketplaces) GetLotByID(w http.ResponseWriter, r *http.Request) {
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

	lot, err := controller.marketplaces.GetLotByID(ctx, id)
	if err != nil {
		controller.log.Error("could not get active lot", ErrMarketplaces.Wrap(err))
		controller.serveError(w, http.StatusInternalServerError, ErrMarketplaces.Wrap(err))
		return
	}

	responseLot := map[string]interface{}{
		"id":           lot.ID,
		"itemId":       lot.ItemID,
		"type":         lot.Type,
		"status":       lot.Status,
		"startPrice":   lot.StartPrice,
		"maxPrice":     lot.MaxPrice,
		"currentPrice": lot.CurrentPrice,
		"startTime":    lot.StartPrice,
		"endTime":      lot.EndTime,
	}

	if err = json.NewEncoder(w).Encode(responseLot); err != nil {
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
