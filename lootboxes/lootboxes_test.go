// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package lootboxes_test

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"ultimatedivision"
	"ultimatedivision/database/dbtesting"
	"ultimatedivision/lootboxes"
	"ultimatedivision/users"
)

func TestLootBox(t *testing.T) {
	user1 := users.User{
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

	userLootBox1 := lootboxes.LootBox{
		UserID:    user1.ID,
		LootBoxID: uuid.New(),
		Type:      lootboxes.RegularBox,
	}

	userLootBox2 := lootboxes.LootBox{
		UserID:    user1.ID,
		LootBoxID: uuid.New(),
		Type:      lootboxes.UDReleaseCelebrationBox,
	}

	dbtesting.Run(t, func(ctx context.Context, t *testing.T, db ultimatedivision.DB) {
		repositoryUsers := db.Users()
		repositoryLootBoxes := db.LootBoxes()

		t.Run("Create", func(t *testing.T) {
			err := repositoryUsers.Create(ctx, user1)
			require.NoError(t, err)

			err = repositoryLootBoxes.Create(ctx, userLootBox1)
			require.NoError(t, err)
		})

		t.Run("List", func(t *testing.T) {
			err := repositoryLootBoxes.Create(ctx, userLootBox2)
			require.NoError(t, err)

			userLootBoxes, err := repositoryLootBoxes.List(ctx)
			require.NoError(t, err)

			compareLootBoxes(t, userLootBoxes, []lootboxes.LootBox{userLootBox1, userLootBox2})
		})

		t.Run("Get loot box type", func(t *testing.T) {
			lootBoxType, err := repositoryLootBoxes.GetTypeByLootBoxID(ctx, userLootBox1.LootBoxID)
			require.NoError(t, err)

			assert.Equal(t, lootBoxType, userLootBox1.Type)
		})

		t.Run("Delete", func(t *testing.T) {
			err := repositoryLootBoxes.Delete(ctx, userLootBox1.LootBoxID)
			require.NoError(t, err)
		})
	})
}

func compareLootBoxes(t *testing.T, userLootBoxesDB, userLootBoxesTest []lootboxes.LootBox) {
	assert.Equal(t, len(userLootBoxesDB), len(userLootBoxesTest))

	for i := 0; i < len(userLootBoxesDB); i++ {
		assert.Equal(t, userLootBoxesDB[i].UserID, userLootBoxesTest[i].UserID)
		assert.Equal(t, userLootBoxesDB[i].LootBoxID, userLootBoxesTest[i].LootBoxID)
		assert.Equal(t, userLootBoxesDB[i].Type, userLootBoxesTest[i].Type)
	}
}
