// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package whitelist

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"strings"

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
	var password []byte

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
		Password: hex.EncodeToString(password),
	}
	return ErrWhitelist.Wrap(service.whitelist.Create(ctx, whitelist))
}

// generatePassword generates password for user's wallet.
func (service *Service) generatePassword(address Hex, privateKey *ecdsa.PrivateKey) ([]byte, error) {
	addressLow := strings.ToLower(string(address))
	addressNFTSale := strings.ToLower(string(service.config.AddressNFTSale))

	addressLowHex, err := hex.DecodeString(addressLow[2:])
	if err != nil {
		return []byte{}, nil
	}
	addressNFTSaleHex, err := hex.DecodeString(addressNFTSale[2:])
	if err != nil {
		return []byte{}, nil
	}

	hashSignature := crypto.Keccak256Hash(append(addressLowHex, addressNFTSaleHex...))

	EthereumSignedMessageHashHex, err := hex.DecodeString(EthereumSignedMessageHash)
	if err != nil {
		return []byte{}, nil
	}

	messageSignature := crypto.Keccak256Hash(append(EthereumSignedMessageHashHex, hashSignature.Bytes()...))

	var resultPassword []byte
	password, err := crypto.Sign(messageSignature.Bytes(), privateKey)
	pass := string(password)[:len(password)-1]
	passLastByte := hex.EncodeToString([]byte(string(password)[len(password)-1:]))

	if passLastByte == "00" {
		resultPassword = append([]byte(pass), []byte{27}...)
	} else if passLastByte == "01" {
		resultPassword = append([]byte(pass), []byte{28}...)
	}

	return resultPassword, err
}

// GetByAddress returns whitelist by address from the database.
func (service *Service) GetByAddress(ctx context.Context, address Hex) (Response, error) {
	whitelist, err := service.whitelist.GetByAddress(ctx, address)

	response := Response{
		Password: Hex(whitelist.Password),
		SmartContract: SmartContract{
			AddressNFT:     service.config.AddressNFT,
			AddressNFTSale: service.config.AddressNFTSale,
			Price:          service.config.Price,
		},
	}

	return response, ErrWhitelist.Wrap(err)
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
func (service *Service) Delete(ctx context.Context, address Hex) error {
	return ErrWhitelist.Wrap(service.whitelist.Delete(ctx, address))
}

// SetPassword generates passwords for all whitelist items.
func (service *Service) SetPassword(ctx context.Context, privateKey Hex) error {
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
			Password: hex.EncodeToString(password),
		}
		if err := service.Update(ctx, whitelist); err != nil {
			return ErrWhitelist.Wrap(err)
		}
	}

	return nil
}
