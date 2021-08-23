package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/helloworld", HelloWorld).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Hello World!")
}
