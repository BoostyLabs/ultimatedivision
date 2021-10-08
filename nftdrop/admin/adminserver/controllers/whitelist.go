// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package controllers

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zeebo/errs"

	"ultimatedivision/internal/logger"
	"ultimatedivision/nftdrop/whitelist"
)

var (
	// ErrWhitelist is an internal error type for whitelist controller.
	ErrWhitelist = errs.Class("whitelist controller error")
)

// WhitelistTemplates holds all whitelist related templates.
type WhitelistTemplates struct {
	List        *template.Template
	Create      *template.Template
	SetPassword *template.Template
}

// Whitelist is a mvc controller that handles all whitelist related views.
type Whitelist struct {
	log logger.Logger

	whitelist *whitelist.Service

	templates WhitelistTemplates
}

// NewWhitelist is a constructor for whitelist controller.
func NewWhitelist(log logger.Logger, whitelist *whitelist.Service, templates WhitelistTemplates) *Whitelist {
	whitelistController := &Whitelist{
		log:       log,
		whitelist: whitelist,
		templates: templates,
	}

	return whitelistController
}

// Create is an endpoint that creates new item in whitelist.
func (controller *Whitelist) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	switch r.Method {
	case http.MethodGet:
		err := controller.templates.Create.Execute(w, nil)
		if err != nil {
			controller.log.Error("could not execute create whitelist item template", ErrWhitelist.Wrap(err))
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	case http.MethodPost:
		err := r.ParseForm()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		var request whitelist.Request
		request.Address = whitelist.Address(r.FormValue("address"))
		if !request.Address.ValidateAddress() {
			http.Error(w, errs.New("invalid wallet address").Error(), http.StatusBadRequest)
			return
		}

		request.Key = r.FormValue("key")

		err = controller.whitelist.Create(ctx, request)
		if err != nil {
			controller.log.Error("could not create whitelist item", ErrWhitelist.Wrap(err))
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		Redirect(w, r, "", http.MethodGet)
	}
}

// List is an endpoint that will provide a web page with all whitelist items.
func (controller *Whitelist) List(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	whitelist, err := controller.whitelist.List(ctx)
	if err != nil {
		controller.log.Error("could not list whitelist", ErrWhitelist.Wrap(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = controller.templates.List.Execute(w, whitelist)
	if err != nil {
		controller.log.Error("could not execute list whitelist template", ErrWhitelist.Wrap(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// Delete is an endpoint that deletes whitelist item.
func (controller *Whitelist) Delete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	params := mux.Vars(r)

	walletAddress := whitelist.Address(params["address"])
	if !walletAddress.ValidateAddress() {
		http.Error(w, errs.New("invalid wallet address").Error(), http.StatusBadRequest)
		return
	}

	err := controller.whitelist.Delete(ctx, walletAddress)
	if err != nil {
		controller.log.Error("could not delete whitelist item", ErrWhitelist.Wrap(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	Redirect(w, r, "/whitelist", http.MethodGet)
}

// SetPassword is an endpoint that sets password for all whitelist item.
func (controller *Whitelist) SetPassword(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	switch r.Method {
	case http.MethodGet:
		err := controller.templates.SetPassword.Execute(w, nil)
		if err != nil {
			controller.log.Error("could not execute set password to whitelist items template", ErrWhitelist.Wrap(err))
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	case http.MethodPost:
		params := mux.Vars(r)

		key := params["key"]

		err := controller.whitelist.SetPassword(ctx, key)
		if err != nil {
			controller.log.Error("could not set password", ErrWhitelist.Wrap(err))
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		Redirect(w, r, "", http.MethodGet)
	}
}
