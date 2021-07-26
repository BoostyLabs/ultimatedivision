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

// CreateClub creates clubs.
func (service *Service) CreateClub(ctx context.Context, userID uuid.UUID) error {
	newClub := Club{
		ID:        uuid.New(),
		OwnerID:   userID,
		CreatedAt: time.Now().UTC(),
	}

	return service.clubs.CreateClub(ctx, newClub)
}

// CreateSquad creates new squad for club.
func (service *Service) CreateSquad(ctx context.Context, clubID uuid.UUID) error {
	newSquad := Squads{
		ID:     uuid.New(),
		ClubID: clubID,
	}
	return service.clubs.CreateSquad(ctx, newSquad)
}

// Add add new card to the squad of the club.
func (service *Service) Add(ctx context.Context, squadID uuid.UUID, cardID uuid.UUID, position Position) error {
	capitan, err := service.clubs.GetCapitan(ctx, squadID)
	if err != nil {
		return ErrClubs.Wrap(err)
	}

	newSquadCard := SquadCards{
		ID:       squadID,
		CardID:   cardID,
		Position: position,
		Capitan:  capitan,
	}

	err = service.clubs.Add(ctx, newSquadCard)

	return ErrClubs.Wrap(err)
}

// UpdateSquad updates tactic and formation of the squad.
func (service *Service) UpdateSquad(ctx context.Context, squadID uuid.UUID, tactic Tactic, formation Formation) error {
	updatedSquad := Squads{
		ID:        squadID,
		Formation: formation,
		Tactic:    tactic,
	}

	err := service.clubs.UpdateTacticFormation(ctx, updatedSquad)

	return ErrClubs.Wrap(err)
}

// UpdateCapitan updates capitan in the club.
func (service *Service) UpdateCapitan(ctx context.Context, squadID uuid.UUID, capitanID uuid.UUID) error {
	err := service.clubs.UpdateCapitan(ctx, capitanID, squadID)

	return ErrClubs.Wrap(err)
}

// UpdateCardPosition updates position of card in the squad.
func (service *Service) UpdateCardPosition(ctx context.Context, squadID uuid.UUID, cardID uuid.UUID, position Position) error {
	err := service.clubs.UpdatePosition(ctx, squadID, cardID, position)

	return ErrClubs.Wrap(err)
}

// GetSquad returns all squads from club.
func (service *Service) GetSquad(ctx context.Context, clubID uuid.UUID) (Squads, []SquadCards, error) {
	squad, err := service.clubs.GetSquad(ctx, clubID)
	if err != nil {
		return Squads{},nil, ErrClubs.Wrap(err)
	}

	squadCards, err := service.clubs.ListSquadCards(ctx, squad.ID)
	if err != nil {
		return Squads{},nil, ErrClubs.Wrap(err)
	}

	return squad, squadCards, nil
}

// ListClub returns users clubs.
func (service *Service) ListClub(ctx context.Context, userID uuid.UUID) ([]Club, error) {
	userClubs, err := service.clubs.ListClubs(ctx, userID)

	return userClubs, err
}
