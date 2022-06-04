package routes

import (
	"api/handler"
	"github.com/gorilla/mux"
	"net/http"
)

func Setup() {
	r := mux.NewRouter()
	r.HandleFunc("/upload", handler.UploadPhoto)
	r.HandleFunc("/{fileName}", handler.PostPhoto)
	http.ListenAndServe(":8080", r)
}
