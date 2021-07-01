// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package controllers

import (
	"html/template"
	"net/http"

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
	ctx := r.Context()

	email := r.FormValue("name")
	password := r.FormValue("password")

	admin := admins.NewAdmin(email, []byte(password))

	err := admin.EncodePassword()
	if err != nil {
		controller.log.Error("could not encode admin password", ErrAdmins.Wrap(err))
		http.Error(w, "could not encode admin password", http.StatusInternalServerError) // status code should depends on error type.
		return
	}

	err = controller.admins.Create(ctx, admin)
	if err != nil {
		controller.log.Error("could not create admin", ErrAdmins.Wrap(err))
		http.Error(w, "could not create admin", http.StatusInternalServerError) // status code should depends on error type.
		return
	}
}

// GenerateForm is endpoint to generate create admin form.
func (controller *Admins) GenerateForm(w http.ResponseWriter, r *http.Request) {
	err := controller.templates.Create.Execute(w, nil)
	if err != nil {
		controller.log.Error("could not execute list admins template", ErrAdmins.Wrap(err))
		http.Error(w, "could not execute list admins template", http.StatusInternalServerError)
		return
	}
}
