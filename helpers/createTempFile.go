package helpers

import (
	"log"
	"os"
)

// CreateUploadFile create file in upload-images/ directory
func CreateUploadFile(fileBytes []byte, filename string) {
	uploadFile, err := os.Create("upload-images/" + filename)
	if err != nil {
		log.Println(err)
	}
	defer uploadFile.Close()
	// write this byte array to our temporary file
	_, err = uploadFile.Write(fileBytes)
	if err != nil {
		log.Fatal("can't write to temporary file", err)
	}
	// return that we have successfully uploaded our file!
	log.Printf("Successfully Uploaded File\n")
}
