package controllers

import (
	"api/rabbit"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

func UploadPhoto(w http.ResponseWriter, r *http.Request) {
	// Parse our multipart form, 10 << 20 specifies a maximum
	// upload of 10 MB files.
	r.ParseMultipartForm(10 << 20)
	// FormFile returns the first file for the given key `myFile`
	// it also returns the FileHeader so we can get the Filename,
	// the Header and the size of the file
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

	// read all of the contents of our uploaded file into a
	// byte array
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	rabbit.SendMsgRabbit(fileBytes, handler.Filename)
	fmt.Fprintf(w, "Successfully Uploaded File %v", handler.Filename)
}

func GetPhoto(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fileName, ok := vars["fileName"]
	if !ok {
		fmt.Println("fileName is missing in parameters")
	}
	//file, err := os.OpenFile("resources/"+fileName, os.O_RDONLY, 0666)

	fileBytes, err := ioutil.ReadFile("resources/" + fileName)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(fileBytes)
	return

	//request := make(map[string]string)
	//reqBody, _ := ioutil.ReadAll(r.Body)
	//err := json.Unmarshal(reqBody, &request)
	//if err != nil {
	//	log.Printf("can't unmarshal'")
	//}
	//file, err := os.OpenFile("resources/"+request["filename"], os.O_RDONLY, 0666)
	//
	//w.WriteHeader(http.StatusOK)
	//w.Header().Set("Content-Type", "application/octet-stream")
	//w.Write(file)
	//return
	//
	//reqImg, err := client.Get("http://www.google.com/intl/en_com/images/srpr/logo3w.png")
	//if err != nil {
	//	fmt.Fprintf(res, "Error %d", err)
	//	return
	//}
	//res.Header().Set("Content-Length", fmt.Sprint(reqImg.ContentLength))
	//res.Header().Set("Content-Type", reqImg.Header.Get("Content-Type"))
	//if _, err = io.Copy(res, reqImg.Body); err != nil {
	//	// handle error
	//}
	//reqImg.Body.Close()
}
