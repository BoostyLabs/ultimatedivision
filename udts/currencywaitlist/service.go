// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package currencywaitlist

import (
	"context"
	"math/big"
	"strings"
	"time"

	"github.com/zeebo/errs"

	"ultimatedivision/pkg/cryptoutils"
	"ultimatedivision/udts"
	"ultimatedivision/users"
)

// ErrCurrencyWaitlist indicated that there was an error in service.
var ErrCurrencyWaitlist = errs.Class("currency waitlist service error")

// Service is handling currency wait list related logic.
//
// architecture: Service
type Service struct {
	config           Config
	currencyWaitList DB
	users            *users.Service
	udts             *udts.Service
}

// NewService is a constructor for currencyWaitList service.
func NewService(config Config, currencyWaitList DB, users *users.Service, udts *udts.Service) *Service {
	return &Service{
		config:           config,
		currencyWaitList: currencyWaitList,
		users:            users,
		udts:             udts,
	}
}

// Create creates item of currency wait list.
func (service *Service) Create(ctx context.Context, walletAddress cryptoutils.Address) (Transaction, error) {
	var transaction Transaction

	walletAddress = cryptoutils.Address(strings.ToLower(string(walletAddress)))
	user, err := service.users.GetByWalletAddress(ctx, walletAddress)
	if err != nil {
		return transaction, ErrCurrencyWaitlist.Wrap(err)
	}

	udt, err := service.udts.GetByUserID(ctx, user.ID)
	if err != nil {
		if !udts.ErrNoUDT.Has(err) {
			return transaction, ErrCurrencyWaitlist.Wrap(err)
		}

		udt = udts.UDT{
			UserID: user.ID,
			Value:  *big.NewInt(0),
			Nonce:  0,
		}
		if err = service.udts.Create(ctx, udt); err != nil {
			return transaction, ErrCurrencyWaitlist.Wrap(err)
		}
	}

	if item, err := service.GetByWalletAddressAndNonce(ctx, walletAddress, udt.Nonce); item.Signature != "" && err == nil {
		transaction = Transaction{
			Signature:    item.Signature,
			GameContract: service.config.GameContract,
			Value:        item.Value,
			Nonce:        item.Nonce,
		}
		return transaction, nil
	}

	var value = new(big.Int)
	value.SetString(service.config.WinValue, 10)
	item := Item{
		WalletAddress: walletAddress,
		Value:         *value,
		Nonce:         udt.Nonce,
	}

	if err = service.currencyWaitList.Create(ctx, item); err != nil {
		return transaction, ErrCurrencyWaitlist.Wrap(err)
	}

	for range time.NewTicker(time.Millisecond * service.config.IntervalSignatureCheck).C {
		if item, err := service.GetByWalletAddressAndNonce(ctx, walletAddress, udt.Nonce); item.Signature != "" && err == nil {
			transaction = Transaction{
				Signature:    item.Signature,
				GameContract: service.config.GameContract,
				Value:        item.Value,
				Nonce:        item.Nonce,
			}
			break
		}
	}

	return transaction, err
}

// GetByWalletAddressAndNonce returns item of currency wait list by wallet address and nonce.
func (service *Service) GetByWalletAddressAndNonce(ctx context.Context, walletAddress cryptoutils.Address, nonce int64) (Item, error) {
	item, err := service.currencyWaitList.GetByWalletAddressAndNonce(ctx, walletAddress, nonce)
	return item, ErrCurrencyWaitlist.Wrap(err)
}

// List returns items of currency wait list.
func (service *Service) List(ctx context.Context) ([]Item, error) {
	items, err := service.currencyWaitList.List(ctx)
	return items, ErrCurrencyWaitlist.Wrap(err)
}

// Update updates signature of item by wallet address and nonce.
func (service *Service) Update(ctx context.Context, signature cryptoutils.Signature, walletAddress cryptoutils.Address, nonce int64) error {
	return ErrCurrencyWaitlist.Wrap(service.currencyWaitList.UpdateSignature(ctx, signature, walletAddress, nonce))
}

// Delete deletes item of currency wait list by wallet address and nonce.
func (service *Service) Delete(ctx context.Context, walletAddress cryptoutils.Address, nonce int64) error {
	return ErrCurrencyWaitlist.Wrap(service.currencyWaitList.Delete(ctx, walletAddress, nonce))
}
