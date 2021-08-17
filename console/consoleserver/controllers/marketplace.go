// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/zeebo/errs"

	"ultimatedivision/internal/logger"
	"ultimatedivision/marketplace"
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
		controller.log.Error("could not get active lots list", marketplace.ErrNoLot.Wrap(err))
		controller.serveError(w, http.StatusInternalServerError, marketplace.ErrNoLot.Wrap(err))
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
		controller.serveError(w, http.StatusBadRequest, ErrMarketplace.Wrap(fmt.Errorf("id parameter is empty")))
		return
	}

	id, err := uuid.Parse(vars["id"])
	if err != nil {
		controller.serveError(w, http.StatusBadRequest, ErrMarketplace.Wrap(err))
		return
	}

	lot, err := controller.marketplace.GetLotByID(ctx, id)
	if err != nil {
		controller.log.Error("could not get lot", marketplace.ErrNoLot.Wrap(err))
		controller.serveError(w, http.StatusInternalServerError, marketplace.ErrNoLot.Wrap(err))
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

	err := r.ParseForm()
	if err != nil {
		controller.serveError(w, http.StatusBadRequest, ErrMarketplace.Wrap(fmt.Errorf("could not parse lot create form")))
		return
	}

	itemIDForm := r.FormValue("itemId")
	if itemIDForm == "" {
		controller.serveError(w, http.StatusBadRequest, ErrMarketplace.Wrap(fmt.Errorf("itemIDForm input is empty")))
		return
	}
	itemID, err := uuid.Parse(itemIDForm)
	if err != nil {
		controller.serveError(w, http.StatusBadRequest, ErrMarketplace.Wrap(fmt.Errorf("could not parse item id")))
		return
	}

	startPriceForm := r.FormValue("startPrice")
	if startPriceForm == "" {
		controller.serveError(w, http.StatusBadRequest, ErrMarketplace.Wrap(fmt.Errorf("startPriceForm input is empty")))
		return
	}
	startPrice, err := strconv.ParseFloat(startPriceForm, 64)
	if err != nil {
		controller.serveError(w, http.StatusBadRequest, ErrMarketplace.Wrap(fmt.Errorf("could not parse start price")))
		return
	}

	maxPriceForm := r.FormValue("maxPrice")
	maxPrice := 0.0
	if maxPriceForm != "" {
		maxPrice, err = strconv.ParseFloat(maxPriceForm, 64)
		if err != nil {
			controller.serveError(w, http.StatusBadRequest, ErrMarketplace.Wrap(fmt.Errorf("could not parse max price")))
			return
		}
	}

	periodForm := r.FormValue("period")
	if periodForm == "" {
		controller.serveError(w, http.StatusBadRequest, ErrMarketplace.Wrap(fmt.Errorf("periodForm input is empty")))
		return
	}
	period, err := strconv.Atoi(periodForm)
	if err != nil {
		controller.serveError(w, http.StatusBadRequest, ErrMarketplace.Wrap(fmt.Errorf("could not parse period")))
		return
	}

	createLot := marketplace.CreateLot{
		ItemID: itemID,
		// TODO: userID with token
		StartPrice: startPrice,
		MaxPrice:   maxPrice,
		Period:     marketplace.Period(period),
	}

	err = controller.marketplace.CreateLot(ctx, createLot)
	if err != nil {
		controller.log.Error("could not create lot", ErrMarketplace.Wrap(err))
		controller.serveError(w, http.StatusInternalServerError, ErrMarketplace.Wrap(err))
		return
	}
}

// PlaceBetLot is an endpoint that returns lot by id.
func (controller *Marketplace) PlaceBetLot(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	w.Header().Set("Content-Type", "application/json")

	IDForm := r.FormValue("Id")
	if IDForm == "" {
		controller.serveError(w, http.StatusBadRequest, ErrMarketplace.Wrap(fmt.Errorf("IDForm input is empty")))
		return
	}
	id, err := uuid.Parse(IDForm)
	if err != nil {
		controller.serveError(w, http.StatusBadRequest, ErrMarketplace.Wrap(fmt.Errorf("could not parse id")))
		return
	}

	betAmountForm := r.FormValue("betAmount")
	if betAmountForm == "" {
		controller.serveError(w, http.StatusBadRequest, ErrMarketplace.Wrap(fmt.Errorf("betAmountForm input is empty")))
		return
	}
	betAmount, err := strconv.ParseFloat(betAmountForm, 64)
	if err != nil {
		controller.serveError(w, http.StatusBadRequest, ErrMarketplace.Wrap(fmt.Errorf("could not parse bet amount")))
		return
	}

	betLot := marketplace.BetLot{
		ID: id,
		// TODO: shopperID with token
		BetAmount: betAmount,
	}

	err = controller.marketplace.PlaceBetLot(ctx, betLot)
	if err != nil {
		controller.log.Error("could not place bet lot", ErrMarketplace.Wrap(err))
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
