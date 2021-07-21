// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package cards

import (
	"context"

	"github.com/google/uuid"
	"github.com/zeebo/errs"
)

// ErrNoCard indicated that card does not exist.
var ErrNoCard = errs.Class("card does not exist")

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
	// Delete deletes card record in the data base.
	Delete(ctx context.Context, id uuid.UUID) error
}

// Card describes card entity.
type Card struct {
	ID               uuid.UUID    `json:"id"`
	PlayerName       string       `json:"playerName"`
	Quality          Quality      `json:"quality"`
	PictureType      int          `json:"pictureType"`
	Height           Height       `json:"height"`
	Weight           Weight       `json:"weight"`
	SkinColor        int          `json:"skinColor"`
	HairStyle        int          `json:"hairStyle"`
	HairColor        int          `json:"hairColor"`
	Accessories      []int        `json:"accessories"`
	DominantFoot     DominantFoot `json:"dominantFoot"`
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

// Height defines the list of possible card heights.
type Height float64

const (
	// MinHeight indicates that card minimum height is 150 conventional units.
	MinHeight Height = 150.0
	// MaxHeight indicates that card maximum height is 210 conventional units.
	MaxHeight Height = 210.0
)

// Weight defines the list of possible card weights.
type Weight float64

const (
	// MinWeight indicates that card minimum weight is 50 conventional units.
	MinWeight Weight = 50.0
	// MaxWeight indicates that card maximum weight is 100 conventional units.
	MaxWeight Weight = 100.0
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

// CreateCards entitty for using generate cards.
type CreateCards struct {
	UserID              uuid.UUID
	CountCard           int
	PercentageQualities []int
}

// DominantFoots defines the list of possible percent dominant foots.
var DominantFoots = map[string]int{
	"left":  50,
	"right": 50,
}

// GroupSkills defines the list of possible group skills.
var GroupSkills = map[string][]int{
	"elementary":     {1, 20},
	"basic":          {21, 40},
	"medium":         {41, 60},
	"upper-medium":   {61, 80},
	"advanced":       {81, 90},
	"upper-advanced": {91, 100},
}

// Skills defines the list of possible percent skills.
var Skills = map[string]map[string]int{
	"wood": {
		"elementary":   30,
		"basic":        30,
		"medium":       30,
		"upper-medium": 9,
		"advanced":     1,
	},
	"silver": {
		"elementary":   15,
		"basic":        20,
		"medium":       45,
		"upper-medium": 15,
		"advanced":     5,
	},
	"gold": {
		"elementary":     5,
		"basic":          15,
		"medium":         25,
		"upper-medium":   45,
		"advanced":       8,
		"upper-advanced": 2,
	},
	"diamond": {
		"basic":          5,
		"medium":         30,
		"upper-medium":   50,
		"advanced":       10,
		"upper-advanced": 5,
	},
}
