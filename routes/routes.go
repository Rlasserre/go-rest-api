package routes

import (
	"go-rest-api/controllers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personages", controllers.AllPersonages).Methods("Get")
	r.HandleFunc("/api/personages/{id}", controllers.GetPersonageById).Methods("Get")
	log.Fatal(http.ListenAndServe(":8000", r))
}
