package controllers

import (
	"encoding/json"
	"fmt"
	"go-rest-api/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func AllPersonages(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Personages)
}

func GetPersonageById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for _, personage := range models.Personages {
		if strconv.Itoa(personage.Id) == id {
			json.NewEncoder(w).Encode(personage)
		}
	}
}
