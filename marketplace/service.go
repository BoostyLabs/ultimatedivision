// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package marketplace

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	"ultimatedivision/cards"
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
func (service *Service) CreateLot(ctx context.Context, lot Lot) error {

	card, err := service.cards.Get(ctx, lot.ItemID)
	if err == nil {
		if card.Status == cards.StatusSale {
			return ErrMarketplace.Wrap(fmt.Errorf("the card is already on sale"))
		}

		if err := service.cards.UpdateStatus(ctx, lot.ItemID, cards.StatusSale); err != nil {
			return ErrMarketplace.Wrap(err)
		}

		lot.Type = TypeCard
	}
	// TODO: check other items

	if lot.Type == "" {
		return ErrMarketplace.Wrap(fmt.Errorf("not found item by id"))
	}

	if _, err := service.users.Get(ctx, lot.UserID); err != nil {
		return ErrMarketplace.Wrap(err)
	}

	if lot.MaxPrice != 0 || lot.MaxPrice < lot.StartPrice {
		return ErrMarketplace.Wrap(fmt.Errorf("max price less start price"))
	}

	if lot.Period < MinPeriod && lot.Period < MaxPeriod {
		return ErrMarketplace.Wrap(fmt.Errorf("period exceed the range from 1 to 120 hours"))
	}

	lot = Lot{
		ID:         uuid.New(),
		ItemID:     lot.ItemID,
		Type:       lot.Type,
		UserID:     lot.UserID,
		Status:     StatusActive,
		StartPrice: lot.StartPrice,
		MaxPrice:   lot.MaxPrice,
		StartTime:  time.Now().UTC(),
		EndTime:    lot.StartTime.Add(time.Duration(lot.Period) * time.Hour),
		Period:     lot.Period,
	}

	return service.marketplace.CreateLot(ctx, lot)
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

// PlaceBet checks the amount of money and makes a bet.
func (service *Service) PlaceBet(ctx context.Context, id, shopperID uuid.UUID, betAmount float64) error {
	if _, err := service.users.Get(ctx, shopperID); err != nil {
		return ErrMarketplace.Wrap(err)
	}
	// TODO: check if the user has the required amount of money.

	lot, err := service.GetLotByID(ctx, id)
	if err != nil {
		return ErrMarketplace.Wrap(err)
	}

	if betAmount > lot.CurrentPrice {
		/** TODO: the transaction may be required for all operations,
		so that an error in the middle does not lead to an unwanted result in the database. **/

		// TODO: update status to `hold` for new user's money.
		// TODO: unhold old user's money if exist.

		if err := service.UpdateShopperIDLot(ctx, id, shopperID); err != nil {
			return ErrMarketplace.Wrap(err)
		}

		if betAmount >= lot.MaxPrice || lot.MaxPrice != 0 {
			if err := service.UpdateCurrentPriceLot(ctx, id, lot.MaxPrice); err != nil {
				return ErrMarketplace.Wrap(err)
			}

			winLot := WinLot{
				ID:        id,
				ItemID:    lot.ItemID,
				Type:      TypeCard,
				UserID:    lot.UserID,
				ShopperID: shopperID,
				Status:    StatusSoldBuynow,
				Amount:    lot.MaxPrice,
			}

			if err := service.WinLot(ctx, winLot); err != nil {
				return ErrMarketplace.Wrap(err)
			}

		} else {
			if err := service.UpdateCurrentPriceLot(ctx, id, betAmount); err != nil {
				return ErrMarketplace.Wrap(err)
			}
			if lot.EndTime.Sub(time.Now().UTC()) < time.Minute {
				if err := service.UpdateEndTimeLot(ctx, id, time.Now().UTC().Add(time.Minute)); err != nil {
					return ErrMarketplace.Wrap(err)
				}
			}
		}

	} else {
		return ErrMarketplace.Wrap(fmt.Errorf("not enough money"))
	}

	return nil
}

// WinLot changes the owner of the item and transfers money.
func (service *Service) WinLot(ctx context.Context, winLot WinLot) error {
	if err := service.UpdateStatusLot(ctx, winLot.ID, winLot.Status); err != nil {
		return ErrMarketplace.Wrap(err)
	}

	// TODO: transfer money to the old cardholder from new user. If userID == shopperID not transfer mb
	// TODO: change userId for item and status to active.

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
	return service.marketplace.UpdateShopperIDLot(ctx, id, shopperID)
}

// UpdateStatusLot updates status of lot.
func (service *Service) UpdateStatusLot(ctx context.Context, id uuid.UUID, status Status) error {
	return service.marketplace.UpdateStatusLot(ctx, id, status)
}

// UpdateCurrentPriceLot updates current price of lot.
func (service *Service) UpdateCurrentPriceLot(ctx context.Context, id uuid.UUID, currentPrice float64) error {
	return service.marketplace.UpdateCurrentPriceLot(ctx, id, currentPrice)
}

// UpdateEndTimeLot updates end time of lot.
func (service *Service) UpdateEndTimeLot(ctx context.Context, id uuid.UUID, endTime time.Time) error {
	return service.marketplace.UpdateEndTimeLot(ctx, id, endTime)
}
