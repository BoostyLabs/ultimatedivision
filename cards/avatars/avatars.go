// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package avatars

import (
	"context"

	"github.com/google/uuid"
	"github.com/zeebo/errs"
)

// ErrNoAvatar indicated that avatar does not exist.
var ErrNoAvatar = errs.Class("avatar does not exist")

// DB is exposing access to avatars db.
//
// architecture: DB
type DB interface {
	// Create adds avatar in the data base.
	Create(ctx context.Context, avatar Avatar) error
	// Get returns avatar by id from the data base.
	Get(ctx context.Context, id uuid.UUID) (Avatar, error)
}

// Avatar entity describes the values that make up the avatar, and contains the photo itself.
type Avatar struct {
	CardID         uuid.UUID   `json:"cardId"`
	PictureType    PictureType `json:"pictureType"`
	FaceColor      int         `json:"faceColor"`
	FaceType       int         `json:"faceType"`
	EyeBrowsType   int         `json:"eyeBrowsType"`
	EyeBrowsColor  int         `json:"eyeBrowsColor"`
	HairstyleColor int         `json:"hairstyleColor"`
	HairstyleType  int         `json:"hairstyleType"`
	Nose           int         `json:"nose"`
	Tshirt         int         `json:"tshirt"`
	Beard          int         `json:"beard"`
	Lips           int         `json:"lips"`
	Tattoo         int         `json:"tattoo"`
	Image          []byte      `json:"image"`
}

// PictureType defines the list of possible type of picture.
type PictureType int

const (
	// PictureTypeFirst indicates the type of photo is the first.
	PictureTypeFirst PictureType = 1
)

// Config defines values needed by generate avatars.
type Config struct {
	PathToAvarars string `json:"pathToAvarars"`

	FaceColorFolder string `json:"faceColorFolder"`

	TattooFolder     string `json:"tattooFolder"`
	TattooTypeFolder string `json:"tattooTypeFolder"`
	TattooFile       string `json:"tattooFile"`

	FaceTypeFolder string `json:"faceTypeFolder"`
	FaceTypeFile   string `json:"faceTypeFile"`

	EyeBrowsFolder     string `json:"eyeBrowsFolder"`
	EyeBrowsTypeFolder string `json:"eyeBrowsTypeFolder"`
	EyeBrowsColorFile  string `json:"eyeBrowsColorFile"`

	HairstyleFolder      string `json:"hairstyleFolder"`
	HairstyleColorFolder string `json:"hairstyleColorFolder"`
	HairstyleTypeFile    string `json:"hairstyleTypeFile"`

	NoseFolder     string `json:"noseFolder"`
	NoseTypeFolder string `json:"noseTypeFolder"`
	NoseFile       string `json:"noseFile"`

	BeardFolder string `json:"beardFolder"`
	BeardFile   string `json:"beardFile"`

	LipsFolder string `json:"lipsFolder"`
	LipsFile   string `json:"lipsFile"`

	TshirtFolder string `json:"tshirtFolder"`
	TshirtFile   string `json:"tshirtFile"`

	PercentageFacialFeatures struct {
		Hairstyle int `json:"hairstyle"`
		Beard     int `json:"beard"`
	} `json:"percentageFacialFeatures"`
}
