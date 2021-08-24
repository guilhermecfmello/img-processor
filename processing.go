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
	json.NewEncoder(w).Encode("Convert an image to black and white channel")
}

func BwProcessing(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Black and white route hit")

	img, _ := ReceiveImage(r)

	// TODO black and white convertion here
	newImage, _ := _CreateNewProcessedImage(img, _PixelConvertionColorfulToBw)

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
			fmt.Println("original pixel: ", originalPixel)
			processedPixel = function(originalPixel)
			// fmt.Println("i: ", i, "j: ", j, "pixel: ", processedPixel)
			newImage.Set(i, j, color.RGBA{processedPixel.r, processedPixel.g, processedPixel.b, processedPixel.a})
		}
	}

	return newImage, nil
}

func _PixelConvertionColorfulToBw(p Pixel) Pixel {
	var newPixel Pixel
	var average = p.r/3 + p.g/3 + p.b/3
	// fmt.Println("original pixel: ", p)
	newPixel.r = average
	// newPixel.r |= newPixel.r << 8
	newPixel.g = average
	// newPixel.g |= newPixel.g << 8
	newPixel.b = average
	// newPixel.b |= newPixel.b << 8
	newPixel.a = p.a
	// newPixel.a |= newPixel.a << 8
	return newPixel
}
