// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package controllers

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/zeebo/errs"

	"ultimatedivision/admin/adminauth"
	"ultimatedivision/clubs"
	"ultimatedivision/internal/logger"
)

var (
	// ErrClubs is an internal error type for clubs controller.
	ErrClubs = errs.Class("clubs controller error")
)

// ClubsTemplates holds all clubs related templates.
type ClubsTemplates struct {
	List           *template.Template
	ListSquads     *template.Template
	UpdateSquad    *template.Template
	ListSquadCards *template.Template
}

// Clubs is a mvc controller that handles all clubs related views.
type Clubs struct {
	log logger.Logger

	clubs *clubs.Service

	templates ClubsTemplates
}

// NewClubs is a constructor for clubs controller.
func NewClubs(log logger.Logger, clubs *clubs.Service, templates ClubsTemplates) *Clubs {
	clubsController := &Clubs{
		log:       log,
		clubs:     clubs,
		templates: templates,
	}

	return clubsController
}

// Create is an endpoint that will create club.
func (controller *Clubs) Create(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idParam := params["userID"]

	id, err := uuid.Parse(idParam)
	if err != nil {
		http.Error(w, "could not parse uuid", http.StatusBadRequest)
		return
	}

	ctx := clubs.SetUserID(r.Context(), id)

	err = controller.clubs.Create(ctx)
	if err != nil {
		controller.log.Error("could not create club", ErrClubs.Wrap(err))

		switch {
		case clubs.ErrClubs.Has(err):
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		case adminauth.ErrUnauthenticated.Has(err):
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusInternalServerError)
		}

		return
	}

	url := fmt.Sprintf("/clubs/%s", id.String())

	Redirect(w, r, url, http.MethodGet)
}

// Get is an endpoint that will provide a web page with users club.
func (controller *Clubs) Get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	params := mux.Vars(r)
	idParam := params["userID"]

	id, err := uuid.Parse(idParam)
	if err != nil {
		http.Error(w, "could not parse uuid", http.StatusBadRequest)
		return
	}

	ctx = clubs.SetUserID(ctx, id)

	clubs, err := controller.clubs.Get(ctx)
	if err != nil {
		controller.log.Error("could not get clubs list", ErrClubs.Wrap(err))
		switch {
		case adminauth.ErrUnauthenticated.Has(err):
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		default:
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}

		return
	}

	err = controller.templates.List.Execute(w, clubs)
	if err != nil {
		controller.log.Error("can not execute list clubs template", ErrClubs.Wrap(err))
		http.Error(w, "can not execute list clubs template", http.StatusInternalServerError)
		return
	}
}

// GetSquad is an endpoint that will provide a web page with users squad of club.
func (controller *Clubs) GetSquad(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	params := mux.Vars(r)
	idParam := params["clubID"]

	fmt.Println(idParam)

	id, err := uuid.Parse(idParam)
	if err != nil {
		http.Error(w, "could not parse uuid", http.StatusBadRequest)
		return
	}

	squad, err := controller.clubs.GetSquad(ctx, id)
	if err != nil {
		controller.log.Error("could not get squad", ErrClubs.Wrap(err))
		switch {
		case clubs.ErrNoSquad.Has(err):
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		case adminauth.ErrUnauthenticated.Has(err):
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		default:
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}

		return
	}

	err = controller.templates.ListSquads.Execute(w, squad)
	if err != nil {
		controller.log.Error("could not execute get squad template", ErrClubs.Wrap(err))
		http.Error(w, "could not execute get squad template", http.StatusInternalServerError)
		return
	}
}

// GetSquadCard is an endpoint that will provide a web page with cards of squad.
func (controller *Clubs) GetSquadCard(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	params := mux.Vars(r)
	idParam := params["squadID"]

	id, err := uuid.Parse(idParam)
	if err != nil {
		http.Error(w, "could not parse uuid", http.StatusBadRequest)
		return
	}

	squadCard, err := controller.clubs.GetSquadCard(ctx, id)
	if err != nil {
		controller.log.Error("could not get squad card", ErrClubs.Wrap(err))
		switch {
		case adminauth.ErrUnauthenticated.Has(err):
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		default:
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}

		return
	}

	err = controller.templates.ListSquadCards.Execute(w, squadCard)
	if err != nil {
		controller.log.Error("could not execute get card template", ErrClubs.Wrap(err))
		http.Error(w, "could not execute get card template", http.StatusInternalServerError)
		return
	}
}

// CreateSquad is an endpoint thar creates squad for club.
func (controller *Clubs) CreateSquad(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	params := mux.Vars(r)
	idParam := params["clubID"]
	fmt.Println(idParam)
	if idParam == "" {
		http.Error(w, "empty id param", http.StatusBadRequest)
	}

	id, err := uuid.Parse(idParam)
	fmt.Println(id, err)
	if err != nil {
		http.Error(w, "could not parse uuid", http.StatusBadRequest)
		return
	}

	err = controller.clubs.CreateSquad(ctx, id)
	if err != nil {
		controller.log.Error("could not create squad", ErrClubs.Wrap(err))

		switch {
		case adminauth.ErrUnauthenticated.Has(err):
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		default:
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}

		return
	}

	url := fmt.Sprintf("/clubs/squad/%s", id.String())

	Redirect(w, r, url, http.MethodGet)
}

// UpdateSquad is an endpoint that updates squad formation, tactic.
func (controller *Clubs) UpdateSquad(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		err := controller.templates.UpdateSquad.Execute(w, nil)
		if err != nil {
			controller.log.Error("could not execute update squad template", ErrClubs.Wrap(err))
			http.Error(w, "could not execute update squad template", http.StatusInternalServerError)
			return
		}
	case http.MethodPost:

	}
}
