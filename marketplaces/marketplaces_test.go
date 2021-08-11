// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package marketplaces_test

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"ultimatedivision"
	"ultimatedivision/database/dbtesting"
	"ultimatedivision/marketplaces"
	"ultimatedivision/users"
)

func TestMarketplaces(t *testing.T) {

	lot1 := marketplaces.Lot{
		ID:           uuid.New(),
		ItemID:       uuid.New(),
		Type:         marketplaces.TypeCard,
		UserID:       uuid.New(),
		ShopperID:    uuid.New(),
		Status:       marketplaces.StatusSoldBuynow,
		StartPrice:   5.0,
		MaxPrice:     30.0,
		CurrentPrice: 30.0,
		StartTime:    time.Now().UTC(),
		EndTime:      time.Now().AddDate(0, 0, 2).UTC(),
	}

	lot2 := marketplaces.Lot{
		ID:           uuid.New(),
		ItemID:       uuid.New(),
		Type:         marketplaces.TypeCard,
		UserID:       uuid.New(),
		Status:       marketplaces.StatusActive,
		StartPrice:   5.0,
		CurrentPrice: 25.0,
		StartTime:    time.Now().UTC(),
		EndTime:      time.Now().AddDate(0, 0, 1).UTC(),
	}

	user1 := users.User{
		ID:           uuid.New(),
		Email:        "tarkovskynik@gmail.com",
		PasswordHash: []byte{0},
		NickName:     "Nik",
		FirstName:    "Nikita",
		LastName:     "Tarkovskyi",
		LastLogin:    time.Now(),
		Status:       0,
		CreatedAt:    time.Now(),
	}

	user2 := users.User{
		ID:           uuid.New(),
		Email:        "3560876@gmail.com",
		PasswordHash: []byte{1},
		NickName:     "qwerty",
		FirstName:    "Stas",
		LastName:     "Isakov",
		LastLogin:    time.Now(),
		Status:       1,
		CreatedAt:    time.Now(),
	}

	dbtesting.Run(t, func(ctx context.Context, t *testing.T, db ultimatedivision.DB) {
		repositoryMarketplaces := db.Marketplaces()
		repositoryUsers := db.Users()
		id := uuid.New()
		t.Run("get sql no rows", func(t *testing.T) {
			_, err := repositoryMarketplaces.Get(ctx, id)
			require.Error(t, err)
			assert.Equal(t, true, marketplaces.ErrNoLot.Has(err))
		})

		t.Run("get", func(t *testing.T) {
			err := repositoryUsers.Create(ctx, user1)
			require.NoError(t, err)

			lot1.UserID = user1.ID
			err = repositoryMarketplaces.Create(ctx, lot1)
			require.NoError(t, err)

			lotFromDB, err := repositoryMarketplaces.Get(ctx, lot1.ID)
			require.NoError(t, err)
			compareLotForGet(t, lot1, lotFromDB)
		})

		t.Run("list active", func(t *testing.T) {
			err := repositoryUsers.Create(ctx, user2)
			require.NoError(t, err)

			lot2.UserID = user2.ID
			err = repositoryMarketplaces.Create(ctx, lot2)
			require.NoError(t, err)

			activeLots, err := repositoryMarketplaces.ListActive(ctx)
			assert.NoError(t, err)
			assert.Equal(t, len(activeLots), 1)
			compareAllFieldsLot(t, lot2, activeLots[0])
		})
	})
}

func compareAllFieldsLot(t *testing.T, lot1, lot2 marketplaces.Lot) {
	assert.Equal(t, lot1.ID, lot2.ID)
	assert.Equal(t, lot1.ItemID, lot2.ItemID)
	assert.Equal(t, lot1.Type, lot2.Type)
	assert.Equal(t, lot1.UserID, lot2.UserID)
	assert.Equal(t, lot1.ShopperID, lot2.ShopperID)
	assert.Equal(t, lot1.Status, lot2.Status)
	assert.Equal(t, lot1.StartPrice, lot2.StartPrice)
	assert.Equal(t, lot1.MaxPrice, lot2.MaxPrice)
	assert.Equal(t, lot1.CurrentPrice, lot2.CurrentPrice)
	// TODO: compare dates in a better way.
	// assert.Equal(t, lot1.StartTime, lot2.StartTime)
	// assert.Equal(t, lot1.EndTime, lot2.EndTime)
}

func compareLotForGet(t *testing.T, lot1, lot2 marketplaces.Lot) {
	assert.Equal(t, lot1.ID, lot2.ID)
	assert.Equal(t, lot1.ItemID, lot2.ItemID)
	assert.Equal(t, lot1.Type, lot2.Type)
	assert.Equal(t, lot1.Status, lot2.Status)
	assert.Equal(t, lot1.StartPrice, lot2.StartPrice)
	assert.Equal(t, lot1.MaxPrice, lot2.MaxPrice)
	assert.Equal(t, lot1.CurrentPrice, lot2.CurrentPrice)
}
