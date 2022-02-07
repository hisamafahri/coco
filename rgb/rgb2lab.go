package coco

import "math"

func Rgb2Lab(r float64, g float64, b float64) [3]float64 {
	xyz := Rgb2Xyz(r, g, b)
	x := xyz[0]
	y := xyz[1]
	z := xyz[2]

	x /= 95.047
	y /= 100
	z /= 108.883

	if x > 0.008856 {
		x = math.Pow(x, (1.0 / 3.0))
	} else {
		x = (7.787 * x) + (16.0 / 116.0)
	}

	if y > 0.008856 {
		y = math.Pow(y, (1.0 / 3.0))
	} else {
		y = (7.787 * y) + (16.0 / 116.0)
	}

	if z > 0.008856 {
		z = math.Pow(z, (1.0 / 3.0))
	} else {
		z = (7.787 * z) + (16.0 / 116.0)
	}

	l := (116 * y) - 16
	a := 500 * (x - y)
	b = 200 * (y - z)

	var result [3]float64
	result[0] = math.Round(l)
	result[1] = math.Round(a)
	result[2] = math.Round(b)

	return result
}
