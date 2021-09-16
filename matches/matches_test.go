// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package matches_test

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
	"ultimatedivision/matches"
	"ultimatedivision/users"
)

func TestMatches(t *testing.T) {
	testUser1 := users.User{
		ID:           uuid.New(),
		Email:        "test@gmail.com",
		PasswordHash: []byte{1},
		NickName:     "testNickName",
		FirstName:    "test",
		LastName:     "test",
		LastLogin:    time.Now(),
		Status:       1,
		CreatedAt:    time.Now(),
	}

	testUser2 := users.User{
		ID:           uuid.New(),
		Email:        "test@gmail.com",
		PasswordHash: []byte{2},
		NickName:     "testNickName",
		FirstName:    "test",
		LastName:     "test",
		LastLogin:    time.Now(),
		Status:       1,
		CreatedAt:    time.Now(),
	}

	testCard := cards.Card{
		ID:     uuid.New(),
		UserID: testUser1.ID,
	}

	testMatch := matches.Match{
		ID:      uuid.New(),
		User1ID: testUser1.ID,
		User2ID: testUser2.ID,
		Score:   "0:0",
	}

	updatedTestMatch := matches.Match{
		ID:      testMatch.ID,
		User1ID: testUser1.ID,
		User2ID: testUser2.ID,
		Score:   "0:1",
	}

	testMatchGoal := matches.MatchGoals{
		ID:      uuid.New(),
		MatchID: updatedTestMatch.ID,
		UserID:  testUser1.ID,
		CardID:  testCard.ID,
		Minute:  25,
	}

	dbtesting.Run(t, func(ctx context.Context, t *testing.T, db ultimatedivision.DB) {
		repositoryCards := db.Cards()
		repositoryUsers := db.Users()
		repositoryMatches := db.Matches()

		t.Run("Create", func(t *testing.T) {
			err := repositoryUsers.Create(ctx, testUser1)
			require.NoError(t, err)

			err = repositoryUsers.Create(ctx, testUser2)
			require.NoError(t, err)

			err = repositoryMatches.Create(ctx, testMatch)
			require.NoError(t, err)
		})

		t.Run("Update", func(t *testing.T) {
			err := repositoryMatches.Update(ctx, testMatch.ID, "0:1")
			require.NoError(t, err)
		})

		t.Run("List matches", func(t *testing.T) {
			allMatchesDB, err := repositoryMatches.ListMatches(ctx)
			require.NoError(t, err)
			compareMatchesSlice(t, allMatchesDB, []matches.Match{updatedTestMatch})
		})

		t.Run("Get", func(t *testing.T) {
			matchDB, err := repositoryMatches.Get(ctx, updatedTestMatch.ID)
			require.NoError(t, err)
			compareMatches(t, matchDB, updatedTestMatch)
		})

		t.Run("Add goal in the match", func(t *testing.T) {
			err := repositoryCards.Create(ctx, testCard)
			require.NoError(t, err)

			err = repositoryMatches.AddGoal(ctx, testMatchGoal)
			require.NoError(t, err)
		})

		t.Run("List match goals", func(t *testing.T) {
			matchGoalsDB, err := repositoryMatches.ListMatchGoals(ctx, updatedTestMatch.ID)
			require.NoError(t, err)
			compareMatchGoals(t, matchGoalsDB, []matches.MatchGoals{testMatchGoal})
		})

		t.Run("delete", func(t *testing.T) {
			err := repositoryMatches.Delete(ctx, updatedTestMatch.ID)
			require.NoError(t, err)
		})
	})
}

func compareMatches(t *testing.T, matchDB, matchTest matches.Match) {
	assert.Equal(t, matchDB.ID, matchTest.ID)
	assert.Equal(t, matchDB.User1ID, matchTest.User1ID)
	assert.Equal(t, matchDB.User2ID, matchTest.User2ID)
	assert.Equal(t, matchDB.Score, matchTest.Score)
}

func compareMatchesSlice(t *testing.T, matchesDB, matchesTest []matches.Match) {
	assert.Equal(t, len(matchesDB), len(matchesTest))

	for i := 0; i < len(matchesDB); i++ {
		assert.Equal(t, matchesDB[i].ID, matchesTest[i].ID)
		assert.Equal(t, matchesDB[i].User1ID, matchesTest[i].User1ID)
		assert.Equal(t, matchesDB[i].User2ID, matchesTest[i].User2ID)
		assert.Equal(t, matchesDB[i].Score, matchesTest[i].Score)
	}
}

func compareMatchGoals(t *testing.T, matchGoalsDB, matchGoalsTest []matches.MatchGoals) {
	assert.Equal(t, len(matchGoalsDB), len(matchGoalsTest))

	for i := 0; i < len(matchGoalsDB); i++ {
		assert.Equal(t, matchGoalsDB[i].ID, matchGoalsTest[i].ID)
		assert.Equal(t, matchGoalsDB[i].MatchID, matchGoalsTest[i].MatchID)
		assert.Equal(t, matchGoalsDB[i].UserID, matchGoalsTest[i].UserID)
		assert.Equal(t, matchGoalsDB[i].CardID, matchGoalsTest[i].CardID)
		assert.Equal(t, matchGoalsDB[i].Minute, matchGoalsTest[i].Minute)
	}
}
