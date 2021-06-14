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
	Create(ctx context.Context, card *Card) error
	Get(ctx context.Context, id uuid.UUID) (*Card, error)
	List(ctx context.Context) ([]*Card, error)
}

// Card describes card entity.
type Card struct {
	Id           uuid.UUID     `json:"id"`
	PlayerName   string        `json:"playerName"`
	Quality      Quality       `json:"quality"`
	PictureType  PictureType   `json:"pictureType"`
	Height       float32       `json:"height"`
	Weight       float32       `json:"weight"`
	SkinColor    SkinColor     `json:"skinColor"`
	HairStyle    HairStyle     `json:"hairStyle"`
	HairColor    HairColor     `json:"hairColor"`
	Accessories  []Accessories `json:"accessories"`
	DominantFoot DominantFoot  `json:"dominantFoot"`
	Tactics
	Physique
	Technique
	Offense
	Defence
	Goalkeeping
}

// Quality defines the list of possible card qualities.
type Quality string

const (
	QualityWood    Quality = "wood"
	QualityBronze  Quality = "bronze"
	QualitySilver  Quality = "silver"
	QualityGold    Quality = "gold"
	QualityDiamond Quality = "diamond"
)

// PictureType defines the list of possible card picture types.
type PictureType string

// SkinColor defines the list of possible card skin colors.
type SkinColor uint8

// HairStyle defines the list of possible card hairstyles.
type HairStyle uint8

// HairColor defines the list of possible card hair colors.
type HairColor uint8

// Accessories defines the list of possible card accessories.
type Accessories uint8

// DominantFoot defines the list of possible card dominant foots.
type DominantFoot string

const (
	Left  DominantFoot = "left"
	Right DominantFoot = "right"
)

// Tactics describes card tactic parameters.
type Tactics struct {
	Positioning uint8 `json:"positioning"`
	Composure   uint8 `json:"composure"`
	Aggression  uint8 `json:"aggression"`
	Vision      uint8 `json:"vision"`
	Awareness   uint8 `json:"awareness"`
	Crosses     uint8 `json:"crosses"`
}

// Physique describes card physique parameters.
type Physique struct {
	Acceleration  uint8 `json:"acceleration"`
	RunningSpeed  uint8 `json:"runningSpeed"`
	ReactionSpeed uint8 `json:"reactionSpeed"`
	Agility       uint8 `json:"agility"`
	Stamina       uint8 `json:"stamina"`
	Strength      uint8 `json:"strength"`
	Jumping       uint8 `json:"jumping"`
	Balance       uint8 `json:"balance"`
}

// Technique describes card technique parameters.
type Technique struct {
	Dribbling    uint8 `json:"dribbling"`
	BallControl  uint8 `json:"ballControl"`
	WeakFoot     uint8 `json:"weakFoot"`
	SkillMoves   uint8 `json:"skillMoves"`
	Finesse      uint8 `json:"finesse"`
	Curve        uint8 `json:"curve"`
	Volleys      uint8 `json:"volleys"`
	ShortPassing uint8 `json:"shortPassing"`
	LongPassing  uint8 `json:"longPassing"`
	ForwardPass  uint8 `json:"forwardPass"`
}

// Offense describes card offense parameters.
type Offense struct {
	FinishingAbility uint8 `json:"finishingAbility"`
	ShotPower        uint8 `json:"shotPower"`
	Accuracy         uint8 `json:"accuracy"`
	Distance         uint8 `json:"distance"`
	Penalty          uint8 `json:"penalty"`
	FreeKicks        uint8 `json:"freeKicks"`
	Corners          uint8 `json:"corners"`
	HeadingAccuracy  uint8 `json:"headingAccuracy"`
}

// Defence describes card defence parameters.
type Defence struct {
	OffsideTrap   uint8 `json:"offsideTrap"`
	Sliding       uint8 `json:"sliding"`
	Tackles       uint8 `json:"tackles"`
	BallFocus     uint8 `json:"ballFocus"`
	Interceptions uint8 `json:"interceptions"`
	Vigilance     uint8 `json:"vigilance"`
}

// Goalkeeping describes card goalkeeping parameters.
type Goalkeeping struct {
	Reflexes uint8 `json:"reflexes"`
	Diving   uint8 `json:"diving"`
	Handling uint8 `json:"handling"`
	Sweeping uint8 `json:"sweeping"`
	Throwing uint8 `json:"throwing"`
}
