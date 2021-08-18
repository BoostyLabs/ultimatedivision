// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package marketplace

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	"ultimatedivision/cards"
	"ultimatedivision/internal/auth"
	"ultimatedivision/users"
)

// Service is handling marketplace related logic.
//
// architecture: Service
type Service struct {
	marketplace DB
	users       *users.Service
	cards       *cards.Service
}

// NewService is a constructor for marketplace service.
func NewService(marketplace DB, users *users.Service, cards *cards.Service) *Service {
	return &Service{
		marketplace: marketplace,
		users:       users,
		cards:       cards,
	}
}

// CreateLot add lot in DB.
func (service *Service) CreateLot(ctx context.Context, createLot CreateLot) error {

	_, err := auth.GetClaims(ctx)
	if err != nil {
		return ErrMarketplace.Wrap(err)
	}

	card, err := service.cards.Get(ctx, createLot.ItemID)
	if err == nil {
		if card.Status == cards.StatusSale {
			return ErrMarketplace.Wrap(fmt.Errorf("the card is already on sale"))
		}

		if err := service.cards.UpdateStatus(ctx, createLot.ItemID, cards.StatusSale); err != nil {
			return ErrMarketplace.Wrap(err)
		}

		createLot.Type = TypeCard
	}
	// TODO: check other items

	if createLot.Type == "" {
		return ErrMarketplace.Wrap(fmt.Errorf("not found item by id"))
	}

	if _, err := service.users.Get(ctx, createLot.UserID); err != nil {
		return ErrMarketplace.Wrap(err)
	}

	if createLot.MaxPrice != 0 || createLot.MaxPrice < createLot.StartPrice {
		return ErrMarketplace.Wrap(fmt.Errorf("max price less start price"))
	}

	if createLot.Period < MinPeriod && createLot.Period < MaxPeriod {
		return ErrMarketplace.Wrap(fmt.Errorf("period exceed the range from 1 to 120 hours"))
	}

	lot := Lot{
		ID:         uuid.New(),
		ItemID:     createLot.ItemID,
		Type:       createLot.Type,
		UserID:     createLot.UserID,
		Status:     StatusActive,
		StartPrice: createLot.StartPrice,
		MaxPrice:   createLot.MaxPrice,
		StartTime:  time.Now().UTC(),
		EndTime:    time.Now().UTC().Add(time.Duration(createLot.Period) * time.Hour),
		Period:     createLot.Period,
	}

	return ErrMarketplace.Wrap(service.marketplace.CreateLot(ctx, lot))
}

// GetLotByID returns lot by id from DB.
func (service *Service) GetLotByID(ctx context.Context, id uuid.UUID) (Lot, error) {
	return service.marketplace.GetLotByID(ctx, id)
}

// ListActiveLots returns active lots from DB.
func (service *Service) ListActiveLots(ctx context.Context) ([]Lot, error) {
	return service.marketplace.ListActiveLots(ctx)
}

// ListActiveLotsWhereEndTimeLTENow returns active lots from DB.
func (service *Service) ListActiveLotsWhereEndTimeLTENow(ctx context.Context) ([]Lot, error) {
	return service.marketplace.ListActiveLotsWhereEndTimeLTENow(ctx)
}

// PlaceBetLot checks the amount of money and makes a bet.
func (service *Service) PlaceBetLot(ctx context.Context, betLot BetLot) error {
	_, err := auth.GetClaims(ctx)
	if err != nil {
		return ErrMarketplace.Wrap(err)
	}

	if _, err := service.users.Get(ctx, betLot.ShopperID); err != nil {
		return ErrMarketplace.Wrap(err)
	}
	// TODO: check if the user has the required amount of money.

	lot, err := service.GetLotByID(ctx, betLot.ID)
	if err != nil {
		return ErrMarketplace.Wrap(err)
	}

	if betLot.BetAmount < lot.CurrentPrice {
		return ErrMarketplace.Wrap(fmt.Errorf("not enough money"))
	}

	/** TODO: the transaction may be required for all operations,
	so that an error in the middle does not lead to an unwanted result in the database. **/

	// TODO: update status to `hold` for new user's money.
	// TODO: unhold old user's money if exist.

	if err := service.UpdateShopperIDLot(ctx, betLot.ID, betLot.ShopperID); err != nil {
		return ErrMarketplace.Wrap(err)
	}

	if betLot.BetAmount >= lot.MaxPrice || lot.MaxPrice != 0 {
		if err := service.UpdateCurrentPriceLot(ctx, betLot.ID, lot.MaxPrice); err != nil {
			return ErrMarketplace.Wrap(err)
		}

		winLot := WinLot{
			ID:        betLot.ID,
			ItemID:    lot.ItemID,
			Type:      TypeCard,
			UserID:    lot.UserID,
			ShopperID: betLot.ShopperID,
			Status:    StatusSoldBuynow,
			Amount:    lot.MaxPrice,
		}

		if err := service.WinLot(ctx, winLot); err != nil {
			return ErrMarketplace.Wrap(err)
		}

	} else {
		if err := service.UpdateCurrentPriceLot(ctx, betLot.ID, betLot.BetAmount); err != nil {
			return ErrMarketplace.Wrap(err)
		}
		if lot.EndTime.Sub(time.Now().UTC()) < time.Minute {
			if err := service.UpdateEndTimeLot(ctx, betLot.ID, time.Now().UTC().Add(time.Minute)); err != nil {
				return ErrMarketplace.Wrap(err)
			}
		}
	}

	return nil
}

// WinLot changes the owner of the item and transfers money.
func (service *Service) WinLot(ctx context.Context, winLot WinLot) error {
	_, err := auth.GetClaims(ctx)
	if err != nil {
		return ErrMarketplace.Wrap(err)
	}

	if err := service.UpdateStatusLot(ctx, winLot.ID, winLot.Status); err != nil {
		return ErrMarketplace.Wrap(err)
	}

	// TODO: transfer money to the old cardholder from new user. If userID == shopperID not transfer mb

	if winLot.Type == TypeCard {
		if err := service.cards.UpdateStatus(ctx, winLot.ItemID, cards.StatusActive); err != nil {
			return ErrMarketplace.Wrap(err)
		}

		if err := service.cards.UpdateUserID(ctx, winLot.ItemID, winLot.ShopperID); err != nil {
			return ErrMarketplace.Wrap(err)
		}
	}
	// TODO: check other items

	return nil
}

// UpdateShopperIDLot updates shopper id of lot.
func (service *Service) UpdateShopperIDLot(ctx context.Context, id, shopperID uuid.UUID) error {
	_, err := auth.GetClaims(ctx)
	if err != nil {
		return ErrMarketplace.Wrap(err)
	}
	return ErrMarketplace.Wrap(service.marketplace.UpdateShopperIDLot(ctx, id, shopperID))
}

// UpdateStatusLot updates status of lot.
func (service *Service) UpdateStatusLot(ctx context.Context, id uuid.UUID, status Status) error {
	_, err := auth.GetClaims(ctx)
	if err != nil {
		return ErrMarketplace.Wrap(err)
	}
	return ErrMarketplace.Wrap(service.marketplace.UpdateStatusLot(ctx, id, status))
}

// UpdateCurrentPriceLot updates current price of lot.
func (service *Service) UpdateCurrentPriceLot(ctx context.Context, id uuid.UUID, currentPrice float64) error {
	_, err := auth.GetClaims(ctx)
	if err != nil {
		return ErrMarketplace.Wrap(err)
	}
	return ErrMarketplace.Wrap(service.marketplace.UpdateCurrentPriceLot(ctx, id, currentPrice))
}

// UpdateEndTimeLot updates end time of lot.
func (service *Service) UpdateEndTimeLot(ctx context.Context, id uuid.UUID, endTime time.Time) error {
	_, err := auth.GetClaims(ctx)
	if err != nil {
		return ErrMarketplace.Wrap(err)
	}
	return ErrMarketplace.Wrap(service.marketplace.UpdateEndTimeLot(ctx, id, endTime))
}
