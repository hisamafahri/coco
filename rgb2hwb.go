package coco

import "math"

func Rgb2Hwb(r float64, g float64, b float64) [3]float64 {
	h := Rgb2Hsl(r, g, b)[0]
	w := 1.0 / 255.0 * math.Min(r, math.Min(g, b))
	b = 1.0 - 1.0/255.0*math.Max(r, math.Max(g, b))

	var result [3]float64
	result[0] = math.Round(h)
	result[1] = math.Round(w * 100)
	result[2] = math.Round(b * 100)

	return result
}
