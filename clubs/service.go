// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package clubs

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/zeebo/errs"

	"ultimatedivision/internal/auth"
	"ultimatedivision/users"
	"ultimatedivision/users/userauth"
)

// ErrClubs indicates that there was an error in the service.
var ErrClubs = errs.Class("clubs service error")

// Service is handling users related logic.
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
func (service *Service) Create(ctx context.Context) error {
	var id uuid.UUID

	// get user id for admin server.
	idParam, ok := getUserID(ctx)

	claims, err := auth.GetClaims(ctx)
	if err != nil {
		if !ok {
			return userauth.ErrUnauthenticated.Wrap(err)
		}

		if ok {
			id = idParam
		}
	} else {
		id = claims.ID
	}

	nickname, err := service.users.GetNickNameByID(ctx, id)
	if err != nil {
		return ErrClubs.Wrap(err)
	}

	newClub := Club{
		ID:        uuid.New(),
		OwnerID:   id,
		Name:      nickname,
		CreatedAt: time.Now().UTC(),
	}

	return ErrClubs.Wrap(service.clubs.Create(ctx, newClub))
}

// CreateSquad creates new squad for club.
func (service *Service) CreateSquad(ctx context.Context, clubID uuid.UUID) error {
	/*_, err := auth.GetClaims(ctx)
	if err != nil {
		return userauth.ErrUnauthenticated.Wrap(err)
	}*/

	newSquad := Squad{
		ID:     uuid.New(),
		ClubID: clubID,
	}

	return ErrClubs.Wrap(service.clubs.CreateSquad(ctx, newSquad))
}

// Add add new card to the squad of the club.
func (service *Service) Add(ctx context.Context, newSquadCard SquadCard) error {
	_, err := auth.GetClaims(ctx)
	if err != nil {
		return userauth.ErrUnauthenticated.Wrap(err)
	}

	return ErrClubs.Wrap(service.clubs.AddSquadCard(ctx, newSquadCard))
}

// Delete deletes card from squad.
func (service *Service) Delete(ctx context.Context, squadID uuid.UUID, cardID uuid.UUID) error {
	/*_, err := auth.GetClaims(ctx)
	if err != nil {
		return userauth.ErrUnauthenticated.Wrap(err)
	}*/

	return ErrClubs.Wrap(service.clubs.DeleteSquadCard(ctx, squadID, cardID))
}

// UpdateSquad updates tactic and formation of the squad.
func (service *Service) UpdateSquad(ctx context.Context, updatedSquad Squad) error {
	/*_, err := auth.GetClaims(ctx)
	if err != nil {
		return userauth.ErrUnauthenticated.Wrap(err)
	}*/

	return ErrClubs.Wrap(service.clubs.UpdateTacticFormationCaptain(ctx, updatedSquad))
}

// UpdateCardPosition updates position of card in the squad.
func (service *Service) UpdateCardPosition(ctx context.Context, squadCard SquadCard) error {
	/*_, err := auth.GetClaims(ctx)
	if err != nil {
		return userauth.ErrUnauthenticated.Wrap(err)
	}*/

	return ErrClubs.Wrap(service.clubs.UpdatePosition(ctx, squadCard.SquadID, squadCard.CardID, squadCard.Position))
}

// GetSquad returns all squads from club.
func (service *Service) GetSquad(ctx context.Context, clubID uuid.UUID) (Squad, error) {
	/*_, err := auth.GetClaims(ctx)
	if err != nil {
		return Squad{}, userauth.ErrUnauthenticated.Wrap(err)
	}*/

	squad, err := service.clubs.GetSquad(ctx, clubID)
	if err != nil {
		return Squad{}, ErrClubs.Wrap(err)
	}

	return squad, nil
}

// GetSquadCard returns all cards from squad.
func (service *Service) GetSquadCard(ctx context.Context, squadID uuid.UUID) ([]SquadCard, error) {
	/*_, err := auth.GetClaims(ctx)
	if err != nil {
		return nil, userauth.ErrUnauthenticated.Wrap(err)
	}*/

	squadCards, err := service.clubs.ListSquadCards(ctx, squadID)
	if err != nil {
		return nil, ErrClubs.Wrap(err)
	}

	return squadCards, nil
}

// Get returns user club.
func (service *Service) Get(ctx context.Context) (Club, error) {
	var id uuid.UUID

	// get user id for admin server.
	idParam, ok := getUserID(ctx)

	claims, err := auth.GetClaims(ctx)
	if err != nil {
		if !ok {
			return Club{}, userauth.ErrUnauthenticated.Wrap(err)
		}

		if ok {
			id = idParam
		}
	} else {
		id = claims.ID
	}

	club, err := service.clubs.GetByUserID(ctx, id)
	return club, ErrClubs.Wrap(err)
}

// Key defines new type used for key in context.
type Key string

// UserIDKey is key for user id.
const UserIDKey Key = "userID"

// SetUserID creates new context with user id.
func SetUserID(ctx context.Context, userID uuid.UUID) context.Context {
	return context.WithValue(ctx, UserIDKey, userID)
}

// getUserID returns user id from context.
func getUserID(ctx context.Context) (uuid.UUID, bool) {
	key, ok := ctx.Value(UserIDKey).(uuid.UUID)
	return key, ok
}
