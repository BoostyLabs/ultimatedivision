// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package managers_test

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"ultimatedivision"
	"ultimatedivision/clubs"
	"ultimatedivision/database/dbtesting"
	"ultimatedivision/divisions"
	"ultimatedivision/managers"
	"ultimatedivision/users"
)

func TestManager(t *testing.T) {
	user1 := users.User{
		ID:           uuid.New(),
		Email:        "tarkovskynik@gmail.com",
		PasswordHash: []byte{0},
		NickName:     "Nik",
		FirstName:    "Nikita",
		LastName:     "Tarkovskyi",
		LastLogin:    time.Now().UTC(),
		Status:       0,
		CreatedAt:    time.Now().UTC(),
	}

	division1 := divisions.Division{
		ID:             uuid.New(),
		Name:           10,
		PassingPercent: 10,
		CreatedAt:      time.Now().UTC(),
	}

	club1 := clubs.Club{
		ID:         uuid.New(),
		OwnerID:    user1.ID,
		Name:       user1.NickName,
		Status:     clubs.StatusActive,
		DivisionID: division1.ID,
		Ownership:  clubs.OwnershipOwner,
		CreatedAt:  time.Now().UTC(),
	}

	club2 := clubs.Club{
		ID:         uuid.New(),
		OwnerID:    user1.ID,
		Name:       user1.NickName,
		Status:     clubs.StatusActive,
		DivisionID: division1.ID,
		Ownership:  clubs.OwnershipOwner,
		CreatedAt:  time.Now().UTC(),
	}

	manager1 := managers.Manager{
		UserID:  user1.ID,
		ClubID:  club1.ID,
		EndedAt: time.Now().Add(5 * time.Minute).UTC(),
	}

	manager2 := managers.Manager{
		UserID:  user1.ID,
		ClubID:  club2.ID,
		EndedAt: time.Now().Add(5 * time.Minute).UTC(),
	}

	fakeManager := managers.Manager{
		UserID:  uuid.New(),
		ClubID:  uuid.New(),
		EndedAt: time.Now().Add(5 * time.Minute).UTC(),
	}

	dbtesting.Run(t, func(ctx context.Context, t *testing.T, db ultimatedivision.DB) {
		repositoryUser := db.Users()
		repositoryDivisions := db.Divisions()
		repositoryClubs := db.Clubs()
		repositoryManagers := db.Managers()

		t.Run("Create", func(t *testing.T) {
			err := repositoryUser.Create(ctx, user1)
			require.NoError(t, err)

			err = repositoryDivisions.Create(ctx, division1)
			require.NoError(t, err)

			_, err = repositoryClubs.Create(ctx, club1)
			require.NoError(t, err)

			_, err = repositoryClubs.Create(ctx, club2)
			require.NoError(t, err)

			err = repositoryManagers.Create(ctx, manager1)
			require.NoError(t, err)

			err = repositoryManagers.Create(ctx, manager2)
			require.NoError(t, err)
		})

		t.Run("List", func(t *testing.T) {
			allManagers, err := repositoryManagers.List(ctx)
			require.NoError(t, err)
			require.Equal(t, 2, len(allManagers))
			compareManagers(t, allManagers[0], manager1)
			compareManagers(t, allManagers[1], manager2)
		})

		t.Run("List by user id", func(t *testing.T) {
			allManagers, err := repositoryManagers.ListByUserID(ctx, user1.ID)
			require.NoError(t, err)
			require.Equal(t, 2, len(allManagers))
			compareManagers(t, allManagers[0], manager1)
			compareManagers(t, allManagers[1], manager2)
		})

		t.Run("Delete sql no rows", func(t *testing.T) {
			err := repositoryManagers.Delete(ctx, fakeManager.UserID, fakeManager.ClubID)
			require.Error(t, err)
			assert.Equal(t, true, managers.ErrNoManager.Has(err))
		})

		t.Run("Delete", func(t *testing.T) {
			err := repositoryManagers.Delete(ctx, manager1.UserID, manager1.ClubID)
			require.NoError(t, err)

			err = repositoryManagers.Delete(ctx, manager2.UserID, manager2.ClubID)
			require.NoError(t, err)
		})
	})
}

func compareManagers(t *testing.T, manager1, manager2 managers.Manager) {
	assert.Equal(t, manager1.UserID, manager2.UserID)
	assert.Equal(t, manager1.ClubID, manager2.ClubID)
	assert.WithinDuration(t, manager1.EndedAt, manager2.EndedAt, time.Second)
}
