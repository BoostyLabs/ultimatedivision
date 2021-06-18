package admins

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"

	"github.com/google/uuid"

	"ultimatedivision"
	"ultimatedivision/database/dbtesting"
)

func TestAdmin(t *testing.T) {
	admin1 := Admin{
		ID:           uuid.New(),
		Email:        "admin1@gmail.com",
		PasswordHash: []byte{1},
		CreatedAt:    time.Now(),
	}

	admin2 := Admin{
		ID:           uuid.New(),
		Email:        "admin2@gmail.com",
		PasswordHash: []byte{1},
		CreatedAt:    time.Now(),
	}

	dbtesting.Run(t, func(ctx context.Context, t *testing.T, db ultimatedivision.DB) {
		repository := db.Admins()
		id := uuid.New()
		t.Run("Get sql no rows", func(t *testing.T) {
			_, err := repository.Get(ctx, id)
			require.Error(t, err)
			assert.Equal(t, true, ErrNoAdmin.Has(err))
		})
		t.Run("Get", func(t *testing.T) {
			err := repository.Create(ctx, admin1)
			require.NoError(t, err)

			adminFromDB, err := repository.Get(ctx, admin1.ID)
			require.NoError(t, err)
			compareAdmins(t,adminFromDB,admin1)
		})
		t.Run("List", func(t *testing.T) {
			err := repository.Create(ctx,admin2)
			require.NoError(t, err)

			allAdmins,err := repository.List(ctx)
			require.NoError(t, err)
			compareAdmins(t,allAdmins[0],admin1)
			compareAdmins(t,allAdmins[1],admin2)
		})
	})
}

func compareAdmins(t *testing.T, adminFromDB Admin, testAdmin Admin) {
	assert.Equal(t, adminFromDB.ID, testAdmin.ID)
	assert.Equal(t, adminFromDB.Email, testAdmin.Email)
	assert.Equal(t, adminFromDB.PasswordHash, testAdmin.PasswordHash)
	assert.Equal(t, adminFromDB.CreatedAt, testAdmin.CreatedAt)
}
