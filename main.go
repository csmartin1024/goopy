package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	// setup port
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	// handle home page requests
	http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(w, "You sir, have been gooped.")
	})
	// serve
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
}
