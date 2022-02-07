package coco

import (
	"math"
)

func Rgb2Hsl(r float64, g float64, b float64) [3]float64 {
	r = r / 255
	g = g / 255
	b = b / 255

	min := math.Min(math.Min(r, g), b)

	max := math.Max(math.Max(r, g), b)

	delta := max - min

	var h float64
	var s float64

	if max == min {
		h = 0
	} else if r == max {
		h = (g - b) / delta
	} else if g == max {
		h = 2 + (b-r)/delta
	} else if b == max {
		h = 4 + (r-g)/delta
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
