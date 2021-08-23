package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Synthesis
// Processing

// TODO: Validate requisitions (headers and bodys, throwing errors)
func main() {
	router := mux.NewRouter()

	router.HandleFunc("/helloworld", HelloWorld).Methods("GET")

	// Black and White
	router.HandleFunc("/processing/bw", BwDescribe).Methods("GET")
	router.HandleFunc("/processing/bw", BwProcessing).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", router))
}
