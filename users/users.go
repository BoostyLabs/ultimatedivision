// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package users

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/zeebo/errs"
	"golang.org/x/crypto/bcrypt"
)

// ErrNoUser indicated that user does not exist.
var ErrNoUser = errs.Class("user does not exist")

// DB exposes access to users db.
//
// architecture: DB
type DB interface {
	// List returns all users from the data base.
	List(ctx context.Context) ([]User, error)
	// Get returns user by id from the data base.
	Get(ctx context.Context, id uuid.UUID) (User, error)
	// GetByEmail returns user by email from the data base.
	GetByEmail(ctx context.Context, email string) (User, error)
	// Create creates a user and writes to the database.
	Create(ctx context.Context, user User) error
	// Update updates a status in the database.
	Update(ctx context.Context, status int, id uuid.UUID) error
	// Delete deletes a user in the database.
	Delete(ctx context.Context, id uuid.UUID) error
}

// Status defines the list of possible user statuses.
type Status int

const (
	// StatusActive indicates that user can login to the account.
	StatusActive Status = 0
	// StatusSuspended indicates that user cannot login to the account.
	StatusSuspended Status = 1
	// StatusVerified indicates that user email is verified.
	StatusVerified Status = 3
)

// User describes user entity.
type User struct {
	ID           uuid.UUID `json:"id"`
	Email        string    `json:"email"`
	PasswordHash []byte    `json:"passwordHash"`
	NickName     string    `json:"nickName"`
	FirstName    string    `json:"firstName"`
	LastName     string    `json:"lastName"`
	LastLogin    time.Time `json:"lastLogin"`
	Status       Status    `json:"status"`
	CreatedAt    time.Time `json:"createdAt"`
}

// EncodePass encode the password and generate "hash" to store from users password.
func (user *User) EncodePass() error {
	hash, err := bcrypt.GenerateFromPassword(user.PasswordHash, bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.PasswordHash = hash
	return nil
}

// RegistrationRequest for body payload.
type RegistrationRequest struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	NickName  string `json:"nickName"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}
