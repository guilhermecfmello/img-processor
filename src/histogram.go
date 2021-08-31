package main

import (
	"fmt"
	"image"
	"image/color"
)

const RGBMAX uint = 256

type Histogram struct {
	r [RGBMAX]uint
	g [RGBMAX]uint
	b [RGBMAX]uint

	rNormalized [RGBMAX]float64
	gNormalized [RGBMAX]float64
	bNormalized [RGBMAX]float64

	rAccumulated [RGBMAX]float64
	gAccumulated [RGBMAX]float64
	bAccumulated [RGBMAX]float64

	pixelQtd int
	width    int
	height   int
}

func (histogram *Histogram) Init(img image.Image) {
	bounds := img.Bounds()
	histogram.width, histogram.height = bounds.Max.X, bounds.Max.Y
	histogram.pixelQtd = histogram.width * histogram.height
	for i := 0; i <= histogram.width; i++ {
		for j := 0; j <= histogram.height; j++ {
			pixel := GetPixelFromImage(img, i, j)
			histogram.r[pixel.r] += 1
			histogram.g[pixel.g] += 1
			histogram.b[pixel.b] += 1
		}
	}
}

func (histogram *Histogram) Normalize() {
	rgbmaxIntCast := int(RGBMAX) - 1
	for i := 0; i <= rgbmaxIntCast; i++ {
		histogram.rNormalized[i] = float64(histogram.r[i]) / float64(histogram.pixelQtd)
		histogram.gNormalized[i] = float64(histogram.g[i]) / float64(histogram.pixelQtd)
		histogram.bNormalized[i] = float64(histogram.b[i]) / float64(histogram.pixelQtd)
	}
}

func (histogram *Histogram) Accumulate() {
	rgbmaxIntCast := int(RGBMAX) - 1
	for i := 1; i <= rgbmaxIntCast; i++ {
		histogram.rAccumulated[i] = histogram.rNormalized[i] + histogram.rAccumulated[i-1]
		histogram.gAccumulated[i] = histogram.gNormalized[i] + histogram.gAccumulated[i-1]
		histogram.bAccumulated[i] = histogram.bNormalized[i] + histogram.bAccumulated[i-1]
	}
}

func (histogram *Histogram) GenerateEqualizedImage(img image.Image) image.Image {
	upLeft := image.Point{0, 0}
	lowRight := image.Point{histogram.width, histogram.height}
	newImage := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	for i := 0; i <= histogram.width; i++ {
		for j := 0; j <= histogram.height; j++ {
			originalPixel := GetPixelFromImage(img, i, j)
			fmt.Printf("\n[%v,%v]\n", i, j)
			newPixel := _EqualizePixel(originalPixel, histogram)
			newImage.Set(i, j, color.RGBA{newPixel.r, newPixel.g, newPixel.b, newPixel.a})
		}
	}

	return newImage
}

func _EqualizePixel(originalPixel *Pixel, histogram *Histogram) *Pixel {
	var newPixel Pixel
	newPixel.r = uint8(histogram.rAccumulated[originalPixel.r] * float64(255))
	newPixel.g = uint8(histogram.gAccumulated[originalPixel.g] * float64(255))
	newPixel.b = uint8(histogram.bAccumulated[originalPixel.b] * float64(255))
	newPixel.a = originalPixel.a
	fmt.Println("\nold: R: ", originalPixel.r, "G: ", originalPixel.g, "B: ", originalPixel.b)
	fmt.Printf("\thist.R: %v hist.G: %v, hist.B: %v\n", histogram.rAccumulated[originalPixel.r], histogram.gAccumulated[originalPixel.g], histogram.bAccumulated[originalPixel.b])
	fmt.Println("new: R: ", newPixel.r, "G: ", newPixel.g, "B: ", newPixel.b)
	return &newPixel
}

// TODO: Terminal print of histogram chart
func (hist Histogram) toString() string {
	return ""
}
