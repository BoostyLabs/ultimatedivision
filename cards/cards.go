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
	Create(ctx context.Context, card Card) error
	Get(ctx context.Context, id uuid.UUID) (Card, error)
	List(ctx context.Context) ([]Card, error)
}

// Card describes card entity.
type Card struct {
	ID               uuid.UUID    `json:"id"`
	PlayerName       string       `json:"playerName"`
	Quality          Quality      `json:"quality"`
	PictureType      int          `json:"pictureType"`
	Height           float32      `json:"height"`
	Weight           float32      `json:"weight"`
	SkinColor        int          `json:"skinColor"`
	HairStyle        int          `json:"hairStyle"`
	HairColor        int          `json:"hairColor"`
	Accessories      []int        `json:"accessories"`
	DominantFoot     DominantFoot `json:"dominantFoot"`
	UserId           uuid.UUID    `json:"userId"`
	ImageId          []byte       `json:"imageId"`
	Positioning      int          `json:"positioning"`
	Composure        int          `json:"composure"`
	Aggression       int          `json:"aggression"`
	Vision           int          `json:"vision"`
	Awareness        int          `json:"awareness"`
	Crosses          int          `json:"crosses"`
	Acceleration     int          `json:"acceleration"`
	RunningSpeed     int          `json:"runningSpeed"`
	ReactionSpeed    int          `json:"reactionSpeed"`
	Agility          int          `json:"agility"`
	Stamina          int          `json:"stamina"`
	Strength         int          `json:"strength"`
	Jumping          int          `json:"jumping"`
	Balance          int          `json:"balance"`
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
	FinishingAbility int          `json:"finishingAbility"`
	ShotPower        int          `json:"shotPower"`
	Accuracy         int          `json:"accuracy"`
	Distance         int          `json:"distance"`
	Penalty          int          `json:"penalty"`
	FreeKicks        int          `json:"freeKicks"`
	Corners          int          `json:"corners"`
	HeadingAccuracy  int          `json:"headingAccuracy"`
	OffsideTrap      int          `json:"offsideTrap"`
	Sliding          int          `json:"sliding"`
	Tackles          int          `json:"tackles"`
	BallFocus        int          `json:"ballFocus"`
	Interceptions    int          `json:"interceptions"`
	Vigilance        int          `json:"vigilance"`
	Reflexes         int          `json:"reflexes"`
	Diving           int          `json:"diving"`
	Handling         int          `json:"handling"`
	Sweeping         int          `json:"sweeping"`
	Throwing         int          `json:"throwing"`
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

// PictureType describes picture typy entity.
type PictureType struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// SkinColor describes skin color entity.
type SkinColor struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// HairStyle describes hairstyle entity.
type HairStyle struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// HairColor describes hair color entity.
type HairColor struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Accessory describes accessory entity.
type Accessory struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Accessory describes accessory entity.
type Image struct {
	Img  []byte `json:"img"`
	Name string `json:"name"`
}

// DominantFoot defines the list of possible card dominant foots.
type DominantFoot string

const (
	Left  DominantFoot = "left"
	Right DominantFoot = "right"
)
