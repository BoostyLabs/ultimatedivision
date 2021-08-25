// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"github.com/zeebo/errs"

	"ultimatedivision/cards"
	"ultimatedivision/internal/logger"
	"ultimatedivision/marketplace"
	"ultimatedivision/users/userauth"
)

var (
	// ErrMarketplace is an internal error type for marketplace controller.
	ErrMarketplace = errs.Class("marketplace controller error")
)

// MarketplaceTemplates holds all marketplace related templates.
type MarketplaceTemplates struct {
	List   *template.Template
	Create *template.Template
}

// Marketplace is a mvc controller that handles all marketplace related views.
type Marketplace struct {
	log logger.Logger

	marketplace *marketplace.Service
	cards       *cards.Service

	templates MarketplaceTemplates
}

// NewMarketplace is a constructor for marketplace controller.
func NewMarketplace(log logger.Logger, marketplace *marketplace.Service, cards *cards.Service, templates MarketplaceTemplates) *Marketplace {
	marketplaceController := &Marketplace{
		log:         log,
		marketplace: marketplace,
		cards:       cards,
		templates:   templates,
	}

	return marketplaceController
}

// ListActiveLots is an endpoint that will provide a web page with active lots.
func (controller *Marketplace) ListActiveLots(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	lots, err := controller.marketplace.ListActiveLots(ctx)
	if err != nil {
		controller.log.Error("could not get lots list", ErrMarketplace.Wrap(err))

		if userauth.ErrUnauthenticated.Has(err) {
			http.Error(w, ErrMarketplace.Wrap(err).Error(), http.StatusUnauthorized)
			return
		}

		if marketplace.ErrNoLot.Has(err) {
			http.Error(w, ErrMarketplace.Wrap(err).Error(), http.StatusNotFound)
			return
		}

		http.Error(w, "could not get lots list", http.StatusInternalServerError)
		return
	}

	err = controller.templates.List.Execute(w, lots)
	if err != nil {
		controller.log.Error("can not execute list lots template", ErrMarketplace.Wrap(err))
		http.Error(w, "can not execute list lots template", http.StatusInternalServerError)
		return
	}
}

// CreateLot is an endpoint that will add card to database.
func (controller *Marketplace) CreateLot(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	switch r.Method {
	case http.MethodGet:
		cardsList, err := controller.cards.List(ctx)
		if err != nil {
			controller.log.Error("could not get list cards", ErrMarketplace.Wrap(err))
			switch {
			case cards.ErrNoCard.Has(err):
				http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			case userauth.ErrUnauthenticated.Has(err):
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			default:
				http.Error(w, "could not get list cards", http.StatusInternalServerError)
			}

			return
		}

		if err = controller.templates.Create.Execute(w, cardsList); err != nil {
			controller.log.Error("could not execute create marketplace template", ErrMarketplace.Wrap(err))
			http.Error(w, "could not execute create marketplace template", http.StatusInternalServerError)
			return
		}
	case http.MethodPost:
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "could not parse admin create form", http.StatusBadRequest)
			return
		}

		itemIDForm := r.FormValue("itemId")
		if itemIDForm == "" {
			http.Error(w, "item id input is empty", http.StatusBadRequest)
			return
		}
		itemID, err := uuid.Parse(itemIDForm)
		if err != nil {
			http.Error(w, "could not parse item id", http.StatusBadRequest)
			return
		}

		startPriceForm := r.FormValue("startPrice")
		if startPriceForm == "" {
			http.Error(w, "start price input is empty", http.StatusBadRequest)
			return
		}
		strings.ToValidUTF8(startPriceForm, "")
		startPrice, err := strconv.ParseFloat(startPriceForm, 64)
		if err == nil {
			http.Error(w, fmt.Sprintf("%s %s", startPriceForm, err.Error()), http.StatusBadRequest)
		}

		maxPriceForm := r.FormValue("maxPrice")
		strings.ToValidUTF8(maxPriceForm, "")
		maxPrice, err := strconv.ParseFloat(maxPriceForm, 64)
		if err == nil {
			http.Error(w, fmt.Sprintf("%s %s", maxPriceForm, err.Error()), http.StatusBadRequest)
		}

		periodForm := r.FormValue("period")
		if periodForm == "" {
			http.Error(w, "period input is empty", http.StatusBadRequest)
			return
		}
		strings.ToValidUTF8(periodForm, "")
		period, err := strconv.Atoi(periodForm)
		if err == nil {
			http.Error(w, fmt.Sprintf("%s %s", periodForm, err.Error()), http.StatusBadRequest)
		}

		createLot := marketplace.CreateLot{
			ItemID:     itemID,
			StartPrice: startPrice,
			MaxPrice:   maxPrice,
			Period:     marketplace.Period(period),
		}

		if err := controller.marketplace.CreateLot(ctx, createLot); err != nil {
			controller.log.Error("could not create lot", ErrMarketplace.Wrap(err))
			switch {
			case userauth.ErrUnauthenticated.Has(err):
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			default:
				http.Error(w, "could not create lot", http.StatusInternalServerError)
			}

			return
		}

		Redirect(w, r, "/marketplace", "GET")
	}
}
