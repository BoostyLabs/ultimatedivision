// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package whitelist

import (
	"context"
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/zeebo/errs"
)

// ErrWhitelist indicated that there was an error in service.
var ErrWhitelist = errs.Class("whitelist service error")

// Service is handling whitelist related logic.
//
// architecture: Service
type Service struct {
	config    Config
	whitelist DB
}

// NewService is a constructor for whitelist service.
func NewService(config Config, whitelist DB) *Service {
	return &Service{
		config:    config,
		whitelist: whitelist,
	}
}

// Create adds whitelist in the database.
func (service *Service) Create(ctx context.Context, request Request) error {
	var password []byte

	if request.Key != "" {
		privateKey, err := crypto.HexToECDSA(request.Key)
		if err != nil {
			return ErrWhitelist.Wrap(err)
		}

		password, err = service.GeneratePassword(request.Address, privateKey)
		if err != nil {
			return ErrWhitelist.Wrap(err)
		}
	}

	whitelist := Whitelist{
		Address:  request.Address,
		Password: password,
	}
	return ErrWhitelist.Wrap(service.whitelist.Create(ctx, whitelist))
}

// GeneratePassword generates password for user's wallet.
func (service *Service) GeneratePassword(address Address, privateKey *ecdsa.PrivateKey) ([]byte, error) {
	dataSignature := []byte(service.config.SmartContract.Address + string(address))
	hashSignature := crypto.Keccak256Hash(dataSignature)

	return crypto.Sign(hashSignature.Bytes(), privateKey)
}

// GetByAddress returns whitelist by address from the database.
func (service *Service) GetByAddress(ctx context.Context, address Address) (Whitelist, error) {
	whitelist, err := service.whitelist.GetByAddress(ctx, address)
	return whitelist, ErrWhitelist.Wrap(err)
}

// List returns all whitelist from the database.
func (service *Service) List(ctx context.Context) ([]Whitelist, error) {
	whitelistRecords, err := service.whitelist.List(ctx)
	return whitelistRecords, ErrWhitelist.Wrap(err)
}

// ListWithoutPassword returns whitelist without password from the database.
func (service *Service) ListWithoutPassword(ctx context.Context) ([]Whitelist, error) {
	whitelistRecords, err := service.whitelist.ListWithoutPassword(ctx)
	return whitelistRecords, ErrWhitelist.Wrap(err)
}

// Update updates whitelist by address.
func (service *Service) Update(ctx context.Context, whitelist Whitelist) error {
	return ErrWhitelist.Wrap(service.whitelist.Update(ctx, whitelist))
}

// Delete deletes whitelist.
func (service *Service) Delete(ctx context.Context, address Address) error {
	return ErrWhitelist.Wrap(service.whitelist.Delete(ctx, address))
}

// SetPassword generates passwords for all whitelist items.
func (service *Service) SetPassword(ctx context.Context, key string) error {
	privateKey, err := crypto.HexToECDSA(key)
	if err != nil {
		return ErrWhitelist.Wrap(err)
	}

	whitelist, err := service.ListWithoutPassword(ctx)
	if err != nil {
		return ErrWhitelist.Wrap(err)
	}

	for _, v := range whitelist {
		password, err := service.GeneratePassword(v.Address, privateKey)
		if err != nil {
			return ErrWhitelist.Wrap(err)
		}

		whitelist := Whitelist{
			Address:  v.Address,
			Password: password,
		}
		if err := service.Update(ctx, whitelist); err != nil {
			return ErrWhitelist.Wrap(err)
		}
	}

	return nil
}
