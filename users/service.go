// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package users

import (
	"context"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/google/uuid"
	"github.com/zeebo/errs"
)

// ErrUsers indicates that there was an error in the service.
var ErrUsers = errs.Class("users service error")

// ErrUnauthenticated should be returned when user performs unauthenticated action.
var ErrUnauthenticated = errs.Class("user unauthenticated error")

// ErrWalletAddressAlreadyInUse should be returned when users wallet address is already in use.
var ErrWalletAddressAlreadyInUse = errs.Class("wallet address is already in use")

// Service is handling users related logic.
//
// architecture: Service.
type Service struct {
	users DB
}

// NewService is a constructor for users service.
func NewService(users DB) *Service {
	return &Service{
		users: users,
	}
}

// Get returns user from DB.
func (service *Service) Get(ctx context.Context, userID uuid.UUID) (User, error) {
	user, err := service.users.Get(ctx, userID)
	return user, ErrUsers.Wrap(err)
}

// GetByEmail returns user by email from DB.
func (service *Service) GetByEmail(ctx context.Context, email string) (User, error) {
	user, err := service.users.GetByEmail(ctx, email)
	return user, ErrUsers.Wrap(err)
}

// GetByWalletAddress returns user by wallet address from the data base.
func (service *Service) GetByWalletAddress(ctx context.Context, walletAddress common.Address, walletType WalletType) (User, error) {
	user, err := service.users.GetByWalletAddress(ctx, walletAddress.String(), walletType)
	return user, ErrUsers.Wrap(err)
}

// GetByCasperWalletAddress returns user by Casper wallet address from the data base.
func (service *Service) GetByCasperWalletAddress(ctx context.Context, walletAddress string, walletType WalletType) (User, error) {
	user, err := service.users.GetByWalletAddress(ctx, walletAddress, walletType)
	return user, ErrUsers.Wrap(err)
}

// GetByCasperHash  returns user by Casper hash from the database.
func (service *Service) GetByCasperHash(ctx context.Context, hash string) (User, error) {
	user, err := service.users.GetByCasperHash(ctx, hash)
	return user, ErrUsers.Wrap(err)
}

// List returns all users from DB.
func (service *Service) List(ctx context.Context) ([]User, error) {
	users, err := service.users.List(ctx)
	return users, ErrUsers.Wrap(err)
}

// Create creates a user and returns user email.
func (service *Service) Create(ctx context.Context, email, password, nickName, firstName, lastName string) error {
	user := User{
		ID:           uuid.New(),
		Email:        email,
		PasswordHash: []byte(password),
		NickName:     nickName,
		FirstName:    firstName,
		LastName:     lastName,
		LastLogin:    time.Time{},
		Status:       StatusCreated,
		CreatedAt:    time.Now().UTC(),
	}
	err := user.EncodePass()
	if err != nil {
		return ErrUsers.Wrap(err)
	}

	return ErrUsers.Wrap(service.users.Create(ctx, user))
}

// Delete deletes a user.
func (service *Service) Delete(ctx context.Context, id uuid.UUID) error {
	return ErrUsers.Wrap(service.users.Delete(ctx, id))
}

// Update updates a users status.
func (service *Service) Update(ctx context.Context, status Status, id uuid.UUID) error {
	return ErrUsers.Wrap(service.users.Update(ctx, status, id))
}

// GetProfile returns user profile.
func (service *Service) GetProfile(ctx context.Context, userID uuid.UUID) (*ProfileWithWallet, error) {
	user, err := service.users.Get(ctx, userID)
	if err != nil {
		return nil, ErrUsers.Wrap(err)
	}

	return &ProfileWithWallet{
		ID:                  user.ID,
		Email:               user.Email,
		NickName:            user.NickName,
		CreatedAt:           user.CreatedAt,
		LastLogin:           user.LastLogin,
		Wallet:              user.Wallet,
		CasperWalletAddress: user.CasperWallet,
		CasperWalletHash:    user.CasperWalletHash,
		WalletType:          user.WalletType,
	}, nil
}

// GetNickNameByID returns nickname of user.
func (service *Service) GetNickNameByID(ctx context.Context, id uuid.UUID) (string, error) {
	nickname, err := service.users.GetNickNameByID(ctx, id)

	return nickname, ErrUsers.Wrap(err)
}

// UpdateWalletAddress updates wallet address.
func (service *Service) UpdateWalletAddress(ctx context.Context, wallet common.Address, id uuid.UUID, walletType WalletType) error {
	_, err := service.GetByWalletAddress(ctx, wallet, walletType)
	if err == nil {
		return ErrWalletAddressAlreadyInUse.New("wallet address already in use")
	}

	return ErrUsers.Wrap(service.users.UpdateWalletAddress(ctx, wallet, walletType, id))
}

// UpdateCasperWalletAddress updates Casper wallet address.
func (service *Service) UpdateCasperWalletAddress(ctx context.Context, wallet string, id uuid.UUID, walletType WalletType) error {
	_, err := service.GetByCasperWalletAddress(ctx, wallet, walletType)
	if err == nil {
		return ErrWalletAddressAlreadyInUse.New("wallet address already in use")
	}

	return ErrUsers.Wrap(service.users.UpdateCasperWalletAddress(ctx, wallet, walletType, id))
}

// ChangeWalletAddress changes wallet address.
func (service *Service) ChangeWalletAddress(ctx context.Context, wallet common.Address, id uuid.UUID) error {
	user, err := service.GetByWalletAddress(ctx, wallet, WalletTypeETH)
	if err != nil {
		return ErrUsers.Wrap(err)
	}
	if user.ID == id {
		return ErrUsers.New("this address is used by you")
	}

	err = service.users.UpdateWalletAddress(ctx, common.Address{}, WalletTypeETH, user.ID)
	if err != nil {
		return ErrUsers.Wrap(err)
	}

	return ErrUsers.Wrap(service.users.UpdateWalletAddress(ctx, wallet, WalletTypeETH, id))
}
