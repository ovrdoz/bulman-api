package main

import (
	"log"
	"net/http"

	"bulman-api/apis"
	"bulman-api/helper"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/projects", apis.GetProjects).Methods("GET")
	r.HandleFunc("/api/projects/{id}", apis.GetProject).Methods("GET")
	r.HandleFunc("/api/projects", apis.CreateProject).Methods("POST")
	r.HandleFunc("/api/projects/{id}", apis.UpdateProject).Methods("PUT")
	r.HandleFunc("/api/projects/{id}", apis.DeleteProject).Methods("DELETE")

	r.HandleFunc("/api/scenarios", apis.GetScenarios).Methods("GET")
	r.HandleFunc("/api/scenarios/{id}", apis.GetScenario).Methods("GET")
	r.HandleFunc("/api/scenarios", apis.CreateScenario).Methods("POST")
	r.HandleFunc("/api/scenarios/{id}", apis.UpdateScenario).Methods("PUT")
	r.HandleFunc("/api/scenarios/{id}", apis.DeleteScenario).Methods("DELETE")

	handler := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"POST", "GET", "OPTIONS", "PUT", "DELETE"},
		AllowedHeaders: []string{"Accept", "content-type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
	}).Handler(r)

	config := helper.GetConfiguration()
	log.Fatal(http.ListenAndServe(":"+config.Port, handler))

}
