// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package matches

import (
	"context"
	"math/rand"
	"sort"
	"time"

	"github.com/google/uuid"
	"github.com/zeebo/errs"

	"ultimatedivision/clubs"
	"ultimatedivision/internal/pagination"
)

// ErrMatches indicates that there was an error in the service.
var ErrMatches = errs.Class("matches service error")

// Service is handling matches related logic.
//
// architecture: Service
type Service struct {
	matches DB
	config  Config
	clubs   *clubs.Service
}

// NewService is a constructor for matches service.
func NewService(matches DB, config Config, clubs *clubs.Service) *Service {
	return &Service{
		matches: matches,
		config:  config,
		clubs:   clubs,
	}
}

// periodBegin defines index of the beginning of period.
const periodBegin = 0

// periodEnd defines index of the ending of period.
const periodEnd = 1

// Play initiates match between users, calls methods to generate result.
func (service *Service) Play(ctx context.Context, matchID uuid.UUID, squadCards1 []clubs.SquadCard, squadCards2 []clubs.SquadCard, user1, user2 uuid.UUID) error {
	periods := map[string][]int{
		"period1":  {service.config.Periods.First.Begin, service.config.Periods.First.End},
		"period2":  {service.config.Periods.Second.Begin, service.config.Periods.Second.End},
		"period3":  {service.config.Periods.Third.Begin, service.config.Periods.Third.End},
		"period4":  {service.config.Periods.Fourth.Begin, service.config.Periods.Fourth.End},
		"period5":  {service.config.Periods.Fifth.Begin, service.config.Periods.Fifth.End},
		"period6":  {service.config.Periods.Sixth.Begin, service.config.Periods.Sixth.End},
		"period7":  {service.config.Periods.Seventh.Begin, service.config.Periods.Seventh.End},
		"period8":  {service.config.Periods.Eighth.Begin, service.config.Periods.Eighth.End},
		"period9":  {service.config.Periods.Ninth.Begin, service.config.Periods.Ninth.End},
		"period10": {service.config.Periods.Tenth.Begin, service.config.Periods.Tenth.End},
	}

	goalProbability := service.config.GoalProbability

	squadPowerAccuracy := service.config.SquadPowerAccuracy

	goalProbabilityByPosition := map[clubs.Position]int{
		clubs.ST:  service.config.GoalProbabilityByPosition.ST,
		clubs.RW:  service.config.GoalProbabilityByPosition.RW,
		clubs.LW:  service.config.GoalProbabilityByPosition.LW,
		clubs.CAM: service.config.GoalProbabilityByPosition.CAM,
		clubs.CM:  service.config.GoalProbabilityByPosition.CM,
		clubs.RM:  service.config.GoalProbabilityByPosition.RM,
		clubs.LM:  service.config.GoalProbabilityByPosition.LM,
		clubs.CDM: service.config.GoalProbabilityByPosition.CDM,
		clubs.CD:  service.config.GoalProbabilityByPosition.CD,
		clubs.LB:  service.config.GoalProbabilityByPosition.LB,
		clubs.RB:  service.config.GoalProbabilityByPosition.RB,
	}

	rand.Seed(time.Now().UTC().UnixNano())

	periodName := sortMapKey(periods)

	for _, key := range periodName {
		randNumber := rand.Intn(100) + 1
		if randNumber > goalProbability {
			continue
		}
		period := periods[key]

		minute := generateMinute(period[periodBegin], period[periodEnd])
		userID, cardID, err := service.chooseSquad(ctx, goalProbabilityByPosition,
			squadPowerAccuracy, user1, user2, squadCards1, squadCards2)
		if err != nil {
			return ErrMatches.Wrap(err)
		}

		err = service.AddGoal(ctx, MatchGoals{
			ID:      uuid.New(),
			MatchID: matchID,
			UserID:  userID,
			CardID:  cardID,
			Minute:  minute,
		})
		if err != nil {
			return ErrMatches.Wrap(err)
		}
	}
	return nil
}

// sortMapKey returns sorted slice names of periods.
func sortMapKey(periodMap map[string][]int) []string {
	periodsName := make([]string, 0, len(periodMap))

	for periodName := range periodMap {
		periodsName = append(periodsName, periodName)
	}

	sort.Strings(periodsName)

	return periodsName
}

// generateMinute generates the minute at which the goal was scored.
func generateMinute(begin, end int) int {
	rand.Seed(time.Now().UnixNano())
	minute := begin + rand.Intn(end-begin+1)

	return minute
}

// choseGoalscorer returns id of cards which scored goal.
func chooseGoalscorer(squadCards []clubs.SquadCard, goalByPosition map[clubs.Position]int) uuid.UUID {
	rand.Seed(time.Now().UTC().UnixNano())
	var cardsByPosition []uuid.UUID
	randNumber := rand.Intn(100) + 1

	// TODO: refactor positions.

	switch {
	case randNumber > 0 && randNumber <= goalByPosition[clubs.ST]:
		for _, card := range squadCards {
			if card.Position == clubs.ST {
				cardsByPosition = append(cardsByPosition, card.CardID)
			}
		}

		if len(cardsByPosition) > 0 {
			break
		}

		fallthrough
	case randNumber > goalByPosition[clubs.ST] &&
		randNumber < goalByPosition[clubs.ST]+goalByPosition[clubs.RW]:
		for _, card := range squadCards {
			if card.Position == clubs.RW || card.Position == clubs.LW ||
				card.Position == clubs.CM || card.Position == clubs.CAM {
				cardsByPosition = append(cardsByPosition, card.CardID)
			}
		}

		if len(cardsByPosition) > 0 {
			break
		}

		fallthrough
	case randNumber > goalByPosition[clubs.ST]+goalByPosition[clubs.RW] &&
		randNumber < 100-goalByPosition[clubs.CD]:
		for _, card := range squadCards {
			if card.Position == clubs.RM || card.Position == clubs.LM ||
				card.Position == clubs.CDM || card.Position == clubs.CAM {
				cardsByPosition = append(cardsByPosition, card.CardID)
			}
		}

		if len(cardsByPosition) > 0 {
			break
		}

		fallthrough
	case randNumber >= 100-goalByPosition[clubs.CD] && randNumber < 100:
		for _, card := range squadCards {
			if card.Position == clubs.CD || card.Position == clubs.LB ||
				card.Position == clubs.RB {
				cardsByPosition = append(cardsByPosition, card.CardID)
			}
		}
	}

	randIndex := rand.Intn(len(squadCards))
	goalscorer := cardsByPosition[randIndex]

	return goalscorer
}

// chooseSquad returns the squad which is stronger in the period.
func (service *Service) chooseSquad(ctx context.Context, goalByPosition map[clubs.Position]int, squadPowerAccuracy int, user1 uuid.UUID, user2 uuid.UUID, squadCards1, squadCards2 []clubs.SquadCard) (uuid.UUID, uuid.UUID, error) {
	squad1Effectiveness, err := service.clubs.CalculateEffectivenessOfSquad(ctx, squadCards1)
	if err != nil {
		return uuid.Nil, uuid.Nil, ErrMatches.Wrap(err)
	}

	squad2Effectiveness, err := service.clubs.CalculateEffectivenessOfSquad(ctx, squadCards2)
	if err != nil {
		return uuid.Nil, uuid.Nil, ErrMatches.Wrap(err)
	}

	randAccuracy1 := float64(rand.Intn(2*squadPowerAccuracy)+1-squadPowerAccuracy) / 100
	randAccuracy2 := float64(rand.Intn(2*squadPowerAccuracy)+1-squadPowerAccuracy) / 100

	squad1Effectiveness += squad1Effectiveness * randAccuracy1
	squad2Effectiveness += squad1Effectiveness * randAccuracy2

	if squad1Effectiveness > squad2Effectiveness {
		return user1, chooseGoalscorer(squadCards1, goalByPosition), nil
	}

	return user2, chooseGoalscorer(squadCards2, goalByPosition), nil
}

// Create creates new match.
func (service *Service) Create(ctx context.Context, squadCards1 []clubs.SquadCard, squadCards2 []clubs.SquadCard, user1ID, user2ID uuid.UUID) error {
	newMatch := Match{
		ID:      uuid.New(),
		User1ID: user1ID,
		User2ID: user2ID,
	}

	err := service.matches.Create(ctx, newMatch)
	if err != nil {
		return ErrMatches.Wrap(err)
	}

	err = service.Play(ctx, newMatch.ID, squadCards1, squadCards2, newMatch.User1ID, newMatch.User2ID)

	return ErrMatches.Wrap(err)
}

// Get returns match by id.
func (service *Service) Get(ctx context.Context, matchID uuid.UUID) (Match, error) {
	match, err := service.matches.Get(ctx, matchID)

	return match, ErrMatches.Wrap(err)
}

// List returns page of matches.
func (service *Service) List(ctx context.Context, cursor pagination.Cursor) (Page, error) {
	if cursor.Limit <= 0 {
		cursor.Limit = service.config.Cursor.Limit
	}
	if cursor.Page <= 0 {
		cursor.Page = service.config.Cursor.Page
	}

	allMatches, err := service.matches.ListMatches(ctx, cursor)

	return allMatches, ErrMatches.Wrap(err)
}

// GetGoalsByUserID returns number of goals scored by user's squad.
func (service *Service) GetGoalsByUserID(ctx context.Context, userID, matchID uuid.UUID) (int, error) {
	number, err := service.matches.GetGoals(ctx, userID, matchID)

	return number, ErrMatches.Wrap(err)
}

// Delete deletes match.
func (service *Service) Delete(ctx context.Context, id uuid.UUID) error {
	return ErrMatches.Wrap(service.matches.Delete(ctx, id))
}

// AddGoal adds goal in the match.
func (service *Service) AddGoal(ctx context.Context, matchGoal MatchGoals) error {
	return ErrMatches.Wrap(service.matches.AddGoal(ctx, matchGoal))
}

// ListMatchGoals returns all goals scored in the match.
func (service *Service) ListMatchGoals(ctx context.Context, matchID uuid.UUID) ([]MatchGoals, error) {
	matchGoals, err := service.matches.ListMatchGoals(ctx, matchID)

	return matchGoals, ErrMatches.Wrap(err)
}
