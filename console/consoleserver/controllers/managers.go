// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/zeebo/errs"

	"ultimatedivision/clubs/managers"
	"ultimatedivision/internal/logger"
)

var (
	// ErrManagers is an internal error type for managers controller.
	ErrManagers = errs.Class("managers controller error")
)

// Managers is a mvc controller that handles all managers related views.
type Managers struct {
	log logger.Logger

	managers *managers.Service
}

// NewManagers is a constructor for managers controller.
func NewManagers(log logger.Logger, managers *managers.Service) *Managers {
	managersController := &Managers{
		log:      log,
		managers: managers,
	}

	return managersController
}

// Create is ann endpoint that creates manager for club.
func (controller *Managers) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx := r.Context()
	params := mux.Vars(r)

	clubID, err := uuid.Parse(params["clubId"])
	if err != nil {
		controller.serveError(w, http.StatusBadRequest, ErrManagers.Wrap(err))
		return
	}

	var manager managers.Manager

	if err = json.NewDecoder(r.Body).Decode(&manager); err != nil {
		controller.serveError(w, http.StatusBadRequest, ErrManagers.Wrap(err))
		return
	}

	err = controller.managers.Create(ctx, manager.EndedAt, manager.UserID, clubID)
	if err != nil {
		controller.log.Error("could not create manager", ErrManagers.Wrap(err))
		controller.serveError(w, http.StatusInternalServerError, ErrManagers.Wrap(err))
		return
	}
}

// serveError replies to the request with specific code and error message.
func (controller *Managers) serveError(w http.ResponseWriter, status int, err error) {
	w.WriteHeader(status)

	var response struct {
		Error string `json:"error"`
	}

	response.Error = err.Error()

	if err = json.NewEncoder(w).Encode(response); err != nil {
		controller.log.Error("failed to write json error response", ErrManagers.Wrap(err))
	}
}
