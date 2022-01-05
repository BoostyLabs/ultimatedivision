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

// PositionInTheField defines position(coordinate) of player/ball in the field.
type PositionInTheField struct {
	X int `json:"x"`
	Y int `json:"y"`
}

// Compare compares two positions in the field.
func (position PositionInTheField) Compare(position1 PositionInTheField) bool {
	return position.X == position1.X && position.Y == position1.Y
}

// Config defines configuration for matches.
type Config struct {
	SizeOfFieldByOX int `json:"sizeOfFieldByOX"`
	SizeOfFieldByOY int `json:"sizeOfFieldByOY"`

	RoundDuration  int `json:"roundDuration"`
	NumberOfRounds int `json:"numberOfRounds"`

	Positions struct {
		GK struct {
			PositionInTheField
		} `json:"gk"`
		LB struct {
			PositionInTheField
		} `json:"lb"`
		LCB struct {
			PositionInTheField
		} `json:"lcb"`
		CCB struct {
			PositionInTheField
		} `json:"ccb"`
		RCB struct {
			PositionInTheField
		} `json:"rcb"`
		RB struct {
			PositionInTheField
		} `json:"rb"`
		LCDM struct {
			PositionInTheField
		} `json:"lcdm"`
		CCDM struct {
			PositionInTheField
		} `json:"ccdm"`
		RCDM struct {
			PositionInTheField
		} `json:"rcdm"`
		LCM struct {
			PositionInTheField
		} `json:"lcm"`
		CCM struct {
			PositionInTheField
		} `json:"ccm"`
		RCM struct {
			PositionInTheField
		} `json:"rcm"`
		LM struct {
			PositionInTheField
		} `json:"lm"`
		RM struct {
			PositionInTheField
		} `json:"rm"`
		LCAM struct {
			PositionInTheField
		} `json:"lcam"`
		CCAM struct {
			PositionInTheField
		} `json:"ccam"`
		RCAM struct {
			PositionInTheField
		} `json:"rcam"`
		LWB struct {
			PositionInTheField
		} `json:"lwb"`
		RWB struct {
			PositionInTheField
		} `json:"rwb"`
		LW struct {
			PositionInTheField
		} `json:"lw"`
		RW struct {
			PositionInTheField
		} `json:"rw"`
		LST struct {
			PositionInTheField
		} `json:"lst"`
		CST struct {
			PositionInTheField
		} `json:"cst"`
		RST struct {
			PositionInTheField
		} `json:"rst"`
	} `json:"positions"`

	pagination.Cursor `json:"pagination"`

	NumberOfPointsForWin    int `json:"numberOfPointsForWin"`
	NumberOfPointsForDraw   int `json:"numberOfPointsForDraw"`
	NumberOfPointsForLosing int `json:"numberOfPointsForLosing"`
}

// Action defines list of possible player action in the field.
type Action string

const (
	// ActionMove defines move action by player.
	ActionMove Action = "move"
	// ActionMoveWithBall defines move action by player with ball.
	ActionMoveWithBall Action = "moveWithBall"
	// ActionPass defines pass by player to another player.
	ActionPass Action = "pass"
	// ActionCrossPass defines passing the ball by throwing it into the air in the direction of a player on his team.
	ActionCrossPass Action = "crossPass"
	// ActionPassThrough defines pass in free zone on the move often between players of the other team.
	ActionPassThrough Action = "passTrough"
	// ActionDirectShot defines direct shot.
	ActionDirectShot Action = "directShot"
	// ActionCurlShot defines curl shot.
	ActionCurlShot Action = "curlShot"
	// ActionTakeawayShot defines powerful shot from the box.
	ActionTakeawayShot Action = "takeawayShot"
	// ActionTackle defines tackling the ball from an opponent.
	ActionTackle Action = "tackle"
	// ActionSlidingTackle defines tackle by sliding on the field.
	ActionSlidingTackle Action = "slidingTackle"
	// ActionDribbling defines action when player move with some feints ot tricks.
	ActionDribbling Action = "dribbling"
	// ActionFeints defines action when player show feints.
	ActionFeints Action = "feints"
)

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
	MatchResults []MatchResult                `json:"matchResults"`
	Transaction  currencywaitlist.Transaction `json:"transaction"`
	Question     string                       `json:"question"`
}

// SquadCardWithPosition contains which card and where it located in the field.
type SquadCardWithPosition struct {
	Card     cards.Card         `json:"card"`
	Position PositionInTheField `json:"position"`
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
}

// ActionRequest defines request for every action.
type ActionRequest struct {
	PlayerID          uuid.UUID   `json:"playerId"`
	Action            Action      `json:"action"`
	Distance          int         `json:"distance"`
	ReceiverPlayerID  uuid.UUID   `json:"receiverPlayerId"`
	OpponentPlayerIDs []uuid.UUID `json:"opponentPlayerIds"`
}

// IsValid checks is action request valid.
func (a ActionRequest) IsValid() bool {
	switch {
	case a.Action == ActionMove || a.Action == ActionMoveWithBall:
		if a.ReceiverPlayerID != uuid.Nil {
			return false
		}
	default:
		if a.ReceiverPlayerID == uuid.Nil {
			return false
		}
	}

	if a.PlayerID == uuid.Nil {
		return false
	}

	if a.Action == ActionMove || a.Action == ActionMoveWithBall || a.Action == ActionPass ||
		a.Action == ActionCrossPass || a.Action == ActionPassThrough || a.Action == ActionDirectShot ||
		a.Action == ActionCurlShot || a.Action == ActionTakeawayShot || a.Action == ActionTackle ||
		a.Action == ActionSlidingTackle || a.Action == ActionDribbling {
		return true
	}

	return false
}

// CardPossibleAction defines in which position card could be placed and which action it could do there.
type CardPossibleAction struct {
	CardID    uuid.UUID            `json:"cardId"`
	Action    Action               `json:"action"`
	Positions []PositionInTheField `json:"positions"`
}
