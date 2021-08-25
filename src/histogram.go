package main

import (
	"image"
)

const RGBMAX uint8 = 255

type Histogram struct {
	r [RGBMAX]int8
	g [RGBMAX]uint8
	b [RGBMAX]uint8
}

func (h *Histogram) Init(img image.Image) {
	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y
	var histogram Histogram
	for i := 0; i <= width; i++ {
		for j := 0; j <= height; j++ {
			pixel := GetPixelFromImage(img, i, j)
			histogram.r[pixel.r+1] += 1
			histogram.g[pixel.g+1] += 1
			histogram.b[pixel.b+1] += 1
		}
	}
	return histogram
}

// TODO: Terminal print of histogram chart
func (hist Histogram) toString() string {
	var print string
	var level int
	print = "======= Histogram =======\n"
	for i := 1; i <= RGBMAX; i += 2 {
		level = (hist.r[i] + hist.for ; level > 0 ; level-- {
		// 	print += ""
		// }r[i-1]) / 2
		// 
	}
	return print
}
