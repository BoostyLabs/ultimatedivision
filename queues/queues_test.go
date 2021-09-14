// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package queues_test

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"ultimatedivision"
	"ultimatedivision/database/dbtesting"
	"ultimatedivision/internal/pagination"
	"ultimatedivision/queues"
	"ultimatedivision/users"
)

func TestQueues(t *testing.T) {
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

	queue1 := queues.Queue{
		UserID: user1.ID,
		Status: queues.StatusSearches,
	}

	queue2 := queues.Queue{
		UserID: user2.ID,
		Status: queues.StatusGames,
	}

	cursor1 := pagination.Cursor{
		Limit: 2,
		Page:  1,
	}

	dbtesting.Run(t, func(ctx context.Context, t *testing.T, db ultimatedivision.DB) {
		repositoryQueues := db.Queues()
		repositoryUsers := db.Users()
		id := uuid.New()
		t.Run("get sql no rows", func(t *testing.T) {
			_, err := repositoryQueues.Get(ctx, id)
			require.Error(t, err)
			assert.Equal(t, true, queues.ErrNoQueue.Has(err))
		})

		t.Run("get", func(t *testing.T) {
			err := repositoryUsers.Create(ctx, user1)
			require.NoError(t, err)

			err = repositoryQueues.Create(ctx, queue1)
			require.NoError(t, err)

			queueFromDB, err := repositoryQueues.Get(ctx, user1.ID)
			require.NoError(t, err)
			compareQueues(t, queue1, queueFromDB)
		})

		t.Run("list paginated", func(t *testing.T) {
			err := repositoryUsers.Create(ctx, user2)
			require.NoError(t, err)

			err = repositoryQueues.Create(ctx, queue2)
			require.NoError(t, err)

			allQueues, err := repositoryQueues.ListPaginated(ctx, cursor1)
			assert.NoError(t, err)
			assert.Equal(t, len(allQueues.Queues), 2)
			compareQueues(t, queue1, allQueues.Queues[0])
			compareQueues(t, queue2, allQueues.Queues[1])
		})

		t.Run("update status", func(t *testing.T) {
			queue1.Status = queues.StatusGames
			err := repositoryQueues.UpdateStatus(ctx, queue1.UserID, queue1.Status)
			require.NoError(t, err)

			allQueues, err := repositoryQueues.ListPaginated(ctx, cursor1)
			assert.NoError(t, err)
			assert.Equal(t, len(allQueues.Queues), 2)
			compareQueues(t, queue1, allQueues.Queues[1])
			compareQueues(t, queue2, allQueues.Queues[0])
		})

		t.Run("delete", func(t *testing.T) {
			err := repositoryQueues.Delete(ctx, queue1.UserID)
			require.NoError(t, err)

			allQueues, err := repositoryQueues.ListPaginated(ctx, cursor1)
			assert.NoError(t, err)
			assert.Equal(t, len(allQueues.Queues), 1)
			compareQueues(t, queue2, allQueues.Queues[0])
		})
	})
}

func compareQueues(t *testing.T, queue1, queue2 queues.Queue) {
	assert.Equal(t, queue1.UserID, queue2.UserID)
	assert.Equal(t, queue1.Status, queue2.Status)
}
