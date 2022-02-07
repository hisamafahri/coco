package coco

import "math"

func Xyz2Rgb(x float64, y float64, z float64) [3]float64 {
	x = x / 100
	y = y / 100
	z = z / 100
	var r float64
	var g float64
	var b float64

	r = (x * 3.2404542) + (y * -1.5371385) + (z * -0.4985314)
	g = (x * -0.969266) + (y * 1.8760108) + (z * 0.041556)
	b = (x * 0.0556434) + (y * -0.2040259) + (z * 1.0572252)

	// Assume sRGB
	if r > 0.0031308 {
		r = ((1.055 * math.Pow(r, (1.0/2.4))) - 0.055)
	} else {
		r = r * 12.92
	}

	if g > 0.0031308 {
		g = ((1.055 * math.Pow(g, (1.0/2.4))) - 0.055)
	} else {
		g = g * 12.92
	}

	if b > 0.0031308 {
		b = ((1.055 * math.Pow(b, (1.0/2.4))) - 0.055)
	} else {
		b = b * 12.92
	}

	r = math.Min(math.Max(0, r), 1)
	g = math.Min(math.Max(0, g), 1)
	b = math.Min(math.Max(0, b), 1)

	var result [3]float64
	result[0] = math.Round(r * 255)
	result[1] = math.Round(g * 255)
	result[2] = math.Round(b * 255)

	return result
}
