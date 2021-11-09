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
	// EndSeason updates a status in the database when season ended.
	EndSeason(ctx context.Context, id int) error
	// List returns all seasons from the data base.
	List(ctx context.Context) ([]Season, error)
	// Get returns season by id from the data base.
	Get(ctx context.Context, id int) (Season, error)
	// Delete deletes a season in the database.
	Delete(ctx context.Context, id int) error
}

// Status defines the list of possible season statuses.
type Status int

const (
	// StatusStarted indicates that season start.
	StatusStarted Status = 0
	// StatusEnded indicates that season ended.
	StatusEnded Status = 1
)

// Season describes seasons entity.
type Season struct {
	ID         int       `json:"id"`
	DivisionID uuid.UUID `json:"divisionID"`
	Status     Status    `json:"status"`
	StartedAt  time.Time `json:"startedAt"`
	EndedAt    time.Time `json:"endedAt"`
}

// Config defines configuration for seasons.
type Config struct {
	SeasonTime time.Duration `json:"seasonTime"`
}
