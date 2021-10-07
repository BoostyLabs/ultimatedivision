// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package avatars

import (
	"context"
	"fmt"
	"image"

	"github.com/google/uuid"
	"github.com/nfnt/resize"
	"github.com/zeebo/errs"

	"ultimatedivision/pkg/imageprocessing"
	"ultimatedivision/pkg/rand"
)

// ErrAvatar indicated that there was an error in service.
var ErrAvatar = errs.Class("avatar service error")

// Service is handling avatars related logic.
//
// architecture: Service
type Service struct {
	avatars DB
	config  Config
}

// NewService is a constructor for avatars service.
func NewService(avatars DB, config Config) *Service {
	return &Service{
		config:  config,
		avatars: avatars,
	}
}

// Create adds avatar in DB.
func (service *Service) Create(ctx context.Context, avatar Avatar) error {
	return ErrAvatar.Wrap(service.avatars.Create(ctx, avatar))
}

// GenerateAvatar generates a common avatar from different layers of photos.
func (service *Service) GenerateAvatar(ctx context.Context, cardID uuid.UUID, isTattoo bool) (Avatar, error) {
	var (
		layer  image.Image
		layers []image.Image
		count  int
		err    error
	)

	avatar := Avatar{
		CardID:      cardID,
		PictureType: PictureTypeFirst,
	}

	// FaceColor
	if count, err = imageprocessing.SearchCountFiles(service.config.PathToAvararsComponents, service.config.FaceColorFolder); err != nil {
		return avatar, ErrNoAvatarFile.Wrap(err)
	}
	if avatar.FaceColor, err = rand.RandomInRange(count); err != nil {
		return avatar, ErrAvatar.Wrap(err)
	}

	// FaceType
	pathToFaceColor := service.config.PathToAvararsComponents + "/" + fmt.Sprintf(service.config.FaceColorFolder, avatar.FaceColor)
	if count, err = imageprocessing.SearchCountFiles(pathToFaceColor, service.config.FaceTypeFolder); err != nil {
		return avatar, ErrNoAvatarFile.Wrap(err)
	}
	if avatar.FaceType, err = rand.RandomInRange(count); err != nil {
		return avatar, ErrAvatar.Wrap(err)
	}

	pathToFaceType := pathToFaceColor + "/" + fmt.Sprintf(service.config.FaceTypeFolder, avatar.FaceType)
	if layer, err = imageprocessing.CreateLayer(pathToFaceType, fmt.Sprintf(service.config.FaceTypeFile, avatar.FaceType)); err != nil {
		return avatar, ErrNoAvatarFile.Wrap(err)
	}
	layers = append(layers, layer)

	// NoseType
	pathToNoseType := pathToFaceType + "/" + service.config.NoseFolder
	if count, err = imageprocessing.SearchCountFiles(pathToNoseType, service.config.NoseTypeFolder); err != nil {
		return avatar, ErrNoAvatarFile.Wrap(err)
	}
	if avatar.Nose, err = rand.RandomInRange(count); err != nil {
		return avatar, ErrAvatar.Wrap(err)
	}

	pathToNoseType += "/" + fmt.Sprintf(service.config.NoseTypeFolder, avatar.Nose)
	if layer, err = imageprocessing.CreateLayer(pathToNoseType, fmt.Sprintf(service.config.NoseFile, avatar.Nose)); err != nil {
		return avatar, ErrNoAvatarFile.Wrap(err)
	}
	layers = append(layers, layer)

	// LipsType
	pathToLipsType := pathToNoseType + "/" + service.config.LipsFolder
	if count, err = imageprocessing.SearchCountFiles(pathToLipsType, service.config.LipsFile); err != nil {
		return avatar, ErrNoAvatarFile.Wrap(err)
	}
	if avatar.Lips, err = rand.RandomInRange(count); err != nil {
		return avatar, ErrAvatar.Wrap(err)
	}

	if layer, err = imageprocessing.CreateLayer(pathToLipsType, fmt.Sprintf(service.config.LipsFile, avatar.Lips)); err != nil {
		return avatar, ErrNoAvatarFile.Wrap(err)
	}
	layers = append(layers, layer)

	// EyeBrowsType
	pathToEyeBrowsType := pathToFaceType + "/" + service.config.EyeBrowsFolder
	if count, err = imageprocessing.SearchCountFiles(pathToEyeBrowsType, service.config.EyeBrowsTypeFolder); err != nil {
		return avatar, ErrNoAvatarFile.Wrap(err)
	}
	if avatar.EyeBrowsType, err = rand.RandomInRange(count); err != nil {
		return avatar, ErrAvatar.Wrap(err)
	}

	// EyeBrowsColor
	pathToBrowsColor := pathToEyeBrowsType + "/" + fmt.Sprintf(service.config.EyeBrowsTypeFolder, avatar.EyeBrowsType)
	if count, err = imageprocessing.SearchCountFiles(pathToBrowsColor, service.config.EyeBrowsColorFile); err != nil {
		return avatar, ErrNoAvatarFile.Wrap(err)
	}
	if avatar.EyeBrowsColor, err = rand.RandomInRange(count); err != nil {
		return avatar, ErrAvatar.Wrap(err)
	}

	if layer, err = imageprocessing.CreateLayer(pathToBrowsColor, fmt.Sprintf(service.config.EyeBrowsColorFile, avatar.EyeBrowsColor)); err != nil {
		return avatar, ErrNoAvatarFile.Wrap(err)
	}
	layers = append(layers, layer)

	// Tattoo
	if isTattoo {
		pathToTattoo := service.config.PathToAvararsComponents + "/" + service.config.TattooFolder + "/" + fmt.Sprintf(service.config.TattooTypeFolder, avatar.FaceType)
		if count, err = imageprocessing.SearchCountFiles(pathToTattoo, service.config.TattooFile); err != nil {
			return avatar, ErrNoAvatarFile.Wrap(err)
		}
		if avatar.Tattoo, err = rand.RandomInRange(count); err != nil {
			return avatar, ErrAvatar.Wrap(err)
		}

		if layer, err = imageprocessing.CreateLayer(pathToTattoo, fmt.Sprintf(service.config.TattooFile, avatar.Tattoo)); err != nil {
			return avatar, ErrNoAvatarFile.Wrap(err)
		}
		layers = append(layers, layer)
	}

	// Hairstyles
	if rand.IsIncludeRange(service.config.PercentageFacialFeatures.Hairstyle) {
		// HairstylesColor
		pathToHairstylesColor := pathToFaceType + "/" + service.config.HairstyleFolder
		if count, err = imageprocessing.SearchCountFiles(pathToHairstylesColor, service.config.HairstyleColorFolder); err != nil {
			return avatar, ErrNoAvatarFile.Wrap(err)
		}
		if avatar.HairstyleColor, err = rand.RandomInRange(count); err != nil {
			return avatar, ErrAvatar.Wrap(err)
		}

		// HairstylesType
		pathToHairstylesType := pathToHairstylesColor + "/" + fmt.Sprintf(service.config.HairstyleColorFolder, avatar.HairstyleColor)
		if count, err = imageprocessing.SearchCountFiles(pathToHairstylesType, service.config.HairstyleTypeFile); err != nil {
			return avatar, ErrNoAvatarFile.Wrap(err)
		}
		if avatar.HairstyleType, err = rand.RandomInRange(count); err != nil {
			return avatar, ErrAvatar.Wrap(err)
		}

		if layer, err = imageprocessing.CreateLayer(pathToHairstylesType, fmt.Sprintf(service.config.HairstyleTypeFile, avatar.HairstyleType)); err != nil {
			return avatar, ErrNoAvatarFile.Wrap(err)
		}
		layers = append(layers, layer)
	}

	// BeardType
	if rand.IsIncludeRange(service.config.PercentageFacialFeatures.Beard) {
		pathToBeardType := pathToNoseType + "/" + service.config.BeardFolder
		if count, err = imageprocessing.SearchCountFiles(pathToBeardType, service.config.BeardFile); err != nil {
			return avatar, ErrNoAvatarFile.Wrap(err)
		}
		if avatar.Beard, err = rand.RandomInRange(count); err != nil {
			return avatar, ErrAvatar.Wrap(err)
		}

		if layer, err = imageprocessing.CreateLayer(pathToBeardType, fmt.Sprintf(service.config.BeardFile, avatar.Beard)); err != nil {
			return avatar, ErrNoAvatarFile.Wrap(err)
		}
		layers = append(layers, layer)
	}

	// T-shirtType
	pathToTshirtType := pathToFaceType + "/" + service.config.TshirtFolder
	if count, err = imageprocessing.SearchCountFiles(pathToTshirtType, service.config.TshirtFile); err != nil {
		return avatar, ErrNoAvatarFile.Wrap(err)
	}
	if avatar.Tshirt, err = rand.RandomInRange(count); err != nil {
		return avatar, ErrAvatar.Wrap(err)
	}

	if layer, err = imageprocessing.CreateLayer(pathToTshirtType, fmt.Sprintf(service.config.TshirtFile, avatar.Tshirt)); err != nil {
		return avatar, ErrNoAvatarFile.Wrap(err)
	}
	layers = append(layers, layer)

	// EyeLaserType
	if rand.IsIncludeRange(service.config.PercentageFacialFeatures.EyeLaser) {
		pathToEyeLaserType := pathToFaceType + "/" + service.config.EyeLaserFolder
		if count, err = imageprocessing.SearchCountFiles(pathToEyeLaserType, service.config.EyeLaserTypeFolder); err != nil {
			return avatar, ErrNoAvatarFile.Wrap(err)
		}
		if avatar.EyeLaserType, err = rand.RandomInRange(count); err != nil {
			return avatar, ErrAvatar.Wrap(err)
		}

		pathToEyeLaserType += "/" + fmt.Sprintf(service.config.EyeLaserTypeFolder, avatar.EyeLaserType)
		if layer, err = imageprocessing.CreateLayer(pathToEyeLaserType, fmt.Sprintf(service.config.EyeLaserTypeFile, avatar.EyeLaserType)); err != nil {
			return avatar, ErrNoAvatarFile.Wrap(err)
		}
		layers = append(layers, layer)
	}

	originalImage := imageprocessing.Layering(layers)
	previewImage := resize.Resize(uint(service.config.SizePreviewImage.Width), uint(service.config.SizePreviewImage.Height), originalImage, resize.Lanczos3)

	avatar.OriginalURL = service.config.PathToOutputAvatars + "/" + avatar.CardID.String() + "_" + FormatImageOriginal + "." + TypeImagePNG
	if err = imageprocessing.SaveImage(avatar.OriginalURL, originalImage); err != nil {
		return avatar, ErrAvatar.Wrap(err)
	}

	avatar.PreviewURL = service.config.PathToOutputAvatars + "/" + avatar.CardID.String() + "_" + FormatImagePreview + "." + TypeImagePNG
	if err = imageprocessing.SaveImage(avatar.PreviewURL, previewImage); err != nil {
		return avatar, ErrAvatar.Wrap(err)
	}

	return avatar, nil
}

// Get returns avatar from DB.
func (service *Service) Get(ctx context.Context, cardID uuid.UUID) (Avatar, error) {
	avatar, err := service.avatars.Get(ctx, cardID)
	if err != nil {
		return Avatar{}, ErrAvatar.Wrap(err)
	}
	return avatar, ErrAvatar.Wrap(err)
}
