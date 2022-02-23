// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package imageprocessing

import (
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"testing"

	"github.com/fogleman/gg"
	"github.com/zeebo/errs"
)

// TypeFile defines the list of possible type of files.
type TypeFile string

const (
	// TypeFilePNG indicates that the type file is png.
	TypeFilePNG TypeFile = "png"
	// TypeFileJSON indicates that the type file is json.
	TypeFileJSON TypeFile = "json"
)

// LayerComponentsCountOLD searches count files in the specified path and by name of file.
func LayerComponentsCountOLD(pathToLayerComponents, nameFile string) (int, error) {
	files, err := ioutil.ReadDir(pathToLayerComponents)
	if err != nil {
		return 0, fmt.Errorf(pathToLayerComponents + " - folder does not exist")
	}
	var count int
	for _, file := range files {
		fmt.Println("Name of the file === ", nameFile, "Bites ===  ", []byte(file.Name()))
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

//LayerComponentsCount testing.
func LayerComponentsCount(pathToLayerComponents, nameFile string) (int, error) {
	files, err := ioutil.ReadDir(pathToLayerComponents)
	if err != nil {
		return 0, fmt.Errorf(pathToLayerComponents + " - folder does not exist")
	}
	var length = len(files)
	var count int

	for i := 0; i < length; i++ {
		isMatched, err := regexp.Match(fmt.Sprintf(nameFile, `\d`), []byte(files[i].Name())) //potential replace by regex.Match
		if err != nil {
			return 0, err
		}
		if isMatched {
			count++
		}
	}
	return count, nil
}

//LayerComponentsCountTEST testing.
func LayerComponentsCountTEST(pathToLayerComponents, nameFile string) (int, error) {
	files, err := ioutil.ReadDir(pathToLayerComponents)
	if err != nil {
		return 0, fmt.Errorf(pathToLayerComponents + " - folder does not exist")
	}
	var length = len(files)
	var count int

	for i := 0; i < length; i++ {
		//count++
		//if isMatched := strings.Contains(nameFile, files[i].Name()); isMatched == true {
		count++
		//}
	}
	return count, nil
}

// CreateLayer searches and decodes image to layer.
func CreateLayer(path, name string) (image.Image, error) {
	image, err := os.Open(filepath.Join(path, name))
	if err != nil {
		return nil, err
	}
	defer func() {
		err = errs.Combine(err, image.Close())
	}()

	layer, err := png.Decode(image)
	if err != nil {
		return nil, err
	}
	return layer, err
}

// Layering overlays image layers on the base image.
func Layering(layers []image.Image, width, height int) *image.RGBA {
	var generalImage *image.RGBA

	for k, layer := range layers {
		if k == 0 {
			baseLayer := layer.Bounds()
			generalImage = image.NewRGBA(baseLayer)
			draw.Draw(generalImage, baseLayer, layer, image.Point{}, draw.Src)
			continue
		}

		if layer != nil {
			offset := image.Pt(width, height)
			draw.Draw(generalImage, layer.Bounds().Add(offset), layer, image.Point{}, draw.Over)
		}
	}
	return generalImage
}

// Layering2 testing overlays image layers on the base image.
func Layering2(layers []image.Image, width, height int) *image.RGBA {
	var generalImage *image.RGBA
	count := len(layers)
	for i := 0; i < count; i++ {
		if i == 0 {
			baseLayer := layers[i].Bounds()
			generalImage = image.NewRGBA(baseLayer)
			draw.Draw(generalImage, baseLayer, layers[i], image.Point{}, draw.Src)
			continue
		}

		if layers[i] != nil {
			offset := image.Pt(width, height)
			draw.Draw(generalImage, layers[i].Bounds().Add(offset), layers[i], image.Point{}, draw.Over)
		}
	}
	return generalImage
}

//BenchmarkSaveImage testing.
func BenchmarkSaveImage(b *testing.B) {
	file, err := os.Open("/Users/ivanbalagus/go/src/ultimatedivision/assets/avatars/Man_1/Face_type_1/EyeBrows/EyeBrows_type_1/EyeBrows_color_1.png")
	if err != nil {
		b.Fatal(err)
	}
	img, err := png.Decode(file)
	if err != nil {
		b.Fatal(err)
	}
	for i := 0; i < b.N; i++ {
		err := SaveImage("/Users/oleksii/WWW/work/ultimdivision2.0/ultimatedivision/pkg/imageprocessing/testdata", filepath.Join("/Users/oleksii/WWW/work/ultimdivision2.0/ultimatedivision/pkg/imageprocessing/testdata", strconv.Itoa(i)+"."+string(TypeFilePNG)), img)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// SaveImage saves image by path.
func SaveImage(path, fullPath string, baseImage image.Image) error {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return err
	}

	resultImage, err := os.Create(fullPath)
	if err != nil {
		return err
	}

	if err = png.Encode(resultImage, baseImage); err != nil {
		return err
	}
	defer func() {
		err = errs.Combine(err, resultImage.Close())
	}()

	return nil
}

// Inscription entity describes values required to apply inscription to the image.
type Inscription struct {
	Img         image.Image
	Width       int
	Height      int
	PathToFonts string
	FontSize    float64
	FontColor   string
	Text        string
	X           float64
	Y           float64
	TextAlign   float64
}

// ApplyInscription overlays the inscription on the image.
func ApplyInscription(inscription Inscription) (image.Image, error) {
	dc := gg.NewContext(inscription.Width, inscription.Height)
	if err := dc.LoadFontFace(inscription.PathToFonts, inscription.FontSize); err != nil {
		return nil, err
	}

	dc.SetHexColor(inscription.FontColor)
	dc.DrawImage(inscription.Img, 0, 0)
	dc.DrawStringAnchored(inscription.Text, inscription.X, inscription.Y, inscription.TextAlign, 0.5)
	dc.Clip()
	return dc.Image(), nil
}
