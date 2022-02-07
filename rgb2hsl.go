package coco

import (
	"math"
)

func Rgb2Hsl(r float64, g float64, b float64) [3]float64 {
	rVal := r / 255
	gVal := g / 255
	bVal := b / 255

	min := math.Min(math.Min(rVal, gVal), bVal)

	max := math.Max(math.Max(rVal, gVal), bVal)

	delta := max - min

	var h float64
	var s float64

	if max == min {
		h = 0
	} else if rVal == max {
		h = (gVal - bVal) / delta
	} else if gVal == max {
		h = 2 + (bVal-rVal)/delta
	} else if bVal == max {
		h = 4 + (rVal-gVal)/delta
	}

	h = math.Min(h*60, 360)

	if h < 0 {
		h += 360
	}

	l := (min + max) / 2

	if max == min {
		s = 0
	} else if l <= 0.5 {
		s = delta / (max + min)
	} else {
		s = delta / (2 - max - min)
	}

	var result [3]float64
	result[0] = math.Round(h)
	result[1] = math.Round(s * 100)
	result[2] = math.Round(l * 100)

	return result
}
