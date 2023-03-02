package routes

import (
	"go-rest-api/controllers"
	"log"
	"net/http"
)

func HandleRequest() {
	http.HandleFunc("/", controllers.Home)
	http.HandleFunc("/api/personages", controllers.AllPersonages)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
