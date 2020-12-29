package main

import (
	"log"
	"net/http"

	"bulman-api/apis"
	"bulman-api/helper"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/projects", apis.GetProjects).Methods("GET")
	r.HandleFunc("/api/projects/{id}", apis.GetProject).Methods("GET")
	r.HandleFunc("/api/projects", apis.CreateProject).Methods("POST")
	r.HandleFunc("/api/projects/{id}", apis.UpdateProject).Methods("PUT")
	r.HandleFunc("/api/projects/{id}", apis.DeleteProject).Methods("DELETE")

	config := helper.GetConfiguration()
	log.Fatal(http.ListenAndServe(config.Port, r))

}
