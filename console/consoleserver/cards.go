// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package consoleserver

import (
	"fmt"
	"net/http"

	"github.com/zeebo/errs"

	"ultimatedivision/cards"
	"ultimatedivision/internal/logger"
)

const (
	// PrefixMinMaxValues the number of first characters for min or max values.
	PrefixMinMaxValues = 4
)

var (
	// ErrCards is an internal error type for cards controller.
	ErrCards = errs.Class("cards controller error")
)

// Cards is a mvc controller that handles all cards related views.
type Cards struct {
	log logger.Logger

	cards *cards.Service
}

// NewCards is a constructor for cards controller.
func NewCards(log logger.Logger, cards *cards.Service) *Cards {
	cardsController := &Cards{
		log:   log,
		cards: cards,
	}

	return cardsController
}

// List is an endpoint that will provide a web page with all cards.
func (controller *Cards) List(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var cardsList []cards.Card
	var err error
	urlQuery := r.URL.Query()
	tactics := urlQuery.Get(string(cards.Tactics))
	minPhysique := urlQuery.Get(string(cards.MinPhysique))
	maxPhysique := urlQuery.Get(string(cards.MaxPhysique))
	playerName := urlQuery.Get(string(cards.PlayerName))
	var filter cards.Filter
	var filters []cards.Filter

	if tactics != "" {
		filter = cards.Filter{
			cards.Tactics: tactics,
		}
		filters = append(filters, filter)
	}

	if minPhysique != "" {
		filter = cards.Filter{
			cards.MinPhysique: minPhysique,
		}
		filters = append(filters, filter)
	}

	if maxPhysique != "" {
		filter = cards.Filter{
			cards.MaxPhysique: maxPhysique,
		}
		filters = append(filters, filter)
	}

	if playerName != "" {
		filter = cards.Filter{
			cards.PlayerName: playerName,
		}
		filters = append(filters, filter)
	}

	if len(filters) > 0 {
		cardsList, err = controller.cards.ListWithFilters(ctx, filters)
	} else {
		cardsList, err = controller.cards.List(ctx)
	}
	if err != nil {
		controller.log.Error("could not get cards list", ErrCards.Wrap(err))
		http.Error(w, "could not get cards list", http.StatusInternalServerError)
		return
	}
	if len(cardsList) != 0 {
		fmt.Fprintf(w, "first card name - "+cardsList[0].PlayerName)
	} else {
		fmt.Fprintf(w, "not found cards")
	}
}
