// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package controllers

import (
	"golang.org/x/crypto/bcrypt"
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
func NewAdmins(log logger.Logger, admins *admins.Service) *Admins {
	managersController := &Admins{
		log:    log,
		admins: admins,
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

// AdminsInput is struct for form.
type AdminsInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// NewAdminInput is constructor for AdminsInput.
func NewAdminInput() *AdminsInput {
	return &AdminsInput{}
}

// Encode is method to encode password.
func (adminInput *AdminsInput) Encode() error {
	hash, err := bcrypt.GenerateFromPassword([]byte(adminInput.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	adminInput.Password = string(hash)
	return nil
}

// Create is an endpoint that creates new admin.
func (controller *Admins) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	adminInput := NewAdminInput()
	adminInput.Email = r.FormValue("name")
	adminInput.Password = r.FormValue("password")

	err := adminInput.Encode()
	if err != nil {
		controller.log.Error("could not encode admins password", ErrAdmins.Wrap(err))
		http.Error(w, "could not encode admins password", http.StatusInternalServerError)
		return
	}

	admin := admins.NewAdmin(adminInput.Email, []byte(adminInput.Password))

	err = controller.admins.Create(ctx, admin)
	if err != nil {
		controller.log.Error("could not create admin", ErrAdmins.Wrap(err))
		http.Error(w, "could not create admin", http.StatusInternalServerError) // status code should depends on error type.
		return
	}
}

// GenerateForm is endpoint to generate create admin form.
func (controller *Admins) GenerateForm(w http.ResponseWriter, r *http.Request) {
	err := controller.templates.List.Execute(w, nil)
	if err != nil {
		controller.log.Error("could not execute list admins template", ErrAdmins.Wrap(err))
		http.Error(w, "could not execute list admins template", http.StatusInternalServerError)
		return
	}
}
