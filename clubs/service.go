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
	cards *cards.Service
}

// NewService is a constructor for clubs service.
func NewService(clubs DB, users *users.Service, cards *cards.Service) *Service {
	return &Service{
		clubs: clubs,
		users: users,
		cards: cards,
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

// AddSquadCards adds cards to the squad.
func (service *Service) AddSquadCards(ctx context.Context, squadID uuid.UUID, newSquadCard SquadCard) error {
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

// CalculateEffectivenessOfSquad calculates effectiveness of user's squad.
func (service *Service) CalculateEffectivenessOfSquad(ctx context.Context, squadCards []SquadCard) (float64, error) {
	var effectiveness float64

	if len(squadCards) == 0 {
		return float64(0), nil
	}

	cardsFromSquad, err := service.cards.GetCardsFromSquadCards(ctx, squadCards[0].SquadID)
	if err != nil {
		return float64(0), ErrClubs.Wrap(err)
	}

	for index, squadCard := range squadCards {
		switch squadCard.Position {
		case GK:
			effectiveness += cardsFromSquad[index].EfficientGK()
		case LB:
			effectiveness += cardsFromSquad[index].EfficientLB()
		case LWB:
			effectiveness += cardsFromSquad[index].EfficientLB()
		case CCD:
			effectiveness += cardsFromSquad[index].EfficientCD()
		case LCD:
			effectiveness += cardsFromSquad[index].EfficientCD()
		case RCD:
			effectiveness += cardsFromSquad[index].EfficientCD()
		case RB:
			effectiveness += cardsFromSquad[index].EfficientLB()
		case RWB:
			effectiveness += cardsFromSquad[index].EfficientLB()
		case CCDM:
			effectiveness += cardsFromSquad[index].EfficientCDM()
		case LCDM:
			effectiveness += cardsFromSquad[index].EfficientCDM()
		case RCDM:
			effectiveness += cardsFromSquad[index].EfficientCDM()
		case CCM:
			effectiveness += cardsFromSquad[index].EfficientCM()
		case LCM:
			effectiveness += cardsFromSquad[index].EfficientCM()
		case RCM:
			effectiveness += cardsFromSquad[index].EfficientCM()
		case CCAM:
			effectiveness += cardsFromSquad[index].EfficientCAM()
		case LCAM:
			effectiveness += cardsFromSquad[index].EfficientCAM()
		case RCAM:
			effectiveness += cardsFromSquad[index].EfficientCAM()
		case LM:
			effectiveness += cardsFromSquad[index].EfficientLM()
		case RM:
			effectiveness += cardsFromSquad[index].EfficientLM()
		case LW:
			effectiveness += cardsFromSquad[index].EfficientLW()
		case RW:
			effectiveness += cardsFromSquad[index].EfficientLW()
		case CST:
			effectiveness += cardsFromSquad[index].EfficientST()
		case RST:
			effectiveness += cardsFromSquad[index].EfficientST()
		case LST:
			effectiveness += cardsFromSquad[index].EfficientST()
		}
	}

	return effectiveness, nil
}
