// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package imageprocessing_test

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"image/png"
	"os"
	"path/filepath"
	"testing"
	"time"
)

//BenchmarkCreateLayer testing CreateLayer
func BenchmarkCreateLayer(t *testing.B) {
	PathsNames1 := map[string]string{
		"/Users/oleksii/WWW/work/ultimdivision2.0/ultimatedivision/assets/avatars/Man_1/Face_type_1":                               "Face_type_1.png",
		"/Users/oleksii/WWW/work/ultimdivision2.0/ultimatedivision/assets/avatars/Man_1/Face_type_1/Noses/Nose_1":                  "Nose_1.png",
		"/Users/oleksii/WWW/work/ultimdivision2.0/ultimatedivision/assets/avatars/Man_1/Face_type_1/Noses/Nose_1/Lips":             "Lips_4.png",
		"/Users/oleksii/WWW/work/ultimdivision2.0/ultimatedivision/assets/avatars/Man_1/Face_type_1/EyeBrows/EyeBrows_type_1":      "EyeBrows_color_1.png",
		"/Users/oleksii/WWW/work/ultimdivision2.0/ultimatedivision/assets/avatars/Man_1/Face_type_1/Hairstyles/Hairstyles_color_1": "Hairstyle_2.png",
		"/Users/oleksii/WWW/work/ultimdivision2.0/ultimatedivision/assets/avatars/Man_1/Face_type_1/Noses/Nose_1/Beards":           "Beard_2.png",
		"/Users/oleksii/WWW/work/ultimdivision2.0/ultimatedivision/assets/avatars/Man_1/Face_type_1/T-shirts":                      "T-shirt_3.png",
		"/Users/oleksii/WWW/work/ultimdivision2.0/ultimatedivision/assets/avatars/Background":                                      "silver.png",
	}
	PathsNames2 := map[string]string{
		"/Users/oleksii/WWW/work/ultimdivision2.0/ultimatedivision/assets/avatars/Man_1/Face_type_1":                               "Face_type_1.png",
		"/Users/oleksii/WWW/work/ultimdivision2.0/ultimatedivision/assets/avatars/Man_1/Face_type_1/Noses/Nose_1":                  "Nose_1.png",
		"/Users/oleksii/WWW/work/ultimdivision2.0/ultimatedivision/assets/avatars/Man_1/Face_type_1/Noses/Nose_1/Lips":             "Lips_4.png",
		"/Users/oleksii/WWW/work/ultimdivision2.0/ultimatedivision/assets/avatars/Man_1/Face_type_1/EyeBrows/EyeBrows_type_1":      "EyeBrows_color_1.png",
		"/Users/oleksii/WWW/work/ultimdivision2.0/ultimatedivision/assets/avatars/Man_1/Face_type_1/Hairstyles/Hairstyles_color_1": "Hairstyle_2.png",
		"/Users/oleksii/WWW/work/ultimdivision2.0/ultimatedivision/assets/avatars/Man_1/Face_type_1/Noses/Nose_1/Beards":           "Beard_2.png",
		"/Users/oleksii/WWW/work/ultimdivision2.0/ultimatedivision/assets/avatars/Man_1/Face_type_1/T-shirts":                      "T-shirt_3.png",
		"/Users/oleksii/WWW/work/ultimdivision2.0/ultimatedivision/assets/avatars/Background":                                      "silver.png",
	}
	PathsNames3 := map[string]string{
		"/Users/oleksii/WWW/work/ultimdivision2.0/ultimatedivision/assets/avatars/Man_1/Face_type_1":                               "Face_type_1.png",
		"/Users/oleksii/WWW/work/ultimdivision2.0/ultimatedivision/assets/avatars/Man_1/Face_type_1/Noses/Nose_1":                  "Nose_1.png",
		"/Users/oleksii/WWW/work/ultimdivision2.0/ultimatedivision/assets/avatars/Man_1/Face_type_1/Noses/Nose_1/Lips":             "Lips_4.png",
		"/Users/oleksii/WWW/work/ultimdivision2.0/ultimatedivision/assets/avatars/Man_1/Face_type_1/EyeBrows/EyeBrows_type_1":      "EyeBrows_color_1.png",
		"/Users/oleksii/WWW/work/ultimdivision2.0/ultimatedivision/assets/avatars/Man_1/Face_type_1/Hairstyles/Hairstyles_color_1": "Hairstyle_2.png",
		"/Users/oleksii/WWW/work/ultimdivision2.0/ultimatedivision/assets/avatars/Man_1/Face_type_1/Noses/Nose_1/Beards":           "Beard_2.png",
		"/Users/oleksii/WWW/work/ultimdivision2.0/ultimatedivision/assets/avatars/Man_1/Face_type_1/T-shirts":                      "T-shirt_3.png",
		"/Users/oleksii/WWW/work/ultimdivision2.0/ultimatedivision/assets/avatars/Background":                                      "silver.png",
	}
	PathsNames4 := map[string]string{
		"/Users/oleksii/WWW/work/ultimdivision2.0/ultimatedivision/assets/avatars/Man_1/Face_type_1":                               "Face_type_1.png",
		"/Users/oleksii/WWW/work/ultimdivision2.0/ultimatedivision/assets/avatars/Man_1/Face_type_1/Noses/Nose_1":                  "Nose_1.png",
		"/Users/oleksii/WWW/work/ultimdivision2.0/ultimatedivision/assets/avatars/Man_1/Face_type_1/Noses/Nose_1/Lips":             "Lips_4.png",
		"/Users/oleksii/WWW/work/ultimdivision2.0/ultimatedivision/assets/avatars/Man_1/Face_type_1/EyeBrows/EyeBrows_type_1":      "EyeBrows_color_1.png",
		"/Users/oleksii/WWW/work/ultimdivision2.0/ultimatedivision/assets/avatars/Man_1/Face_type_1/Hairstyles/Hairstyles_color_1": "Hairstyle_2.png",
		"/Users/oleksii/WWW/work/ultimdivision2.0/ultimatedivision/assets/avatars/Man_1/Face_type_1/Noses/Nose_1/Beards":           "Beard_2.png",
		"/Users/oleksii/WWW/work/ultimdivision2.0/ultimatedivision/assets/avatars/Man_1/Face_type_1/T-shirts":                      "T-shirt_3.png",
		"/Users/oleksii/WWW/work/ultimdivision2.0/ultimatedivision/assets/avatars/Background":                                      "silver.png",
	}
	PathsNames5 := map[string]string{
		"/Users/oleksii/WWW/work/ultimdivision2.0/ultimatedivision/assets/avatars/Man_1/Face_type_1":                               "Face_type_1.png",
		"/Users/oleksii/WWW/work/ultimdivision2.0/ultimatedivision/assets/avatars/Man_1/Face_type_1/Noses/Nose_1":                  "Nose_1.png",
		"/Users/oleksii/WWW/work/ultimdivision2.0/ultimatedivision/assets/avatars/Man_1/Face_type_1/Noses/Nose_1/Lips":             "Lips_4.png",
		"/Users/oleksii/WWW/work/ultimdivision2.0/ultimatedivision/assets/avatars/Man_1/Face_type_1/EyeBrows/EyeBrows_type_1":      "EyeBrows_color_1.png",
		"/Users/oleksii/WWW/work/ultimdivision2.0/ultimatedivision/assets/avatars/Man_1/Face_type_1/Hairstyles/Hairstyles_color_1": "Hairstyle_2.png",
		"/Users/oleksii/WWW/work/ultimdivision2.0/ultimatedivision/assets/avatars/Man_1/Face_type_1/Noses/Nose_1/Beards":           "Beard_2.png",
		"/Users/oleksii/WWW/work/ultimdivision2.0/ultimatedivision/assets/avatars/Man_1/Face_type_1/T-shirts":                      "T-shirt_3.png",
		"/Users/oleksii/WWW/work/ultimdivision2.0/ultimatedivision/assets/avatars/Background":                                      "silver.png",
	}

	fmt.Println(time.Now().String())

	for path, name := range PathsNames1 {
		image, err := os.Open(filepath.Join(path, name))
		require.NotEqual(t, image, nil)
		require.NoError(t, err)

		layer, err := png.Decode(image)
		require.NoError(t, err)
		require.NotEqual(t, layer, nil)
	}
	for path, name := range PathsNames2 {
		image, err := os.Open(filepath.Join(path, name))
		require.NotEqual(t, image, nil)
		require.NoError(t, err)

		layer, err := png.Decode(image)
		require.NoError(t, err)
		require.NotEqual(t, layer, nil)
	}
	for path, name := range PathsNames3 {
		image, err := os.Open(filepath.Join(path, name))
		require.NotEqual(t, image, nil)
		require.NoError(t, err)

		layer, err := png.Decode(image)
		require.NoError(t, err)
		require.NotEqual(t, layer, nil)
	}
	for path, name := range PathsNames4 {
		image, err := os.Open(filepath.Join(path, name))
		require.NotEqual(t, image, nil)
		require.NoError(t, err)

		layer, err := png.Decode(image)
		require.NoError(t, err)
		require.NotEqual(t, layer, nil)
	}
	for path, name := range PathsNames5 {
		image, err := os.Open(filepath.Join(path, name))
		require.NotEqual(t, image, nil)
		require.NoError(t, err)

		layer, err := png.Decode(image)
		require.NoError(t, err)
		require.NotEqual(t, layer, nil)
	}

	fmt.Println(time.Now().String())
}

//func TestSaveImage(t *testing.T) {
//	PathToAvararsComponents := "/oleksii/WWW/work/ultimdivision2.0/ultimatedivision/assets/avatars"
//	FaceColorFolder := "Man_1"
//
//	pathToFaceColor := filepath.Join(PathToAvararsComponents, FaceColorFolder)
//	baceImage1 := imageprocessing.Inscription{0xc000120800, 1574, 2566, "./assets/fonts/font-avatars.ttf", 118, "6B4538", 19, 960, 1725, 1}
//	baceImage2 := {0xc000090800, 1574, 2566, "./assets/fonts/font-avatars.ttf", 118, "6B4538", 32, 960, 1725, 1}
//	baceImage3 := {0xc0005040c0, 1574, 2566, "./assets/fonts/font-avatars.ttf", 118, "6B4538", 9, 960, 1725, 1}
//	baceImage4 := {0xc0002f20c0, 1574, 2566, "./assets/fonts/font-avatars.ttf", 118, "6B4538", 24, 960, 1725, 1}
//	baceImage5 := {0xc000436080, 1574, 2566, "./assets/fonts/font-avatars.ttf", 118, "6B4538", 28, 960, 1725, 1}
//	//baseImage := image.Image
//	path1 := "./assets/output"
//	path2 := "./assets/output"
//	path3 := "./assets/output"
//	path4 := "./assets/output"
//	path5 := "./assets/output"
//
//	fullPaths := []string{
//		"assets/output/5bb6da2b-056b-4de2-8e25-5d9eb7cd06b0.png",
//		"assets/output/985ce2e0-2f96-45fb-a5ec-0a2fedc4eb0d.png",
//		"assets/output/034a7b5c-fbae-4b13-9bf9-830dbb7a0af1.png",
//		"assets/output/0e25b14a-c22d-42bb-917d-46c8bc727804.png",
//		"assets/output/baf78363-55bc-4399-9ce7-ac50a82dfc1f.png",
//	}
//	baseImage = "5bb6da2b-056b-4de2-8e25-5d9eb7cd06b0"
//	err := os.MkdirAll(path1, os.ModePerm)
//	require.NoError(t, err)
//
//	for _, path := range fullPaths {
//		resultImage, err := os.Create(path)
//		require.NoError(t, err)
//		if err = png.Encode(resultImage, baseImage); err != nil {
//			return err
//		}
//		defer func() {
//			err = errs.Combine(err, resultImage.Close())
//		}()
//	}
//	//if err != nil {
//	//	return err
//	//}
//
//}

//func TestLayering(t *testing.T) {
//	//var layers []image.Image
//	//layers[0] := 0xc000346180
//	layers := []image.Image
//
//	layers := append(layers,0xc000346180)layers1 := []image.Image{0xc000346180, 0xc00038c300, 0xc000092400, 0xc000346640, 0xc0007161c0, 0xc000716300}
//	layers2 := []image.Image{0xc000788000, 0xc0002921c0, 0xc000292800, 0xc0004800c0, 0xc000480440, 0xc000480800}
////	layers3 := []image.Image{0xc000716000, 0xc000536000, 0xc000536340, 0xc000788040, 0xc000347040, 0xc000347440, 0xc000716180}
////	layers4 := []image.Image{0xc00038c0c0, 0xc000092000, 0xc00038c540, 0xc000716040, 0xc0005363c0, 0xc000536540, 0xc000480280}
////	layers5 := []image.Image{0xc000480080, 0xc000346100, 0xc000346740, 0xc0002923c0, 0xc0002927c0, 0xc000292c00, 0xc0003465c0}
////	layers6 := []image.Image{0xc000480300, 0xc000346840}
////	layers7 := []image.Image{0xc000536040, 0xc000292600}
////	layers8 := []image.Image{0xc000536080, 0xc0007f8040}
////	layers9 := []image.Image{0xc000480340, 0xc000293080}
//
//	layers10 := []image.Image{0xc000092140, 0xc000536100}
//
//	width1 := 0
//	width2 := 0
//	width3 := 0
//	width4 := 0
//	width5 := 0
//	width6 := 380
//	width7 := 380
//	width8 := 380
//	width9 := 380
//	width10 := 380
//
//	height1 := 0
//	height2 := 0
//	height3 := 0
//	height4 := 0
//	height5 := 0
//	height6 := 237
//	height7 := 237
//	height8 := 237
//	height9 := 237
//	height10 := 237
//
//	var generalImage *image.RGBA
//	for k, layer := range layers1 {
//		if k == 0 {
//			baseLayer := layer.Bounds()
//			generalImage = image.NewRGBA(baseLayer)
//			draw.Draw(generalImage, baseLayer, layer, image.Point{}, draw.Src)
//			continue
//		}
//
//		if layer != nil {
//			offset := image.Pt(width1, height1)
//			draw.Draw(generalImage, layer.Bounds().Add(offset), layer, image.Point{}, draw.Over)
//		}
//	}
//	for k, layer := range layers2 {
//		if k == 0 {
//			baseLayer := layer.Bounds()
//			generalImage = image.NewRGBA(baseLayer)
//			draw.Draw(generalImage, baseLayer, layer, image.Point{}, draw.Src)
//			continue
//		}
//
//		if layer != nil {
//			offset := image.Pt(width2, height2)
//			draw.Draw(generalImage, layer.Bounds().Add(offset), layer, image.Point{}, draw.Over)
//		}
//	}
//	for k, layer := range layers3 {
//		if k == 0 {
//			baseLayer := layer.Bounds()
//			generalImage = image.NewRGBA(baseLayer)
//			draw.Draw(generalImage, baseLayer, layer, image.Point{}, draw.Src)
//			continue
//		}
//
//		if layer != nil {
//			offset := image.Pt(width3, height3)
//			draw.Draw(generalImage, layer.Bounds().Add(offset), layer, image.Point{}, draw.Over)
//		}
//	}
//	for k, layer := range layers4 {
//		if k == 0 {
//			baseLayer := layer.Bounds()
//			generalImage = image.NewRGBA(baseLayer)
//			draw.Draw(generalImage, baseLayer, layer, image.Point{}, draw.Src)
//			continue
//		}
//
//		if layer != nil {
//			offset := image.Pt(width4, height4)
//			draw.Draw(generalImage, layer.Bounds().Add(offset), layer, image.Point{}, draw.Over)
//		}
//	}
//	for k, layer := range layers5 {
//		if k == 0 {
//			baseLayer := layer.Bounds()
//			generalImage = image.NewRGBA(baseLayer)
//			draw.Draw(generalImage, baseLayer, layer, image.Point{}, draw.Src)
//			continue
//		}
//
//		if layer != nil {
//			offset := image.Pt(width5, height5)
//			draw.Draw(generalImage, layer.Bounds().Add(offset), layer, image.Point{}, draw.Over)
//		}
//	}
//	for k, layer := range layers6 {
//		if k == 0 {
//			baseLayer := layer.Bounds()
//			generalImage = image.NewRGBA(baseLayer)
//			draw.Draw(generalImage, baseLayer, layer, image.Point{}, draw.Src)
//			continue
//		}
//
//		if layer != nil {
//			offset := image.Pt(width6, height6)
//			draw.Draw(generalImage, layer.Bounds().Add(offset), layer, image.Point{}, draw.Over)
//		}
//	}
//	for k, layer := range layers7 {
//		if k == 0 {
//			baseLayer := layer.Bounds()
//			generalImage = image.NewRGBA(baseLayer)
//			draw.Draw(generalImage, baseLayer, layer, image.Point{}, draw.Src)
//			continue
//		}
//
//		if layer != nil {
//			offset := image.Pt(width7, height7)
//			draw.Draw(generalImage, layer.Bounds().Add(offset), layer, image.Point{}, draw.Over)
//		}
//	}
//	for k, layer := range layers8 {
//		if k == 0 {
//			baseLayer := layer.Bounds()
//			generalImage = image.NewRGBA(baseLayer)
//			draw.Draw(generalImage, baseLayer, layer, image.Point{}, draw.Src)
//			continue
//		}
//
//		if layer != nil {
//			offset := image.Pt(width8, height8)
//			draw.Draw(generalImage, layer.Bounds().Add(offset), layer, image.Point{}, draw.Over)
//		}
//	}
//	for k, layer := range layers9 {
//		if k == 0 {
//			baseLayer := layer.Bounds()
//			generalImage = image.NewRGBA(baseLayer)
//			draw.Draw(generalImage, baseLayer, layer, image.Point{}, draw.Src)
//			continue
//		}
//
//		if layer != nil {
//			offset := image.Pt(width9, height9)
//			draw.Draw(generalImage, layer.Bounds().Add(offset), layer, image.Point{}, draw.Over)
//		}
//	}
//	for k, layer := range layers10 {
//		if k == 0 {
//			baseLayer := layer.Bounds()
//			generalImage = image.NewRGBA(baseLayer)
//			draw.Draw(generalImage, baseLayer, layer, image.Point{}, draw.Src)
//			continue
//		}
//
//		if layer != nil {
//			offset := image.Pt(width10, height10)
//			draw.Draw(generalImage, layer.Bounds().Add(offset), layer, image.Point{}, draw.Over)
//		}
//	}
//
//}
