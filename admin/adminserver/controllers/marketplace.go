// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package controllers

import (
	"html/template"
	"net/http"

	"github.com/zeebo/errs"

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
	List *template.Template
}

// Marketplace is a mvc controller that handles all marketplace related views.
type Marketplace struct {
	log logger.Logger

	marketplace *marketplace.Service

	templates CardTemplates
}

// NewMarketplace is a constructor for marketplace controller.
func NewMarketplace(log logger.Logger, marketplace *marketplace.Service, templates CardTemplates) *Marketplace {
	marketplaceController := &Marketplace{
		log:         log,
		marketplace: marketplace,
		templates:   templates,
	}

	return marketplaceController
}

// List is an endpoint that will provide a web page with active lots.
func (controller *Marketplace) List(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	w.Header().Set("Content-Type", "application/json")

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
