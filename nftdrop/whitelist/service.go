// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package whitelist

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"

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
func (service *Service) Create(ctx context.Context, request CreateWallet) error {
	var password string

	if request.PrivateKey != "" {
		privateKeyECDSA, err := crypto.HexToECDSA(string(request.PrivateKey))
		if err != nil {
			return ErrWhitelist.Wrap(err)
		}

		password, err = service.generatePassword(request.Address, privateKeyECDSA)
		if err != nil {
			return ErrWhitelist.Wrap(err)
		}
	}

	whitelist := Wallet{
		Address:  request.Address,
		Password: Signature(password),
	}
	return ErrWhitelist.Wrap(service.whitelist.Create(ctx, whitelist))
}

// generatePassword generates password for user's wallet.
func (service *Service) generatePassword(address Address, privateKey *ecdsa.PrivateKey) (string, error) {
	addressNFT, err := hex.DecodeString(string(address)[2:])
	if err != nil {
		return "", ErrWhitelist.Wrap(err)
	}

	addressNFTSale, err := hex.DecodeString(string(service.config.NFTSale)[2:])
	if err != nil {
		return "", ErrWhitelist.Wrap(err)
	}

	EthereumSignedMessage, err := hex.DecodeString(EthereumSignedMessageHash)
	if err != nil {
		return "", ErrWhitelist.Wrap(err)
	}

	dataSignature := crypto.Keccak256Hash(append(EthereumSignedMessage, crypto.Keccak256Hash(append(addressNFT, addressNFTSale...)).Bytes()...))
	password, err := crypto.Sign(dataSignature.Bytes(), privateKey)
	if err != nil {
		return "", ErrWhitelist.Wrap(err)
	}

	passwordWithoutEnd := string(password)[:len(password)-1]
	passwordString := hex.EncodeToString(password)
	passwordLastSymbol := passwordString[len(passwordString)-1:]

	var resultPassword []byte
	if passwordLastSymbol == fmt.Sprintf("%d", PrivateKeyVZero) {
		resultPassword = append([]byte(passwordWithoutEnd), []byte{byte(PrivateKeyVTwentySeven)}...)
	} else if passwordLastSymbol == fmt.Sprintf("%d", PrivateKeyVOne) {
		resultPassword = append([]byte(passwordWithoutEnd), []byte{byte(PrivateKeyVTwentySeven)}...)
	}

	return hex.EncodeToString(resultPassword), nil
}

// GetByAddress returns whitelist by address from the database.
func (service *Service) GetByAddress(ctx context.Context, address Address) (Transaction, error) {
	whitelist, err := service.whitelist.GetByAddress(ctx, address)

	transactionValue := Transaction{
		Password: whitelist.Password,
		SmartContractAddress: SmartContractAddress{
			NFT:     service.config.NFT,
			NFTSale: service.config.NFTSale,
		},
	}

	return transactionValue, ErrWhitelist.Wrap(err)
}

// List returns all whitelist from the database.
func (service *Service) List(ctx context.Context) ([]Wallet, error) {
	whitelistRecords, err := service.whitelist.List(ctx)
	return whitelistRecords, ErrWhitelist.Wrap(err)
}

// ListWithoutPassword returns whitelist without password from the database.
func (service *Service) ListWithoutPassword(ctx context.Context) ([]Wallet, error) {
	whitelistRecords, err := service.whitelist.ListWithoutPassword(ctx)
	return whitelistRecords, ErrWhitelist.Wrap(err)
}

// Update updates whitelist by address.
func (service *Service) Update(ctx context.Context, whitelist Wallet) error {
	return ErrWhitelist.Wrap(service.whitelist.Update(ctx, whitelist))
}

// Delete deletes whitelist.
func (service *Service) Delete(ctx context.Context, address Address) error {
	return ErrWhitelist.Wrap(service.whitelist.Delete(ctx, address))
}

// SetPassword generates passwords for all whitelist items.
func (service *Service) SetPassword(ctx context.Context, privateKey PrivateKey) error {
	privateKeyECDSA, err := crypto.HexToECDSA(string(privateKey))
	if err != nil {
		return ErrWhitelist.Wrap(err)
	}

	whitelist, err := service.ListWithoutPassword(ctx)
	if err != nil {
		return ErrWhitelist.Wrap(err)
	}

	for _, v := range whitelist {
		password, err := service.generatePassword(v.Address, privateKeyECDSA)
		if err != nil {
			return ErrWhitelist.Wrap(err)
		}

		whitelist := Wallet{
			Address:  v.Address,
			Password: Signature(password),
		}
		if err := service.Update(ctx, whitelist); err != nil {
			return ErrWhitelist.Wrap(err)
		}
	}

	return nil
}
