// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package matches

import (
	"context"
	"math/rand"
	"sort"
	"time"
	"ultimatedivision/cards"
	"ultimatedivision/clubs"

	"github.com/google/uuid"
	"github.com/zeebo/errs"

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
}

// NewService is a constructor for matches service.
func NewService(clubs DB, config Config) *Service {
	return &Service{
		matches: clubs,
		config:  config,
	}
}

// periodBegin defines index of the beginning of period.
const periodBegin = 0

// periodEnd defines index of the ending of period.
const periodEnd = 1

// Play initiates match between users, calls methods to generate result.
func (service *Service) Play(ctx context.Context, formationSquad1 clubs.Formation, formationSquad2 clubs.Formation, squadCards1 []clubs.SquadCard, squadCards2 []clubs.SquadCard, user1, user2 uuid.UUID) {
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
		if randNumber > 0 && randNumber <= goalProbability {
			period := periods[key]
			minute := generateMinute(period[periodBegin], period[periodEnd])
			// TODO: call choose squad func
			chooseGoalscorer(squadCards1, goalProbabilityByPosition)
		}
	}

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
func chooseGoalscorer(squadCards []clubs.SquadCard, probabilityByPosition map[clubs.Position]int) cards.Card {
	rand.Seed(time.Now().UTC().UnixNano())
	var goalscorer []cards.Card
	randNumber := rand.Intn(100) + 1

	switch {
	case randNumber > 0 && randNumber <= probabilityByPosition[clubs.ST]:
		for _, card := range squadCards{
			var squadCard cards.Card
			if card.Position == clubs.ST {
				goalscorer = append(goalscorer, squadCard)
			}
		}
	case randNumber > probabilityByPosition[clubs.ST] && randNumber < probabilityByPosition[clubs.ST] + probabilityByPosition[clubs.RW]:
		for _, card := range squadCards{
			var squadCard cards.Card
			if card.Position == clubs.RW || card.Position == clubs.LW || card.Position == clubs.CM || card.Position == clubs.CAM  {
				goalscorer = append(goalscorer, squadCard)
			}
		}
	}
	return cards.Card{}
}

// chooseSquad returns the squad which is stronger in the period.
func chooseSquad(squadPowerAccuracy int, squad1, squad2 clubs.Squad) clubs.Squad {
	// TODO: call method to count power of squads, generate accuracy.
	return clubs.Squad{}
}

// Create creates new match.
func (service *Service) Create(ctx context.Context, user1ID, user2ID uuid.UUID) error {
	newMatch := Match{
		ID:      uuid.New(),
		User1ID: user1ID,
		User2ID: user2ID,
	}

	return ErrMatches.Wrap(service.matches.Create(ctx, newMatch))
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
