// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/zeebo/errs"

	"ultimatedivision/clubs"
	"ultimatedivision/internal/logger"
)

var (
	// ErrClubs is an internal error type for clubs controller.
	ErrClubs = errs.Class("clubs controller error")
)

// ClubsTemplates holds all clubs related templates.
type ClubsTemplates struct {
	List               *template.Template
	ListSquads         *template.Template
	UpdateSquad        *template.Template
	ListSquadCards     *template.Template
	UpdateCardPosition *template.Template
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
	ctx := r.Context()

	params := mux.Vars(r)
	if params["userID"] == "" {
		http.Error(w, ErrClubs.New("empty id parameter").Error(), http.StatusBadRequest)
	}

	id, err := uuid.Parse(params["userID"])
	if err != nil {
		http.Error(w, ErrClubs.New("could not parse uuid").Error(), http.StatusBadRequest)
		return
	}

	err = controller.clubs.Create(ctx, id)
	if err != nil {
		controller.log.Error("could not create club", ErrClubs.Wrap(err))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	url := fmt.Sprintf("/clubs/%s", id.String())

	Redirect(w, r, url, http.MethodGet)
}

// Get is an endpoint that will provide a web page with users club.
func (controller *Clubs) Get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	params := mux.Vars(r)
	if params["userID"] == "" {
		http.Error(w, "empty user id", http.StatusBadRequest)
		return
	}

	id, err := uuid.Parse(params["userID"])
	if err != nil {
		http.Error(w, "could not parse uuid", http.StatusBadRequest)
		return
	}

	club, err := controller.clubs.Get(ctx, id)
	if err != nil {
		controller.log.Error("could not get clubs list", ErrClubs.Wrap(err))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	err = controller.templates.List.Execute(w, club)
	if err != nil {
		controller.log.Error("can not execute list clubs template", ErrClubs.Wrap(err))
		http.Error(w, "can not execute list clubs template", http.StatusInternalServerError)
		return
	}
}

// CreateSquad is an endpoint thar creates squad for club.
func (controller *Clubs) CreateSquad(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	params := mux.Vars(r)
	if params["clubID"] == "" {
		http.Error(w, "empty club id", http.StatusBadRequest)
	}

	id, err := uuid.Parse(params["clubID"])
	fmt.Println(id, err)
	if err != nil {
		http.Error(w, "could not parse uuid", http.StatusBadRequest)
		return
	}

	err = controller.clubs.CreateSquad(ctx, id)
	if err != nil {
		controller.log.Error("could not create squad", ErrClubs.Wrap(err))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	url := fmt.Sprintf("/clubs/squad/%s", id.String())

	Redirect(w, r, url, http.MethodGet)
}

// GetSquad is an endpoint that will provide a web page with users squad of club.
func (controller *Clubs) GetSquad(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	params := mux.Vars(r)
	if params["clubID"] == "" {
		http.Error(w, ErrClubs.New("could not parse uuid").Error(), http.StatusBadRequest)
		return
	}

	id, err := uuid.Parse(params["clubID"])
	if err != nil {
		http.Error(w, ErrClubs.Wrap(err).Error(), http.StatusBadRequest)
		return
	}

	squad, err := controller.clubs.GetSquad(ctx, id)
	if err != nil {
		controller.log.Error("could not get squad", ErrClubs.Wrap(err))
		switch {
		case clubs.ErrNoSquad.Has(err):
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
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

// GetSquadCard is an endpoint that will provide a web page with cards of squad.
func (controller *Clubs) GetSquadCard(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	params := mux.Vars(r)
	if params["squadID"] == "" {
		http.Error(w, "empty squad id", http.StatusBadRequest)
		return
	}

	id, err := uuid.Parse(params["squadID"])
	if err != nil {
		http.Error(w, "could not parse uuid", http.StatusBadRequest)
		return
	}

	squadCard, err := controller.clubs.GetSquadCard(ctx, id)
	if err != nil {
		controller.log.Error("could not get squad card", ErrClubs.Wrap(err))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	err = controller.templates.ListSquadCards.Execute(w, squadCard)
	if err != nil {
		controller.log.Error("could not execute get card template", ErrClubs.Wrap(err))
		http.Error(w, "could not execute get card template", http.StatusInternalServerError)
		return
	}
}

// UpdatePosition is an endpoint that updates position of card in squad.
func (controller *Clubs) UpdatePosition(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	params := mux.Vars(r)
	if params["squadID"] == "" || params["cardID"] == "" {
		http.Error(w, ErrClubs.New("empty id parameter").Error(), http.StatusBadRequest)
		return
	}

	squadID, err := uuid.Parse(params["squadID"])
	if err != nil {
		http.Error(w, ErrClubs.Wrap(err).Error(), http.StatusBadRequest)
		return
	}

	cardID, err := uuid.Parse(params["cardID"])
	if err != nil {
		http.Error(w, ErrClubs.Wrap(err).Error(), http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		err = controller.templates.UpdateCardPosition.Execute(w, nil)
		if err != nil {
			controller.log.Error("could not execute update squad template", ErrClubs.Wrap(err))
			http.Error(w, "could not execute update squad template", http.StatusInternalServerError)
			return
		}
	case http.MethodPost:
		newPosition, err := strconv.Atoi(r.FormValue("position"))
		if err != nil {
			http.Error(w, "invalid new position value", http.StatusInternalServerError)
			return
		}

		err = controller.clubs.UpdateCardPosition(ctx, clubs.Position(newPosition), cardID)
		if err != nil {
			controller.log.Error("could not execute update squad template", ErrClubs.Wrap(err))
			switch {
			case clubs.ErrNoSquad.Has(err):
				http.Error(w, "squad does noe exists", http.StatusNotFound)
			default:
				http.Error(w, "could not update position of card", http.StatusInternalServerError)
			}
			return
		}

		// TODO: update url.
		url := fmt.Sprintf("/squad/%s", squadID)

		Redirect(w, r, url, http.MethodGet)
	}
}

func (controller *Clubs) Add(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()


}


// Users - get club
// Clubs - get all clubs / create club / create squad
// Cards - Add card to the squad
