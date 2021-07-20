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
	"ultimatedivision/cards"
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

	RegularBox := lootboxes.LootBox{
		ID:       uuid.New(),
		Name:     "Regular Box",
		Cost:     1000,
		CardsNum: 5,
		Wood:     0.8,
		Silver:   0.15,
		Gold:     0.04,
		Diamond:  0.01,
	}

	userLootBox := lootboxes.UserLootBox{
		ID:        uuid.New(),
		UserID:    user1.ID,
		LootBoxID: RegularBox.ID,
	}

	card1 := cards.Card{
		ID:               uuid.New(),
		PlayerName:       "test",
		Quality:          "bronze",
		PictureType:      1,
		Height:           178.8,
		Weight:           72.2,
		SkinColor:        1,
		HairStyle:        1,
		HairColor:        1,
		Accessories:      []int{1, 2},
		DominantFoot:     "left",
		UserID:           user1.ID,
		Tactics:          1,
		Positioning:      2,
		Composure:        3,
		Aggression:       4,
		Vision:           5,
		Awareness:        6,
		Crosses:          7,
		Physique:         8,
		Acceleration:     9,
		RunningSpeed:     10,
		ReactionSpeed:    11,
		Agility:          12,
		Stamina:          13,
		Strength:         14,
		Jumping:          15,
		Balance:          16,
		Technique:        17,
		Dribbling:        18,
		BallControl:      19,
		WeakFoot:         20,
		SkillMoves:       21,
		Finesse:          22,
		Curve:            23,
		Volleys:          24,
		ShortPassing:     25,
		LongPassing:      26,
		ForwardPass:      27,
		Offense:          28,
		FinishingAbility: 29,
		ShotPower:        30,
		Accuracy:         31,
		Distance:         32,
		Penalty:          33,
		FreeKicks:        34,
		Corners:          35,
		HeadingAccuracy:  36,
		Defence:          37,
		OffsideTrap:      38,
		Sliding:          39,
		Tackles:          40,
		BallFocus:        41,
		Interceptions:    42,
		Vigilance:        43,
		Goalkeeping:      44,
		Reflexes:         45,
		Diving:           46,
		Handling:         47,
		Sweeping:         48,
		Throwing:         49,
	}

	userLoot := lootboxes.UserLoot{
		ID:     userLootBox.ID,
		CardID: card1.ID,
	}

	dbtesting.Run(t, func(ctx context.Context, t *testing.T, db ultimatedivision.DB) {
		repositoryCard := db.Cards()
		repositoryUsers := db.Users()
		repositoryLootBoxes := db.LootBoxes()

		t.Run("Create", func(t *testing.T) {
			err := repositoryUsers.Create(ctx, user1)
			require.NoError(t, err)

			err = repositoryLootBoxes.Create(ctx, userLootBox)
			require.NoError(t, err)
		})

		t.Run("Create loot", func(t *testing.T) {
			err := repositoryCard.Create(ctx, card1)
			require.NoError(t, err)

			err = repositoryLootBoxes.CreateCards(ctx, userLoot)
			require.NoError(t, err)
		})

		t.Run("Get id of opened loot boxes", func(t *testing.T) {
			id, err := repositoryLootBoxes.Get(ctx, user1.ID)
			require.NoError(t, err)

			assert.Equal(t, id, []uuid.UUID{userLootBox.ID})
		})

		t.Run("Get id of all cards form loot box", func(t *testing.T) {
			id, err := repositoryLootBoxes.GetCards(ctx, userLootBox.ID)
			require.NoError(t, err)

			assert.Equal(t, id, []uuid.UUID{card1.ID})
		})
	})
}
