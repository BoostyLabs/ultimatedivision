// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/zeebo/errs"

	"ultimatedivision/internal/logger"
	"ultimatedivision/lootboxes"
)

var (
	// ErrLootBoxes is an internal error type for lootboxes controller.
	ErrLootBoxes = errs.Class("lootboxes controller error")
)

func (controller *LootBoxes) serveError(w http.ResponseWriter, status int, err error) {
	w.WriteHeader(status)

	var response struct {
		Error string `json:"error"`
	}

	response.Error = err.Error()

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		controller.log.Error("failed to write json error response", ErrLootBoxes.Wrap(err))
	}
}

// LootBoxes is a mvc controller that handles all lootboxes related views.
type LootBoxes struct {
	log logger.Logger

	lootBoxes *lootboxes.Service
}

// NewLootBoxes is a constructor for lootboxes controller.
func NewLootBoxes(log logger.Logger, lootBoxes *lootboxes.Service) *LootBoxes {
	lootBoxesController := &LootBoxes{
		log:       log,
		lootBoxes: lootBoxes,
	}

	return lootBoxesController
}

// Create is an endpoint that creates new lootbox for user.
func (controller *LootBoxes) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	ctx := r.Context()

	var lootBox lootboxes.LootBox

	if err := json.NewDecoder(r.Body).Decode(&lootBox); err != nil {
		controller.serveError(w, http.StatusBadRequest, ErrLootBoxes.Wrap(err))
		return
	}

	err := controller.lootBoxes.Create(ctx, lootBox)
	if err != nil {
		controller.log.Error("could not create lootbox for user", ErrLootBoxes.Wrap(err))
		controller.serveError(w, http.StatusInternalServerError, ErrLootBoxes.Wrap(err))
		return
	}
}

// Open is an endpoint that opens user lootbox.
func (controller *LootBoxes) Open(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	ctx := r.Context()

	var lootBox lootboxes.LootBox

	if err := json.NewDecoder(r.Body).Decode(&lootBox); err != nil {
		controller.serveError(w, http.StatusBadRequest, ErrLootBoxes.Wrap(err))
		return
	}

	cards, err := controller.lootBoxes.Open(ctx, lootBox)
	if err != nil {
		controller.log.Error("could not open loot box", ErrLootBoxes.Wrap(err))
		controller.serveError(w, http.StatusInternalServerError, ErrLootBoxes.Wrap(err))
		return
	}

	if err = json.NewEncoder(w).Encode(cards); err != nil {
		controller.log.Error("could not encode cards", ErrLootBoxes.Wrap(err))
		controller.serveError(w, http.StatusInternalServerError, ErrLootBoxes.Wrap(err))
		return
	}
}
