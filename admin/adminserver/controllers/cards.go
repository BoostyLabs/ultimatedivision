// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package controllers

import (
	"html/template"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/zeebo/errs"

	"ultimatedivision/cards"
	"ultimatedivision/internal/logger"
	"ultimatedivision/users/userauth"
)

var (
	// ErrCards is an internal error type for cards controller.
	ErrCards = errs.Class("cards controller error")
)

// CardTemplates holds all cards related templates.
type CardTemplates struct {
	List *template.Template
}

// Cards is a mvc controller that handles all cards related views.
type Cards struct {
	log logger.Logger

	cards *cards.Service

	templates CardTemplates

	percentageQualities cards.PercentageQualities
}

// NewCards is a constructor for cards controller.
func NewCards(log logger.Logger, cards *cards.Service, templates CardTemplates, percentageQualities cards.PercentageQualities) *Cards {
	cardsController := &Cards{
		log:                 log,
		cards:               cards,
		templates:           templates,
		percentageQualities: percentageQualities,
	}

	return cardsController
}

// List is an endpoint that will provide a web page with all cards.
func (controller *Cards) List(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	listCards, err := controller.cards.List(ctx, cards.Pagination{})
	if err != nil {
		controller.log.Error("could not get cards list", ErrCards.Wrap(err))
		switch {
		case userauth.ErrUnauthenticated.Has(err):
			http.Error(w, ErrCards.Wrap(err).Error(), http.StatusUnauthorized)
		case cards.ErrNoCard.Has(err):
			http.Error(w, ErrCards.Wrap(err).Error(), http.StatusNotFound)
		default:
			http.Error(w, ErrCards.Wrap(err).Error(), http.StatusInternalServerError)
		}
		return
	}

	err = controller.templates.List.Execute(w, listCards)
	if err != nil {
		controller.log.Error("can not execute list cards template", ErrCards.Wrap(err))
		http.Error(w, ErrCards.Wrap(err).Error(), http.StatusInternalServerError)
		return
	}
}

// Create is an endpoint that will add card to database.
func (controller *Cards) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)
	if vars["userID"] == "" {
		http.Error(w, "userID parameter is empty", http.StatusBadRequest)
		return
	}
	userID, err := uuid.Parse(vars["userID"])
	if err != nil {
		http.Error(w, "could not parse user id", http.StatusBadRequest)
		return
	}

	percentageQualities := []int{controller.percentageQualities.Wood, controller.percentageQualities.Silver, controller.percentageQualities.Gold, controller.percentageQualities.Diamond}

	if _, err := controller.cards.Create(ctx, userID, percentageQualities); err != nil {
		controller.log.Error("could not create card", ErrCards.Wrap(err))

		if userauth.ErrUnauthenticated.Has(err) {
			http.Error(w, ErrCards.Wrap(err).Error(), http.StatusUnauthorized)
			return
		}

		http.Error(w, ErrCards.Wrap(err).Error(), http.StatusInternalServerError)
		return
	}

	Redirect(w, r, "/cards", "GET")
}

// Delete is an endpoint that will destroy record card to database.
func (controller *Cards) Delete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)
	if vars["id"] == "" {
		http.Error(w, "id parameter is empty", http.StatusBadRequest)
		return
	}
	id, err := uuid.Parse(vars["id"])
	if err != nil {
		http.Error(w, "could not parse card id", http.StatusBadRequest)
		return
	}
	if err := controller.cards.Delete(ctx, id); err != nil {
		controller.log.Error("could not delete card", ErrCards.Wrap(err))

		if userauth.ErrUnauthenticated.Has(err) {
			http.Error(w, ErrCards.Wrap(err).Error(), http.StatusUnauthorized)
			return
		}

		http.Error(w, ErrCards.Wrap(err).Error(), http.StatusInternalServerError)
		return
	}

	Redirect(w, r, "/cards", "GET")
}
