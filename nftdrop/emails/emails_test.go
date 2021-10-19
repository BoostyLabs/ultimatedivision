// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package emails_test

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"ultimatedivision/nftdrop"
	"ultimatedivision/nftdrop/database/dbtesting"
	"ultimatedivision/nftdrop/emails"
)

func TestEmails(t *testing.T) {
	email1 := emails.Email{
		ID:        uuid.New(),
		Name:      "tarkovskynik@gmail.com",
		CreatedAt: time.Now(),
	}

	email2 := emails.Email{
		ID:        uuid.New(),
		Name:      "3560876@gmail.com",
		CreatedAt: time.Now(),
	}

	dbtesting.Run(t, func(ctx context.Context, t *testing.T, db nftdrop.DB) {
		repository := db.Emails()
		id := uuid.New()
		t.Run("get sql no rows", func(t *testing.T) {
			_, err := repository.Get(ctx, id)
			require.Error(t, err)
			assert.Equal(t, true, emails.ErrNoEmail.Has(err))
		})

		t.Run("get", func(t *testing.T) {
			err := repository.Create(ctx, email1)
			require.NoError(t, err)

			emailFromDB, err := repository.Get(ctx, email1.ID)
			require.NoError(t, err)
			compareEmails(t, email1, emailFromDB)
		})

		t.Run("getByEmail", func(t *testing.T) {
			emailFromDB, err := repository.GetByEmail(ctx, email1.Name)
			require.NoError(t, err)
			compareEmails(t, email1, emailFromDB)
		})

		t.Run("list", func(t *testing.T) {
			err := repository.Create(ctx, email2)
			require.NoError(t, err)

			allUsers, err := repository.List(ctx)
			assert.NoError(t, err)
			assert.Equal(t, len(allUsers), 2)
			compareEmails(t, email1, allUsers[0])
			compareEmails(t, email2, allUsers[1])
		})

		t.Run("update", func(t *testing.T) {
			err := repository.Update(ctx, email2.Name, email1.ID)
			require.NoError(t, err)

			userFromDB, err := repository.Get(ctx, email1.ID)
			require.NoError(t, err)
			assert.Equal(t, email2.Name, userFromDB.Name)
		})

		t.Run("delete", func(t *testing.T) {
			err := repository.Delete(ctx, email1.ID)
			require.NoError(t, err)
		})
	})
}

func compareEmails(t *testing.T, email1, email2 emails.Email) {
	assert.Equal(t, email1.ID, email2.ID)
	assert.Equal(t, email1.Name, email2.Name)
}
