// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package whitelist

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/zeebo/errs"
)

// ErrWhitelist indicated that there was an error in service.
var ErrWhitelist = errs.Class("whitelist service error")

// Service is handling whitelist related logic.
//
// architecture: Service
type Service struct {
	whitelist DB
}

// NewService is a constructor for whitelist service.
func NewService(whitelist DB) *Service {
	return &Service{whitelist: whitelist}
	// github.com/ethereum/go-ethereum/crypto

}

const randomString = "0x1c8aff950685c2ed4bc3174f347228"

// Create adds whitelist in the data base.
func (service *Service) Create(ctx context.Context, whitelistRequest Request) error {
	privateKey, err := crypto.HexToECDSA("fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19")
	if err != nil {
		return err
	}

	signature, err := service.CreateSignature(service.ReturnHash(ctx).Bytes(), privateKey)

	sigPublicKey, err := crypto.SigToPub(service.ReturnHash(ctx).Bytes(), signature)
	if err != nil {
		return ErrWhitelist.Wrap(err)
	}
	fmt.Println(sigPublicKey)
	whitelist := Whitelist{
		Address: whitelistRequest.Address,
		// TODO: generate password
		Password: []byte{},
	}

	// TODO: transaction

	return ErrWhitelist.Wrap(service.whitelist.Create(ctx, whitelist))
}

func (service *Service) CreateSignature(hash []byte, privateKey *ecdsa.PrivateKey) ([]byte, error){
	signature, err := crypto.Sign(hash, privateKey)
	if err != nil {
		return []byte(""),ErrWhitelist.Wrap(err)
	}
	return signature, nil
}

// Get returns whitelist by address from the data base.
func (service *Service) Get(ctx context.Context, address string) (Whitelist, error) {
	whitelist, err := service.whitelist.Get(ctx, address)
	return whitelist, ErrWhitelist.Wrap(err)
}

// ReturnHash returns random hash.
func (service *Service) ReturnHash(ctx context.Context) common.Hash {
	hash := crypto.Keccak256Hash([]byte(randomString))

	return hash
}

// List returns all whitelist from the data base.
func (service *Service) List(ctx context.Context) ([]Whitelist, error) {
	whitelistRecords, err := service.whitelist.List(ctx)
	return whitelistRecords, ErrWhitelist.Wrap(err)
}
