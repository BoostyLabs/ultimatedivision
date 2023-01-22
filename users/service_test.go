// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package users_test

import (
	"context"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"ultimatedivision"
	"ultimatedivision/internal/servicetesting"
	"ultimatedivision/users"
)

func TestUsersService(t *testing.T) {
	password := "12345"
	user1 := users.User{
		ID:               uuid.New(),
		Email:            "1234@gmail.com",
		PasswordHash:     []byte(password),
		NickName:         "qwerty",
		FirstName:        "Oleksii",
		LastName:         "Prysiazhniuk",
		Wallet:           common.HexToAddress("0xb2cdC7EB2F9d2E629ee97BB91700622A42e688b8"),
		CasperWallet:     "01a4db357602c3d45a2b7b68110e66440ac2a2e792cebffbce83eaefb73e65aef1",
		CasperWalletHash: "4bfcd0ebd44c3de9d1e6556336cbb73259649b7d6b344bc1499d40652fd5781a",
		WalletType:       users.WalletTypeVelas,
		LastLogin:        time.Now().UTC(),
		Status:           1,
		CreatedAt:        time.Now().UTC(),
	}

	userProfile := users.Profile{
		ID:        user1.ID,
		Email:     "",
		NickName:  "",
		CreatedAt: user1.CreatedAt,
		LastLogin: time.Time{},
		Wallet:    user1.Wallet,
	}

	profileWithWallet := users.ProfileWithWallet{
		ID:             uuid.UUID{},
		Email:          "",
		NickName:       "",
		CreatedAt:      time.Time{},
		LastLogin:      time.Time{},
		Wallet:         common.Address{},
		CasperWallet:   "",
		CasperWalletID: "",
		WalletType:     "",
	}

	servicetesting.Run(t, func() ultimatedivision.Config { return ultimatedivision.Config{} }, func(ctx context.Context, t *testing.T, ultimatedivision *ultimatedivision.Peer) {
		service := ultimatedivision.Users.Service

		t.Run("negative get", func(t *testing.T) {
			_, err := service.Get(ctx, uuid.New())
			require.Error(t, err)
			assert.True(t, users.ErrNoUser.Has(err))
		})

		t.Run("create", func(t *testing.T) {
			err := service.Create(ctx, user1.Email, password, user1.NickName, user1.FirstName, user1.LastName)
			require.NoError(t, err)
		})

		t.Run("get", func(t *testing.T) {
			actualUser, err := service.Get(ctx, user1.ID)
			require.NoError(t, err)
			compareUsers(t, user1, actualUser)
		})

		t.Run("negative get profile", func(t *testing.T) {
			_, err := service.GetProfile(ctx, uuid.New())
			require.Error(t, err)
			assert.True(t, users.ErrNoUser.Has(err))
		})

		t.Run("get profile", func(t *testing.T) {
			profile, err := service.GetProfile(ctx, userProfile.ID)
			require.NoError(t, err)

			compareProfile(t, profileWithWallet, *profile)
		})

	})
}

func compareProfile(t *testing.T, expected, actual users.ProfileWithWallet) {
	assert.Equal(t, expected.ID, actual.ID)
	assert.Equal(t, expected.Email, actual.Email)
	assert.Equal(t, expected.NickName, actual.NickName)
	assert.Equal(t, expected.Wallet, actual.Wallet)
	assert.Equal(t, expected.CasperWallet, actual.CasperWallet)
	assert.Equal(t, expected.CasperWalletID, actual.CasperWalletID)
	assert.Equal(t, expected.WalletType, actual.WalletType)
	assert.Equal(t, expected.CreatedAt, actual.CreatedAt)
	assert.Equal(t, expected.LastLogin, actual.LastLogin)
}
