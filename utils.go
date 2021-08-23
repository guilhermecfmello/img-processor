package main

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
)

// type Pixel struct {
// 	r, g, b int
// }

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Hello World!")
}

func ReceiveFile(r *http.Request) (bytes.Buffer, *multipart.FileHeader) {
	// Max file length = 20MB TODO (Transfer this parameter to database configuration)
	r.ParseMultipartForm(20 << 20)
	var buf bytes.Buffer

	// Getting file from form-data
	file, header, err := r.FormFile("img")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	// Copy the file data to my buffer
	io.Copy(&buf, file)

	// contents := buf.String()
	return buf, header
}

// func getPixel() {

// }
