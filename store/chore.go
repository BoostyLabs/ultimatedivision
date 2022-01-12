// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package store

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/zeebo/errs"

	"ultimatedivision/cards"
	"ultimatedivision/pkg/sync"
)

var (
	// ChoreError represents store chore error type.
	ChoreError = errs.Class("expiration store chore error")
)

// Chore requests access token for contis api calls, re-requests it after token's expiration time.
//
// architecture: Chore
type Chore struct {
	config Config
	Loop   *sync.Cycle
	store  *Service
	cards  *cards.Service
}

// NewChore instantiates Chore.
func NewChore(config Config, store *Service, cards *cards.Service) *Chore {
	return &Chore{
		config: config,
		Loop:   sync.NewCycle(config.StoreRenewalInterval),
		store:  store,
		cards:  cards,
	}
}

// Run runs the renewal of cards in store.
func (chore *Chore) Run(ctx context.Context) error {
	if _, err := chore.store.Get(ctx, 1); err != nil {
		if !ErrNoSetting.Has(err) {
			return ChoreError.Wrap(err)
		}

		setting := Setting{
			ID:          1,
			CardsAmount: 10,
			IsRenewal:   true,
			HourRenewal: 0,
		}
		if err = chore.store.Create(ctx, setting); err != nil {
			return ChoreError.Wrap(err)
		}
	}

	t := time.Now().UTC()
	return chore.Loop.Run(ctx, func(ctx context.Context) error {
		setting, err := chore.store.Get(ctx, 1)
		if err != nil {
			return ChoreError.Wrap(err)
		}

		timeRenewal := time.Date(t.Year(), t.Month(), t.Day(), setting.HourRenewal, 0, 0, 0, time.UTC)
		duration := time.Until(timeRenewal)
		if duration > 0 {
			return nil
		}
		t = t.Add(24 * time.Hour)

		if !setting.IsRenewal {
			return nil
		}

		cardsList, err := chore.cards.ListByTypeOrdered(ctx)
		if err != nil {
			return ChoreError.Wrap(err)
		}

		cardsAmount := setting.CardsAmount - len(cardsList)
		percentageQualities := []int{
			chore.config.PercentageQualities.Wood,
			chore.config.PercentageQualities.Silver,
			chore.config.PercentageQualities.Gold,
			chore.config.PercentageQualities.Diamond,
		}

		for i := 0; i < cardsAmount; i++ {
			if _, err := chore.cards.Create(ctx, uuid.Nil, percentageQualities, cards.TypeOrdered); err != nil {
				return ChoreError.Wrap(err)
			}
		}

		return ChoreError.Wrap(err)
	})
}

// Close closes the chore for renewal cards in store.
func (chore *Chore) Close() {
	chore.Loop.Close()
}
