// Copyright (C) 2023 Creditor Corp. Group.
// See LICENSE for copying information.

package servicetesting

import (
	"context"
	"testing"

	"ultimatedivision"
	"ultimatedivision/database/dbtesting"
	"ultimatedivision/internal/logger/zaplog"
	"ultimatedivision/users"
)

// Run method will establish connection with db, create services, allows reconfiguring config and run tests.
func Run(t *testing.T, setupConfig func() ultimatedivision.Config, test func(ctx context.Context, t *testing.T, ultimatedivision *ultimatedivision.Peer)) {
	dbtesting.Run(t, func(ctx context.Context, t *testing.T, db ultimatedivision.DB) {
		ultimatedivision := &ultimatedivision.Peer{
			Log:      zaplog.NewLog(),
			Database: db,
		}

		var err error

		ultimatedivision.Users.Service = users.NewService(db.Users())
		if err != nil {
			t.Fatal(err)
		}

		test(ctx, t, ultimatedivision)
	})
}
