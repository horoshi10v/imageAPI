package handler

import (
	"api/rabbit"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

// UploadPhoto GET upload photo from request and send to RabbitMQ
func UploadPhoto(w http.ResponseWriter, r *http.Request) {
	// Parse our multipart form, 10 << 20 specifies a maximum
	// upload of 10 MB files.
	r.ParseMultipartForm(10 << 20)
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		log.Println("Error Retrieving the File")
		log.Println(err)
		return
	}
	defer file.Close()
	log.Printf("Uploaded File: %+v\n", handler.Filename)
	log.Printf("File Size: %+v\n", handler.Size)
	log.Printf("MIME Header: %+v\n", handler.Header)

	//transfer file to []byte for send to RabbitMQ
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	rabbit.SendMsgRabbit(fileBytes, handler.Filename)
	fmt.Fprintf(w, "Successfully Uploaded File %v", handler.Filename)
}

// PostPhoto POST to server file from ./resources
// allowed Query params "?quality=100/75/50/25"
func PostPhoto(w http.ResponseWriter, r *http.Request) {
	var fileBytes []byte
	vars := mux.Vars(r)
	fileName := vars["fileName"]
	params := r.URL.Query().Get("quality")

	//http://localhost:8080/{filename}
	if params == "" {
		fileBytes, _ = ioutil.ReadFile("resources/" + "q100." + fileName)
	} else {
		//http://localhost:8080/{filename}?quality=100/75/50/25
		fileBytes, _ = ioutil.ReadFile("resources/" + "q" + params + "." + fileName)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(fileBytes)
	return
}
