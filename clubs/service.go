// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package clubs

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/zeebo/errs"

	"ultimatedivision/cards"
	"ultimatedivision/users"
)

// ErrClubs indicates that there was an error in the service.
var ErrClubs = errs.Class("clubs service error")

// squadSize defines number of cards in the full squad.
const squadSize = 11

// Service is handling clubs related logic.
//
// architecture: Service
type Service struct {
	clubs DB
	users *users.Service
	card  Cards
}

// Cards is exposing access to method CardsWithNewPositions and methods to count effectiveness of card in cards service.
type Cards interface {
	CardsWithNewPositions(ctx context.Context, cards []SquadCard, positions []Position) (map[Position]uuid.UUID, error)
	GetCardsFromSquadCards(ctx context.Context, id uuid.UUID) ([]cards.Card, error)
	EffectivenessGK(card cards.Card) float64
	EffectivenessCD(card cards.Card) float64
	EffectivenessLBorRB(card cards.Card) float64
	EffectivenessCDM(card cards.Card) float64
	EffectivenessCM(card cards.Card) float64
	EffectivenessCAM(card cards.Card) float64
	EffectivenessRMorLM(card cards.Card) float64
	EffectivenessRWorLW(card cards.Card) float64
	EffectivenessST(card cards.Card) float64
}

// NewService is a constructor for clubs service.
func NewService(clubs DB, users *users.Service, card Cards) *Service {
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

// AddSquadCard adds card to the squad.
func (service *Service) AddSquadCard(ctx context.Context, squadID uuid.UUID, newSquadCard SquadCard) error {
	squadCards, err := service.clubs.ListSquadCards(ctx, squadID)
	if err != nil {
		return ErrClubs.Wrap(err)
	}

	if len(squadCards) == squadSize {
		return ErrClubs.New("squad is full")
	}

	formation, err := service.clubs.GetFormation(ctx, squadID)
	if err != nil {
		return ErrClubs.Wrap(err)
	}

	newSquadCard.SquadID = squadID
	newSquadCard.Position = FormationToPosition[formation][newSquadCard.Position]

	for _, card := range squadCards {
		if card.Position != newSquadCard.Position {
			continue
		}

		err = service.clubs.DeleteSquadCard(ctx, squadID, card.CardID)
		if err != nil {
			return ErrClubs.Wrap(err)
		}
		break
	}

	return ErrClubs.Wrap(service.clubs.AddSquadCard(ctx, newSquadCard))
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

	return ErrClubs.Wrap(service.clubs.UpdatePositions(ctx, updatedCards))
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

	if len(squadCards) < squadSize {
		for i := len(squadCards); i < squadSize; i++ {
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
				break
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

	err = service.clubs.UpdatePositions(ctx, squadCardsWithNewPositions)
	if err != nil {
		return nil, ErrClubs.Wrap(err)
	}

	return cardsWithNewPositions, nil
}

// CalculateEffectivenessOfSquad calculates effectiveness of user's squad.
func (service *Service) CalculateEffectivenessOfSquad(ctx context.Context, squadCards []SquadCard) (float64, error) {
	var effectiveness float64

	if len(squadCards) == 0 {
		return float64(0), nil
	}

	cardsFromSquad, err := service.card.GetCardsFromSquadCards(ctx, squadCards[0].SquadID)
	if err != nil {
		return float64(0), ErrClubs.Wrap(err)
	}

	for index, squadCard := range squadCards {
		switch squadCard.Position {
		case GK:
			effectiveness += service.card.EffectivenessGK(cardsFromSquad[index])
		case LB:
			effectiveness += service.card.EffectivenessLBorRB(cardsFromSquad[index])
		case LWB:
			effectiveness += service.card.EffectivenessLBorRB(cardsFromSquad[index])
		case CCD:
			effectiveness += service.card.EffectivenessCD(cardsFromSquad[index])
		case LCD:
			effectiveness += service.card.EffectivenessCD(cardsFromSquad[index])
		case RCD:
			effectiveness += service.card.EffectivenessCD(cardsFromSquad[index])
		case RB:
			effectiveness += service.card.EffectivenessLBorRB(cardsFromSquad[index])
		case RWB:
			effectiveness += service.card.EffectivenessLBorRB(cardsFromSquad[index])
		case CCDM:
			effectiveness += service.card.EffectivenessCDM(cardsFromSquad[index])
		case LCDM:
			effectiveness += service.card.EffectivenessCDM(cardsFromSquad[index])
		case RCDM:
			effectiveness += service.card.EffectivenessCDM(cardsFromSquad[index])
		case CCM:
			effectiveness += service.card.EffectivenessCM(cardsFromSquad[index])
		case LCM:
			effectiveness += service.card.EffectivenessCM(cardsFromSquad[index])
		case RCM:
			effectiveness += service.card.EffectivenessCM(cardsFromSquad[index])
		case CCAM:
			effectiveness += service.card.EffectivenessCAM(cardsFromSquad[index])
		case LCAM:
			effectiveness += service.card.EffectivenessCAM(cardsFromSquad[index])
		case RCAM:
			effectiveness += service.card.EffectivenessCAM(cardsFromSquad[index])
		case LM:
			effectiveness += service.card.EffectivenessRMorLM(cardsFromSquad[index])
		case RM:
			effectiveness += service.card.EffectivenessRMorLM(cardsFromSquad[index])
		case LW:
			effectiveness += service.card.EffectivenessRWorLW(cardsFromSquad[index])
		case RW:
			effectiveness += service.card.EffectivenessRWorLW(cardsFromSquad[index])
		case CST:
			effectiveness += service.card.EffectivenessST(cardsFromSquad[index])
		case RST:
			effectiveness += service.card.EffectivenessST(cardsFromSquad[index])
		case LST:
			effectiveness += service.card.EffectivenessST(cardsFromSquad[index])
		}
	}

	return effectiveness, nil
}
