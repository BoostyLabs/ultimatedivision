// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package marketplaces

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/zeebo/errs"
)

// ErrNoLot indicated that lot does not exist.
var ErrNoLot = errs.Class("lot does not exist")

// DB is exposing access to lots db.
//
// architecture: DB
type DB interface {
	// Create add lot in the data base.
	Create(ctx context.Context, lot Lot) error
	// Get returns lot by id from the data base.
	Get(ctx context.Context, id uuid.UUID) (Lot, error)
	// ListActive returns active lots from the data base.
	ListActive(ctx context.Context) ([]Lot, error)
}

// Lot describes lot entity.
type Lot struct {
	ID           uuid.UUID `json:"id"`
	ItemID       uuid.UUID `json:"itemId"`
	Type         Type      `json:"type"`
	UserID       uuid.UUID `json:"userId"`
	ShopperID    uuid.UUID `json:"shopperId"`
	Status       Status    `json:"status"`
	StartPrice   float64   `json:"startPrice"`
	MaxPrice     float64   `json:"maxPrice"`
	CurrentPrice float64   `json:"currentPrice"`
	StartTime    time.Time `json:"startTime"`
	EndTime      time.Time `json:"endTime"`
	Period       Period    `json:"period"`
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
	// StatusActive indicates that lot status is active.
	StatusActive Status = "active"
	// StatusExpired indicates that lot status is expired.
	StatusExpired Status = "expired"
	// StatusSold indicates that lot status is sold.
	StatusSold Status = "sold"
	// StatusSoldBuynow indicates that lot status is sold buynow.
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
