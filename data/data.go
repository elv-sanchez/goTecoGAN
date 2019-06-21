package data

import (
	"image"
	"image/png"
	"os"
)

func DecodeImage(imagePath string) (image.Image, error) {
	imageFile, err := os.Open(imagePath)
	if err != nil {
		return nil, err
	}
	defer imageFile.Close()
	im, err := png.Decode(imageFile)
	// fmt.Println(im)
	if err != nil {
		return nil, err
	}

	// TODO: TEST !!!

	return im, nil
}
