// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package users

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/zeebo/errs"
)

var ErrNoUser = errs.Class("user does not exist")

type DB interface {
	List(ctx context.Context) ([]User, error)
	Get(id uuid.UUID) (User, error)
	GetByEmail(email string) (User, error)
}

// UserStatus defined the list of possible user statuses
type userStatus string

const (
	userStatus_active    userStatus = "active"
	userStatus_suspended userStatus = "suspended"
)

// User structure
type User struct {
	ID        uuid.UUID  `json:"id"`
	Email     string     `json:"email"`
	Password  string     `json:"password"`
	NickName  string     `json:"nick_name"`
	FirstName string     `json:"first_name"`
	LastName  string     `json:"last_name"`
	LastLogin time.Time  `json:"last_login"`
	Status    userStatus `json:"status"`
	CreatedAt time.Time  `json:"created_at"`
}
