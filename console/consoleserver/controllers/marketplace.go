// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/zeebo/errs"

	"ultimatedivision/internal/logger"
	"ultimatedivision/marketplace"
	"ultimatedivision/users/userauth"
)

var (
	// ErrMarketplace is an internal error type for marketplace controller.
	ErrMarketplace = errs.Class("marketplace controller error")
)

// Marketplace is a mvc controller that handles all marketplace related views.
type Marketplace struct {
	log         logger.Logger
	marketplace *marketplace.Service
}

// NewMarketplace is a constructor for marketplace controller.
func NewMarketplace(log logger.Logger, marketplace *marketplace.Service) *Marketplace {
	marketplaceController := &Marketplace{
		log:         log,
		marketplace: marketplace,
	}

	return marketplaceController
}

// ResponseLot describes the values required to response for get lot by id.
type ResponseLot struct {
	ID           uuid.UUID          `json:"id"`
	ItemID       uuid.UUID          `json:"itemId"`
	Type         marketplace.Type   `json:"type"`
	Status       marketplace.Status `json:"status"`
	StartPrice   float64            `json:"startPrice"`
	MaxPrice     float64            `json:"maxPrice"`
	CurrentPrice float64            `json:"currentPrice"`
	StartTime    time.Time          `json:"startTime"`
	EndTime      time.Time          `json:"endTime"`
	Period       marketplace.Period `json:"period"`
}

// ListActiveLots is an endpoint that returns active lots list.
func (controller *Marketplace) ListActiveLots(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	w.Header().Set("Content-Type", "application/json")

	listActiveLots, err := controller.marketplace.ListActiveLots(ctx)
	if err != nil {
		controller.log.Error("could not get active lots list", ErrMarketplace.Wrap(err))

		if userauth.ErrUnauthenticated.Has(err) {
			controller.serveError(w, http.StatusUnauthorized, ErrMarketplace.Wrap(err))
			return
		}

		if marketplace.ErrNoLot.Has(err) {
			controller.serveError(w, http.StatusNotFound, ErrMarketplace.Wrap(err))
			return
		}

		controller.serveError(w, http.StatusInternalServerError, ErrMarketplace.Wrap(err))
		return
	}

	if err = json.NewEncoder(w).Encode(listActiveLots); err != nil {
		controller.log.Error("failed to write json response", ErrMarketplace.Wrap(err))
		return
	}
}

// GetLotByID is an endpoint that returns lot by id.
func (controller *Marketplace) GetLotByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")

	if vars["id"] == "" {
		controller.serveError(w, http.StatusBadRequest, ErrMarketplace.New("id parameter is empty"))
		return
	}

	id, err := uuid.Parse(vars["id"])
	if err != nil {
		controller.serveError(w, http.StatusBadRequest, ErrMarketplace.Wrap(err))
		return
	}

	lot, err := controller.marketplace.GetLotByID(ctx, id)
	if err != nil {
		controller.log.Error("could not get lot", ErrMarketplace.Wrap(err))

		if userauth.ErrUnauthenticated.Has(err) {
			controller.serveError(w, http.StatusUnauthorized, ErrMarketplace.Wrap(err))
			return
		}

		if marketplace.ErrNoLot.Has(err) {
			controller.serveError(w, http.StatusNotFound, ErrMarketplace.Wrap(err))
			return
		}

		controller.serveError(w, http.StatusInternalServerError, ErrMarketplace.Wrap(err))
		return
	}

	responseLot := ResponseLot{
		ID:           lot.ID,
		ItemID:       lot.ItemID,
		Type:         lot.Type,
		Status:       lot.Status,
		StartPrice:   lot.StartPrice,
		MaxPrice:     lot.MaxPrice,
		CurrentPrice: lot.CurrentPrice,
		StartTime:    lot.StartTime,
		EndTime:      lot.EndTime,
		Period:       lot.Period,
	}

	if err = json.NewEncoder(w).Encode(responseLot); err != nil {
		controller.log.Error("failed to write json response", ErrMarketplace.Wrap(err))
		return
	}
}

// CreateLot is an endpoint that returns lot by id.
func (controller *Marketplace) CreateLot(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	w.Header().Set("Content-Type", "application/json")

	var createLot marketplace.CreateLot
	if err := json.NewDecoder(r.Body).Decode(&createLot); err != nil {
		controller.serveError(w, http.StatusBadRequest, ErrMarketplace.Wrap(err))
		return
	}

	if err := createLot.Validate(); err != nil {
		controller.serveError(w, http.StatusBadRequest, ErrMarketplace.Wrap(err))
	}

	if err := controller.marketplace.CreateLot(ctx, createLot); err != nil {
		controller.log.Error("could not create lot", ErrMarketplace.Wrap(err))

		if userauth.ErrUnauthenticated.Has(err) {
			controller.serveError(w, http.StatusUnauthorized, ErrMarketplace.Wrap(err))
			return
		}

		controller.serveError(w, http.StatusInternalServerError, ErrMarketplace.Wrap(err))
		return
	}
}

// PlaceBetLot is an endpoint that returns lot by id.
func (controller *Marketplace) PlaceBetLot(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	w.Header().Set("Content-Type", "application/json")

	var betLot marketplace.BetLot
	if err := json.NewDecoder(r.Body).Decode(&betLot); err != nil {
		controller.serveError(w, http.StatusBadRequest, ErrMarketplace.Wrap(err))
		return
	}

	if betLot.ID.String() == "" {
		controller.serveError(w, http.StatusBadRequest, ErrMarketplace.New("id lot is empty"))
		return
	}

	if betLot.BetAmount == 0 {
		controller.serveError(w, http.StatusBadRequest, ErrMarketplace.New("bet amount is empty"))
		return
	}

	if err := controller.marketplace.PlaceBetLot(ctx, betLot); err != nil {
		controller.log.Error("could not place bet lot", ErrMarketplace.Wrap(err))

		if userauth.ErrUnauthenticated.Has(err) {
			controller.serveError(w, http.StatusUnauthorized, ErrMarketplace.Wrap(err))
			return
		}

		controller.serveError(w, http.StatusInternalServerError, ErrMarketplace.Wrap(err))
		return
	}
}

// serveError replies to the request with specific code and error message.
func (controller *Marketplace) serveError(w http.ResponseWriter, status int, err error) {
	w.WriteHeader(status)

	var response struct {
		Error string `json:"error"`
	}

	response.Error = err.Error()

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		controller.log.Error("failed to write json error response", ErrMarketplace.Wrap(err))
	}
}
