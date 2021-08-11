// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package marketplaces

import (
	"context"

	"github.com/google/uuid"
)

// Service is handling marketplaces related logic.
//
// architecture: Service
type Service struct {
	marketplaces DB
}

// NewService is a constructor for marketplaces service.
func NewService(marketplaces DB) *Service {
	return &Service{
		marketplaces: marketplaces,
	}
}

// Create add lot in DB.
func (service *Service) Create(ctx context.Context, lot Lot) error {
	return service.marketplaces.Create(ctx, lot)
}

// Get returns lot by id from DB.
func (service *Service) Get(ctx context.Context, id uuid.UUID) (Lot, error) {
	return service.marketplaces.Get(ctx, id)
}

// ListActive returns active lots from DB.
func (service *Service) ListActive(ctx context.Context) ([]Lot, error) {
	return service.marketplaces.ListActive(ctx)
}
