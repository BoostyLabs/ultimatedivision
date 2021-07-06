// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package controllers

import (
	"html/template"
	"net/http"
	"net/url"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/zeebo/errs"

	"ultimatedivision/cards"
	"ultimatedivision/internal/logger"
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
}

// NewCards is a constructor for cards controller.
func NewCards(log logger.Logger, cards *cards.Service, templates CardTemplates) *Cards {
	managersController := &Cards{
		log:       log,
		cards:     cards,
		templates: templates,
	}

	return managersController
}

// List is an endpoint that will provide a web page with all cards.
func (controller *Cards) List(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	cards, err := controller.cards.List(ctx)
	if err != nil {
		controller.log.Error("could not get cards list", ErrCards.Wrap(err))
		http.Error(w, "could not get cards list", http.StatusInternalServerError)
		return
	}

	err = controller.templates.List.Execute(w, cards)
	if err != nil {
		controller.log.Error("can not execute list cards template", ErrCards.Wrap(err))
		http.Error(w, "can not execute list cards template", http.StatusInternalServerError)
		return
	}
}

// Create is an endpoint that will add card to database.
func (controller *Cards) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// TODO:func generate card
	card := cards.Card{
		ID:               uuid.New(),
		PlayerName:       "Dmytro",
		Quality:          "bronze",
		PictureType:      1,
		Height:           178.8,
		Weight:           72.2,
		SkinColor:        1,
		HairStyle:        1,
		HairColor:        1,
		Accessories:      []int{1, 2},
		DominantFoot:     "left",
		UserID:           uuid.MustParse("7549b8d8-8937-4af4-be0b-1be2a5bdb4c9"),
		Tactics:          1,
		Positioning:      2,
		Composure:        3,
		Aggression:       4,
		Vision:           5,
		Awareness:        6,
		Crosses:          7,
		Physique:         8,
		Acceleration:     9,
		RunningSpeed:     10,
		ReactionSpeed:    11,
		Agility:          12,
		Stamina:          13,
		Strength:         14,
		Jumping:          15,
		Balance:          16,
		Technique:        17,
		Dribbling:        18,
		BallControl:      19,
		WeakFoot:         20,
		SkillMoves:       21,
		Finesse:          22,
		Curve:            23,
		Volleys:          24,
		ShortPassing:     25,
		LongPassing:      26,
		ForwardPass:      27,
		Offense:          28,
		FinishingAbility: 29,
		ShotPower:        30,
		Accuracy:         31,
		Distance:         32,
		Penalty:          33,
		FreeKicks:        34,
		Corners:          35,
		HeadingAccuracy:  36,
		Defence:          37,
		OffsideTrap:      38,
		Sliding:          39,
		Tackles:          40,
		BallFocus:        41,
		Interceptions:    42,
		Vigilance:        43,
		Goalkeeping:      44,
		Reflexes:         45,
		Diving:           46,
		Handling:         47,
		Sweeping:         48,
		Throwing:         49,
	}

	if err := controller.cards.Create(ctx, card); err != nil {
		controller.log.Error("could not create card", ErrCards.Wrap(err))
		http.Error(w, "could not create card", http.StatusInternalServerError)
		return
	}

	controller.Redirect(w, r, "", "GET")
}

// Delete is an endpoint that will destroy record card to database.
func (controller *Cards) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := uuid.MustParse(vars["id"])
	ctx := r.Context()

	if err := controller.cards.Delete(ctx, id); err != nil {
		controller.log.Error("could not delete card", ErrCards.Wrap(err))
		http.Error(w, "could not delete card", http.StatusInternalServerError)
		return
	}

	controller.Redirect(w, r, "/cards", "GET")
}

// Redirect redirects to specific url.
func (controller *Cards) Redirect(w http.ResponseWriter, r *http.Request, urlString, method string) {
	newRequest := new(http.Request)
	*newRequest = *r
	newRequest.URL = new(url.URL)
	*newRequest.URL = *r.URL
	newRequest.Method = method

	http.Redirect(w, newRequest, urlString, http.StatusFound)
}
