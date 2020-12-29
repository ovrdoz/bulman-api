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

	handler := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowedHeaders: []string{"a_custom_header", "content_type"},
	}).Handler(r)

	config := helper.GetConfiguration()
	log.Fatal(http.ListenAndServe(config.Port, r))

	http.ListenAndServe(config.Port, handler)

}
