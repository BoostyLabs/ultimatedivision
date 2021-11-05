// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package seasons

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/zeebo/errs"

	"ultimatedivision/divisions"
)

// ErrSeasons indicates that there was an error in the service.
var ErrSeasons = errs.Class("seasons service error")

// Service is handling seasons related logic.
//
// architecture: Service
type Service struct {
	seasons   DB
	divisions divisions.Service
	config    Config
}

// NewService is a constructor for seasons service.
func NewService(seasons DB, config Config) *Service {
	return &Service{
		seasons: seasons,
		config:  config,
	}
}

// Create creates a season.
func (service *Service) Create(ctx context.Context) error {
	seasons, err := service.divisions.List(ctx)
	if err != nil {
		return ErrSeasons.Wrap(err)
	}

	for _, division := range seasons {
		season := Season{
			DivisionID: division.ID,
			StartedAt:  time.Now().UTC(),
			EndedAt:    time.Time{},
		}
		if err = service.seasons.Create(ctx, season); err != nil {
			return ErrSeasons.Wrap(err)
		}
	}

	return nil
}

// List returns all seasons from DB.
func (service *Service) List(ctx context.Context) ([]Season, error) {
	seasons, err := service.seasons.List(ctx)
	return seasons, ErrSeasons.Wrap(err)
}

// Get returns season from DB.
func (service *Service) Get(ctx context.Context, seasonID uuid.UUID) (Season, error) {
	season, err := service.seasons.Get(ctx, seasonID)
	return season, ErrSeasons.Wrap(err)
}

// Delete deletes a season.
func (service *Service) Delete(ctx context.Context, id uuid.UUID) error {
	return ErrSeasons.Wrap(service.seasons.Delete(ctx, id))
}
