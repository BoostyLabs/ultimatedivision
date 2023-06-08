// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package marketplace

import (
	"context"
	"math/big"
	"time"

	"github.com/google/uuid"
	"github.com/zeebo/errs"

	"ultimatedivision/cards"
	"ultimatedivision/pkg/pagination"
	"ultimatedivision/users"
)

// ErrNoLot indicated that lot does not exist.
var ErrNoLot = errs.Class("lot does not exist")

// DB is exposing access to lots db.
//
// architecture: DB
type DB interface {
	// CreateLot add lot in the data base.
	CreateLot(ctx context.Context, lot Lot) error
	// GetLotByID returns lot by id from the data base.
	GetLotByID(ctx context.Context, id uuid.UUID) (Lot, error)
	// GetLotEndTimeByID returns lot end time by id from data base.
	GetLotEndTimeByID(ctx context.Context, id uuid.UUID) (time.Time, error)
	// GetCurrentPriceByCardID returns current price by card id from the data base.
	GetCurrentPriceByCardID(ctx context.Context, cardID uuid.UUID) (big.Int, error)
	// ListActiveLots returns active lots from the data base.
	ListActiveLots(ctx context.Context, cursor pagination.Cursor) (Page, error)
	// ListActiveLotsByCardID returns active lots from the data base by card id.
	ListActiveLotsByCardID(ctx context.Context, cardIds []uuid.UUID, cursor pagination.Cursor) (Page, error)
	// ListExpiredLot returns lots where end time lower than or equal to time now UTC from the data base.
	ListExpiredLot(ctx context.Context) ([]Lot, error)
	// UpdateShopperIDLot updates shopper id of lot in the database.
	UpdateShopperIDLot(ctx context.Context, id, shopperID uuid.UUID) error
	// UpdateStatusLot updates status of lot in the database.
	UpdateStatusLot(ctx context.Context, id uuid.UUID, status Status) error
	// UpdateCurrentPriceLot updates current price of lot in the database.
	UpdateCurrentPriceLot(ctx context.Context, id uuid.UUID, currentPrice big.Int) error
	// UpdateEndTimeLot updates end time of lot in the database.
	UpdateEndTimeLot(ctx context.Context, id uuid.UUID, endTime time.Time) error
	// Delete deletes lot in the database.
	Delete(ctx context.Context, cardID uuid.UUID) error
}

// Lot describes lot entity.
type Lot struct {
	CardID       uuid.UUID  `json:"cardId"`
	Type         Type       `json:"type"`
	UserID       uuid.UUID  `json:"userId"`
	ShopperID    uuid.UUID  `json:"shopperId"`
	Status       Status     `json:"status"`
	StartPrice   big.Int    `json:"startPrice"`
	MaxPrice     big.Int    `json:"maxPrice"`
	CurrentPrice big.Int    `json:"currentPrice"`
	StartTime    time.Time  `json:"startTime"`
	EndTime      time.Time  `json:"endTime"`
	Period       Period     `json:"period"`
	Card         cards.Card `json:"card"`
}

// Type defines the list of possible lot types.
type Type string

const (
	// TypeCard indicates that lot type is card.
	TypeCard Type = "card"
)

// Status defines the list of possible lot statuses.
type Status string

const (
	// StatusActive indicates that lot is active that is, the lot is sold at auction.
	StatusActive Status = "active"
	// StatusExpired indicates that the time of lot has expired.
	StatusExpired Status = "expired"
	// StatusSold indicates that the lot was sold after the expiration of the lot at the highest rate.
	StatusSold Status = "sold"
	// StatusSoldBuynow indicates that the lot was purchased at the maximum price.
	StatusSoldBuynow Status = "soldBuynow"
)

// Period defines the list of possible lot periods.
type Period int

const (
	// MinPeriod indicates that lot minimal period time is 1 hour.
	MinPeriod Period = 1
	// MaxPeriod indicates that lot maximal period time is 120 hour.
	MaxPeriod Period = 120
)

// Config defines configuration for marketplace.
type Config struct {
	LotRenewalInterval time.Duration `json:"lotRenewalInterval"`
	pagination.Cursor  `json:"cursor"`

	RPCNodeAddress string `json:"rpcNodeAddress"`

	NFTContractAddress    string `json:"nftContractAddress"`
	MarketContractAddress string `json:"marketContractAddress"`
	TokenContractAddress  string `json:"tokenContractAddress"`

	MarketContractPackageAddress string `json:"marketContractPackageAddress"`
	NFTApprovePrefix             string `json:"nftApprovePrefix"`
	CreateListingPrefix          string `json:"createListingPrefix"`

	TokenAmountForApproving int64 `json:"tokenAmountForApproving"`

	PrivateOwnerAccountKey string `json:"privateOwnerAccountKey"`
	PublicOwnerAccountKey  string `json:"publicOwnerAccountKey"`
}

// CreateLot entity that contains the values required to create the lot.
type CreateLot struct {
	CardID     uuid.UUID `json:"cardId"`
	Type       Type      `json:"type"`
	UserID     uuid.UUID `json:"userId"`
	StartPrice big.Int   `json:"startPrice"`
	MaxPrice   big.Int   `json:"maxPrice"`
	Period     Period    `json:"period"`
}

// BetLot entity that contains the values required to place bet the lot.
type BetLot struct {
	CardID    uuid.UUID `json:"cardId"`
	UserID    uuid.UUID `json:"userId"`
	BetAmount big.Int   `json:"betAmount"`
}

// WinLot entity that contains the values required to win the lot.
type WinLot struct {
	CardID    uuid.UUID `json:"cardId"`
	Type      Type      `json:"type"`
	UserID    uuid.UUID `json:"userId"`
	ShopperID uuid.UUID `json:"shopperID"`
	Status    Status    `json:"status"`
	Amount    big.Int   `json:"amount"`
}

// ValidateCreateLot check is empty fields of create lot entity.
func (createLot CreateLot) ValidateCreateLot() error {
	if createLot.CardID.String() == "" {
		return ErrMarketplace.New("item id is empty")
	}

	if createLot.StartPrice.BitLen() == 0 {
		return ErrMarketplace.New("start price is empty")
	}

	if createLot.Period == 0 {
		return ErrMarketplace.New("period is empty")
	}

	return nil
}

// ValidateBetLot check is empty fields of bet lot entity.
func (betLot BetLot) ValidateBetLot() error {
	if betLot.CardID.String() == "" {
		return ErrMarketplace.New("lot id is empty")
	}

	if betLot.BetAmount.BitLen() == 0 {
		return ErrMarketplace.New("bet amount is empty")
	}

	return nil
}

// ResponseCreateLot entity describes the values required to response for create lot in admin.
type ResponseCreateLot struct {
	Cards cards.Page
	Users []users.User
}

// ResponsePlaceBetLot entity describes the values required to response for place bet lot in admin.
type ResponsePlaceBetLot struct {
	ID    uuid.UUID
	Users []users.User
}

// Page holds lot page entity which is used to show listed page of lots.
type Page struct {
	Lots []Lot           `json:"lots"`
	Page pagination.Page `json:"page"`
}
