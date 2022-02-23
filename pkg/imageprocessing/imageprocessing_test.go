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

//BenchmarkCreateLayer testing CreateLayer.
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
