package coco

import "math"

func Lab2Xyz(l float64, a float64, b float64) [3]float64 {
	var x float64
	var y float64
	var z float64

	y = (l + 16) / 116
	x = a/500 + y
	z = y - b/200

	x2 := math.Pow(x, 3)
	y2 := math.Pow(y, 3)
	z2 := math.Pow(z, 3)

	if x2 > 0.008856 {
		x = x2
	} else {
		x = (x - 16.0/116.0) / 7.787
	}

	if y2 > 0.008856 {
		y = y2
	} else {
		y = (y - 16.0/116.0) / 7.787
	}

	if z2 > 0.008856 {
		z = z2
	} else {
		z = (z - 16.0/116.0) / 7.787
	}

	x *= 95.047
	y *= 100
	z *= 108.883

	var result [3]float64
	result[0] = math.Round(x)
	result[1] = math.Round(y)
	result[2] = math.Round(z)

	return result
}

func Lab2Lch(l float64, a float64, b float64) [3]float64 {
	var h float64

	hr := math.Atan2(b, a)
	h = hr * 360.0 / 2.0 / math.Pi

	if h < 0 {
		h += 360
	}

	c := math.Sqrt(a*a + b*b)

	var result [3]float64
	result[0] = math.Round(l)
	result[1] = math.Round(c)
	result[2] = math.Round(h)

	return result
}
