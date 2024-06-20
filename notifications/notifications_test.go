// Copyright (C) 2022 Creditor Corp. Group.
// See LICENSE for copying information.

package notifications_test

import (
	"context"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"ultimatedivision"
	"ultimatedivision/database/dbtesting"
	"ultimatedivision/divisions"
	"ultimatedivision/notifications"
	"ultimatedivision/seasons"
	"ultimatedivision/users"
)

func TestNotifications(t *testing.T) {
	user1ID := uuid.New()
	notification1ID := uuid.New()
	user1 := users.User{
		ID:               user1ID,
		Email:            "oleksii@gmail.com",
		PasswordHash:     []byte{0},
		NickName:         "Free",
		FirstName:        "Oleksii",
		LastName:         "Prysiazhniuk",
		Wallet:           common.HexToAddress("0xb2cdC7EB2F9d2E629ee97BB91700622A42e688b7"),
		CasperWallet:     "01a4db357602c3d45a2b7b68110e66440ac2a2e792cebffbce83eaefb73e65aef1",
		CasperWalletHash: "4bfcd0ebd44c3de9d1e6556336cbb73259649b7d6b344bc1499d40652fd5781a",
		WalletType:       users.WalletTypeETH,
		LastLogin:        time.Now().UTC(),
		Status:           0,
		CreatedAt:        time.Now().UTC(),
	}

	user2 := users.User{
		ID:               uuid.New(),
		Email:            "oleksii2@gmail.com",
		PasswordHash:     []byte{0},
		NickName:         "Free2",
		FirstName:        "Oleksii2",
		LastName:         "Prysiazhniuk2",
		Wallet:           common.HexToAddress("0xb2cdC7EB2F9d2E629ee97BB91700622A42e688b7"),
		CasperWallet:     "01a4db357602c3d45a2b7b68110e66440ac2a2e792cebffbce83eaefb73e65aef1",
		CasperWalletHash: "4bfcd0ebd44c3de9d1e6556336cbb73259649b7d6b344bc1499d40652fd5781a",
		WalletType:       users.WalletTypeETH,
		LastLogin:        time.Now().UTC(),
		Status:           0,
		CreatedAt:        time.Now().UTC(),
	}

	division1 := divisions.Division{
		ID:             uuid.New(),
		Name:           10,
		PassingPercent: 10,
		CreatedAt:      time.Now().UTC(),
	}

	season1 := seasons.Season{
		ID:         1,
		DivisionID: division1.ID,
		StartedAt:  time.Now().UTC(),
		EndedAt:    time.Time{},
	}
	value := *big.NewInt(100)

	reward1 := seasons.Reward{
		ID:                  uuid.New(),
		UserID:              user1ID,
		SeasonID:            season1.ID,
		WalletAddress:       common.Address{},
		CasperWalletAddress: user1.CasperWallet,
		WalletType:          user1.WalletType,
		Value:               value,
		Status:              1,
	}

	reward2 := seasons.Reward{
		ID:                  uuid.New(),
		UserID:              uuid.New(),
		SeasonID:            season1.ID,
		WalletAddress:       common.Address{},
		CasperWalletAddress: user1.CasperWallet,
		WalletType:          user1.WalletType,
		Value:               value,
		Status:              1,
	}

	notification1 := notifications.Notification{
		ID:                notification1ID,
		UserID:            user1ID,
		Status:            notifications.StatusNotRead,
		RelatedObjectType: notifications.Token,
		RelatedObjectID:   reward1.ID,
		Type:              notifications.TypeInfo,
		Title:             notifications.TitleSeasonReward,
		Message:           notifications.MessageTypeSeasonRewards,
		CreatedAt:         time.Now().UTC(),
		UpdatedAt:         time.Time{},
	}

	notification2 := notifications.Notification{
		ID:                uuid.New(),
		UserID:            user1ID,
		Status:            notifications.StatusNotRead,
		RelatedObjectType: notifications.Token,
		RelatedObjectID:   reward2.ID,
		Type:              notifications.TypeInfo,
		Title:             notifications.TitleSeasonReward,
		Message:           notifications.MessageTypeSeasonRewards,
		CreatedAt:         time.Now().UTC(),
		UpdatedAt:         time.Time{},
	}

	notification3 := notifications.Notification{
		ID:                uuid.New(),
		UserID:            user2.ID,
		Status:            notifications.StatusNotRead,
		RelatedObjectType: notifications.Bid,
		RelatedObjectID:   reward2.ID,
		Type:              notifications.TypeInfo,
		Title:             notifications.TitleSeasonReward,
		Message:           notifications.MessageTypeSeasonRewards,
		CreatedAt:         time.Now().UTC(),
		UpdatedAt:         time.Time{},
	}

	dbtesting.Run(t, func(ctx context.Context, t *testing.T, db ultimatedivision.DB) {
		repositoryNotifications := db.Notifications()
		repositoryUsers := db.Users()
		repositorySeasonReward := db.Seasons()

		t.Run("create", func(t *testing.T) {
			err := repositoryUsers.Create(ctx, user1)
			require.NoError(t, err)
			err = repositoryUsers.Create(ctx, user2)
			require.NoError(t, err)
		})

		t.Run("Create Reward", func(t *testing.T) {
			err := repositorySeasonReward.CreateReward(ctx, reward1)
			require.NoError(t, err)
			err = repositorySeasonReward.CreateReward(ctx, reward2)
			require.NoError(t, err)
		})

		t.Run("Create notification", func(t *testing.T) {
			err := repositoryNotifications.Create(ctx, notification1)
			require.NoError(t, err)
			err = repositoryNotifications.Create(ctx, notification2)
			require.NoError(t, err)
			err = repositoryNotifications.Create(ctx, notification3)
			require.NoError(t, err)
		})

		t.Run("List Season Rewards By User", func(t *testing.T) {
			allNotifications, err := repositoryNotifications.ListSeasonRewardsByUser(ctx, notification1.UserID)
			assert.NoError(t, err)

			require.Equal(t, 2, len(allNotifications))
			require.Equal(t, reward1.Value, allNotifications[0].Reward)
		})

		t.Run("get total unread notifications", func(t *testing.T) {
			allNotifications, err := repositoryNotifications.ListSeasonRewardsByUser(ctx, notification1.UserID)
			require.NoError(t, err)
			total, err := repositoryNotifications.GetTotalUnreadNotifications(ctx, notification1.UserID)
			require.NoError(t, err)
			assert.Equal(t, len(allNotifications), total)
		})

		t.Run("update status", func(t *testing.T) {
			err := repositoryNotifications.ChangeNotificationStatus(ctx, notification2.ID, notifications.StatusRead)
			require.NoError(t, err)

			notificationsFromDB, err := repositoryNotifications.ListSeasonRewardsByUser(ctx, notification2.UserID)
			require.NoError(t, err)
			assert.Equal(t, notifications.StatusRead, notificationsFromDB[0].Status)
		})
	})
}
