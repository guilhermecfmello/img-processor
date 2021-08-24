package main

import (
	"encoding/json"
	"fmt"
	"image"
	"net/http"
)

type Pixel struct {
	r, g, b int
}

func noneDescribe(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("None image processing is applied, only for tests")
}

func noneProcessing(w http.ResponseWriter, r *http.Request) {
	img, _ := ReceiveImage(r)
	ReturnImage(img, w)
}

func BwDescribe(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Convert an image to black and white channel")
}

func BwProcessing(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Black and white route hit")

	img, header := ReceiveImage(r)
	fmt.Printf("Uploaded File: %+v\n", header.Filename)
	fmt.Printf("File Size: %+v\n", header.Size)
	fmt.Printf("MIME Header: %+v\n", header.Header)

	// TODO black and white convertion here
	imgMatrix, _ := _ImageToMatrix(img)
	fmt.Println(imgMatrix)
	ReturnImage(img, w)

}

func _ImageToMatrix(img image.Image) ([][]Pixel, error) {
	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	for i := 0; i <= width; i++ {
		for j := 0; j <= height; j++ {
			// fmt.Printf("content: %v\n", img.At(i, j).RGBA())
		}
	}
	var imgMatrix [][]Pixel
	fmt.Printf("ImgToMatrix\n\tImg Bounds: %v", bounds)
	fmt.Println(bounds)

	return imgMatrix, nil
}

// func rgbaToPixel(r uint32, g uint32, b uint32, a uint32) Pixel {
// 	return Pixel{int(r / 257), int(g / 257), int(b / 257), int(a / 257)}
// }
