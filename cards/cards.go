// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package cards

import (
	"context"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"github.com/zeebo/errs"

	"ultimatedivision/pkg/sqlsearchoperators"
)

// ErrNoCard indicated that card does not exist.
var ErrNoCard = errs.Class("card does not exist")

// ErrCards indicated that there was an error in service.
var ErrCards = errs.Class("cards service error")

// ErrInvalidFilter indicated that filter does not valid.
var ErrInvalidFilter = errs.Class("invalid filter")

// DB is exposing access to cards db.
//
// architecture: DB
type DB interface {
	// Create add card in the data base.
	Create(ctx context.Context, card Card) error
	// Get returns card by id from the data base.
	Get(ctx context.Context, id uuid.UUID) (Card, error)
	// List returns all cards from the data base.
	List(ctx context.Context) ([]Card, error)
	// ListWithFilters returns all cards from the data base with filters.
	ListWithFilters(ctx context.Context, filters []Filters) ([]Card, error)
	// ListByPlayerName returns cards from DB by player name.
	ListByPlayerName(ctx context.Context, filters Filters) ([]Card, error)
	// UpdateStatus updates status card in the database.
	UpdateStatus(ctx context.Context, id uuid.UUID, status Status) error
	// UpdateUserID updates user id card in the database.
	UpdateUserID(ctx context.Context, id, userID uuid.UUID) error
	// Delete deletes card record in the data base.
	Delete(ctx context.Context, id uuid.UUID) error
}

// Card describes card entity.
type Card struct {
	ID               uuid.UUID    `json:"id"`
	PlayerName       string       `json:"playerName"`
	Quality          Quality      `json:"quality"`
	PictureType      int          `json:"pictureType"`
	Height           float64      `json:"height"`
	Weight           float64      `json:"weight"`
	SkinColor        int          `json:"skinColor"`
	HairStyle        int          `json:"hairStyle"`
	HairColor        int          `json:"hairColor"`
	Accessories      []int        `json:"accessories"`
	DominantFoot     DominantFoot `json:"dominantFoot"`
	IsTattoos        bool         `json:"isTattoos"`
	Status           Status       `json:"status"`
	Type             Type         `json:"type"`
	UserID           uuid.UUID    `json:"userId"`
	Tactics          int          `json:"tactics"`
	Positioning      int          `json:"positioning"`
	Composure        int          `json:"composure"`
	Aggression       int          `json:"aggression"`
	Vision           int          `json:"vision"`
	Awareness        int          `json:"awareness"`
	Crosses          int          `json:"crosses"`
	Physique         int          `json:"physique"`
	Acceleration     int          `json:"acceleration"`
	RunningSpeed     int          `json:"runningSpeed"`
	ReactionSpeed    int          `json:"reactionSpeed"`
	Agility          int          `json:"agility"`
	Stamina          int          `json:"stamina"`
	Strength         int          `json:"strength"`
	Jumping          int          `json:"jumping"`
	Balance          int          `json:"balance"`
	Technique        int          `json:"technique"`
	Dribbling        int          `json:"dribbling"`
	BallControl      int          `json:"ballControl"`
	WeakFoot         int          `json:"weakFoot"`
	SkillMoves       int          `json:"skillMoves"`
	Finesse          int          `json:"finesse"`
	Curve            int          `json:"curve"`
	Volleys          int          `json:"volleys"`
	ShortPassing     int          `json:"shortPassing"`
	LongPassing      int          `json:"longPassing"`
	ForwardPass      int          `json:"forwardPass"`
	Offense          int          `json:"offense"`
	FinishingAbility int          `json:"finishingAbility"`
	ShotPower        int          `json:"shotPower"`
	Accuracy         int          `json:"accuracy"`
	Distance         int          `json:"distance"`
	Penalty          int          `json:"penalty"`
	FreeKicks        int          `json:"freeKicks"`
	Corners          int          `json:"corners"`
	HeadingAccuracy  int          `json:"headingAccuracy"`
	Defence          int          `json:"defence"`
	OffsideTrap      int          `json:"offsideTrap"`
	Sliding          int          `json:"sliding"`
	Tackles          int          `json:"tackles"`
	BallFocus        int          `json:"ballFocus"`
	Interceptions    int          `json:"interceptions"`
	Vigilance        int          `json:"vigilance"`
	Goalkeeping      int          `json:"goalkeeping"`
	Reflexes         int          `json:"reflexes"`
	Diving           int          `json:"diving"`
	Handling         int          `json:"handling"`
	Sweeping         int          `json:"sweeping"`
	Throwing         int          `json:"throwing"`
}

// Quality defines the list of possible card qualities.
type Quality string

const (
	// QualityWood indicates that card quality is wood.
	QualityWood Quality = "wood"
	// QualitySilver indicates that card quality is silver.
	QualitySilver Quality = "silver"
	// QualityGold indicates that card quality is gold.
	QualityGold Quality = "gold"
	// QualityDiamond indicates that card quality is diamond.
	QualityDiamond Quality = "diamond"
)

// PictureType defines the list of possible card picture types.
var PictureType = map[int]string{
	1: "https://drive.google.com/file/d/1ESKPpiCoMUkOEpaa40VBFl4O1bPrDntS/view?usp=sharing",
	2: "https://drive.google.com/file/d/1baFCTjDVzIy5ucdcz-jMCb2FPSKyIRU2/view?usp=sharing",
}

// SkinColor defines the list of possible card skin colors.
var SkinColor = map[int]string{
	1: "https://drive.google.com/file/d/1ESKPpiCoMUkOEpaa40VBFl4O1bPrDntS/view?usp=sharing",
	2: "https://drive.google.com/file/d/1baFCTjDVzIy5ucdcz-jMCb2FPSKyIRU2/view?usp=sharing",
}

// HairStyle defines the list of possible card hairstyles.
var HairStyle = map[int]string{
	1: "https://drive.google.com/file/d/1ESKPpiCoMUkOEpaa40VBFl4O1bPrDntS/view?usp=sharing",
	2: "https://drive.google.com/file/d/1baFCTjDVzIy5ucdcz-jMCb2FPSKyIRU2/view?usp=sharing",
}

// HairColor defines the list of possible card hair colors.
var HairColor = map[int]string{
	1: "https://drive.google.com/file/d/1ESKPpiCoMUkOEpaa40VBFl4O1bPrDntS/view?usp=sharing",
	2: "https://drive.google.com/file/d/1baFCTjDVzIy5ucdcz-jMCb2FPSKyIRU2/view?usp=sharing",
}

// Accessory defines the list of possible card accessories.
var Accessory = map[int]string{
	1: "https://drive.google.com/file/d/1ESKPpiCoMUkOEpaa40VBFl4O1bPrDntS/view?usp=sharing",
	2: "https://drive.google.com/file/d/1baFCTjDVzIy5ucdcz-jMCb2FPSKyIRU2/view?usp=sharing",
}

// DominantFoot defines the list of possible card dominant foots.
type DominantFoot string

const (
	// DominantFootLeft indicates that dominant foot of the footballer is left.
	DominantFootLeft DominantFoot = "left"
	// DominantFootRight indicates that dominant foot of the footballer is right.
	DominantFootRight DominantFoot = "right"
)

// Status defines the list of possible card statuses.
type Status int

const (
	// StatusActive indicates that the card can be used in a team and sold.
	StatusActive Status = 0
	// StatusSale indicates that the card is sold and can't used by the team.
	StatusSale Status = 1
)

// Type defines the list of possible card Typees.
type Type string

const (
	// TypeWon indicates that the card won in a lootbox.
	TypeWon Type = "won"
	// TypeBought indicates that the card bought on the marketplaced.
	TypeBought Type = "bought"
)

// RangeValueForSkills defines the list of possible group skills.
var RangeValueForSkills = map[string][]int{}

// Config defines values needed by generate cards.
type Config struct {
	Height struct {
		Min float64 `json:"min"`
		Max float64 `json:"max"`
	} `json:"height"`

	Weight struct {
		Min float64 `json:"min"`
		Max float64 `json:"max"`
	} `json:"weight"`

	DominantFoots struct {
		Left  int `json:"left"`
		Right int `json:"right"`
	} `json:"dominantFoots"`

	Skills struct {
		Wood struct {
			Elementary  int `json:"elementary"`
			Basic       int `json:"basic"`
			Medium      int `json:"medium"`
			UpperMedium int `json:"upperMedium"`
			Advanced    int `json:"advanced"`
		} `json:"wood"`
		Silver struct {
			Elementary  int `json:"elementary"`
			Basic       int `json:"basic"`
			Medium      int `json:"medium"`
			UpperMedium int `json:"upperMedium"`
			Advanced    int `json:"advanced"`
		} `json:"silver"`
		Gold struct {
			Elementary    int `json:"elementary"`
			Basic         int `json:"basic"`
			Medium        int `json:"medium"`
			UpperMedium   int `json:"upperMedium"`
			Advanced      int `json:"advanced"`
			UpperAdvanced int `json:"upperAdvanced"`
		} `json:"gold"`
		Diamond struct {
			Basic         int `json:"basic"`
			Medium        int `json:"medium"`
			UpperMedium   int `json:"upperMedium"`
			Advanced      int `json:"advanced"`
			UpperAdvanced int `json:"upperAdvanced"`
		} `json:"diamond"`
	} `json:"skills"`

	RangeValueForSkills struct {
		MinElementary    int `json:"minElementary"`
		MaxElementary    int `json:"maxElementary"`
		MinBasic         int `json:"minBasic"`
		MaxBasic         int `json:"maxBasic"`
		MinMedium        int `json:"minMedium"`
		MaxMedium        int `json:"maxMedium"`
		MinUpperMedium   int `json:"minUpperMedium"`
		MaxUpperMedium   int `json:"maxUpperMedium"`
		MinAdvanced      int `json:"minAdvanced"`
		MaxAdvanced      int `json:"maxAdvanced"`
		MinUpperAdvanced int `json:"minUpperAdvanced"`
		MaxUpperAdvanced int `json:"maxUpperAdvanced"`
	} `json:"rangeValueForSkills"`

	Tattoos struct {
		Gold    int `json:"gold"`
		Diamond int `json:"diamond"`
	} `json:"tattoos"`
}

// PercentageQualities entity for probabilities generate cards.
type PercentageQualities struct {
	Wood    int `json:"wood"`
	Silver  int `json:"silver"`
	Gold    int `json:"gold"`
	Diamond int `json:"diamond"`
}

// Filters entity for using filter cards.
type Filters struct {
	Name           Filter
	Value          string
	SearchOperator sqlsearchoperators.SearchOperator
}

// Filter defines the list of possible filters.
type Filter string

const (
	// FilterTactics indicates filtering by card tactics.
	FilterTactics Filter = "tactics"
	// FilterPositioning indicates filtering by card positioning.
	FilterPositioning Filter = "positioning"
	// FilterComposure indicates filtering by card composure.
	FilterComposure Filter = "composure"
	// FilterAggression indicates filtering by card aggression.
	FilterAggression Filter = "aggression"
	// FilterVision indicates filtering by card vision.
	FilterVision Filter = "vision"
	// FilterAwareness indicates filtering by card awareness.
	FilterAwareness Filter = "awareness"
	// FilterCrosses indicates filtering by card crosses.
	FilterCrosses Filter = "crosses"
	// FilterPhysique indicates filtering by card physique.
	FilterPhysique Filter = "physique"
	// FilterAcceleration indicates filtering by card acceleration.
	FilterAcceleration Filter = "acceleration"
	// FilterRunningSpeed indicates filtering by card running speed.
	FilterRunningSpeed Filter = "running_speed"
	// FilterReactionSpeed indicates filtering by card reaction speed.
	FilterReactionSpeed Filter = "reaction_speed"
	// FilterAgility indicates filtering by card agility.
	FilterAgility Filter = "agility"
	// FilterStamina indicates filtering by card stamina.
	FilterStamina Filter = "stamina"
	// FilterStrength indicates filtering by card strength.
	FilterStrength Filter = "strength"
	// FilterJumping indicates filtering by card jumping.
	FilterJumping Filter = "jumping"
	// FilterBalance indicates filtering by card balance.
	FilterBalance Filter = "balance"
	// FilterTechnique indicates filtering by card technique.
	FilterTechnique Filter = "technique"
	// FilterDribbling indicates filtering by card dribbling.
	FilterDribbling Filter = "dribbling"
	// FilterBallControl indicates filtering by card ball control.
	FilterBallControl Filter = "ball_control"
	// FilterWeakFoot indicates filtering by card weak foot.
	FilterWeakFoot Filter = "weak_foot"
	// FilterSkillMoves indicates filtering by card skill moves.
	FilterSkillMoves Filter = "skill_moves"
	// FilterFinesse indicates filtering by card finesse.
	FilterFinesse Filter = "finesse"
	// FilterCurve indicates filtering by card curve.
	FilterCurve Filter = "curve"
	// FilterVolleys indicates filtering by card volleys.
	FilterVolleys Filter = "volleys"
	// FilterShortPassing indicates filtering by card short passing.
	FilterShortPassing Filter = "short_passing"
	// FilterLongPassing indicates filtering by card long passing.
	FilterLongPassing Filter = "long_passing"
	// FilterForwardPass indicates filtering by card forward pass.
	FilterForwardPass Filter = "forward_pass"
	// FilterOffense indicates filtering by card offense.
	FilterOffense Filter = "offense"
	// FilterFinishingAbility indicates filtering by card finishing ability.
	FilterFinishingAbility Filter = "finishing_ability"
	// FilterShotPower indicates filtering by card shot power.
	FilterShotPower Filter = "shot_power"
	// FilterAccuracy indicates filtering by card accuracy.
	FilterAccuracy Filter = "accuracy"
	// FilterDistance indicates filtering by card distance.
	FilterDistance Filter = "distance"
	// FilterPenalty indicates filtering by card penalty.
	FilterPenalty Filter = "penalty"
	// FilterFreeKicks indicates filtering by card free kicks.
	FilterFreeKicks Filter = "free_kicks"
	// FilterCorners indicates filtering by card corners.
	FilterCorners Filter = "corners"
	// FilterHeadingAccuracy indicates filtering by card heading accuracy.
	FilterHeadingAccuracy Filter = "heading_accuracy"
	// FilterDefence indicates filtering by card defence.
	FilterDefence Filter = "defence"
	// FilterOffsideTrap indicates filtering by card offside trap.
	FilterOffsideTrap Filter = "offside_trap"
	// FilterSliding indicates filtering by card sliding.
	FilterSliding Filter = "sliding"
	// FilterTackles indicates filtering by card tackles.
	FilterTackles Filter = "tackles"
	// FilterBallFocus indicates filtering by card ball focus.
	FilterBallFocus Filter = "ball_focus"
	// FilterInterceptions indicates filtering by card interceptions.
	FilterInterceptions Filter = "interceptions"
	// FilterVigilance indicates filtering by card vigilance.
	FilterVigilance Filter = "vigilance"
	// FilterGoalkeeping indicates filtering by card goalkeeping.
	FilterGoalkeeping Filter = "goalkeeping"
	// FilterReflexes indicates filtering by card reflexes.
	FilterReflexes Filter = "reflexes"
	// FilterDiving indicates filtering by card diving.
	FilterDiving Filter = "diving"
	// FilterHandling indicates filtering by card handling.
	FilterHandling Filter = "handling"
	// FilterSweeping indicates filtering by card sweeping.
	FilterSweeping Filter = "sweeping"
	// FilterThrowing indicates filtering by card throwing.
	FilterThrowing Filter = "throwing"
	// FilterQuality indicates filtering by card quality.
	FilterQuality Filter = "quality"
	// FilterHeight indicates filtering by card height.
	FilterHeight Filter = "height"
	// FilterWeight indicates filtering by card weight.
	FilterWeight Filter = "weight"
	// FilterDominantFoot indicates filtering by card dominant foot.
	FilterDominantFoot Filter = "dominant_foot"
	// FilterType indicates filtering by card type.
	FilterType Filter = "type"
	// FilterPrice indicates filtering by card price.
	FilterPrice Filter = "price"
	// FilterPlayerName indicates filtering by card player name.
	FilterPlayerName Filter = "player_name"
)

// Validate check of valid UTF-8 bytes and type.
func (f Filters) Validate() error {
	if f.Name == FilterTactics || f.Name == FilterPositioning || f.Name == FilterComposure || f.Name == FilterAggression ||
		f.Name == FilterVision || f.Name == FilterAwareness || f.Name == FilterCrosses || f.Name == FilterPhysique ||
		f.Name == FilterAcceleration || f.Name == FilterRunningSpeed || f.Name == FilterReactionSpeed || f.Name == FilterAgility ||
		f.Name == FilterStamina || f.Name == FilterStrength || f.Name == FilterJumping || f.Name == FilterBalance ||
		f.Name == FilterTechnique || f.Name == FilterDribbling || f.Name == FilterBallControl || f.Name == FilterWeakFoot ||
		f.Name == FilterSkillMoves || f.Name == FilterFinesse || f.Name == FilterCurve || f.Name == FilterVolleys ||
		f.Name == FilterShortPassing || f.Name == FilterLongPassing || f.Name == FilterForwardPass || f.Name == FilterOffense ||
		f.Name == FilterFinishingAbility || f.Name == FilterShotPower || f.Name == FilterAccuracy || f.Name == FilterDistance ||
		f.Name == FilterPenalty || f.Name == FilterFreeKicks || f.Name == FilterCorners || f.Name == FilterHeadingAccuracy ||
		f.Name == FilterDefence || f.Name == FilterOffsideTrap || f.Name == FilterSliding || f.Name == FilterTackles ||
		f.Name == FilterBallFocus || f.Name == FilterInterceptions || f.Name == FilterVigilance || f.Name == FilterGoalkeeping ||
		f.Name == FilterReflexes || f.Name == FilterDiving || f.Name == FilterHandling || f.Name == FilterSweeping || f.Name == FilterThrowing {
		strings.ToValidUTF8(f.Value, "")

		_, err := strconv.Atoi(f.Value)
		if err != nil {
			return ErrInvalidFilter.New("%s %s", f.Value, err)
		}
		return nil
	}

	if f.Name == FilterHeight || f.Name == FilterWeight || f.Name == FilterPrice {
		strings.ToValidUTF8(f.Value, "")

		_, err := strconv.ParseFloat(f.Value, 64)
		if err != nil {
			return ErrInvalidFilter.New("%s %s", f.Value, err)
		}
		return nil
	}

	if f.Name == FilterQuality {
		strings.ToValidUTF8(f.Value, "")

		if f.SearchOperator != sqlsearchoperators.EQ {
			return ErrInvalidFilter.New("'%s' not suitable for %s", f.SearchOperator, f.Name)
		}

		quality := Quality(f.Value)
		if quality == QualityWood || quality == QualitySilver || quality == QualityGold || quality == QualityDiamond {
			return nil
		}
		return ErrInvalidFilter.New("%s %s", f.Value, "is not an indicator of quality card")
	}

	if f.Name == FilterDominantFoot {
		strings.ToValidUTF8(f.Value, "")

		if f.SearchOperator != sqlsearchoperators.EQ {
			return ErrInvalidFilter.New("'%s' not suitable for %s", f.SearchOperator, f.Name)
		}

		dominantFoot := DominantFoot(f.Value)
		if dominantFoot == DominantFootLeft || dominantFoot == DominantFootRight {
			return nil
		}
		return ErrInvalidFilter.New("%s %s", f.Value, "is not an indicator of dominant foot card")
	}

	if f.Name == FilterType {
		strings.ToValidUTF8(f.Value, "")

		if f.SearchOperator != sqlsearchoperators.EQ {
			return ErrInvalidFilter.New("'%s' not suitable for %s", f.SearchOperator, f.Name)
		}

		filterType := Type(f.Value)
		if filterType == TypeWon || filterType == TypeBought {
			return nil
		}
		return ErrInvalidFilter.New("%s %s", f.Value, "is not an indicator of type card")
	}

	return ErrInvalidFilter.New("invalid name parameter - %s", f.Name)
}
