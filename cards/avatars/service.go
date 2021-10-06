// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package avatars

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"io/ioutil"
	"math/rand"
	"os"
	"regexp"

	"github.com/google/uuid"
	"github.com/zeebo/errs"
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
func (service *Service) Create(ctx context.Context, cardID uuid.UUID, isTattoo bool) (Avatar, error) {
	var err error
	avatar := Avatar{
		CardID:      cardID,
		PictureType: PictureTypeFirst,
	}

	// FaceColor
	pathToFaceColor := service.config.PathToAvararsComponents
	nameFile := service.config.FaceColorFolder
	if avatar.FaceColor, err = randomNumber(pathToFaceColor, nameFile); err != nil {
		return Avatar{}, ErrAvatar.New("search random number in %s, error - %s", nameFile, err)
	}

	// FaceType
	pathToFaceType := pathToFaceColor + "/" + fmt.Sprintf(nameFile, avatar.FaceColor)
	nameFile = service.config.FaceTypeFolder
	if avatar.FaceType, err = randomNumber(pathToFaceType, nameFile); err != nil {
		return Avatar{}, ErrAvatar.New("search random number in %s, error - %s", nameFile, err)
	}
	pathToFaceType += "/" + fmt.Sprintf(nameFile, avatar.FaceType)

	// Tattoo
	if isTattoo {
		pathToTattooType := pathToFaceColor + "/" + service.config.TattooFolder + "/" + fmt.Sprintf(service.config.TattooTypeFolder, avatar.FaceType)
		nameFile = service.config.TattooFile
		if avatar.Tattoo, err = randomNumber(pathToTattooType, nameFile); err != nil {
			return Avatar{}, ErrAvatar.New("search random number in %s, error - %s", nameFile, err)
		}
	}

	// EyeBrowsType
	pathToEyeBrowsType := pathToFaceType + "/" + service.config.EyeBrowsFolder
	nameFile = service.config.EyeBrowsTypeFolder
	if avatar.EyeBrowsType, err = randomNumber(pathToEyeBrowsType, nameFile); err != nil {
		return Avatar{}, ErrAvatar.New("search random number in %s, error - %s", nameFile, err)
	}

	// EyeBrowsColor
	pathToBrowsColor := pathToEyeBrowsType + "/" + fmt.Sprintf(nameFile, avatar.EyeBrowsType)
	nameFile = service.config.EyeBrowsColorFile
	if avatar.EyeBrowsColor, err = randomNumber(pathToBrowsColor, nameFile); err != nil {
		return Avatar{}, ErrAvatar.New("search random number in %s, error - %s", nameFile, err)
	}

	// EyeLaserType
	pathToEyeLaserType := pathToFaceType + "/" + service.config.EyeLaserFolder
	nameFile = service.config.EyeLaserTypeFolder
	if avatar.EyeBrowsType, err = randomNumber(pathToEyeLaserType, nameFile); err != nil {
		return Avatar{}, ErrAvatar.New("search random number in %s, error - %s", nameFile, err)
	}

	if (rand.Intn(99) + 1) <= service.config.PercentageFacialFeatures.Hairstyle {
		// HairstylesColor
		pathToHairstylesColor := pathToFaceType + "/" + service.config.HairstyleFolder
		nameFile = service.config.HairstyleColorFolder
		if avatar.HairstyleColor, err = randomNumber(pathToHairstylesColor, nameFile); err != nil {
			return Avatar{}, ErrAvatar.New("search random number in %s, error - %s", nameFile, err)
		}

		// HairstylesType
		pathToHairstylesType := pathToHairstylesColor + "/" + fmt.Sprintf(nameFile, avatar.HairstyleColor)
		nameFile = service.config.HairstyleTypeFile
		if avatar.HairstyleType, err = randomNumber(pathToHairstylesType, nameFile); err != nil {
			return Avatar{}, ErrAvatar.New("search random number in %s, error - %s", nameFile, err)
		}
	}

	// NoseType
	pathToNoseType := pathToFaceType + "/" + service.config.NoseFolder
	nameFile = service.config.NoseTypeFolder
	if avatar.Nose, err = randomNumber(pathToNoseType, nameFile); err != nil {
		return Avatar{}, ErrAvatar.New("search random number in %s, error - %s", nameFile, err)
	}
	pathToNoseType += "/" + fmt.Sprintf(nameFile, avatar.Nose)

	// BeardType
	if (rand.Intn(99) + 1) <= service.config.PercentageFacialFeatures.Beard {
		pathToBeardType := pathToNoseType + "/" + service.config.BeardFolder
		nameFile = service.config.BeardFile
		if avatar.Beard, err = randomNumber(pathToBeardType, nameFile); err != nil {
			return Avatar{}, ErrAvatar.New("search random number in %s, error - %s", nameFile, err)
		}
	}

	// LipsType
	pathToLipsType := pathToNoseType + "/" + service.config.LipsFolder
	nameFile = service.config.LipsFile
	if avatar.Lips, err = randomNumber(pathToLipsType, nameFile); err != nil {
		return Avatar{}, ErrAvatar.New("search random number in %s, error - %s", nameFile, err)
	}

	// T-shirtType
	pathToTshirtType := pathToFaceType + "/" + service.config.TshirtFolder
	nameFile = service.config.TshirtFile
	if avatar.Tshirt, err = randomNumber(pathToTshirtType, nameFile); err != nil {
		return Avatar{}, ErrAvatar.New("search random number in %s, error - %s", nameFile, err)
	}

	if avatar.Image, err = avatar.generateAvatar(service.config); err != nil {
		return Avatar{}, ErrAvatar.Wrap(err)

	}

	return avatar, ErrAvatar.Wrap(service.avatars.Create(ctx, avatar))
}

// randomNumber randomizes numbers in the specified range.
func randomNumber(pathToAvararsCPathToAvararsComponents, nameFile string) (int, error) {
	count, err := searchCountFiles(pathToAvararsCPathToAvararsComponents, nameFile)
	if err != nil {
		return 0, ErrAvatar.Wrap(err)
	}

	switch {
	case count == 1:
		return 1, nil
	case count > 0:
		return rand.Intn(count-1) + 1, nil
	default:
		return 0, ErrNoAvatarFile.New(nameFile)
	}
}

// searchCountFiles searches count files in the specified path and by name of file.
func searchCountFiles(pathToAvararsCPathToAvararsComponents, nameFile string) (int, error) {
	files, _ := ioutil.ReadDir(pathToAvararsCPathToAvararsComponents)

	var count int
	for _, file := range files {
		isMatched, err := regexp.Match(fmt.Sprintf(nameFile, `\d`), []byte(file.Name()))
		if err != nil {
			return 0, err
		}
		if isMatched {
			count++
		}
	}

	return count, nil
}

// generateAvatar generates a common avatar from different layers of photos.
func (avatar *Avatar) generateAvatar(config Config) (string, error) {
	var (
		faceTypeLayer, tattooLayer, eyeBrowsLayer, eyeLaserLayer, hairstylesLayer, noseLayer, beardLayer, lipsLayer, tshirtLayer image.Image
	)

	// FaceColor
	pathToFaceColor := config.PathToAvararsComponents + "/" + fmt.Sprintf(config.FaceColorFolder, avatar.FaceColor)

	// FaceType
	pathToFaceType := pathToFaceColor + "/" + fmt.Sprintf(config.FaceTypeFolder, avatar.FaceType)
	nameFaceTypeFile := fmt.Sprintf(config.FaceTypeFile, avatar.FaceType)
	faceTypeFile, err := os.Open(pathToFaceType + "/" + nameFaceTypeFile)
	if err != nil {
		return "", ErrNoAvatarFile.New("generate avatar with %s, error - %s", nameFaceTypeFile, err)

	}
	faceTypeLayer, err = png.Decode(faceTypeFile)
	if err != nil {
		return "", ErrAvatar.Wrap(err)
	}
	defer func() {
		err = errs.Combine(err, ErrAvatar.Wrap(faceTypeFile.Close()))
	}()

	// Tattoo
	if avatar.Tattoo != 0 {
		pathToTattoo := config.PathToAvararsComponents + "/" + config.TattooFolder + "/" + fmt.Sprintf(config.TattooTypeFolder, avatar.Tattoo)
		nameTattooFile := fmt.Sprintf(config.TattooFile, avatar.Tattoo)
		tattooFile, err := os.Open(pathToTattoo + "/" + nameTattooFile)
		if err != nil {
			return "", ErrNoAvatarFile.New("generate avatar with %s, error - %s", nameTattooFile, err)
		}
		tattooLayer, err = png.Decode(tattooFile)
		if err != nil {
			return "", ErrAvatar.Wrap(err)
		}
		defer func() {
			err = errs.Combine(err, ErrAvatar.Wrap(tattooFile.Close()))
		}()
	}

	// EyeBrows
	pathToEyeBrows := pathToFaceType + "/" + config.EyeBrowsFolder + "/" + fmt.Sprintf(config.EyeBrowsTypeFolder, avatar.EyeBrowsType)
	nameEyeBrowsColorFile := fmt.Sprintf(config.EyeBrowsColorFile, avatar.EyeBrowsColor)
	eyeBrowsFile, err := os.Open(pathToEyeBrows + "/" + nameEyeBrowsColorFile)
	if err != nil {
		return "", ErrNoAvatarFile.New("generate avatar with %s, error - %s", nameEyeBrowsColorFile, err)
	}
	eyeBrowsLayer, err = png.Decode(eyeBrowsFile)
	if err != nil {
		return "", ErrAvatar.Wrap(err)
	}
	defer func() {
		err = errs.Combine(err, ErrAvatar.Wrap(eyeBrowsFile.Close()))
	}()

	// EyeLaser
	pathToEyeLaser := pathToFaceType + "/" + config.EyeLaserFolder + "/" + fmt.Sprintf(config.EyeLaserTypeFolder, avatar.EyeLaserType)
	nameEyeLaserTypeFile := fmt.Sprintf(config.EyeLaserTypeFile, avatar.EyeLaserType)
	eyeLaserFile, err := os.Open(pathToEyeLaser + "/" + nameEyeLaserTypeFile)
	if err != nil {
		return "", ErrNoAvatarFile.New("generate avatar with %s, error - %s", nameEyeLaserTypeFile, err)
	}
	eyeLaserLayer, err = png.Decode(eyeLaserFile)
	if err != nil {
		return "", ErrAvatar.Wrap(err)
	}
	defer func() {
		err = errs.Combine(err, ErrAvatar.Wrap(eyeLaserFile.Close()))
	}()

	// Hairstyles
	if avatar.HairstyleColor != 0 {
		pathToHairstyles := pathToFaceType + "/" + config.HairstyleFolder + "/" + fmt.Sprintf(config.HairstyleColorFolder, avatar.HairstyleColor)
		nameHairstylesTypeFile := fmt.Sprintf(config.HairstyleTypeFile, avatar.HairstyleType)
		hairstylesTypeFile, err := os.Open(pathToHairstyles + "/" + nameHairstylesTypeFile)
		if err != nil {
			return "", ErrNoAvatarFile.New("generate avatar with %s, error - %s", nameHairstylesTypeFile, err)
		}
		hairstylesLayer, err = png.Decode(hairstylesTypeFile)
		if err != nil {
			return "", ErrAvatar.Wrap(err)
		}
		defer func() {
			err = errs.Combine(err, ErrAvatar.Wrap(hairstylesTypeFile.Close()))
		}()
	}

	// Nose
	pathToNose := pathToFaceType + "/" + config.NoseFolder + "/" + fmt.Sprintf(config.NoseTypeFolder, avatar.Nose)
	nameNoseTypeFile := fmt.Sprintf(config.NoseFile, avatar.Nose)
	noseFile, err := os.Open(pathToNose + "/" + nameNoseTypeFile)
	if err != nil {
		return "", ErrNoAvatarFile.New("generate avatar with %s, error - %s", nameNoseTypeFile, err)
	}
	noseLayer, err = png.Decode(noseFile)
	if err != nil {
		return "", ErrAvatar.Wrap(err)
	}
	defer func() {
		err = errs.Combine(err, ErrAvatar.Wrap(noseFile.Close()))
	}()

	// Beard
	if avatar.Beard != 0 {
		pathToBeard := pathToNose + "/" + config.BeardFolder
		nameBeardTypeFile := fmt.Sprintf(config.BeardFile, avatar.Beard)
		beardFile, err := os.Open(pathToBeard + "/" + nameBeardTypeFile)
		if err != nil {
			return "", ErrNoAvatarFile.New("generate avatar with %s, error - %s", nameBeardTypeFile, err)
		}
		beardLayer, err = png.Decode(beardFile)
		if err != nil {
			return "", ErrAvatar.Wrap(err)
		}
		defer func() {
			err = errs.Combine(err, ErrAvatar.Wrap(beardFile.Close()))
		}()
	}

	// Lips
	pathToLips := pathToNose + "/" + config.LipsFolder
	nameLipsTypeFile := fmt.Sprintf(config.LipsFile, avatar.Lips)
	lipsFile, err := os.Open(pathToLips + "/" + nameLipsTypeFile)
	if err != nil {
		return "", ErrNoAvatarFile.New("generate avatar with %s, error - %s", nameLipsTypeFile, err)
	}
	lipsLayer, err = png.Decode(lipsFile)
	if err != nil {
		return "", ErrAvatar.Wrap(err)
	}
	defer func() {
		err = errs.Combine(err, ErrAvatar.Wrap(lipsFile.Close()))
	}()

	// T-shirt
	pathToTshirt := pathToFaceType + "/" + config.TshirtFolder
	nameTshirtTypeFile := fmt.Sprintf(config.TshirtFile, avatar.Tshirt)
	tshirtFile, err := os.Open(pathToTshirt + "/" + nameTshirtTypeFile)
	if err != nil {
		return "", ErrNoAvatarFile.New("generate avatar with %s, error - %s", nameTshirtTypeFile, err)
	}
	tshirtLayer, err = png.Decode(tshirtFile)
	if err != nil {
		return "", ErrAvatar.Wrap(err)
	}
	defer func() {
		err = errs.Combine(err, ErrAvatar.Wrap(tshirtFile.Close()))
	}()

	baseLayer := faceTypeLayer.Bounds()
	baseImage := image.NewRGBA(baseLayer)
	draw.Draw(baseImage, baseLayer, faceTypeLayer, image.Point{}, draw.Src)
	draw.Draw(baseImage, noseLayer.Bounds(), noseLayer, image.Point{}, draw.Over)
	draw.Draw(baseImage, lipsLayer.Bounds(), lipsLayer, image.Point{}, draw.Over)
	draw.Draw(baseImage, eyeBrowsLayer.Bounds(), eyeBrowsLayer, image.Point{}, draw.Over)
	if tattooLayer != nil {
		draw.Draw(baseImage, tattooLayer.Bounds(), tattooLayer, image.Point{}, draw.Over)
	}
	if hairstylesLayer != nil {
		draw.Draw(baseImage, hairstylesLayer.Bounds(), hairstylesLayer, image.Point{}, draw.Over)
	}
	if beardLayer != nil {
		draw.Draw(baseImage, beardLayer.Bounds(), beardLayer, image.Point{}, draw.Over)
	}
	draw.Draw(baseImage, tshirtLayer.Bounds(), tshirtLayer, image.Point{}, draw.Over)
	draw.Draw(baseImage, eyeLaserLayer.Bounds(), eyeLaserLayer, image.Point{}, draw.Over)

	if err = saveImage(config.PathToOutputAvatars, avatar.CardID, baseImage); err != nil {
		return "", ErrAvatar.Wrap(err)
	}

	buf := new(bytes.Buffer)
	if err = png.Encode(buf, baseImage); err != nil {
		return "", ErrAvatar.Wrap(err)
	}

	return base64.StdEncoding.EncodeToString(buf.Bytes()), nil
}

// saveImage saves image by path.
func saveImage(path string, cardID uuid.UUID, baseImage *image.RGBA) error {
	resultImage, err := os.Create(path + "/" + cardID.String() + "." + TypeImagePNG)
	if err != nil {
		return ErrAvatar.Wrap(err)
	}

	if err = png.Encode(resultImage, baseImage); err != nil {
		return ErrAvatar.Wrap(err)
	}
	defer func() {
		err = errs.Combine(err, ErrAvatar.Wrap(resultImage.Close()))
	}()

	return nil
}

// Get returns avatar from DB.
func (service *Service) Get(ctx context.Context, cardID uuid.UUID) (Avatar, error) {
	avatar, err := service.avatars.Get(ctx, cardID)
	if err != nil {
		return Avatar{}, ErrAvatar.Wrap(err)
	}
	avatar.Image, err = readImage(service.config.PathToOutputAvatars, avatar.CardID)
	return avatar, ErrAvatar.Wrap(err)
}

// readImage saves image by path.
func readImage(path string, cardID uuid.UUID) (string, error) {
	image, err := os.Open(path + "/" + cardID.String() + "." + TypeImagePNG)
	if err != nil {
		return "", ErrNoAvatarFile.Wrap(err)
	}
	decodeImage, err := png.Decode(image)
	if err != nil {
		return "", ErrAvatar.Wrap(err)
	}
	defer func() {
		err = errs.Combine(err, ErrAvatar.Wrap(image.Close()))
	}()

	buf := new(bytes.Buffer)
	if err = png.Encode(buf, decodeImage); err != nil {
		return "", ErrAvatar.Wrap(err)
	}

	return base64.StdEncoding.EncodeToString(buf.Bytes()), nil
}
