// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/zeebo/errs"

	"ultimatedivision/internal/logger"
	"ultimatedivision/pkg/auth"
	"ultimatedivision/users"
)

var (
	// ErrUsers is an internal error type for users controller.
	ErrUsers = errs.Class("users controller error")
)

// Users is a mvc controller that handles all users related views.
type Users struct {
	log logger.Logger

	users *users.Service
}

// NewUsers is a constructor for users controller.
func NewUsers(log logger.Logger, users *users.Service) *Users {
	usersController := &Users{
		log:   log,
		users: users,
	}

	return usersController
}

// GetProfile returns the current user profile with all relevant information.
func (controller *Users) GetProfile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx := r.Context()

	claims, err := auth.GetClaims(ctx)
	if err != nil {
		controller.serveError(w, http.StatusUnauthorized, ErrUsers.Wrap(err))
		return
	}

	profile, err := controller.users.GetProfile(ctx, claims.UserID)
	if err != nil {
		controller.serveError(w, http.StatusInternalServerError, ErrUsers.Wrap(err))
		return
	}

	if err = json.NewEncoder(w).Encode(profile); err != nil {
		controller.log.Error("failed to write json response", ErrUsers.Wrap(err))
		return
	}
}

// serveError replies to request with specific code and error.
func (controller *Users) serveError(w http.ResponseWriter, status int, err error) {
	w.WriteHeader(status)

	var response struct {
		Error string `json:"error"`
	}

	response.Error = err.Error()

	if err = json.NewEncoder(w).Encode(response); err != nil {
		controller.log.Error("failed to write json error response", ErrUsers.Wrap(err))
	}
}
