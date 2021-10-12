// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package cards

import (
	"context"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/zeebo/errs"

	"ultimatedivision/cards/avatars"
	"ultimatedivision/internal/pagination"
)

// ErrCards indicated that there was an error in service.
var ErrCards = errs.Class("cards service error")

// Service is handling cards related logic.
//
// architecture: Service
type Service struct {
	cards   DB
	config  Config
	avatars *avatars.Service
}

// NewService is a constructor for cards service.
func NewService(cards DB, config Config, avatars *avatars.Service) *Service {
	return &Service{
		cards:   cards,
		config:  config,
		avatars: avatars,
	}
}

// Create adds card in DB.
func (service *Service) Create(ctx context.Context, userID uuid.UUID, percentageQualities []int, imageName string) (Card, error) {
	qualities := map[string]int{
		"wood":    percentageQualities[0],
		"silver":  percentageQualities[1],
		"gold":    percentageQualities[2],
		"diamond": percentageQualities[3],
	}

	minHeight := service.config.Height.Min
	maxHeight := service.config.Height.Max
	minWeight := service.config.Weight.Min
	maxWeight := service.config.Weight.Max

	var skills = map[string]map[string]int{
		"wood": {
			"elementary":  service.config.Skills.Wood.Elementary,
			"basic":       service.config.Skills.Wood.Basic,
			"medium":      service.config.Skills.Wood.Medium,
			"upperMedium": service.config.Skills.Wood.UpperMedium,
			"advanced":    service.config.Skills.Wood.Advanced,
		},
		"silver": {
			"elementary":  service.config.Skills.Silver.Elementary,
			"basic":       service.config.Skills.Silver.Basic,
			"medium":      service.config.Skills.Silver.Medium,
			"upperMedium": service.config.Skills.Silver.UpperMedium,
			"advanced":    service.config.Skills.Silver.Advanced,
		},
		"gold": {
			"elementary":    service.config.Skills.Gold.Elementary,
			"basic":         service.config.Skills.Gold.Basic,
			"medium":        service.config.Skills.Gold.Medium,
			"upperMedium":   service.config.Skills.Gold.UpperMedium,
			"advanced":      service.config.Skills.Gold.Advanced,
			"upperAdvanced": service.config.Skills.Gold.UpperMedium,
		},
		"diamond": {
			"basic":         service.config.Skills.Diamond.Basic,
			"medium":        service.config.Skills.Diamond.Medium,
			"upperMedium":   service.config.Skills.Diamond.UpperMedium,
			"advanced":      service.config.Skills.Diamond.Advanced,
			"upperAdvanced": service.config.Skills.Diamond.UpperAdvanced,
		},
	}

	RangeValueForSkills = map[string][]int{
		"elementary":    {service.config.RangeValueForSkills.MinElementary, service.config.RangeValueForSkills.MaxElementary},
		"basic":         {service.config.RangeValueForSkills.MinBasic, service.config.RangeValueForSkills.MaxBasic},
		"medium":        {service.config.RangeValueForSkills.MinMedium, service.config.RangeValueForSkills.MaxMedium},
		"upperMedium":   {service.config.RangeValueForSkills.MinUpperMedium, service.config.RangeValueForSkills.MaxUpperMedium},
		"advanced":      {service.config.RangeValueForSkills.MinAdvanced, service.config.RangeValueForSkills.MaxAdvanced},
		"upperAdvanced": {service.config.RangeValueForSkills.MinUpperAdvanced, service.config.RangeValueForSkills.MaxUpperAdvanced},
	}

	var dominantFoots = map[string]int{
		"left":  service.config.DominantFoots.Left,
		"right": service.config.DominantFoots.Right,
	}

	var isTattoo bool
	var tattoos = map[string]int{
		"gold":    service.config.Tattoos.Gold,
		"diamond": service.config.Tattoos.Diamond,
	}

	rand.Seed(time.Now().UTC().UnixNano())

	quality := searchValueByPercent(qualities)
	tactics := generateGroupSkill(skills[quality])
	physique := generateGroupSkill(skills[quality])
	technique := generateGroupSkill(skills[quality])
	offense := generateGroupSkill(skills[quality])
	defence := generateGroupSkill(skills[quality])
	goalkeeping := generateGroupSkill(skills[quality])

	if result := searchValueByPercent(tattoos); result != "" {
		isTattoo = true
	}

	card := Card{
		ID: uuid.New(),
		// TODO: change it.
		PlayerName:       "Dmytro",
		Quality:          Quality(quality),
		Height:           round(rand.Float64()*(maxHeight-minHeight)+minHeight, 0.01),
		Weight:           round(rand.Float64()*(maxWeight-minWeight)+minWeight, 0.01),
		DominantFoot:     DominantFoot(searchValueByPercent(dominantFoots)),
		IsTattoo:         isTattoo,
		Status:           StatusActive,
		Type:             TypeWon,
		UserID:           userID,
		Tactics:          tactics,
		Positioning:      generateSkill(tactics),
		Composure:        generateSkill(tactics),
		Aggression:       generateSkill(tactics),
		Vision:           generateSkill(tactics),
		Awareness:        generateSkill(tactics),
		Crosses:          generateSkill(tactics),
		Physique:         physique,
		Acceleration:     generateSkill(physique),
		RunningSpeed:     generateSkill(physique),
		ReactionSpeed:    generateSkill(physique),
		Agility:          generateSkill(physique),
		Stamina:          generateSkill(physique),
		Strength:         generateSkill(physique),
		Jumping:          generateSkill(physique),
		Balance:          generateSkill(physique),
		Technique:        technique,
		Dribbling:        generateSkill(technique),
		BallControl:      generateSkill(technique),
		WeakFoot:         generateSkill(technique),
		SkillMoves:       generateSkill(technique),
		Finesse:          generateSkill(technique),
		Curve:            generateSkill(technique),
		Volleys:          generateSkill(technique),
		ShortPassing:     generateSkill(technique),
		LongPassing:      generateSkill(technique),
		ForwardPass:      generateSkill(technique),
		Offense:          offense,
		FinishingAbility: generateSkill(offense),
		ShotPower:        generateSkill(offense),
		Accuracy:         generateSkill(offense),
		Distance:         generateSkill(offense),
		Penalty:          generateSkill(offense),
		FreeKicks:        generateSkill(offense),
		Corners:          generateSkill(offense),
		HeadingAccuracy:  generateSkill(offense),
		Defence:          defence,
		OffsideTrap:      generateSkill(defence),
		Sliding:          generateSkill(defence),
		Tackles:          generateSkill(defence),
		BallFocus:        generateSkill(defence),
		Interceptions:    generateSkill(defence),
		Vigilance:        generateSkill(defence),
		Goalkeeping:      goalkeeping,
		Reflexes:         generateSkill(goalkeeping),
		Diving:           generateSkill(goalkeeping),
		Handling:         generateSkill(goalkeeping),
		Sweeping:         generateSkill(goalkeeping),
		Throwing:         generateSkill(goalkeeping),
	}

	var err error
	if err = service.cards.Create(ctx, card); err != nil {
		return card, ErrCards.Wrap(err)
	}

	var avatar avatars.Avatar
	if avatar, err = service.avatars.GenerateAvatar(ctx, card.ID, card.IsTattoo, imageName); err != nil {
		return card, ErrCards.Wrap(err)
	}

	return card, ErrCards.Wrap(service.avatars.Create(ctx, avatar))
}

// searchValueByPercent search value string by percent.
func searchValueByPercent(generateMap map[string]int) string {
	rand := rand.Intn(99) + 1
	var sum int

	for k, v := range generateMap {
		sum += v
		if rand <= sum {
			return k
		}
	}
	return ""
}

// generateGroupSkill search value string by percent and generate assessment in the appropriate range.
func generateGroupSkill(generateMap map[string]int) int {
	skillValue := RangeValueForSkills[searchValueByPercent(generateMap)]
	difference := skillValue[1] - skillValue[0]
	rand := rand.Intn(difference) + 1
	return skillValue[0] + rand
}

// generateSkill generate assessment in the range +-10.
func generateSkill(value int) int {
	rand := rand.Intn(20) - 10
	result := value + rand
	if result < 1 {
		result = 1
	} else if result > 100 {
		result = 100
	}
	return result
}

// round rounds float64 the specified range.
func round(x, unit float64) float64 {
	return math.Round(x/unit) * unit
}

// Get returns card from DB.
func (service *Service) Get(ctx context.Context, cardID uuid.UUID) (Card, error) {
	card, err := service.cards.Get(ctx, cardID)
	return card, ErrCards.Wrap(err)
}

// List returns all cards from DB.
func (service *Service) List(ctx context.Context, cursor pagination.Cursor) (Page, error) {
	if cursor.Limit <= 0 {
		cursor.Limit = service.config.Cursor.Limit
	}
	if cursor.Page <= 0 {
		cursor.Page = service.config.Cursor.Page
	}

	cardsListPage, err := service.cards.List(ctx, cursor)
	return cardsListPage, ErrCards.Wrap(err)
}

// ListWithFilters returns all cards from DB, taking the necessary filters.
func (service *Service) ListWithFilters(ctx context.Context, filters []Filters, cursor pagination.Cursor) (Page, error) {
	var cardsListPage Page

	for _, v := range filters {
		err := v.Validate()
		if err != nil {
			return cardsListPage, err
		}
	}

	if cursor.Limit <= 0 {
		cursor.Limit = service.config.Cursor.Limit
	}
	if cursor.Page <= 0 {
		cursor.Page = service.config.Cursor.Page
	}

	cardsListPage, err := service.cards.ListWithFilters(ctx, filters, cursor)
	return cardsListPage, ErrCards.Wrap(err)
}

// ListCardIDsWithFiltersWhereActiveLot returns card ids where active lots from DB, taking the necessary filters.
func (service *Service) ListCardIDsWithFiltersWhereActiveLot(ctx context.Context, filters []Filters) ([]uuid.UUID, error) {
	for _, v := range filters {
		err := v.Validate()
		if err != nil {
			return nil, err
		}
	}
	cardsList, err := service.cards.ListCardIDsWithFiltersWhereActiveLot(ctx, filters)
	return cardsList, ErrCards.Wrap(err)
}

// ListByPlayerName returns cards from DB by player name.
func (service *Service) ListByPlayerName(ctx context.Context, filter Filters, cursor pagination.Cursor) (Page, error) {
	var cardsListPage Page
	strings.ToValidUTF8(filter.Value, "")

	// TODO: add best check
	_, err := strconv.Atoi(filter.Value)
	if err == nil {
		return cardsListPage, ErrInvalidFilter.New("%s %s", filter.Value, err)
	}

	if cursor.Limit <= 0 {
		cursor.Limit = service.config.Cursor.Limit
	}
	if cursor.Page <= 0 {
		cursor.Page = service.config.Cursor.Page
	}

	cardsListPage, err = service.cards.ListByPlayerName(ctx, filter, cursor)
	return cardsListPage, ErrCards.Wrap(err)
}

// ListCardIDsByPlayerNameWhereActiveLot returns card ids where active lot from DB by player name.
func (service *Service) ListCardIDsByPlayerNameWhereActiveLot(ctx context.Context, filter Filters) ([]uuid.UUID, error) {
	strings.ToValidUTF8(filter.Value, "")

	// TODO: add best check
	_, err := strconv.Atoi(filter.Value)
	if err == nil {
		return nil, ErrInvalidFilter.New("%s %s", filter.Value, err)
	}
	cardIdsList, err := service.cards.ListCardIDsByPlayerNameWhereActiveLot(ctx, filter)
	return cardIdsList, ErrCards.Wrap(err)
}

// ListByUserID returns all user`s cards in database.
func (service *Service) ListByUserID(ctx context.Context, userID uuid.UUID) ([]Card, error) {
	userCards, err := service.cards.ListByUserID(ctx, userID)
	return userCards, ErrCards.Wrap(err)
}

// UpdateStatus updates status of card in database.
func (service *Service) UpdateStatus(ctx context.Context, id uuid.UUID, status Status) error {
	return ErrCards.Wrap(service.cards.UpdateStatus(ctx, id, status))
}

// UpdateUserID updates user's id for card in database.
func (service *Service) UpdateUserID(ctx context.Context, id, userID uuid.UUID) error {
	return ErrCards.Wrap(service.cards.UpdateUserID(ctx, id, userID))
}

// Delete deletes card record in database.
func (service *Service) Delete(ctx context.Context, cardID uuid.UUID) error {
	return ErrCards.Wrap(service.cards.Delete(ctx, cardID))
}
