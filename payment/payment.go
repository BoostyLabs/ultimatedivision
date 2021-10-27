// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package payment

import (
	"context"

	"github.com/google/uuid"

	"ultimatedivision/lootboxes"
	"ultimatedivision/pkg/cryptoutils"
)
// Crypto provides access to all operations related to crypto payment.
type Crypto interface {
	// TopUpBalance tops up balance of user.
	TopUpBalance(ctx context.Context, userID uuid.UUID, walletAddress cryptoutils.Address) error
	// Withdraw withdraws money to a crypto wallet.
	Withdraw(ctx context.Context, userID uuid.UUID, walletAddress cryptoutils.Address) error
}

// UDTToken provides access to all operations related to UDT token payment.
type UDTToken interface {
	// BuyLootbox buys certain lootbox.
	BuyLootbox(ctx context.Context, userID uuid.UUID, lootboxType lootboxes.Type) error
	// BuyCard buys certain card.
	BuyCard(ctx context.Context, uuid2, lotID uuid.UUID) error
	// SellCard sells card.
	SellCard(ctx context.Context, userID uuid.UUID, cardID uuid.UUID) error
	// MakeBid makes bid on lot in the marketplace.
	MakeBid(ctx context.Context, userID, lotID uuid.UUID) error
}

// Payment provides access to all operations related to currency.
type Payment interface {
	Crypto
	UDTToken
}
