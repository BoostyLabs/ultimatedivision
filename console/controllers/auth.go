package controllers

import (
	"html/template"
	"net/http"
	"ultimatedivision/database"
	"ultimatedivision/users"
	"ultimatedivision/users/userauth"
	"unicode"

	"github.com/zeebo/errs"

	"ultimatedivision/internal/auth"
	"ultimatedivision/internal/logger"
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

	err = r.ParseForm()
	if err != nil {
		http.Error(w, "could not get users form", http.StatusBadRequest)
		return
	}
	email := r.FormValue("email")
	if email == "" {
		http.Error(w, "email is empty", http.StatusBadRequest)
		return
	}
	password := r.FormValue("password")
	if password == "" {
		http.Error(w, "password is empty", http.StatusBadRequest)
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

	email = database.NormalizeEmail(email)

	// check if the user email address already exists.
	_, err = auth.service.GetUserByEmail(ctx, email)
	if err == nil {
		http.Error(w, "This email address is already in use.", http.StatusBadRequest)
		return
	}

	// check the password is valid.
	if !isPasswordValid(password) {
		http.Error(w, "The password must contain at least one lowercase (a-z) letter, one uppercase (A-Z) letter, one digit (0-9) and one special character.", http.StatusBadRequest)
		return
	}

	// create the new user in the database.
	err = auth.service.RegisterUser(ctx, email, password, nickName, firstName, lastName)
	if err != nil {
		auth.log.Error("Could not register user", AuthError.Wrap(err))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// @todo sending email function still have to finalize.
	//// launch a goroutine that sends the email verification.
	//go func() {
	//	_, err := auth.service.GenerateAndSendEmailConfirmation(user.Email)
	//	if err != nil {
	//		auth.log.Error("Unable to send account activation email", AuthError.Wrap(err))
	//	}
	//}()
}

func isPasswordValid(s string) bool {
	var number, upper, special bool
	letters := 0
	for _, c := range s {
		switch {
		case unicode.IsNumber(c):
			number = true
		case unicode.IsUpper(c):
			upper = true
		case unicode.IsPunct(c) || unicode.IsSymbol(c) || unicode.IsMark(c):
			special = true
		case unicode.IsLetter(c) || c == ' ':
			letters++
		}
	}
	return len(s) >= 8 && letters >= 1 && number && upper && special
}

// Login is an endpoint to authorize user and set auth cookie in browser.
func (auth *Auth) Login(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var err error

	err = r.ParseForm()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	email := r.Form["email"]
	password := r.Form["password"]
	if len(email) == 0 || len(password) == 0 {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	if email[0] == "" || password[0] == "" {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	response, err := auth.service.Token(ctx, email[0], password[0])
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
