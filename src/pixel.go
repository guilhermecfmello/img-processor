package main

type Pixel struct {
	r uint8
	g uint8
	b uint8
	a uint8
}

func (p *Pixel) setPixelValuesFromUint32(r, g, b, a uint32) {
	normalizedValues := _Uint32ToUint8([]uint32{r, g, b, a})
	p.r = normalizedValues[0]
	p.g = normalizedValues[1]
	p.b = normalizedValues[2]
	p.a = normalizedValues[3]
}

func _Uint32ToUint8(numbers []uint32) []uint8 {
	normalizedValues := []uint8{}

	for _, number := range numbers {
		normalizedValues = append(normalizedValues, uint8((float32(number)/65535)*255))
	}
	return normalizedValues
}
