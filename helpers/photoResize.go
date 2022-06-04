package helpers

import (
	"github.com/h2non/bimg"
)

func PhotoResize(fileBytes []byte, fileName string) {
	//uploadImage := bimg.NewImage(fileBytes)
	bimg.Write("resources/"+fileName, fileBytes)

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
