// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package seasons

import (
	"context"
	"sort"
	"time"

	"github.com/google/uuid"
	"github.com/zeebo/errs"

	"ultimatedivision/clubs"
	"ultimatedivision/divisions"
	"ultimatedivision/gameplay/matches"
)

// ErrSeasons indicates that there was an error in the service.
var ErrSeasons = errs.Class("seasons service error")

// Service is handling seasons related logic.
//
// architecture: Service
type Service struct {
	seasons   DB
	divisions *divisions.Service
	matches   *matches.Service
	config    Config
	clubs     *clubs.Service
}

// NewService is a constructor for seasons service.
func NewService(seasons DB, config Config, divisions *divisions.Service, matches *matches.Service, clubs *clubs.Service) *Service {
	return &Service{
		seasons:   seasons,
		divisions: divisions,
		config:    config,
		matches:   matches,
		clubs:     clubs,
	}
}

// Create creates a season.
func (service *Service) Create(ctx context.Context) error {
	divisions, err := service.divisions.List(ctx)
	if err != nil {
		return ErrSeasons.Wrap(err)
	}

	for _, division := range divisions {
		season := Season{
			DivisionID: division.ID,
			StartedAt:  time.Now().UTC(),
			EndedAt:    time.Time{},
		}
		if err = service.seasons.Create(ctx, season); err != nil {
			return ErrSeasons.Wrap(err)
		}
	}

	return nil
}

// EndSeason changes status when season end.
func (service *Service) EndSeason(ctx context.Context, id int) error {
	return ErrSeasons.Wrap(service.seasons.EndSeason(ctx, id))
}

// List returns all seasons from DB.
func (service *Service) List(ctx context.Context) ([]Season, error) {
	seasons, err := service.seasons.List(ctx)
	return seasons, ErrSeasons.Wrap(err)
}

// GetCurrentSeasons returns all current seasons from DB.
func (service *Service) GetCurrentSeasons(ctx context.Context) ([]Season, error) {
	seasons, err := service.seasons.GetCurrentSeasons(ctx)
	return seasons, ErrSeasons.Wrap(err)
}

// Get returns season from DB.
func (service *Service) Get(ctx context.Context, seasonID int) (Season, error) {
	season, err := service.seasons.Get(ctx, seasonID)
	return season, ErrSeasons.Wrap(err)
}

// Delete deletes a season.
func (service *Service) Delete(ctx context.Context, id int) error {
	return ErrSeasons.Wrap(service.seasons.Delete(ctx, id))
}

// GetSeasonByDivisionID returns season by division id.
func (service *Service) GetSeasonByDivisionID(ctx context.Context, divisionID uuid.UUID) (Season, error) {
	season, err := service.seasons.GetSeasonByDivisionID(ctx, divisionID)
	return season, ErrSeasons.Wrap(err)
}

// GetAllClubsStatistics returns all clubs statistics by division.
func (service *Service) GetAllClubsStatistics(ctx context.Context) (map[divisions.Division][]matches.Statistic, error) {
	currentSeasons, err := service.GetCurrentSeasons(ctx)
	if err != nil {
		return nil, ErrSeasons.Wrap(err)
	}
	clubs, err := service.clubs.List(ctx)
	if err != nil {
		return nil, ErrSeasons.Wrap(err)
	}

	statisticsMap := make(map[divisions.Division][]matches.Statistic)
	var statistics []matches.Statistic
	for _, currentSeason := range currentSeasons {
		for _, club := range clubs {
			statistic, err := service.matches.GetStatistic(ctx, club.ID, currentSeason.ID)
			if err != nil {
				return nil, ErrSeasons.Wrap(err)
			}
			division, err := service.divisions.Get(ctx, currentSeason.DivisionID)
			if err != nil {
				return nil, ErrSeasons.Wrap(err)
			}
			if statistic.MatchPlayed > matches.MinNumberOfMatches {
				if division.ID == statistic.Club.DivisionID {
					statistics = append(statistics, statistic)
					statisticsMap[division] = statistics
				}
			}
		}
	}

	return statisticsMap, nil
}

// UpdatesClubsToNewDivision updates clubs to new division.
func (service *Service) UpdatesClubsToNewDivision(ctx context.Context) error {
	clubsStatisticsByDivision, err := service.GetAllClubsStatistics(ctx)
	if err != nil {
		return ErrSeasons.Wrap(err)
	}

	var totalPassingClubs float64
	for division, statistics := range clubsStatisticsByDivision {
		var percent float64
		percent = 100 / float64(len(statistics))
		if percent < float64(division.PassingPercent) {
			totalPassingClubs = float64(division.PassingPercent) / percent
		} else {
			totalPassingClubs = 1
		}
		sortStatistics := statistics
		sort.Slice(sortStatistics, func(i, j int) bool {
			return sortStatistics[i].Points < sortStatistics[j].Points
		})
		topStatisticsClubs := sortStatistics[int(totalPassingClubs):]
		lowStatisticsClubs := sortStatistics[:int(totalPassingClubs)]

		divisionHigher, err := service.divisions.GetByName(ctx, division.Name-1)
		if err != nil {
			return ErrSeasons.Wrap(err)
		}
		for _, statistic := range topStatisticsClubs {
			err := service.clubs.UpdateClubToNewDivision(ctx, statistic.Club.ID, divisionHigher.ID)
			if err != nil {
				return ErrSeasons.Wrap(err)
			}
		}

		lastDivision, err := service.divisions.GetLastDivision(ctx)
		if err != nil {
			return ErrSeasons.Wrap(err)
		}
		if division.Name > lastDivision.Name {
			divisionLower, err := service.divisions.GetByName(ctx, division.Name+1)
			if err != nil {
				return ErrSeasons.Wrap(err)

			}
			for _, statistic := range lowStatisticsClubs {
				err := service.clubs.UpdateClubToNewDivision(ctx, statistic.Club.ID, divisionLower.ID)
				if err != nil {
					return ErrSeasons.Wrap(err)
				}
			}
		}
	}

	return nil
}
