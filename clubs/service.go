// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package clubs

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/zeebo/errs"

	"ultimatedivision/users"
)

// ErrClubs indicates that there was an error in the service.
var ErrClubs = errs.Class("clubs service error")

// totalSquadCards defines number of players in the full squad.
const totalSquadCards = 11

// Service is handling clubs related logic.
//
// architecture: Service
type Service struct {
	clubs DB
	users *users.Service
}

// NewService is a constructor for clubs service.
func NewService(clubs DB, users *users.Service) *Service {
	return &Service{
		clubs: clubs,
		users: users,
	}
}

// Create creates clubs.
func (service *Service) Create(ctx context.Context, userID uuid.UUID) (uuid.UUID, error) {
	nickname, err := service.users.GetNickNameByID(ctx, userID)
	if err != nil {
		return uuid.New(), ErrClubs.Wrap(err)
	}

	newClub := Club{
		ID:        uuid.New(),
		OwnerID:   userID,
		Name:      nickname,
		CreatedAt: time.Now().UTC(),
	}

	clubID, err := service.clubs.Create(ctx, newClub)

	return clubID, ErrClubs.Wrap(err)
}

// CreateSquad creates new squad for club.
func (service *Service) CreateSquad(ctx context.Context, clubID uuid.UUID) (uuid.UUID, error) {
	newSquad := Squad{
		ID:     uuid.New(),
		ClubID: clubID,
	}

	squadID, err := service.clubs.CreateSquad(ctx, newSquad)

	return squadID, ErrClubs.Wrap(err)
}

// AddSquadCards adds cards to the squad.
func (service *Service) AddSquadCards(ctx context.Context, squadID uuid.UUID, squadCards SquadCard) error {
	formation, err := service.clubs.GetFormation(ctx, squadID)
	if err != nil {
		return ErrClubs.Wrap(err)
	}

	squadCards.SquadID = squadID
	squadCards.Position = FormationToPosition[formation][squadCards.Position]

	return ErrClubs.Wrap(service.clubs.AddSquadCard(ctx, squadCards))
}

// Delete deletes card from squad.
func (service *Service) Delete(ctx context.Context, squadID, cardID uuid.UUID) error {
	return ErrClubs.Wrap(service.clubs.DeleteSquadCard(ctx, squadID, cardID))
}

// UpdateSquad updates tactic and formation of the squad.
func (service *Service) UpdateSquad(ctx context.Context, squadID uuid.UUID, formation Formation, tactic Tactic, captainID uuid.UUID) error {
	updatedSquad := Squad{
		ID:        squadID,
		Tactic:    tactic,
		Formation: formation,
		CaptainID: captainID,
	}

	return ErrClubs.Wrap(service.clubs.UpdateTacticFormationCaptain(ctx, updatedSquad))
}

// UpdateCardPosition updates position of card in the squad.
func (service *Service) UpdateCardPosition(ctx context.Context, squadID uuid.UUID, cardID uuid.UUID, newPosition Position) error {
	formation, err := service.clubs.GetFormation(ctx, squadID)
	if err != nil {
		return ErrClubs.Wrap(err)
	}

	newPosition = FormationToPosition[formation][newPosition]

	return ErrClubs.Wrap(service.clubs.UpdatePosition(ctx, newPosition, squadID, cardID))
}

// GetSquad returns squad of club.
func (service *Service) GetSquad(ctx context.Context, clubID uuid.UUID) (Squad, error) {
	squad, err := service.clubs.GetSquad(ctx, clubID)
	return squad, ErrClubs.Wrap(err)
}

// GetSquadCards returns al cards from squad.
func (service *Service) GetSquadCards(ctx context.Context, squadID uuid.UUID) ([]SquadCard, error) {
	squadCards, err := service.clubs.ListSquadCards(ctx, squadID)
	if err != nil {
		return squadCards, ErrClubs.Wrap(err)
	}

	if len(squadCards) < totalSquadCards {
		for i := len(squadCards); i < totalSquadCards; i++ {
			var squadCard = SquadCard{
				SquadID: squadID,
			}

			squadCards = append(squadCards, squadCard)
		}
	}

	formation, err := service.clubs.GetFormation(ctx, squadID)
	if err != nil {
		return nil, ErrClubs.Wrap(err)
	}

	for i := 0; i < len(squadCards); i++ {
		for j := 0; j < len(FormationToPosition[formation]); j++ {
			if squadCards[i].Position == FormationToPosition[formation][j] {
				squadCards[i].Position = Position(j)
			}
		}
	}

	return squadCards, ErrClubs.Wrap(err)
}

// Get returns user club.
func (service *Service) Get(ctx context.Context, userID uuid.UUID) (Club, error) {
	club, err := service.clubs.GetByUserID(ctx, userID)
	return club, ErrClubs.Wrap(err)
}
