package main

import (
	"fmt"
	"image"
	"image/png"
	"io/ioutil"
	"os"
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
	fmt.Println(PreprocessTest(imageList[1]))
}

func PreprocessTest(imagename string) image.Image {
	imageFile, err := os.Open(imagename)
	if err != nil {
		panic(err)
	}
	im, err := png.Decode(imageFile)
	fmt.Println(im)
	if err != nil {
		panic(err)
	}
	return im

}