// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package consoleserver

import (
	"encoding/json"
	"github.com/google/uuid"
	"net/http"

	"github.com/zeebo/errs"

	"ultimatedivision/clubs"
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

// team is a struct for response clubs, squad and squadCards.
type team struct {
	clubs      clubs.Club
	squad      clubs.Squads
	squadCards []clubs.SquadCards
}

// Get returns club, squad and squad cards by user id.
func (controller *Clubs) Get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	urlQuery := r.URL.Query()
	userID := urlQuery.Get("userId")

	id, err := uuid.Parse(userID)
	if err != nil {
		http.Error(w, "could not parse uuid", http.StatusBadRequest)
		return
	}

	club, err := controller.clubs.Get(ctx, id)
	if err != nil {
		controller.log.Error("could not get user club", ErrClubs.Wrap(err))
		http.Error(w, "could not get user club", http.StatusInternalServerError)
		return
	}

	squad, squadCards, err := controller.clubs.GetSquad(ctx, club.ID)
	if err != nil {
		controller.log.Error("could not get squad and squad cards", ErrClubs.Wrap(err))
		http.Error(w, "could not get squad and squad cards", http.StatusInternalServerError)
		return
	}

	userTeam := team{
		clubs:      club,
		squad:      squad,
		squadCards: squadCards,
	}

	if err = json.NewEncoder(w).Encode(userTeam); err != nil {
		controller.log.Error("failed to write json response", ErrClubs.Wrap(err))
		return
	}
}

// UpdatePosition updates card position in the squad.
func (controller *Clubs) UpdatePosition(w http.ResponseWriter, r *http.Request) {

}

// UpdateSquad updates squad tactic, capitan and formation.
func (controller *Clubs) UpdateSquad(w http.ResponseWriter, r *http.Request) {

}

// Add add new card to the squad.
func (controller *Clubs) Add(w http.ResponseWriter, r *http.Request) {

}

// Delete removes card from squad.
func (controller *Clubs) Delete(w http.ResponseWriter, r *http.Request) {

}
