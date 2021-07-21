// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package cards

import (
	"context"
	"math"
	"math/rand"
	"time"

	"github.com/google/uuid"
)

// Service is handling cards related logic.
//
// architecture: Service
type Service struct {
	cards DB
}

// NewService is a constructor for cards service.
func NewService(cards DB) *Service {
	return &Service{
		cards: cards,
	}
}

// Create add card in DB.
func (service *Service) Create(ctx context.Context, createCards CreateCards) error {

	Qualities := map[string]int{
		"wood":    createCards.PercentageQualities[0],
		"silver":  createCards.PercentageQualities[1],
		"gold":    createCards.PercentageQualities[2],
		"diamond": createCards.PercentageQualities[3],
	}
	rand.Seed(time.Now().UTC().UnixNano())

	quality := searchValueByPercent(Qualities)
	tactics := generateGroupSkill(Skills[quality])
	physique := generateGroupSkill(Skills[quality])
	technique := generateGroupSkill(Skills[quality])
	offense := generateGroupSkill(Skills[quality])
	defence := generateGroupSkill(Skills[quality])
	goalkeeping := generateGroupSkill(Skills[quality])

	card := Card{
		ID:               uuid.New(),
		PlayerName:       "Dmytro",
		Quality:          Quality(quality),
		PictureType:      1,
		Height:           Height(Round(rand.Float64()*(float64(MaxHeight)-float64(MinHeight))+float64(MinHeight), 0.01)),
		Weight:           Weight(Round(rand.Float64()*(float64(MaxWeight)-float64(MinWeight))+float64(MinWeight), 0.01)),
		SkinColor:        1,
		HairStyle:        1,
		HairColor:        1,
		Accessories:      []int{1, 2},
		DominantFoot:     DominantFoot(searchValueByPercent(DominantFoots)),
		UserID:           createCards.UserID,
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
	return service.cards.Create(ctx, card)
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
	skillValue := GroupSkills[searchValueByPercent(generateMap)]
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

// Round rounds float64 the specified range.
func Round(x, unit float64) float64 {
	return math.Round(x/unit) * unit
}

// Get returns card from DB.
func (service *Service) Get(ctx context.Context, cardID uuid.UUID) (Card, error) {
	return service.cards.Get(ctx, cardID)
}

// List returns all cards from DB.
func (service *Service) List(ctx context.Context) ([]Card, error) {
	return service.cards.List(ctx)
}

// Delete destroy card in DB.
func (service *Service) Delete(ctx context.Context, cardID uuid.UUID) error {
	return service.cards.Delete(ctx, cardID)
}
