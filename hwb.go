package coco

import "math"

func Hwb2Rgb(h float64, w float64, b float64) [3]float64 {
	// w = w / 100
	// b = b / 100
	// var result [3]float64

	// if w+b >= 1 {
	// 	g := w / (w + b)
	// 	result[0] = math.Round(g)
	// 	result[1] = math.Round(g)
	// 	result[2] = math.Round(g)
	// 	return result
	// }

	// rgb := Hsl2Rgb(h, 100, 50)
	// for i := 0; i < 3; i++ {
	// 	rgb[i] *= (1 - w - b)
	// 	rgb[i] += w
	// }

	// result[0] = math.Round(rgb[0])
	// result[1] = math.Round(rgb[1])
	// result[2] = math.Round(rgb[2])
	// return result

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
