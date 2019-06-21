package main

import (
	"fmt"
	tf "github.com/tensorflow/tensorflow/tensorflow/go"
	"github.com/tensorflow/tensorflow/tensorflow/go/op"

	"github.com/elv-sanchez/goTecoGAN/data"
	zshape "github.com/elv-sanchez/goTecoGAN/shape"
)

func main() {
	inputDir := "/Users/sanchezoleary/dev/TecoGAN/LR/calendar" // TODO: param!

	dataTemp,_ := data.Dataloading(inputDir)

	shape := zshape.ShapeHelper{zshape.ImageShapes(dataTemp)[0]}

	inputShape := tf.MakeShape(1, shape.GetDim(0), shape.GetDim(1), shape.GetDim(2))
	outputShape := tf.MakeShape(1, shape.GetDim(0) * 4, shape.GetDim(1) * 4, 3)
	sh := zshape.ShapeHelper{&inputShape}
	oh := sh.GetDim(1) - sh.GetDim(1)/8 * 8
	ow := sh.GetDim(2) - sh.GetDim(2)/8 * 8

	paddings := [][]int64{{0,0}, {0,oh}, {0,ow}, {0,0}}
	fmt.Println(outputShape)
	fmt.Println(inputShape)
	root := op.NewScope()
	_ = paddings

	inputs_raw := op.Placeholder(root, tf.Float, op.PlaceholderShape(inputShape))
	_ = inputs_raw

	pre_inputs, _ := tf.NewTensor(inputShape)
	pre_gen, _ := tf.NewTensor(outputShape)
	pre_warp, _ := tf.NewTensor(outputShape)
	_ = pre_inputs
	_ = pre_gen
	_ = pre_warp

	//transpose_pre := op.SpaceToDepth(root, pre_warp, 4)
	//inputs_all := op.Concat()


}