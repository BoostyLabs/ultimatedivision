// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package controllers

import (
	"html/template"
	"net/http"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/zeebo/errs"

	"ultimatedivision/internal/logger"
	"ultimatedivision/users"
)

var (
	// ErrUsers is an internal error type for users controller.
	ErrUsers = errs.Class("users controller error")
)

// UserTemplates holds all users related templates.
type UserTemplates struct {
	List       *template.Template
	Create     *template.Template
	Get        *template.Template
	GetByEmail *template.Template
	Update     *template.Template
}

// Users is a mvc controller that handles all admins related views.
type Users struct {
	log logger.Logger

	users *users.Service

	templates UserTemplates
}

// NewUsers is a constructor for users controller.
func NewUsers(log logger.Logger, users *users.Service, templates UserTemplates) *Users {
	managersController := &Users{
		log:       log,
		users:     users,
		templates: templates,
	}

	return managersController
}

// List is an endpoint that will provide a web page with all users.
func (controller *Users) List(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	users, err := controller.users.List(ctx)
	if err != nil {
		controller.log.Error("could not get users list", ErrUsers.Wrap(err))
		http.Error(w, "could not get users list", http.StatusInternalServerError) // status code should depends on error type.
		return
	}

	err = controller.templates.List.Execute(w, users)
	if err != nil {
		controller.log.Error("can not execute list users template", ErrUsers.Wrap(err))
		http.Error(w, "can not execute list users template", http.StatusInternalServerError)
		return
	}
}

// Create is an endpoint that will create a new user.
func (controller *Users) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	err := r.ParseForm()
	if err != nil {
		controller.log.Error("could not get users form", ErrUsers.Wrap(err))
		http.Error(w, "could not get users form", http.StatusInternalServerError)
		return
	}
	email := r.FormValue("email")
	if email == "" {
		http.Error(w, "email is empty", http.StatusBadRequest)
		return
	}
	password := r.FormValue("password")
	if password == "" {
		http.Error(w, "email is empty", http.StatusBadRequest)
		return
	}
	nickName := r.FormValue("nickName")
	if nickName == "" {
		http.Error(w, "nick name is empty", http.StatusBadRequest)
		return
	}
	firstName := r.FormValue("firstName")
	if firstName == "" {
		http.Error(w, "first name is empty", http.StatusBadRequest)
		return
	}
	lastName := r.FormValue("lastName")
	if lastName == "" {
		http.Error(w, "last name is empty", http.StatusBadRequest)
		return
	}

	user := users.User{
		ID:           uuid.UUID{},
		Email:        email,
		PasswordHash: []byte(password),
		NickName:     nickName,
		FirstName:    firstName,
		LastName:     lastName,
		LastLogin:    time.Time{},
		Status:       0,
		CreatedAt:    time.Now(),
	}

	err = controller.users.Create(ctx, user)
	if err != nil {
		controller.log.Error("could not get users list", ErrUsers.Wrap(err))
		http.Error(w, "could not get users list", http.StatusInternalServerError)
		return
	}

	err = controller.templates.Create.Execute(w, nil)
	if err != nil {
		controller.log.Error("can not execute list users template", ErrUsers.Wrap(err))
		http.Error(w, "can not execute list users template", http.StatusInternalServerError)
		return
	}
}

// CreateUserForm is an endpoint that will provide a web page with create user form.
func (controller *Users) CreateUserForm(w http.ResponseWriter, r *http.Request) {
	if err := controller.templates.Create.Execute(w, nil); err != nil {
		controller.log.Error("cannot execute add user template", ErrAdmins.Wrap(err))
		http.Error(w, "cannot execute add user template", http.StatusInternalServerError)
		return
	}
}

// Get is an endpoint that will provide a web page with user by id.
func (controller *Users) Get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	err := r.ParseForm()
	if err != nil {
		controller.log.Error("could not get users form", ErrUsers.Wrap(err))
		http.Error(w, "could not get users form", http.StatusInternalServerError)
		return
	}
	id := r.FormValue("id")
	if id == "" {
		http.Error(w, "id is empty", http.StatusBadRequest)
		return
	}

	user, err := controller.users.Get(ctx, uuid.MustParse(id))
	if err != nil {
		controller.log.Error("could not get user", ErrUsers.Wrap(err))
		http.Error(w, "could not get user", http.StatusInternalServerError)
		return
	}

	err = controller.templates.Get.Execute(w, user)
	if err != nil {
		controller.log.Error("can not execute get user template", ErrUsers.Wrap(err))
		http.Error(w, "can not execute get user template", http.StatusInternalServerError)
		return
	}
}

// Update is an endpoint that will update users status.
func (controller *Users) Update(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	err := r.ParseForm()
	if err != nil {
		controller.log.Error("could not get users form", ErrUsers.Wrap(err))
		http.Error(w, "could not get users form", http.StatusInternalServerError)
		return
	}
	email := r.FormValue("email")
	if email == "" {
		http.Error(w, "email is empty", http.StatusBadRequest)
		return
	}

	status := r.FormValue("status")
	if status == "" {
		http.Error(w, "status is empty", http.StatusBadRequest)
		return
	}
	s, err := strconv.Atoi(status)
	if err != nil {
		controller.log.Error("could not converted to type int", ErrUsers.Wrap(err))
		http.Error(w, "could not converted to type int", http.StatusInternalServerError)
		return
	}

	err = controller.users.Update(ctx, s, email)
	if err != nil {
		controller.log.Error("could not update users status", ErrUsers.Wrap(err))
		http.Error(w, "could not update users status", http.StatusInternalServerError)
		return
	}

	err = controller.templates.Update.Execute(w, nil)
	if err != nil {
		controller.log.Error("can not execute get user template", ErrUsers.Wrap(err))
		http.Error(w, "can not execute get user template", http.StatusInternalServerError)
		return
	}
}

// GetByEmail is an endpoint that will provide a web page with user by email.
func (controller *Users) GetByEmail(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	err := r.ParseForm()
	if err != nil {
		controller.log.Error("could not get users form", ErrUsers.Wrap(err))
		http.Error(w, "could not get users form", http.StatusInternalServerError)
		return
	}
	email := r.FormValue("email")
	if email == "" {
		http.Error(w, "email is empty", http.StatusBadRequest)
		return
	}

	user, err := controller.users.GetByEmail(ctx, email)
	if err != nil {
		controller.log.Error("could not get user", ErrUsers.Wrap(err))
		http.Error(w, "could not get user", http.StatusInternalServerError)
		return
	}

	err = controller.templates.GetByEmail.Execute(w, user)
	if err != nil {
		controller.log.Error("can not execute get user template", ErrUsers.Wrap(err))
		http.Error(w, "can not execute get user template", http.StatusInternalServerError)
		return
	}
}

// Delete is an endpoint that will delete a user by email.
func (controller *Users) Delete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	err := r.ParseForm()
	if err != nil {
		controller.log.Error("could not get users form", ErrUsers.Wrap(err))
		http.Error(w, "could not get users form", http.StatusInternalServerError)
		return
	}
	email := r.FormValue("email")
	if email == "" {
		http.Error(w, "email is empty", http.StatusBadRequest)
		return
	}

	err = controller.users.Delete(ctx, email)
	if err != nil {
		controller.log.Error("could not delete user", ErrUsers.Wrap(err))
		http.Error(w, "could not delete user", http.StatusInternalServerError)
		return
	}
}
