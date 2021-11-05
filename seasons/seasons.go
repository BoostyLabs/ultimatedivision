package seasons

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/zeebo/errs"
)

// ErrNoSeasons indicated that seasons does not exist.
var ErrNoSeasons = errs.Class("seasons does not exist")

// DB exposes access to seasons db.
//
// architecture: DB
type DB interface {
	// Create creates a season and writes to the database.
	Create(ctx context.Context, season Season) error
	// List returns all seasons from the data base.
	List(ctx context.Context) ([]Season, error)
	// Get returns season by id from the data base.
	Get(ctx context.Context, id uuid.UUID) (Season, error)
	// Delete deletes a season in the database.
	Delete(ctx context.Context, id uuid.UUID) error
}

// Season describes seasons entity.
type Season struct {
	ID         int       `json:"id"`
	DivisionID uuid.UUID `json:"divisionID"`
	StartedAt  time.Time `json:"startedAt"`
	EndedAt    time.Time `json:"endedAt"`
}

// Config defines configuration for seasons.
type Config struct {
	SeasonsRenewalInterval time.Duration `json:"seasonsRenewalInterval"`
}
