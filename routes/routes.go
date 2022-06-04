package routes

import (
	"api/handler"
	"github.com/gorilla/mux"
	"net/http"
)

func Setup() {
	r := mux.NewRouter()
	r.HandleFunc("/upload", handler.UploadPhoto)
	//r.HandleFunc("/{fileName}", handler.GetPhoto)
	r.HandleFunc("/{fileName}", handler.GetQualityPhoto)
	http.ListenAndServe(":8080", r)

}
