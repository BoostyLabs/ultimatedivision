// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package users

import (
	"context"
	"time"
	"unicode"

	"github.com/BoostyLabs/evmsignature"
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
	// GetByWalletAddress returns user by wallet address from the data base.
	GetByWalletAddress(ctx context.Context, walletAddress evmsignature.Address) (User, error)
	// Create creates a user and writes to the database.
	Create(ctx context.Context, user User) error
	// UpdateStatus updates a status in the database.
	UpdateStatus(ctx context.Context, status Status, id uuid.UUID) error
	// UpdatePassword updates a password in the database.
	UpdatePassword(ctx context.Context, passwordHash []byte, id uuid.UUID) error
	// UpdateWalletAddress updates user's address of wallet in the database.
	UpdateWalletAddress(ctx context.Context, wallet evmsignature.Address, id uuid.UUID) error
	// UpdateNonce updates nonce by user.
	UpdateNonce(ctx context.Context, id uuid.UUID, nonce []byte) error
	// Delete deletes a user in the database.
	Delete(ctx context.Context, id uuid.UUID) error
	// GetNickNameByID returns nickname by user id from the database.
	GetNickNameByID(ctx context.Context, id uuid.UUID) (string, error)
	// UpdateLastLogin updates last login time.
	UpdateLastLogin(ctx context.Context, id uuid.UUID) error
	// UpdateEmail updates an email address in the database.
	UpdateEmail(ctx context.Context, id uuid.UUID, newEmail string) error
}

// Status defines the list of possible user statuses.
type Status int

const (
	// StatusCreated indicates that user email is created.
	StatusCreated Status = 0
	// StatusActive indicates that user can login to the account.
	StatusActive Status = 1
	// StatusSuspended indicates that user cannot login to the account.
	StatusSuspended Status = 2
)

// Blockchain defines the list of possible blockchain.
type Blockchain string

const (
	// BlockchainPolygon indicates that blockchain is polygon.
	BlockchainPolygon Blockchain = "polygon"
	// BlockchainVelas indicates that blockchain is velas.
	BlockchainVelas Blockchain = "velas"
)

// DefaultMessageForRegistration use for registration user by metamask.
const DefaultMessageForRegistration = "Register with metamask"

// User describes user entity.
type User struct {
	ID           uuid.UUID            `json:"id"`
	Email        string               `json:"email"`
	PasswordHash []byte               `json:"passwordHash"`
	NickName     string               `json:"nickName"`
	FirstName    string               `json:"firstName"`
	LastName     string               `json:"lastName"`
	Wallet       evmsignature.Address `json:"wallet"`
	Nonce        []byte               `json:"nonce"`
	LastLogin    time.Time            `json:"lastLogin"`
	Status       Status               `json:"status"`
	Blockchain   Blockchain           `json:"blockchain"`
	CreatedAt    time.Time            `json:"createdAt"`
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

// CreateUserFields for crete user.
type CreateUserFields struct {
	Email     string               `json:"email"`
	Password  string               `json:"password"`
	NickName  string               `json:"nickName"`
	FirstName string               `json:"firstName"`
	LastName  string               `json:"lastName"`
	Wallet    evmsignature.Address `json:"wallet"`
}

// Profile for user profile.
type Profile struct {
	Email     string               `json:"email"`
	NickName  string               `json:"nickName"`
	CreatedAt time.Time            `json:"registerDate"`
	LastLogin time.Time            `json:"lastLogin"`
	Wallet    evmsignature.Address `json:"wallet"`
}

// Password for old/new passwords.
type Password struct {
	Password    string `json:"password"`
	NewPassword string `json:"newPassword"`
}

// IsPasswordValid check the password for all conditions.
func IsPasswordValid(s string) bool {
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

// IsValid check the request for all conditions.
func (createUserFields *CreateUserFields) IsValid() bool {
	switch {
	case createUserFields.Email == "":
		return false
	case createUserFields.Password == "":
		return false
	case createUserFields.FirstName == "":
		return false
	case createUserFields.LastName == "":
		return false
	case createUserFields.NickName == "":
		return false
	default:
		return true
	}
}
