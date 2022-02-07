package coco

import "math"

func Rgb2Xyz(r float64, g float64, b float64) [3]float64 {
	r = r / 255
	g = g / 255
	b = b / 255

	// Assume sRGB
	if r > 0.04045 {
		r = math.Pow(((r + 0.055) / 1.055), 2.4)
	} else {
		r = (r / 12.92)
	}

	if g > 0.04045 {
		g = math.Pow(((g + 0.055) / 1.055), 2.4)
	} else {
		g = (g / 12.92)
	}

	if b > 0.04045 {
		b = math.Pow(((b + 0.055) / 1.055), 2.4)
	} else {
		b = (b / 12.92)
	}

	x := (r * 0.4124564) + (g * 0.3575761) + (b * 0.1804375)
	y := (r * 0.2126729) + (g * 0.7151522) + (b * 0.072175)
	z := (r * 0.0193339) + (g * 0.119192) + (b * 0.9503041)

	var result [3]float64
	result[0] = math.Round(x * 100)
	result[1] = math.Round(y * 100)
	result[2] = math.Round(z * 100)

	return result
}
