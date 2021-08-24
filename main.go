package main

import (
	"fmt"
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

	// Return Image test
	router.HandleFunc("/processing/none", noneDescribe).Methods("GET")
	router.HandleFunc("/processing/none", noneProcessing).Methods("POST")

	// Black and White
	router.HandleFunc("/processing/bw", BwDescribe).Methods("GET")
	router.HandleFunc("/processing/bw", BwProcessing).Methods("POST")

	fmt.Println("Server listening in localhost, port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
