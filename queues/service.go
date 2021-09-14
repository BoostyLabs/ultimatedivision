// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package queues

import (
	"context"

	"github.com/google/uuid"

	"ultimatedivision/internal/pagination"
)

// Service is handling queues related logic.
//
// architecture: Service
type Service struct {
	config Config
	queues DB
}

// NewService is a constructor for queues service.
func NewService(config Config, queues DB) *Service {
	return &Service{
		config: config,
		queues: queues,
	}
}

// Create adds queue in database.
func (service *Service) Create(ctx context.Context, queue Queue) error {
	return ErrQueues.Wrap(service.queues.Create(ctx, queue))
}

// Get returns queue from database.
func (service *Service) Get(ctx context.Context, id uuid.UUID) (Queue, error) {
	queue, err := service.queues.Get(ctx, id)
	return queue, ErrQueues.Wrap(err)
}

// ListPaginated returns queues in page from database.
func (service *Service) ListPaginated(ctx context.Context, cursor pagination.Cursor) (Page, error) {
	if cursor.Limit <= 0 {
		cursor.Limit = service.config.Cursor.Limit
	}
	if cursor.Page <= 0 {
		cursor.Page = service.config.Cursor.Page
	}

	queuesListPage, err := service.queues.ListPaginated(ctx, cursor)
	return queuesListPage, ErrQueues.Wrap(err)
}

// UpdateStatus updates queue status in database.
func (service *Service) UpdateStatus(ctx context.Context, id uuid.UUID, status Status) error {
	return ErrQueues.Wrap(service.queues.UpdateStatus(ctx, id, status))
}

// Delete deletes queue record in database.
func (service *Service) Delete(ctx context.Context, id uuid.UUID) error {
	return ErrQueues.Wrap(service.queues.Delete(ctx, id))
}
