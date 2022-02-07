package coco

import (
	"math"
)

func Rgb2Hsv(r float64, g float64, b float64) [3]float64 {
	var rdif float64
	var gdif float64
	var bdif float64
	var h float64
	var s float64

	r = r / 255
	g = g / 255
	b = b / 255

	min := math.Min(math.Min(r, g), b)

	v := math.Max(math.Max(r, g), b)

	diff := v - min

	diffc := func(c float64) float64 {
		return (v-c)/6.0/diff + (1.0 / 2.0)
	}

	if diff == 0 {
		h = 0
		s = 0
	} else {
		s = diff / v
		rdif = diffc(r)
		gdif = diffc(g)
		bdif = diffc(b)

		if r == v {
			h = bdif - gdif
		} else if g == v {
			h = (1.0 / 3.0) + rdif - bdif
		} else if b == v {
			h = (2.0 / 3.0) + gdif - rdif
		}

		if h < 0 {
			h += 1
		} else if h > 1 {
			h -= 1
		}
	}

	var result [3]float64
	result[0] = math.Round(h * 360)
	result[1] = math.Round(s * 100)
	result[2] = math.Round(v * 100)

	return result
}
