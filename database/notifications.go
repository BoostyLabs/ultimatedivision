// Copyright (C) 2022 Creditor Corp. Group.
// See LICENSE for copying information.

package database

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/zeebo/errs"

	"ultimatedivision/notifications"
)

// ErrNotifications indicates that there was an error in the database.
var ErrNotifications = errs.Class("notifications repository error")

// notificationsDB provides access to notifications db.
//
// architecture: Database
type notificationsDB struct {
	conn *sql.DB
}

// Create creates a notification and writes to the database.
func (notificationsDB *notificationsDB) Create(ctx context.Context, notification notifications.Notification) error {
	query := `INSERT INTO notifications(id, user_id, status, related_object_type, 
	                      related_object_id, type, title, message, created_at, updated_at)
	          VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`
	_, err := notificationsDB.conn.ExecContext(ctx, query, notification.ID, notification.UserID, notification.Status, notification.RelatedObjectType,
		notification.RelatedObjectID, notification.Type, notification.Title, notification.Message, notification.CreatedAt, notification.UpdatedAt)

	return ErrNotifications.Wrap(err)
}

// ListSeasonRewardsByUser returns all season notifications from the database by user.
func (notificationsDB *notificationsDB) ListSeasonRewardsByUser(ctx context.Context, userID uuid.UUID) (_ []notifications.NotificationWithSeasonReward, err error) {
	query := `SELECT DISTINCT n.id, n.user_id, n.status, n.related_object_type, n.related_object_id, n.type, n.title, n.message,
       n.created_at, n.updated_at, sr.value
              FROM notifications AS n
              LEFT JOIN season_rewards AS sr ON sr.user_id = n.user_id
              WHERE n.user_id = $1
			  ORDER BY created_at DESC`

	rows, err := notificationsDB.conn.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, ErrNotifications.Wrap(err)
	}
	defer func() {
		err = errs.Combine(err, rows.Close())
	}()

	var notificationsData []notifications.NotificationWithSeasonReward
	for rows.Next() {
		var rewards []uint8
		var notification notifications.NotificationWithSeasonReward
		err = rows.Scan(&notification.ID, &notification.UserID, &notification.Status, &notification.RelatedObjectType,
			&notification.RelatedObjectID, &notification.Type, &notification.Title, &notification.Message, &notification.CreatedAt,
			&notification.UpdatedAt, &rewards)
		if err != nil {
			return nil, ErrNotifications.Wrap(err)
		}

		notification.Reward.SetBytes(rewards)
		notificationsData = append(notificationsData, notification)
	}
	if err = rows.Err(); err != nil {
		return nil, ErrNotifications.Wrap(err)
	}

	return notificationsData, nil
}

// GetTotalUnreadNotifications returns number of unread notifications from the data base.
func (notificationsDB *notificationsDB) GetTotalUnreadNotifications(ctx context.Context, userID uuid.UUID) (int, error) {
	var count int
	query := "SELECT count(*) as total FROM notifications WHERE user_id = $1 AND status = $2"

	err := notificationsDB.conn.QueryRowContext(ctx, query, userID, notifications.StatusNotRead).Scan(&count)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, ErrNotifications.New("notification does not exist")
		}
		return 0, ErrNotifications.Wrap(err)
	}

	return count, nil
}

// ChangeNotificationStatus changes notification status.
func (notificationsDB *notificationsDB) ChangeNotificationStatus(ctx context.Context, id uuid.UUID, status notifications.Status) error {
	query := "UPDATE notifications SET status = $1, updated_at = $2 WHERE id = $3"

	result, err := notificationsDB.conn.ExecContext(ctx, query, status, time.Now().UTC(), id)
	if err != nil {
		return ErrNotifications.Wrap(err)
	}

	rowNum, err := result.RowsAffected()
	if rowNum == 0 {
		return notifications.ErrNoNotification.New("notification does not exist")
	}

	return ErrNotifications.Wrap(err)
}
