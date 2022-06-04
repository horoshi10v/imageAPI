package handler

import (
	"api/rabbit"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
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
	fileBytes, err := ioutil.ReadFile("resources/" + "q" + strconv.Itoa(0) + "." + fileName)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(fileBytes)
	return
}

func GetQualityPhoto(w http.ResponseWriter, r *http.Request) {
	var fileBytes []byte
	vars := mux.Vars(r)
	fileName := vars["fileName"]
	params := r.URL.Query().Get("quality")
	if params == "" {
		fileBytes, _ = ioutil.ReadFile("resources/" + "q100." + fileName)
	} else {
		fileBytes, _ = ioutil.ReadFile("resources/" + "q" + params + "." + fileName)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(fileBytes)
	return

	//param := r.FormValue("quality")
	//if param == "" {
	//	fileBytes, _ = ioutil.ReadFile("resources/" + "q0." + fileName)
	//} else {
	//	fileBytes, _ = ioutil.ReadFile("resources/" + "q" + param + "." + fileName)
	//}

	//params := make(map[string]int)
	//params["100"] = 0
	//params["75"] = 1
	//params["50"] = 2
	//params["25"] = 3
	//vars := mux.Vars(r)
	//fileName := vars["fileName"]
	//quality := vars["quality"]
	//request, ok := params[quality]
	//if ok {
	//	fileBytes, _ := ioutil.ReadFile("resources/" + "q" + strconv.Itoa(request) + "." + fileName)
	//	w.WriteHeader(http.StatusOK)
	//	w.Header().Set("Content-Type", "application/octet-stream")
	//	w.Write(fileBytes)
	//	return
	//}
}
