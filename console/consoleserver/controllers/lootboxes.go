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
	ctx := r.Context()

	var userLootBox lootboxes.UserLootBoxes

	if err := json.NewDecoder(r.Body).Decode(&userLootBox); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	err := controller.lootBoxes.Create(ctx, userLootBox)
	if err != nil {
		controller.log.Error("could not create lootbox for user", ErrLootBoxes.Wrap(err))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
}

// Open is an endpoint that opens user lootbox.
func (controller *LootBoxes) Open(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var userLootBox lootboxes.UserLootBoxes

	if err := json.NewDecoder(r.Body).Decode(&userLootBox); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	cards, err := controller.lootBoxes.Open(ctx, userLootBox)
	if err != nil {
		controller.log.Error("could not open loot box", ErrLootBoxes.Wrap(err))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if err = json.NewEncoder(w).Encode(cards); err != nil {
		controller.log.Error("could not encode cards", ErrLootBoxes.Wrap(err))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
}
