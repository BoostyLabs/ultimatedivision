// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package controllers

import (
	"encoding/json"
	"html/template"
	"net/http"
	"strconv"
	"ultimatedivision/pkg/pagination"

	"github.com/zeebo/errs"

	"ultimatedivision/internal/logger"
	"ultimatedivision/nftdrop/subscribers"
)

var (
	// ErrSubscribers is an internal error type for subscribers controller.
	ErrSubscribers = errs.Class("subscribers controller error")
)

// SubscribersTemplates holds all subscribers related templates.
type SubscribersTemplates struct {
	Create     *template.Template
	List       *template.Template
	GetByEmail *template.Template
	Delete     *template.Template
}

// Subscribers is a mvc controller that handles all Subscribers related views.
type Subscribers struct {
	log logger.Logger

	subscribers *subscribers.Service

	templates SubscribersTemplates
}

// NewSubscribers is a constructor for subscribers controller.
func NewSubscribers(log logger.Logger, subscribers *subscribers.Service, templates SubscribersTemplates) *Subscribers {
	subscribersController := &Subscribers{
		log:         log,
		subscribers: subscribers,
		templates:   templates,
	}

	return subscribersController
}

// List is an endpoint that will provide a web page with subscribers page.
func (controller *Subscribers) List(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var (
		err         error
		limit, page int
	)
	urlQuery := r.URL.Query()
	limitQuery := urlQuery.Get("limit")
	pageQuery := urlQuery.Get("page")

	if limitQuery != "" {
		limit, err = strconv.Atoi(limitQuery)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}

	if pageQuery != "" {
		page, err = strconv.Atoi(pageQuery)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}

	cursor := pagination.Cursor{
		Limit: limit,
		Page:  page,
	}

	subscribersPage, err := controller.subscribers.List(ctx, cursor)
	if err != nil {
		controller.log.Error("could not list subscribers", ErrSubscribers.Wrap(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err = json.NewEncoder(w).Encode(subscribersPage); err != nil {
		controller.log.Error("failed to write json response", ErrSubscribers.Wrap(err))
		return
	}
}

// serveError replies to the request with specific code and error message.
func (controller *Subscribers) serveError(w http.ResponseWriter, status int, err error) {
	w.WriteHeader(status)
	var response struct {
		Error string `json:"error"`
	}
	response.Error = err.Error()

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		controller.log.Error("failed to write json error response", ErrWhitelist.Wrap(err))
	}
}
