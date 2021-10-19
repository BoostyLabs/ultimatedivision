// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/zeebo/errs"

	"ultimatedivision/internal/logger"
	"ultimatedivision/nftdrop/emails"
)

var (
	// ErrEmails is an internal error type for emails controller.
	ErrEmails = errs.Class("emails controller error")
)

// Emails is a mvc controller that handles all emails related views.
type Emails struct {
	log logger.Logger

	emails *emails.Service
}

// NewEmails is a constructor for emails controller.
func NewEmails(log logger.Logger, emails *emails.Service) *Emails {
	emailsController := &Emails{
		log:    log,
		emails: emails,
	}

	return emailsController
}

// Create is an endpoint that writes email.
func (controller *Emails) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx := r.Context()

	var err error
	var request emails.CreateEmailFields

	if err = json.NewDecoder(r.Body).Decode(&request); err != nil {
		controller.serveError(w, http.StatusBadRequest, ErrEmails.Wrap(err))
		return
	}

	err = controller.emails.Create(ctx, request.Email)
	if err != nil {
		controller.log.Error("Unable to write new email", ErrEmails.Wrap(err))
		controller.serveError(w, http.StatusInternalServerError, ErrEmails.Wrap(err))
		return
	}
}

// serveError replies to the request with specific code and error message.
func (controller *Emails) serveError(w http.ResponseWriter, status int, err error) {
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
