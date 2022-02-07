package coco

import (
	"math"
)

func Rgb2Cmyk(r float64, g float64, b float64) [4]float64 {
	rVal := r / 255
	gVal := g / 255
	bVal := b / 255

	k := math.Min(math.Min(1-rVal, 1-gVal), 1-bVal)

	c := (1 - rVal - k) / (1 - k)
	if math.IsNaN(c) {
		c = 0
	}

	m := (1 - gVal - k) / (1 - k)
	if math.IsNaN(m) {
		m = 0
	}

	y := (1 - bVal - k) / (1 - k)
	if math.IsNaN(y) {
		y = 0
	}

	var result [4]float64
	result[0] = math.Round(c * 100)
	result[1] = math.Round(m * 100)
	result[2] = math.Round(y * 100)
	result[3] = math.Round(k * 100)

	return result
}
