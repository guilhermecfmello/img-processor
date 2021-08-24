package main

import (
	"encoding/json"
	"fmt"
	"image"
	"image/color"
	"net/http"
)

func noneDescribe(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("None image processing is applied, only for tests")
}

func noneProcessing(w http.ResponseWriter, r *http.Request) {
	img, _ := ReceiveImage(r)
	SendImage(img, w)
}

func BwDescribe(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Convert an image to greyscale")
}

func BwProcessing(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Black and white route hit")

	img, _ := ReceiveImage(r)
	newImage, _ := _CreateNewProcessedImage(img, _PixelConvertionColorfulToBw)

	SendImage(newImage, w)
}

func RedDescribe(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Convert an image to red channel")
}

func RedProcessing(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Black and white route hit")

	img, _ := ReceiveImage(r)
	newImage, _ := _CreateNewProcessedImage(img, _PixelConvertionColorfulToRead)

	SendImage(newImage, w)
}

func _CreateNewProcessedImage(img image.Image, function func(p Pixel) Pixel) (image.Image, error) {
	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y
	// TODO, modularize the image synteshis
	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}
	var newImage = image.NewRGBA(image.Rectangle{upLeft, lowRight})
	var originalPixel Pixel
	var processedPixel Pixel

	for i := 0; i <= width; i++ {
		for j := 0; j <= height; j++ {
			r, g, b, a := img.At(i, j).RGBA()

			originalPixel.setPixelValuesFromUint32(r, g, b, a)
			processedPixel = function(originalPixel)
			newImage.Set(i, j, color.RGBA{processedPixel.r, processedPixel.g, processedPixel.b, processedPixel.a})
		}
	}

	return newImage, nil
}

func _PixelConvertionColorfulToBw(p Pixel) Pixel {
	var newPixel Pixel
	var average = p.r/3 + p.g/3 + p.b/3
	newPixel.r = average
	newPixel.g = average
	newPixel.b = average
	newPixel.a = p.a
	return newPixel
}

func _PixelConvertionColorfulToRead(p Pixel) Pixel {
	var newPixel Pixel
	newPixel.r = p.r
	newPixel.g = 0
	newPixel.b = 0
	newPixel.a = p.a
	return newPixel
}
