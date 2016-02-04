package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"snow_api/lib"
	"snow_api/models"
)

func main() {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/", lib.Decorator(status)).Methods("GET")
	r.HandleFunc("/v1/api/", lib.Decorator(list)).Methods("GET", "POST")
	r.HandleFunc("/v1/api/{email}", lib.Decorator(show)).Methods("GET")
	fmt.Println("Listening on http://localhost:5000")
	http.ListenAndServe(":5000", r)
}

func status(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Status{Status: "healthy"})
}

func list(w http.ResponseWriter, r *http.Request) {
	sysparm_limit := r.URL.Query().Get("sysparm_limit")
	snowUsers := lib.FetchListOfSnowUsers(sysparm_limit)
	json.NewEncoder(w).Encode(snowUsers)

}

func show(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	email := vars["email"]
	snowUser := lib.FetchOneUser(email)
	json.NewEncoder(w).Encode(snowUser)

}

func create(w http.ResponseWriter, r *http.Request) {

}
