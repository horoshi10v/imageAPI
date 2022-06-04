package routes

import (
	"api/controllers"
	"github.com/gorilla/mux"
	"net/http"
)

func Setup() {
	r := mux.NewRouter()
	r.HandleFunc("/upload", controllers.UploadPhoto)
	r.HandleFunc("/{fileName}", controllers.GetPhoto)
	http.ListenAndServe(":8080", r)

}
