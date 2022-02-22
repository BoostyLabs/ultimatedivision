package imageprocessing_test

import (
	_ "net/http/pprof"
	"testing"

	"ultimatedivision/pkg/imageprocessing"
)

//func BenchmarkSaveImage(b *testing.B) {
//	file, err := os.Open("/Users/oleksii/WWW/work/ultimdivision2.0/ultimatedivision/assets/avatars/Man_1/Face_type_1/EyeBrows/EyeBrows_type_1/EyeBrows_color_1.png")
//	if err != nil {
//		b.Fatal(err)
//	}
//	img, err := png.Decode(file)
//	if err != nil {
//		b.Fatal(err)
//	}
//	for i := 0; i < b.N; i++ {
//		err := imageprocessing.SaveImage("/Users/oleksii/WWW/work/ultimdivision2.0/ultimatedivision/pkg/imageprocessing/testdata", filepath.Join("/Users/oleksii/WWW/work/ultimdivision2.0/ultimatedivision/pkg/imageprocessing/testdata", strconv.Itoa(i)+".png"), img)
//		if err != nil {
//			b.Fatal(err)
//		}
//	}
//
//}

//func BenchmarkLayering(b *testing.B) {
//	for i := 0; i < 5; i++ {
//		pathToNoseType := "/Users/oleksii/WWW/work/ultimdivision2.0/ultimatedivision/assets/avatars/Man_1/Face_type_1/Noses/Nose_1/"
//		NoseFile := "Nose_1.png"
//
//		layer1, err := imageprocessing.CreateLayer(pathToNoseType, NoseFile)
//		if err != nil {
//
//		}
//		layer2, err := imageprocessing.CreateLayer(pathToNoseType, NoseFile)
//		if err != nil {
//
//		}
//		layer3, err := imageprocessing.CreateLayer(pathToNoseType, NoseFile)
//		if err != nil {
//
//		}
//		layer4, err := imageprocessing.CreateLayer(pathToNoseType, NoseFile)
//		if err != nil {
//
//		}
//		layer5, err := imageprocessing.CreateLayer(pathToNoseType, NoseFile)
//		if err != nil {
//
//		}
//
//		layers := []image.Image{}
//
//		layers = append(layers, layer1, layer2, layer3, layer4, layer5)
//
//		_ = imageprocessing.Layering(layers, 0, 0)
//	}
//
//}

//func BenchmarkLayering2(b *testing.B) {
//	for i := 0; i < 5; i++ {
//		pathToNoseType := "/Users/oleksii/WWW/work/ultimdivision2.0/ultimatedivision/assets/avatars/Man_1/Face_type_1/Noses/Nose_1/"
//		NoseFile := "Nose_1.png"
//
//		layer1, err := imageprocessing.CreateLayer(pathToNoseType, NoseFile)
//		if err != nil {
//
//		}
//		layer2, err := imageprocessing.CreateLayer(pathToNoseType, NoseFile)
//		if err != nil {
//
//		}
//		layer3, err := imageprocessing.CreateLayer(pathToNoseType, NoseFile)
//		if err != nil {
//
//		}
//		layer4, err := imageprocessing.CreateLayer(pathToNoseType, NoseFile)
//		if err != nil {
//
//		}
//		layer5, err := imageprocessing.CreateLayer(pathToNoseType, NoseFile)
//		if err != nil {
//
//		}
//
//		layers := []image.Image{}
//
//		layers = append(layers, layer1, layer2, layer3, layer4, layer5)
//
//		_ = imageprocessing.Layering2(layers, 0, 0)
//	}
//
//}

func BenchmarkLayerComponentsCount(b *testing.B) {
	PathToAvararsComponents := "/oleksii/WWW/work/ultimdivision2.0/ultimatedivision/assets/avatars"
	FaceColorFolder := "Man_1"

	for i := 0; i < b.N; i++ {
		_, _ = imageprocessing.LayerComponentsCount(PathToAvararsComponents, FaceColorFolder)
	}

}
func BenchmarkLayerComponentsCount2(b *testing.B) {
	PathToAvararsComponents := "/oleksii/WWW/work/ultimdivision2.0/ultimatedivision/assets/avatars"
	FaceColorFolder := "Man_1"

	for i := 0; i < b.N; i++ {
		_, _ = imageprocessing.LayerComponentsCountTEST(PathToAvararsComponents, FaceColorFolder)
	}
}

func BenchmarkLayerComponentsCount3(b *testing.B) {
	PathToAvararsComponents := "/oleksii/WWW/work/ultimdivision2.0/ultimatedivision/assets/avatars"
	FaceColorFolder := "Man_1"

	for i := 0; i < b.N; i++ {
		_, _ = imageprocessing.LayerComponentsCountOLD(PathToAvararsComponents, FaceColorFolder)
	}
}
