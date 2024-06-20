// Copyright (C) 2022 Creditor Corp. Group.
// See LICENSE for copying information.

package notifications

import (
	"context"
	"math/big"
	"time"

	"github.com/google/uuid"
	"github.com/zeebo/errs"
)

var (
	// ErrNoNotification indicated that notification does not exist.
	ErrNoNotification = errs.Class("notification does not exist")
)

// Notifications is an interface with notifications methods.
type Notifications interface {
	// Send creates a notification and send to user.
	Send(ctx context.Context, userID uuid.UUID, notificationType Type, title Title,
		message MessageType, relatedObjectType RelatedObjectType, relatedObjectID uuid.UUID) (uuid.UUID, error)
	// SendAmountOfLotsChanges sends message when num of lots changes in the marketplaces.
	SendAmountOfLotsChanges(marketplaceType string)
	// ListSeasonRewardsByUser returns all season rewards notifications by user.
	ListSeasonRewardsByUser(ctx context.Context, userID uuid.UUID) ([]NotificationWithSeasonReward, error)
	// GetTotalUnreadNotifications returns number of unread notifications.
	GetTotalUnreadNotifications(ctx context.Context, userID uuid.UUID) (int, error)
	// ChangeNotificationStatus changes notification status.
	ChangeNotificationStatus(ctx context.Context, id uuid.UUID, status Status) error
	// GetSocketNotificationChan returns a channel to send new notification.
	GetSocketNotificationChan() chan *NotificationWithTotalUnread
}

// MailService exposes access to mail service.
type MailService interface {
	// SendOutbidEmail sends email when someone outbid user.
	SendOutbidEmail(email string, username string) error
	// SendWonLotEmail sends email when user wins lot.
	SendWonLotEmail(email string, username, artistName string) error
}

// DB exposes access to notifications db.
//
// architecture: DB
type DB interface {
	// Create creates a notification and writes to the database.
	Create(ctx context.Context, notification Notification) error
	// ListSeasonRewardsByUser returns all notifications from the database by user.
	ListSeasonRewardsByUser(ctx context.Context, userID uuid.UUID) ([]NotificationWithSeasonReward, error)
	// GetTotalUnreadNotifications returns number of unread notifications from the database.
	GetTotalUnreadNotifications(ctx context.Context, userID uuid.UUID) (int, error)
	// ChangeNotificationStatus changes notification status.
	ChangeNotificationStatus(ctx context.Context, id uuid.UUID, status Status) error
}

// ChannelSize indicates notification channel size.
const ChannelSize = 1000

// TimeFormat indicates notification created at format for filtering by date.
const TimeFormat = "2006/01/02"

// Status defines the list of possible notification statuses.
type Status string

const (
	// StatusRead indicates that notification is read.
	StatusRead Status = "read"
	// StatusNotRead indicates that notification not read.
	StatusNotRead Status = "not read"
	// StatusDeleted indicates that notification is deleted.
	StatusDeleted Status = "deleted"
)

// IsValid checks the notifications' status for all conditions.
func (s Status) IsValid() bool {
	switch s {
	case StatusRead,
		StatusNotRead,
		StatusDeleted:
		return true
	default:
		return false
	}
}

// RelatedObjectType defines the list of possible notification related object type.
type RelatedObjectType string

const (
	// Bid indicates that notification related object type 'bid'.
	Bid RelatedObjectType = "bid"
	// Outbid indicates that notification related object type 'outbid'.
	Outbid RelatedObjectType = "outbid"
	// Token indicates that notification related object type 'token'.
	Token RelatedObjectType = "token"
	// Mint indicates that notification related object type 'mint'.
	Mint RelatedObjectType = "mint"
)

// Title defines the list of possible notification titles.
type Title string

const (
	// TitleSeasonReward season reward.
	TitleSeasonReward Title = "SeasonRewards"
)

// MessageType defines the list of possible notification messages.
type MessageType string

const (
	// @todo place this messages to notifications.Config.

	// MessageTypeSeasonRewards Something.
	MessageTypeSeasonRewards MessageType = "Something"
)

// SelectionType defines the list of possible notification selection types.
type SelectionType string

const (
	// SelectionTypeWeek indicates that selection of notifications by week.
	SelectionTypeWeek SelectionType = "week"
	// SelectionTypeMonth indicates that selection of notifications by month.
	SelectionTypeMonth SelectionType = "month"
	// SelectionTypeAll indicates that selection of notifications all time.
	SelectionTypeAll SelectionType = ""
)

// IsValid checks the notifications' selection type for all conditions.
func (s SelectionType) IsValid() bool {
	return s == SelectionTypeMonth || s == SelectionTypeWeek || s == SelectionTypeAll
}

// Weekday defines the list of possible weekdays from filters by week.
type Weekday int

const (
	// Monday indicates that weekday is monday.
	Monday Weekday = 1
	// Tuesday indicates that weekday is tuesday.
	Tuesday Weekday = 2
	// Wednesday indicates that weekday is wednesday.
	Wednesday Weekday = 3
	// Thursday indicates that weekday is thursday.
	Thursday Weekday = 4
	// Friday indicates that weekday is friday.
	Friday Weekday = 5
	// Saturday indicates that weekday is saturday.
	Saturday Weekday = 6
	// Sunday indicates that weekday is sunday.
	Sunday Weekday = 7
)

// Int returns weekday in type int.
func (w Weekday) Int() int {
	return int(w)
}

// Type defines the list of possible notification types.
type Type string

const (
	// TypeInfo indicates that notification info type.
	TypeInfo Type = "info"
	// TypeWarning indicates that notification warning type.
	TypeWarning Type = "warning"
	// TypeSystem indicates that notification system type.
	TypeSystem Type = "system"
)

// Notification describes notification entity.
type Notification struct {
	ID                uuid.UUID         `json:"id"`
	UserID            uuid.UUID         `json:"userId"`
	Status            Status            `json:"status"`
	RelatedObjectType RelatedObjectType `json:"relatedObjectType"`
	RelatedObjectID   uuid.UUID         `json:"relatedObjectId"`
	Type              Type              `json:"type"`
	Title             Title             `json:"title"`
	Message           MessageType       `json:"message"`
	CreatedAt         time.Time         `json:"createdAt"`
	UpdatedAt         time.Time         `json:"updatedAt"`
}

// NotificationWithSeasonReward describes notification with season rewards.
type NotificationWithSeasonReward struct {
	Notification
	Reward big.Int `json:"reward"`
}

// NotificationWithTotalUnread describes notification entity with total unread number.
type NotificationWithTotalUnread struct {
	Notification
	TotalUnreadNotifications int `json:"totalUnreadNotifications"`
}
