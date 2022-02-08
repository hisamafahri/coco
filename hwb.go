package coco

import "math"

func Hwb2Rgb(h float64, w float64, b float64) [3]float64 {
	h = h / 360
	w = w / 100
	b = b / 100
	ratio := w + b
	var f float64

	// w + b cannot be > 1
	if ratio > 1 {
		w /= ratio
		b /= ratio
	}

	i := math.Floor(6 * h)
	v := 1 - b
	f = 6*h - i

	if i+0x01 != 0 {
		f = 1 - f
	}

	// if ((i & 0x01) !== 0) {
	// 	f = 1 - f;
	// }

	n := w + f*(v-w) // Linear interpolation
	var result [3]float64

	switch i {
	default:
	case 6:
	case 0:
		result[0] = math.Round(v * 255)
		result[1] = math.Round(n * 255)
		result[2] = math.Round(w * 255)
	case 1:
		result[0] = math.Round(n * 255)
		result[1] = math.Round(v * 255)
		result[2] = math.Round(w * 255)
	case 2:
		result[0] = math.Round(w * 255)
		result[1] = math.Round(v * 255)
		result[2] = math.Round(n * 255)
	case 3:
		result[0] = math.Round(w * 255)
		result[1] = math.Round(n * 255)
		result[2] = math.Round(v * 255)
	case 4:
		result[0] = math.Round(n * 255)
		result[1] = math.Round(w * 255)
		result[2] = math.Round(v * 255)
	case 5:
		result[0] = math.Round(v * 255)
		result[1] = math.Round(w * 255)
		result[2] = math.Round(n * 255)
	}

	return result
}

func Hwb2Hcg(h float64, w float64, b float64) [3]float64 {
	w = w / 100
	b = b / 100

	v := 1 - b
	c := v - w
	var g float64 = 0

	if c < 1 {
		g = (v - c) / (1 - c)
	}

	var result [3]float64
	result[0] = math.Round(h)
	result[1] = math.Round(c * 100)
	result[2] = math.Round(g * 100)

	return result
}
