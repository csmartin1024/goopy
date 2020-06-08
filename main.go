package main

import (
	"log"
	"net/http"
	"os"

	"github.com/csmartin1024/goopy/helper"
	"github.com/csmartin1024/goopy/routes"
	"github.com/gorilla/mux"
)

func main() {

	helper.InitDatabase()

	r := mux.NewRouter()

	r.HandleFunc("/api/goops", routes.ListGoops).Methods("GET")
	r.HandleFunc("/api/goops/{id}", routes.GetGoop).Methods("GET")
	r.HandleFunc("/api/goops", routes.CreateGoop).Methods("POST")
	// r.HandleFunc("/api/goops/{id}", routes.UpdateGoop).Methods("PUT")
	r.HandleFunc("/api/goops/{id}", routes.DeleteGoop).Methods("DELETE")
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./web/dist")))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	log.Fatal(http.ListenAndServe(":"+port, r))

}
