// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package controllers

import (
	"html/template"
	"net/http"
	"net/url"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/zeebo/errs"

	"ultimatedivision/admin/admins"
	"ultimatedivision/internal/logger"
)

var (
	// ErrAdmins is an internal error type for admins controller.
	ErrAdmins = errs.Class("admins controller error")
)

// AdminTemplates holds all admins related templates.
type AdminTemplates struct {
	List   *template.Template
	Create *template.Template
	Update *template.Template
}

// Admins is a mvc controller that handles all admins related views.
type Admins struct {
	log logger.Logger

	admins *admins.Service

	templates AdminTemplates
}

// NewAdmins is a constructor for admins controller.
func NewAdmins(log logger.Logger, admins *admins.Service, templates AdminTemplates) *Admins {
	managersController := &Admins{
		log:       log,
		admins:    admins,
		templates: templates,
	}

	return managersController
}

// Redirect redirects to specific url.
func (controller *Admins) Redirect(w http.ResponseWriter, r *http.Request, urlString, method string) {
	newRequest := new(http.Request)
	*newRequest = *r
	newRequest.URL = new(url.URL)
	*newRequest.URL = *r.URL
	newRequest.Method = method

	http.Redirect(w, newRequest, urlString, http.StatusMovedPermanently)
}

// List is an endpoint that will provide a web page with all admins.
func (controller *Admins) List(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	admins, err := controller.admins.List(ctx)
	if err != nil {
		controller.log.Error("could not get admins list", ErrAdmins.Wrap(err))
		http.Error(w, "could not get admins list", http.StatusInternalServerError) // status code should depends on error type.
		return
	}

	err = controller.templates.List.Execute(w, admins)
	if err != nil {
		controller.log.Error("can not execute list admins template", ErrAdmins.Wrap(err))
		http.Error(w, "can not execute list admins template", http.StatusInternalServerError)
		return
	}
}

// Create is an endpoint that creates new admin.
func (controller *Admins) Create(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		err := controller.templates.Create.Execute(w, nil)
		if err != nil {
			controller.log.Error("could not execute create admins template", ErrAdmins.Wrap(err))
			http.Error(w, "could not execute create admins template", http.StatusInternalServerError)
			return
		}
	case http.MethodPost:
		ctx := r.Context()

		err := r.ParseForm()
		if err != nil {
			controller.log.Error("could not parse admin create form", ErrAdmins.Wrap(err))
			http.Error(w, "could not parse admin create form", http.StatusInternalServerError)
			return
		}

		email := r.FormValue("email")
		password := r.FormValue("password")
		if email == "" || password == "" {
			http.Error(w, "login or password input is empty", http.StatusBadRequest) // status code should depends on error type.
			return
		}

		admin := admins.Admins{Admin: admins.Admin{
			ID:           uuid.New(),
			Email:        email,
			PasswordHash: []byte(password),
			CreatedAt:    time.Now(),
		}}

		err = admin.EncodePassword()
		if err != nil {
			controller.log.Error("could not encode admin password", ErrAdmins.Wrap(err))
			http.Error(w, "could not encode admin password", http.StatusInternalServerError) // status code should depends on error type.
			return
		}

		err = controller.admins.Create(ctx, admin.Admin)
		if err != nil {
			controller.log.Error("could not create admin", ErrAdmins.Wrap(err))
			http.Error(w, "could not create admin", http.StatusInternalServerError) // status code should depends on error type.
			return
		}

		controller.Redirect(w, r, "", http.MethodGet)
	}
}

// Update is an endpoint that updates admin.
func (controller *Admins) Update(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	params := mux.Vars(r)

	idOfAdmin := params["id"]
	uuidOfAdmin, err := uuid.Parse(idOfAdmin)
	if err != nil {
		controller.log.Error("could not parse uuid", ErrAdmins.Wrap(err))
		http.Error(w, "could not parse uuid", http.StatusInternalServerError)
		return
	}

	switch r.Method {
	case http.MethodGet:
		admin, err := controller.admins.Get(ctx, uuidOfAdmin)
		if err != nil {
			controller.log.Error("could not get admins list", ErrAdmins.Wrap(err))
			http.Error(w, "could not get admins list", http.StatusInternalServerError) // status code should depends on error type.
			return
		}

		err = controller.templates.Update.Execute(w, admin)
		if err != nil {
			controller.log.Error("could not execute update admins template", ErrAdmins.Wrap(err))
			http.Error(w, "could not execute update admins template", http.StatusInternalServerError)
			return
		}
	case http.MethodPost:
		err = r.ParseForm()
		if err != nil {
			controller.log.Error("could not parse admin create form", ErrAdmins.Wrap(err))
			http.Error(w, "could not parse admin create form", http.StatusInternalServerError)
			return
		}

		password := r.FormValue("password")
		passwordAgain := r.FormValue("password-again")
		if password == "" || passwordAgain == "" {
			http.Error(w, "empty field", http.StatusBadRequest) // status code should depends on error type.
			return
		}

		if password != passwordAgain {
			http.Error(w, "password mismatch", http.StatusBadRequest) // status code should depends on error type.
			return
		}

		updatedAdmin := admins.Admins{Admin: admins.Admin{
			ID:           uuidOfAdmin,
			PasswordHash: []byte(password),
		}}

		err = controller.admins.Update(ctx, updatedAdmin.Admin)
		if err != nil {
			controller.log.Error("could not update admin", ErrAdmins.Wrap(err))
			http.Error(w, "could not update admin", http.StatusInternalServerError) // status code should depends on error type.
			return
		}

		controller.Redirect(w,r,"/admins",http.MethodGet)
	}
}
