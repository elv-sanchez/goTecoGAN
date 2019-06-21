package shape

import(
	tf "github.com/tensorflow/tensorflow/tensorflow/go"
)

type ShapeHelper struct {
	*tf.Shape
}

func (sh *ShapeHelper) GetDim(dim int) int64 {
	if dim >= sh.NumDimensions() {
		panic("you a dumbass")
	}
	v, _ := sh.ToSlice()
	return v[dim]
}

func ImageShapes(imageList []string) []*tf.Shape {
	var imageShapes []*tf.Shape
	for _, image := range imageList {
		imageTemp, _  := imageShape(image)
		imageShapes = append(imageShapes, imageTemp)
	}

	return imageShapes
}
