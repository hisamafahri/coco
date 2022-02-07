package coco

import "math"

func Cmyk2Rgb(c float64, m float64, y float64, k float64) [3]float64 {
	c = c / 100
	m = m / 100
	y = y / 100
	k = k / 100

	r := 1 - math.Min(1, c*(1-k)+k)
	g := 1 - math.Min(1, m*(1-k)+k)
	b := 1 - math.Min(1, y*(1-k)+k)

	var result [3]float64
	result[0] = math.Round(r * 255)
	result[1] = math.Round(g * 255)
	result[2] = math.Round(b * 255)

	return result
}
