package coco

import "math"

func Hsv2Rgb(h float64, s float64, v float64) [3]float64 {
	h = h / 60
	s = s / 100
	v = v / 100
	hi := float64(int(math.Floor(h)) % (6))
	f := h - math.Floor(h)
	p := 255 * v * (1 - s)
	q := 255 * v * (1 - (s * f))
	t := 255 * v * (1 - (s * (1 - f)))
	v *= 255

	var result [3]float64

	switch hi {
	case 0:
		result[0] = math.Round(v)
		result[1] = math.Round(t)
		result[2] = math.Round(p)
	case 1:
		result[0] = math.Round(q)
		result[1] = math.Round(v)
		result[2] = math.Round(p)
	case 2:
		result[0] = math.Round(p)
		result[1] = math.Round(v)
		result[2] = math.Round(t)
	case 3:
		result[0] = math.Round(p)
		result[1] = math.Round(q)
		result[2] = math.Round(v)
	case 4:
		result[0] = math.Round(t)
		result[1] = math.Round(p)
		result[2] = math.Round(v)
	case 5:
		result[0] = math.Round(v)
		result[1] = math.Round(p)
		result[2] = math.Round(q)
	}
	return result
}
