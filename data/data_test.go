package data

import (
	"fmt"
	tf "github.com/tensorflow/tensorflow/tensorflow/go"
	"io/ioutil"
	"path"
	"strings"
	"testing"
)

func TestDataLoading(t *testing.T) {

	testInputDirLR := "/Users/sanchezoleary/dev/TecoGAN/LR/calendar"

	imageListLrTemp, err := ioutil.ReadDir(testInputDirLR)
	if err != nil {
		t.Errorf("ReadDir failed: %v", err)
	}

	var imageList []string
	for _, file := range imageListLrTemp {
		if strings.HasSuffix(file.Name(), ".png") {
			imageList = append(imageList, path.Join(testInputDirLR, file.Name()))
		}
	}
	fmt.Println(imageList)
	fmt.Println(DecodeImage(imageList[1]))
}

func TestDataShaping(t *testing.T) {

	test0001 := "/Users/sanchezoleary/dev/TecoGAN/LR/calendar/0001.png"

	img, err := DecodeImage(test0001)
	if err != nil {
		panic(err)
	}

	imgShape := tf.MakeShape(int64(img.Bounds().Max.X), int64(img.Bounds().Max.Y), 3)
	fmt.Println(imgShape)
}


