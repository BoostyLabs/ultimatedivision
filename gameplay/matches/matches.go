// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package matches

import (
	"context"

	"github.com/google/uuid"
	"github.com/zeebo/errs"

	"ultimatedivision/cards"
	"ultimatedivision/clubs"
	"ultimatedivision/pkg/pagination"
	"ultimatedivision/udts/currencywaitlist"
)

// ErrNoMatch indicated that match does not exist.
var ErrNoMatch = errs.Class("match does not exist")

// DB is exposing access to matches db.
//
// architecture: DB
type DB interface {
	// Create creates match in the database.
	Create(ctx context.Context, match Match) error
	// Get returns match from the database.
	Get(ctx context.Context, id uuid.UUID) (Match, error)
	// ListMatches returns page of matches from the database.
	ListMatches(ctx context.Context, cursor pagination.Cursor) (Page, error)
	// ListSquadMatches returns all matches played by squad in season.
	ListSquadMatches(ctx context.Context, seasonID int) ([]Match, error)
	// UpdateMatch updates the number of points that users received for a played match.
	UpdateMatch(ctx context.Context, match Match) error
	// Delete deletes match from the database.
	Delete(ctx context.Context, id uuid.UUID) error
	// AddGoals adds new goal in the match.
	AddGoals(ctx context.Context, matchGoals []MatchGoals) error
	// ListMatchGoals returns all goals from the match from the database.
	ListMatchGoals(ctx context.Context, matchID uuid.UUID) ([]MatchGoals, error)
	// GetMatchResult returns goals of each user in the match from db.
	GetMatchResult(ctx context.Context, matchID uuid.UUID) ([]MatchResult, error)
}

// Config defines configuration for matches.
type Config struct {
	Periods struct {
		First struct {
			Begin int `json:"begin"`
			End   int `json:"end"`
		} `json:"first"`
		Second struct {
			Begin int `json:"begin"`
			End   int `json:"end"`
		} `json:"second"`
		Third struct {
			Begin int `json:"begin"`
			End   int `json:"end"`
		} `json:"third"`
		Fourth struct {
			Begin int `json:"begin"`
			End   int `json:"end"`
		} `json:"fourth"`
		Fifth struct {
			Begin int `json:"begin"`
			End   int `json:"end"`
		} `json:"fifth"`
		Sixth struct {
			Begin int `json:"begin"`
			End   int `json:"end"`
		} `json:"sixth"`
		Seventh struct {
			Begin int `json:"begin"`
			End   int `json:"end"`
		} `json:"seventh"`
		Eighth struct {
			Begin int `json:"begin"`
			End   int `json:"end"`
		} `json:"eighth"`
		Ninth struct {
			Begin int `json:"begin"`
			End   int `json:"end"`
		} `json:"ninth"`
		Tenth struct {
			Begin int `json:"begin"`
			End   int `json:"end"`
		} `json:"tenth"`
	} `json:"periods"`

	GoalProbability int `json:"goalProbability"`

	SquadPowerAccuracy int `json:"squadPowerAccuracy"`

	GoalProbabilityByPosition struct {
		ST  int `json:"st"`
		RW  int `json:"rw"`
		LW  int `json:"lw"`
		CAM int `json:"cam"`
		CM  int `json:"cm"`
		RM  int `json:"rm"`
		LM  int `json:"lm"`
		CDM int `json:"cdm"`
		CD  int `json:"cd"`
		LB  int `json:"lb"`
		RB  int `json:"rb"`
	} `json:"goalProbabilityByPosition"`

	pagination.Cursor `json:"pagination"`

	NumberOfPointsForWin    int `json:"numberOfPointsForWin"`
	NumberOfPointsForDraw   int `json:"numberOfPointsForDraw"`
	NumberOfPointsForLosing int `json:"numberOfPointsForLosing"`
}

// Match describes match entity.
type Match struct {
	ID          uuid.UUID `json:"id"`
	User1ID     uuid.UUID `json:"user1Id"`
	Squad1ID    uuid.UUID `json:"squad1Id"`
	User1Points int       `json:"user1Points"`
	User2ID     uuid.UUID `json:"user2Id"`
	Squad2ID    uuid.UUID `json:"squad2Id"`
	User2Points int       `json:"user2Points"`
	SeasonID    int       `json:"seasonId"`
}

// MatchGoals defines goals scored by clubs.
type MatchGoals struct {
	ID      uuid.UUID `json:"id"`
	MatchID uuid.UUID `json:"matchId"`
	UserID  uuid.UUID `json:"userId"`
	CardID  uuid.UUID `json:"cardId"`
	Minute  int       `json:"minute"`
}

// MatchResult defines quantity goals of each user in the match
// and which cards of user's squad scored in which minute.
type MatchResult struct {
	UserID        uuid.UUID    `json:"userId"`
	QuantityGoals int          `json:"quantityGoals"`
	Goalscorers   []Goalscorer `json:"goals"`
}

// GameResult entity describes values which send to user after game.
type GameResult struct {
	MatchResults      []MatchResult                      `json:"matchResults"`
	Transaction       currencywaitlist.Transaction       `json:"transaction"`
	CasperTransaction currencywaitlist.CasperTransaction `json:"casperTransaction"`
	Question          string                             `json:"question"`
	RPCNodeAddress    string                             `json:"rpcNodeAddress"`
}

// Swap swaps match results.
func Swap(matchResults []MatchResult) []MatchResult {
	matchResults[0], matchResults[1] = matchResults[1], matchResults[0]

	return matchResults
}

// Goalscorer defines which card scored goal in which minute.
type Goalscorer struct {
	Card   cards.Card `json:"card"`
	Minute int        `json:"minute"`
}

// Page holds match page entity which is used to show listed page of matches.
type Page struct {
	Matches []Match         `json:"matches"`
	Page    pagination.Page `json:"page"`
}

// Statistic defined statistic of club in season.
type Statistic struct {
	Club           clubs.Club `json:"club"`
	MatchPlayed    int        `json:"matchPlayed"`
	Wins           int        `json:"wins"`
	Losses         int        `json:"losses"`
	Draws          int        `json:"draws"`
	GoalDifference int        `json:"goalDifference"`
	Points         int        `json:"points"`
	SeasonID       int        `json:"season_id"`
}
