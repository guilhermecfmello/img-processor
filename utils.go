package main

import (
	"bytes"
	"encoding/json"
	"image"
	"image/jpeg"
	"log"
	"mime/multipart"
	"net/http"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Hello World!")
}

func ReceiveImage(r *http.Request) (image.Image, *multipart.FileHeader) {
	// Max file length = 20MB TODO (Transfer this parameter to database configuration)
	r.ParseMultipartForm(20 << 20)

	// Getting file from form-data
	file, header, err := r.FormFile("img")

	if err != nil {
		panic(err)
	}
	// Copy the file data to my buffer
	img, _ := jpeg.Decode(file)
	// contents := buf.String()
	return img, header
}

func SendImage(img image.Image, w http.ResponseWriter) {
	buffer := new(bytes.Buffer)
	if err := jpeg.Encode(buffer, img, nil); err != nil {
		log.Println("unable to encode image.")
	}
	w.Header().Set("Content-Type", "image/jpeg")

	if _, err := w.Write(buffer.Bytes()); err != nil {
		log.Println("unable to write image.")
	}
}
