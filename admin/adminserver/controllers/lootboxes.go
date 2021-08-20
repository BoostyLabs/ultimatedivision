// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package controllers

import (
	"html/template"
	"net/http"

	"github.com/zeebo/errs"

	"ultimatedivision/internal/logger"
	"ultimatedivision/lootboxes"
)

// ErrLootBoxes is an internal error type for loot boxes controller.
var ErrLootBoxes = errs.Class("lootboxes controller error")

// LootBoxesTemplates holds all lootboxes related templates.
type LootBoxesTemplates struct {
	List   *template.Template
	Create *template.Template
	Update *template.Template
}

// LootBoxes is a mvc controller that handles all lootboxes related views.
type LootBoxes struct {
	log logger.Logger

	lootboxes *lootboxes.Service

	templates LootBoxesTemplates
}

// NewLootBoxes is a constructor for loot boxes controller.
func NewLootBoxes(log logger.Logger, lootboxes *lootboxes.Service, templates LootBoxesTemplates) *LootBoxes {
	lootBoxesController := &LootBoxes{
		log:       log,
		lootboxes: lootboxes,
		templates: templates,
	}

	return lootBoxesController
}

// Update is an endpoint that updates loot boxes.
func (controller *LootBoxes) Update(w http.ResponseWriter, r *http.Request){
	switch r.Method {
	case http.MethodGet:
		err := controller.templates.Update.Execute(w, nil)
		if err != nil {
			controller.log.Error("could not execute update loot boxes template", ErrLootBoxes.Wrap(err))
			http.Error(w, "could not execute update loot boxes template", http.StatusInternalServerError)
			return
		}
	case http.MethodPost:

	}
}

// Create is an endpoint that creates loot boxes.
func (controller *LootBoxes) Create(w http.ResponseWriter, r *http.Request){
	switch r.Method {
	case http.MethodGet:
		err := controller.templates.Create.Execute(w, nil)
		if err != nil {
			controller.log.Error("could not execute create lootboxes template", ErrLootBoxes.Wrap(err))
			http.Error(w, "could not execute create lootboxes template", http.StatusInternalServerError)
			return
		}
	case http.MethodPost:

	}
}

