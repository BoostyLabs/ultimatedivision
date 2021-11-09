// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package seasons

import (
	"context"
	"time"

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
	divisions *divisions.Service
	config    Config
}

// NewService is a constructor for seasons service.
func NewService(seasons DB, config Config, divisions *divisions.Service) *Service {
	return &Service{
		seasons:   seasons,
		divisions: divisions,
		config:    config,
	}
}

// Create creates a season.
func (service *Service) Create(ctx context.Context) error {
	divisions, err := service.divisions.List(ctx)
	if err != nil {
		return ErrSeasons.Wrap(err)
	}

	for _, division := range divisions {
		season := Season{
			DivisionID: division.ID,
			Status:     StatusStarted,
			StartedAt:  time.Now().UTC(),
			EndedAt:    time.Time{},
		}
		if err = service.seasons.Create(ctx, season); err != nil {
			return ErrSeasons.Wrap(err)
		}
	}

	return nil
}

// EndSeason change status when season end.
func (service *Service) EndSeason(ctx context.Context) error {
	return ErrSeasons.Wrap(service.seasons.EndSeason(ctx))
}

// List returns all seasons from DB.
func (service *Service) List(ctx context.Context) ([]Season, error) {
	seasons, err := service.seasons.List(ctx)
	return seasons, ErrSeasons.Wrap(err)
}

// Get returns season from DB.
func (service *Service) Get(ctx context.Context, seasonID int) (Season, error) {
	season, err := service.seasons.Get(ctx, seasonID)
	return season, ErrSeasons.Wrap(err)
}

// Delete deletes a season.
func (service *Service) Delete(ctx context.Context, id int) error {
	return ErrSeasons.Wrap(service.seasons.Delete(ctx, id))
}
