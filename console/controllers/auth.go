package controllers

import (
	"encoding/json"
	"html/template"
	"net/http"

	"github.com/zeebo/errs"

	"ultimatedivision/internal/auth"
	"ultimatedivision/internal/logger"
	"ultimatedivision/users"
	"ultimatedivision/users/userauth"
)

// AuthError is a internal error for auth controller.
var AuthError = errs.Class("auth controller error")

// AuthTemplates holds all auth related templates.
type AuthTemplates struct {
	Login *template.Template
}

// Auth login authentication entity.
type Auth struct {
	log     logger.Logger
	service *userauth.Service
	cookie  *auth.CookieAuth

	loginTemplate *template.Template
}

// NewAuth returns new instance of Auth.
func NewAuth(log logger.Logger, service *userauth.Service, authCookie *auth.CookieAuth, templates AuthTemplates) *Auth {
	return &Auth{
		log:           log,
		service:       service,
		cookie:        authCookie,
		loginTemplate: templates.Login,
	}
}

// Register a new user account.
func (auth *Auth) Register(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var err error
	var request users.RegistrationRequest

	if err = json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, AuthError.Wrap(err).Error(), http.StatusBadRequest)
		return
	}

	// create the new user in the database.
	err = auth.service.RegisterUser(ctx, request.Email, request.Password, request.NickName, request.FirstName, request.LastName)
	if err != nil {
		auth.log.Error("Unable to register new user", AuthError.Wrap(err))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

// Login is an endpoint to authorize user and set auth cookie in browser.
func (auth *Auth) Login(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var err error

	var request users.RegistrationRequest

	if err = json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, AuthError.Wrap(err).Error(), http.StatusBadRequest)
		return
	}

	if request.Email == "" || request.Password == "" {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	response, err := auth.service.Token(ctx, request.Email, request.Password)
	if err != nil {
		auth.log.Error("could not get auth token", AuthError.Wrap(err))
		switch {
		case users.ErrNoUser.Has(err):
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		case userauth.ErrUnauthenticated.Has(err):
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		default:
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}

		return
	}

	auth.cookie.SetTokenCookie(w, response)
}

// Logout is an endpoint to log out and remove auth cookie from browser.
func (auth *Auth) Logout(w http.ResponseWriter, r *http.Request) {
	auth.cookie.RemoveTokenCookie(w)
}
