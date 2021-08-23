package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func BwDescribe(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Convert an image to black and white channel")
}

func BwProcessing(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Black and white route hit")

	buf, header := ReceiveFile(r)
	fmt.Printf("Buffer size: %+v\n", buf.Len())
	fmt.Printf("Uploaded File: %+v\n", header.Filename)
	fmt.Printf("File Size: %+v\n", header.Size)
	fmt.Printf("MIME Header: %+v\n", header.Header)
}
