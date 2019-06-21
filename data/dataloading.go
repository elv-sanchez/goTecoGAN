package data

import (
	"io/ioutil"
	"path"
	"strings"
)

func Dataloading (inputDir string) ([]string, error){
	imageListLrTemp, err := ioutil.ReadDir(inputDir)
	if err != nil {
		return nil, err
	}

	var imageList []string
	for _, file := range imageListLrTemp {
		if strings.HasSuffix(file.Name(), ".png") {
			imageList = append(imageList, path.Join(inputDir, file.Name()))
		}
	}
	return imageList, nil
}