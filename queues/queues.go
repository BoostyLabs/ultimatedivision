// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package queues

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/zeebo/errs"

	"ultimatedivision/internal/pagination"
)

// ErrNoQueue indicated that queue does not exist.
var ErrNoQueue = errs.Class("queue does not exist")

// ErrQueues indicated that there was an error in service.
var ErrQueues = errs.Class("queues service error")

// DB is exposing access to queues database.
//
// architecture: DB
type DB interface {
	// Create adds queue in database.
	Create(ctx context.Context, queue Queue) error
	// Get returns queue from database.
	Get(ctx context.Context, id uuid.UUID) (Queue, error)
	// ListPaginated returns page of queues from database.
	ListPaginated(ctx context.Context, cursor pagination.Cursor) (Page, error)
	// UpdateStatus updates status queue in database.
	UpdateStatus(ctx context.Context, id uuid.UUID, status Status) error
	// Delete deletes queue record in database.
	Delete(ctx context.Context, id uuid.UUID) error
}

// Queue describes queue entity.
type Queue struct {
	UserID uuid.UUID `json:"userId"`
	Status Status    `json:"status"`
}

// Status defines list of possible queue statuses.
type Status string

const (
	// StatusSearches indicates that user in queue searches game.
	StatusSearches Status = "searches"
	// StatusGames indicates that in queue plays game.
	StatusGames Status = "games"
)

// Config defines configuration for queues.
type Config struct {
	QueueRenewalInterval time.Duration     `json:"queueRenewalInterval"`
	Cursor               pagination.Cursor `json:"cursor"`
}

// Page holds queue page entity which is used to show listed page of queues.
type Page struct {
	Queues []Queue         `json:"queues"`
	Page   pagination.Page `json:"page"`
}
