// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package controllers

import (
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/zeebo/errs"
	"html/template"
	"net/http"
	"ultimatedivision/users/userauth"

	"ultimatedivision/internal/logger"
	"ultimatedivision/lootboxes"
)

// ErrLootBoxes is an internal error type for loot boxes controller.
var ErrLootBoxes = errs.Class("lootboxes controller error")

// LootBoxesTemplates holds all lootboxes related templates.
type LootBoxesTemplates struct {
	List   *template.Template
	Create *template.Template
	ListCards *template.Template
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

//  <a href="/lootboxes/create/{{.ID}}">Create lootboxes</a>

// Create is an endpoint that creates loot box for user.
func (controller *LootBoxes) Create(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		err := controller.templates.Create.Execute(w, nil)
		if err != nil {
			controller.log.Error("could not execute template", ErrLootBoxes.Wrap(err))
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
	case http.MethodPost:
		// TODO: set userID in context.

		ctx := r.Context()

		lootBoxType := r.FormValue("lootbox")

		err := controller.lootboxes.Create(ctx, lootBoxType)
		if err != nil {
			controller.log.Error("could not create loot box", ErrLootBoxes.Wrap(err))
			switch {
			case userauth.ErrUnauthenticated.Has(err):
				http.Error(w, ErrLootBoxes.Wrap(err).Error(), http.StatusUnauthorized)
			default:
				http.Error(w, ErrLootBoxes.Wrap(err).Error(), http.StatusInternalServerError)
				return
			}
		}

		Redirect(w, r , "/lootboxes", http.MethodGet)
	}
}

// Open is an endpoint that opens loot box by user.
func (controller *LootBoxes) Open(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	vars := mux.Vars(r)
	if vars["id"] == "" {
		http.Error(w, ErrLootBoxes.New("id parameter is empty").Error(), http.StatusBadRequest)
		return
	}

	id, err := uuid.Parse(vars["id"])
	if err != nil {
		http.Error(w, ErrLootBoxes.Wrap(err).Error(), http.StatusBadRequest)
		return
	}

	cards, err := controller.lootboxes.Open(ctx, id)
	if err != nil {
		controller.log.Error("could not create loot box", ErrLootBoxes.Wrap(err))
		switch {
		case userauth.ErrUnauthenticated.Has(err):
			http.Error(w, ErrLootBoxes.Wrap(err).Error(), http.StatusUnauthorized)
		case lootboxes.ErrNoLootBox.Has(err):
			http.Error(w, ErrLootBoxes.Wrap(err).Error(), http.StatusNotFound)
		default:
			http.Error(w, ErrLootBoxes.Wrap(err).Error(), http.StatusInternalServerError)
			return
		}
	}

	err = controller.templates.ListCards.Execute(w, cards)
	if err != nil {
		controller.log.Error("could not execute template", ErrLootBoxes.Wrap(err))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

// List is an endpoint that will provide a web page with all users loot boxes.
func (controller *LootBoxes) List(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	vars := mux.Vars(r)
	if vars["id"] == "" {
		http.Error(w, ErrLootBoxes.New("id parameter is empty").Error(), http.StatusBadRequest)
		return
	}

	id, err := uuid.Parse(vars["id"])
	if err != nil {
		http.Error(w, ErrLootBoxes.Wrap(err).Error(), http.StatusBadRequest)
		return
	}

	lootBoxes, err := controller.lootboxes.GetByUserID(ctx,id)
	if err != nil {
		controller.log.Error("could not create loot box", ErrLootBoxes.Wrap(err))
		switch {
		case userauth.ErrUnauthenticated.Has(err):
			http.Error(w, ErrLootBoxes.Wrap(err).Error(), http.StatusUnauthorized)
		default:
			http.Error(w, ErrLootBoxes.Wrap(err).Error(), http.StatusInternalServerError)
			return
		}
	}

	err = controller.templates.List.Execute(w, lootBoxes)
	if err != nil {
		controller.log.Error("could not execute template", ErrLootBoxes.Wrap(err))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
