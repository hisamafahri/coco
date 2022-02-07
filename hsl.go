package coco

import "math"

func Hsl2Rgb(h float64, s float64, l float64) [3]float64 {
	h = h / 360
	s = s / 100
	l = l / 100

	var t2 float64
	var t3 float64
	var val float64
	var result [3]float64

	if s == 0 {
		val = l * 255
		result[0] = math.Round(val)
		result[1] = math.Round(val)
		result[2] = math.Round(val)

		return result
	}

	if l < 0.5 {
		t2 = l * (1 + s)
	} else {
		t2 = l + s - l*s
	}

	t1 := 2*l - t2

	rgb := [3]float64{0, 0, 0}

	for i := 0; i < 3; i++ {
		t3 = h + 1.0/3.0*(-(float64(i) - 1))

		if t3 < 0 {
			t3++
		}

		if t3 > 1 {
			t3--
		}

		if (6 * t3) < 1 {
			val = t1 + (t2-t1)*6*t3
		} else if (2 * t3) < 1 {
			val = t2
		} else if (3 * t3) < 2 {
			val = t1 + (t2-t1)*(2.0/3.0-t3)*6
		} else {
			val = t1
		}
		rgb[i] = val * 255
	}

	result[0] = math.Round(rgb[0])
	result[1] = math.Round(rgb[1])
	result[2] = math.Round(rgb[2])

	return result
}

func Hsl2Hsv(h float64, s float64, l float64) [3]float64 {
	s = s / 100
	l = l / 100
	smin := s
	lmin := math.Max(l, 0.01)

	l *= 2
	if l <= 1 {
		s *= l
	} else {
		s *= 2 - l
	}

	if lmin <= 1 {
		smin *= lmin
	} else {
		smin *= 2 - lmin
	}

	v := (l + s) / 2
	var sv float64

	if l == 0 {
		sv = (2 * smin) / (lmin + smin)
	} else {
		sv = (2 * s) / (l + s)
	}

	var result [3]float64
	result[0] = math.Round(h)
	result[1] = math.Round(sv * 100)
	result[2] = math.Round(v * 100)

	return result
}
