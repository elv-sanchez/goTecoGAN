package shape

import (
	"github.com/elv-sanchez/goTecoGAN/data"
	tf "github.com/tensorflow/tensorflow/tensorflow/go"

)

func imageShape(input string)(*tf.Shape, error) {
	img, err := data.DecodeImage(input)
	if err != nil {
		return nil, err
	}

	imgShape := tf.MakeShape(int64(img.Bounds().Max.X), int64(img.Bounds().Max.Y), 3)

	return &imgShape, nil

}
