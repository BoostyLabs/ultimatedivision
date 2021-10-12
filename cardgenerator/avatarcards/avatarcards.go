// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package avatarcards

import (
	"ultimatedivision/cards"
)

// AvatarCards describes avatar card entity.
type AvatarCards struct {
	cards.Card
	OriginalURL string `json:"originalUrl"`
}

// Config defines values needed by generate avatar cards.
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

	PercentageQualities struct {
		Wood    int `json:"wood"`
		Silver  int `json:"silver"`
		Gold    int `json:"gold"`
		Diamond int `json:"diamond"`
	} `json:"percentageQualities"`
}
