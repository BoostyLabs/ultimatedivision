// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/zeebo/errs"

	"ultimatedivision/internal/logger"
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
	ctx := r.Context()

	profile, err := controller.users.GetProfile(ctx)
	if err != nil {
		controller.log.Error("could not get users profile", ErrUsers.Wrap(err))
		http.Error(w, "could not get users profile", http.StatusInternalServerError)
		return
	}

	if err = json.NewEncoder(w).Encode(profile); err != nil {
		controller.log.Error("failed to write json response", ErrUsers.Wrap(err))
		return
	}
}
