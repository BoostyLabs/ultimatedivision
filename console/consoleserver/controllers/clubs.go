// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/zeebo/errs"

	"ultimatedivision/clubs"
	"ultimatedivision/internal/auth"
	"ultimatedivision/internal/logger"
)

var (
	// ErrClubs is an internal error type for clubs controller.
	ErrClubs = errs.Class("clubs controller error")
)

// Clubs is a mvc controller that handles all clubs related views.
type Clubs struct {
	log logger.Logger

	clubs *clubs.Service
}

// NewClubs is a constructor for clubs controller.
func NewClubs(log logger.Logger, clubs *clubs.Service) *Clubs {
	clubsController := &Clubs{
		log:   log,
		clubs: clubs,
	}

	return clubsController
}

// Create is an endpoint that creates new club.
func (controller *Clubs) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx := r.Context()

	claims, err := auth.GetClaims(ctx)
	if err != nil {
		controller.serveError(w, http.StatusUnauthorized, ErrClubs.Wrap(err))
		return
	}

	err = controller.clubs.Create(ctx, claims.UserID)
	if err != nil {
		controller.log.Error("could not create club", ErrClubs.Wrap(err))
		controller.serveError(w, http.StatusInternalServerError, ErrClubs.Wrap(err))
		return
	}
}

// Get is an endpoint that returns club, squad and squad cards by user id.
func (controller *Clubs) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx := r.Context()

	claims, err := auth.GetClaims(ctx)
	if err != nil {
		controller.serveError(w, http.StatusUnauthorized, ErrClubs.Wrap(err))
		return
	}

	club, err := controller.clubs.Get(ctx, claims.UserID)
	if err != nil {
		controller.log.Error("could not get club", ErrClubs.Wrap(err))
		switch {
		case clubs.ErrNoClub.Has(err):
			controller.serveError(w, http.StatusNotFound, ErrClubs.Wrap(err))
		default:
			controller.serveError(w, http.StatusInternalServerError, ErrClubs.Wrap(err))
		}
		return
	}

	squad, err := controller.clubs.GetSquad(ctx, club.ID)
	if err != nil {
		controller.log.Error("could not get squad", ErrClubs.Wrap(err))
		switch {
		case clubs.ErrNoClub.Has(err):
			controller.serveError(w, http.StatusNotFound, ErrClubs.Wrap(err))
		default:
			controller.serveError(w, http.StatusInternalServerError, ErrClubs.Wrap(err))
		}
		return
	}

	squadCards, err := controller.clubs.GetSquadCard(ctx, squad.ID)
	if err != nil {
		controller.log.Error("could not create squad", ErrClubs.Wrap(err))
		controller.serveError(w, http.StatusInternalServerError, ErrClubs.Wrap(err))
		return
	}

	clubResponse := ClubResponse{
		Clubs:      club,
		Squad:      squad,
		SquadCards: squadCards,
	}

	if err = json.NewEncoder(w).Encode(clubResponse); err != nil {
		controller.log.Error("could not response with json", ErrClubs.Wrap(err))
		return
	}
}

// CreateSquad is an endpoint that creates new squad for club.
func (controller *Clubs) CreateSquad(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx := r.Context()

	params := mux.Vars(r)
	if params["clubId"] == "" {
		controller.serveError(w, http.StatusBadRequest, ErrClubs.New("empty id parameter"))
		return
	}

	id, err := uuid.Parse(params["clubId"])
	if err != nil {
		controller.serveError(w, http.StatusBadRequest, ErrClubs.Wrap(err))
		return
	}

	err = controller.clubs.CreateSquad(ctx, id)
	if err != nil {
		controller.log.Error("could not create squad", ErrClubs.Wrap(err))
		controller.serveError(w, http.StatusInternalServerError, ErrClubs.Wrap(err))
		return
	}
}

// UpdateSquad is an endpoint that updates squad.
func (controller *Clubs) UpdateSquad(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx := r.Context()

	err := controller.clubs.UpdateSquad()
}

// ClubResponse is a struct for response clubs, squad and squadCards.
type ClubResponse struct {
	Clubs      clubs.Club        `json:"clubs"`
	Squad      clubs.Squad       `json:"squad"`
	SquadCards []clubs.SquadCard `json:"squadCards"`
}

// serveError replies to the request with specific code and error message.
func (controller *Clubs) serveError(w http.ResponseWriter, status int, err error) {
	w.WriteHeader(status)
	var response struct {
		Error string `json:"error"`
	}
	response.Error = err.Error()

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		controller.log.Error("failed to write json error response", ErrCards.Wrap(err))
	}
}
