package marketplace

import (
	"context"
	sync2 "sync"

	"github.com/zeebo/errs"

	"ultimatedivision/cards"
	"ultimatedivision/internal/logger"
	"ultimatedivision/internal/sync"
	"ultimatedivision/users"
)

var (
	// ChoreError represents chore error type.
	ChoreError = errs.Class("expiration lot chore error")
)

// Chore requests access token for contis api calls, re-requests it after token's expiration time.
//
// architecture: Chore
type Chore struct {
	log logger.Logger

	service *Service
	Loop    *sync.Cycle

	Mutex sync2.Mutex
}

// NewChore instantiates Chore.
func NewChore(log logger.Logger, config Config, marketplace DB, users *users.Service, cards *cards.Service) *Chore {

	return &Chore{
		log: log,
		service: NewService(
			marketplace,
			users,
			cards,
		),
		Loop: sync.NewCycle(config.LotRenewalInterval),
	}
}

// Run starts the chore.
func (chore *Chore) Run(ctx context.Context) (err error) {
	return chore.Loop.Run(ctx, func(ctx context.Context) error {
		chore.Mutex.Lock()
		lots, err := chore.service.ListActiveLotsWhereEndTimeLTENow(ctx)

		// TODO: the transaction may be required for all operations.
		for _, lot := range lots {
			if lot.CurrentPrice != 0 {
				// TODO: unhold old user's money.

				winLot := WinLot{
					ID:        lot.ID,
					ItemID:    lot.ItemID,
					Type:      TypeCard,
					UserID:    lot.UserID,
					ShopperID: lot.ShopperID,
					Status:    StatusSold,
					Amount:    lot.MaxPrice,
				}

				err := chore.service.WinLot(ctx, winLot)
				if err != nil {
					return ChoreError.Wrap(err)
				}

			} else {
				err := chore.service.UpdateStatusLot(ctx, lot.ID, StatusExpired)
				if err != nil {
					return ChoreError.Wrap(err)
				}

				if lot.Type == TypeCard {
					if err := chore.service.cards.UpdateStatus(ctx, lot.ItemID, cards.StatusActive); err != nil {
						return ErrMarketplace.Wrap(err)
					}

				}
				// TODO: check other items
			}
		}

		chore.Mutex.Unlock()

		return ChoreError.Wrap(err)
	})
}
