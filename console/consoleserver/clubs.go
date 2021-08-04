// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package consoleserver

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
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

	userTeam := clubs.ClubResponse{
		Clubs:      club,
		Squad:      squad,
		SquadCards: squadCards,
	}

	if err = json.NewEncoder(w).Encode(userTeam); err != nil {
		controller.log.Error("failed to write json response", ErrClubs.Wrap(err))
		return
	}
}

// UpdatePosition updates card position in the squad.
func (controller *Clubs) UpdatePosition(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var squadCard clubs.SquadCards

	if err := json.NewDecoder(r.Body).Decode(&squadCard); err != nil {
		http.Error(w, "could not decode json", http.StatusBadRequest)
		return
	}

	err := controller.clubs.UpdateCardPosition(ctx, squadCard.ID, squadCard.CardID, squadCard.Position)
	if err != nil {
		controller.log.Error("could not update card position", ErrClubs.Wrap(err))
		http.Error(w, "could not update card position", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// UpdateSquad updates squad tactic, capitan and formation.
func (controller *Clubs) UpdateSquad(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var updatedSquad clubs.UpdateRequest

	if err := json.NewDecoder(r.Body).Decode(&updatedSquad); err != nil {
		http.Error(w, "could not decode json", http.StatusBadRequest)
		return
	}

	err := controller.clubs.UpdateSquad(ctx, updatedSquad.ID, updatedSquad.Tactic, updatedSquad.Formation)
	if err != nil {
		controller.log.Error("could not update squad", ErrClubs.Wrap(err))
		http.Error(w, "could not update squad", http.StatusInternalServerError)
		return
	}

	err = controller.clubs.UpdateCapitan(ctx, updatedSquad.ID, updatedSquad.Capitan)
	if err != nil {
		controller.log.Error("could not update squad", ErrClubs.Wrap(err))
		http.Error(w, "could not update squad", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// Add add new card to the squad.
func (controller *Clubs) Add(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var newSquadCard clubs.SquadCards

	if err := json.NewDecoder(r.Body).Decode(&newSquadCard); err != nil {
		http.Error(w, "could not parse json", http.StatusBadRequest)
		return
	}

	err := controller.clubs.Add(ctx, newSquadCard)
	if err != nil {
		controller.log.Error("could not add card to the squad", ErrClubs.Wrap(err))
		http.Error(w, "could not add card to the squad", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// Delete removes card from squad.
func (controller *Clubs) Delete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var squadCard clubs.SquadCards

	if err := json.NewDecoder(r.Body).Decode(&squadCard); err != nil {
		http.Error(w, "could not parse json", http.StatusBadRequest)
		return
	}

	err := controller.clubs.Delete(ctx, squadCard.ID, squadCard.CardID)
	if err != nil {
		controller.log.Error("could not delete card from the squad", ErrClubs.Wrap(err))
		http.Error(w, "could not delete card from the squad", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
