// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package clubs

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/zeebo/errs"
)

// ErrClubs indicates that there was an error in the service.
var ErrClubs = errs.Class("clubs service error")

// Service is handling users related logic.
//
// architecture: Service
type Service struct {
	clubs DB
}

// NewService is a constructor for clubs service.
func NewService(clubs DB) *Service {
	return &Service{
		clubs: clubs,
	}
}

// Create creates clubs.
func (service *Service) Create(ctx context.Context, userID uuid.UUID) error {
	newClub := Club{
		ID:        uuid.New(),
		OwnerID:   userID,
		CreatedAt: time.Now().UTC(),
	}

	return service.clubs.Create(ctx, newClub)
}

// CreateSquad creates new squad for club.
func (service *Service) CreateSquad(ctx context.Context, clubID uuid.UUID) error {
	newSquad := Squad{
		ID:     uuid.New(),
		ClubID: clubID,
	}

	return service.clubs.CreateSquad(ctx, newSquad)
}

// Add add new card to the squad of the club.
func (service *Service) Add(ctx context.Context, newSquadCard SquadCard) error {
	capitan, err := service.clubs.GetCapitan(ctx, newSquadCard.ID)
	if err != nil {
		return ErrClubs.Wrap(err)
	}

	newSquadCard.Capitan = capitan

	return service.clubs.Add(ctx, newSquadCard)
}

// Delete deletes card from squad.
func (service *Service) Delete(ctx context.Context, squadID uuid.UUID, cardID uuid.UUID) error {
	return service.clubs.DeleteSquadCard(ctx, squadID, cardID)
}

// UpdateSquad updates tactic and formation of the squad.
func (service *Service) UpdateSquad(ctx context.Context, squadID uuid.UUID, tactic Tactic, formation Formation) error {
	updatedSquad := Squad{
		ID:        squadID,
		Tactic:    tactic,
		Formation: formation,
	}
	return service.clubs.UpdateTacticFormation(ctx, updatedSquad)
}

// UpdateCapitan updates capitan in the club.
func (service *Service) UpdateCapitan(ctx context.Context, squadID uuid.UUID, capitanID uuid.UUID) error {
	return service.clubs.UpdateCapitan(ctx, capitanID, squadID)
}

// UpdateCardPosition updates position of card in the squad.
func (service *Service) UpdateCardPosition(ctx context.Context, squadID uuid.UUID, cardID uuid.UUID, position Position) error {
	return service.clubs.UpdatePosition(ctx, squadID, cardID, position)
}

// GetSquad returns all squads from club.
func (service *Service) GetSquad(ctx context.Context, clubID uuid.UUID) (Squad, []SquadCard, error) {
	squad, err := service.clubs.GetSquad(ctx, clubID)
	if err != nil {
		return Squad{}, nil, ErrClubs.Wrap(err)
	}

	squadCards, err := service.clubs.ListSquadCards(ctx, squad.ID)
	if err != nil {
		return Squad{}, nil, ErrClubs.Wrap(err)
	}

	return squad, squadCards, nil
}

// Get returns user club.
func (service *Service) Get(ctx context.Context, userID uuid.UUID) (Club, error) {
	return service.clubs.Get(ctx, userID)
}
