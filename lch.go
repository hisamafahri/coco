package coco

import "math"

func Lch2Lab(l float64, c float64, h float64) [3]float64 {
	hr := h / 360.0 * 2 * math.Pi
	a := c * math.Cos(hr)
	b := c * math.Sin(hr)

	var result [3]float64
	result[0] = math.Round(l)
	result[1] = math.Round(a)
	result[2] = math.Round(b)

	return result
}
