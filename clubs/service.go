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

// Service is handling clubs related logic.
//
// architecture: Service
type Service struct {
	clubs DB
	users *users.Service
	card  CardsWithNewPositions
}

// CardsWithNewPositions is exposing access to method CardsWithNewPositions in cards service.
type CardsWithNewPositions interface {
	CardsWithNewPositions(ctx context.Context, cards []SquadCard, positions []Position) (map[Position]uuid.UUID, error)
}

// NewService is a constructor for clubs service.
func NewService(clubs DB, users *users.Service, card CardsWithNewPositions) *Service {
	return &Service{
		clubs: clubs,
		users: users,
		card:  card,
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
		ID:        uuid.New(),
		ClubID:    clubID,
		Formation: FourFourTwo,
		Tactic:    Balanced,
	}

	squadID, err := service.clubs.CreateSquad(ctx, newSquad)

	return squadID, ErrClubs.Wrap(err)
}

// AddSquadCard adds cards to the squad.
func (service *Service) AddSquadCard(ctx context.Context, squadID uuid.UUID, squadCard SquadCard) error {
	formation, err := service.clubs.GetFormation(ctx, squadID)
	if err != nil {
		return ErrClubs.Wrap(err)
	}
	squadCard.SquadID = squadID
	squadCard.Position = FormationToPosition[formation][squadCard.Position]

	return ErrClubs.Wrap(service.clubs.AddSquadCard(ctx, squadCard))
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
	squadCards, err := service.clubs.ListSquadCards(ctx, squadID)
	if err != nil {
		return ErrClubs.Wrap(err)
	}

	var oldPosition Position

	for _, card := range squadCards {
		if card.CardID == cardID {
			oldPosition = card.Position
			break
		}
	}

	formation, err := service.clubs.GetFormation(ctx, squadID)
	if err != nil {
		return ErrClubs.Wrap(err)
	}

	newPosition = FormationToPosition[formation][newPosition]

	updatedCards := make([]SquadCard, 0, 2)

	updatedSquadCard := SquadCard{
		SquadID:  squadID,
		CardID:   cardID,
		Position: newPosition,
	}

	updatedCards = append(updatedCards, updatedSquadCard)

	for _, card := range squadCards {
		if card.Position != newPosition {
			continue
		}

		card.Position = oldPosition
		updatedCards = append(updatedCards, card)
		break
	}

	return ErrClubs.Wrap(service.clubs.UpdatePosition(ctx, updatedCards))
}

// GetSquad returns squad of club.
func (service *Service) GetSquad(ctx context.Context, clubID uuid.UUID) (Squad, error) {
	squad, err := service.clubs.GetSquad(ctx, clubID)
	return squad, ErrClubs.Wrap(err)
}

// GetSquadCards returns al cards from squad.
func (service *Service) GetSquadCards(ctx context.Context, squadID uuid.UUID) ([]SquadCard, error) {
	squadCards, err := service.clubs.ListSquadCards(ctx, squadID)
	return squadCards, ErrClubs.Wrap(err)
}

// Get returns user club.
func (service *Service) Get(ctx context.Context, userID uuid.UUID) (Club, error) {
	club, err := service.clubs.GetByUserID(ctx, userID)
	return club, ErrClubs.Wrap(err)
}

// ChangeFormation is a method that change formation and card position.
func (service *Service) ChangeFormation(ctx context.Context, newFormation Formation, squadID uuid.UUID) (map[Position]uuid.UUID, error) {
	var cardsWithNewPositions map[Position]uuid.UUID

	squadCards, err := service.clubs.ListSquadCards(ctx, squadID)
	if err != nil {
		return nil, ErrClubs.Wrap(err)
	}

	err = service.clubs.UpdateFormation(ctx, newFormation, squadID)
	if err != nil {
		return nil, ErrClubs.Wrap(err)
	}

	cardsWithNewPositions, err = service.card.CardsWithNewPositions(ctx, squadCards, FormationToPosition[newFormation])
	if err != nil {
		return nil, ErrClubs.Wrap(err)
	}

	var squadCardsWithNewPositions []SquadCard
	for position, card := range cardsWithNewPositions {
		squadCard := SquadCard{
			Position: position,
			SquadID:  squadID,
			CardID:   card,
		}

		squadCardsWithNewPositions = append(squadCardsWithNewPositions, squadCard)
	}

	err = service.clubs.UpdatePosition(ctx, squadCardsWithNewPositions)
	if err != nil {
		return nil, ErrClubs.Wrap(err)
	}

	return cardsWithNewPositions, nil
}
