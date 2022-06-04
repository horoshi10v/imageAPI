package helpers

import (
	"fmt"
	"github.com/h2non/bimg"
	"os"
	"strconv"
)

//quality multiplier
const (
	q100 = 1
	q75  = 0.75
	q50  = 0.50
	q25  = 0.25
)

// Resize file resized by multiplied const with use bimg
func Resize(fileBytes []byte, fileName string) {
	var quality = [4]float64{q100, q75, q50, q25}
	var idQuality = make([]int, len(quality))
	for j := range idQuality {
		idQuality[j] = int(quality[j] * 100)
	}

	sizeImage, _ := bimg.NewImage(fileBytes).Size()
	for i, item := range quality {
		newWidth := item * float64(sizeImage.Width)
		newHeight := item * float64(sizeImage.Height)
		newImage, err := bimg.NewImage(fileBytes).Resize(int(newWidth), int(newHeight))
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		//resources/q.100.imageName.png
		bimg.Write("resources/"+"q"+strconv.Itoa(idQuality[i])+"."+fileName, newImage)
	}
}
