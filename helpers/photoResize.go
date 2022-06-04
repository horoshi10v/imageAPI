package helpers

import (
	"fmt"
	"github.com/h2non/bimg"
	"os"
	"strconv"
)

const (
	q100 = 1
	q75  = 0.75
	q50  = 0.50
	q25  = 0.25
)

func Resize(fileBytes []byte, fileName string) {
	var quality = [4]float64{q100, q75, q50, q25}
	sizeImage, _ := bimg.NewImage(fileBytes).Size()
	for i, item := range quality {
		newWidth := item * float64(sizeImage.Width)
		newHeight := item * float64(sizeImage.Height)
		newImage, err := bimg.NewImage(fileBytes).Resize(int(newWidth), int(newHeight))
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		bimg.Write("resources/"+fileName+"q="+strconv.Itoa(i)+"."+fileName, newImage)
	}
}

func PhotoResize(fileBytes []byte, fileName string) {
	//uploadImage := bimg.NewImage(fileBytes)
	bimg.Write("resources/"+fileName, fileBytes)
	//sizeImage, _ := bimg.NewImage(fileBytes).Size()
	//resize(sizeImage)
	//newImage, err := bimg.NewImage(fileBytes).Resize(sizeImage.Width*0.7, sizeImage.Height)
	//if err != nil {
	//	fmt.Fprintln(os.Stderr, err)
	//}
	//buffer, err := bimg.Read("upload-images/4.JPG")
	//if err != nil {
	//	fmt.Fprintln(os.Stderr, err)
	//}
	//
	//newImage, err := bimg.NewImage(buffer).Resize(800, 600)
	//if err != nil {
	//	fmt.Fprintln(os.Stderr, err)
	//}
	//
	//size, err := bimg.NewImage(newImage).Size()
	//if size.Width == 800 && size.Height == 600 {
	//	fmt.Println("The image size is valid")
	//}
	//
	//bimg.Write("new.jpg", newImage)
}

func resizeImage(fileBytes []byte, fileName string) {
	//sizeImage, _ := bimg.NewImage(fileBytes).Size()
	//bimg.Resize(fileBytes, )
}
