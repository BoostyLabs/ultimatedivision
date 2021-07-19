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
	urlQuery := r.URL.Query()
	var filter cards.Filter
	var filters []cards.Filter

	for k, v := range urlQuery {
		if k == "player_name" {
			if err := ValidateKey(k); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			filter = cards.Filter{
				Key:    k,
				Action: "LIKE",
				Value:  ValidateValue(v[0]),
			}
		} else if k[:PrefixMinMaxValues] == "min_" {
			if err := ValidateKey(k[PrefixMinMaxValues:]); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			filter = cards.Filter{
				Key:    k[PrefixMinMaxValues:],
				Action: ">=",
				Value:  ValidateValue(v[0]),
			}
		} else if k[:PrefixMinMaxValues] == "max_" {
			if err := ValidateKey(k[PrefixMinMaxValues:]); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			filter = cards.Filter{
				Key:    k[PrefixMinMaxValues:],
				Action: "<=",
				Value:  ValidateValue(v[0]),
			}
		} else {
			if err := ValidateKey(k); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			filter = cards.Filter{
				Key:    k,
				Action: "=",
				Value:  ValidateValue(v[0]),
			}
		}
		filters = append(filters, filter)
	}

	cards, err := controller.cards.ListWithFilters(ctx, filters)
	if err != nil {
		controller.log.Error("could not get cards list", ErrCards.Wrap(err))
		http.Error(w, "could not get cards list", http.StatusInternalServerError)
		return
	}
	if len(cards) != 0 {
		fmt.Fprintf(w, "first card name - "+cards[0].PlayerName)
	} else {
		fmt.Fprintf(w, "not found cards")
	}
}
