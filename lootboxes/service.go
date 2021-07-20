// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package lootboxes

// Service is handling lootboxes related logic.
//
// architecture: Service
type Service struct {
	lootboxes DB
}

// NewService is a constructor for lootboxes service.
func NewService(lootboxes DB) *Service {
	return &Service{
		lootboxes: lootboxes,
	}
}