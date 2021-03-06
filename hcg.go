package coco

import "math"

func Hcg2Rgb(h float64, c float64, g float64) [3]float64 {
	h = h / 360
	c = c / 100
	g = g / 100

	var result [3]float64

	if c == 0 {
		result[0] = math.Round(g * 255)
		result[1] = math.Round(g * 255)
		result[2] = math.Round(g * 255)
		return result
	}

	pure := [3]float64{0, 0, 0}
	hi := math.Mod(h, 1) * 6
	v := math.Mod(hi, 1)
	w := 1 - v
	var mg float64 = 0

	switch math.Floor(hi) {
	case 0:
		pure[0] = 1
		pure[1] = v
		pure[2] = 0
	case 1:
		pure[0] = w
		pure[1] = 1
		pure[2] = 0
	case 2:
		pure[0] = 0
		pure[1] = 1
		pure[2] = v
	case 3:
		pure[0] = 0
		pure[1] = w
		pure[2] = 1
	case 4:
		pure[0] = v
		pure[1] = 0
		pure[2] = 1
	default:
		pure[0] = 1
		pure[1] = 0
		pure[2] = w
	}

	mg = (1.0 - c) * g

	result[0] = math.Round((c*pure[0] + mg) * 255)
	result[1] = math.Round((c*pure[1] + mg) * 255)
	result[2] = math.Round((c*pure[2] + mg) * 255)

	return result
}

func Hcg2Hsv(h float64, c float64, g float64) [3]float64 {
	c = c / 100
	g = g / 100

	v := c + g*(1.0-c)
	var s float64 = 0

	if v > 0.0 {
		s = c / v
	}

	var result [3]float64
	result[0] = math.Round(h)
	result[1] = math.Round(s * 100)
	result[2] = math.Round(v * 100)

	return result
}

func Hcg2Hsl(h float64, c float64, g float64) [3]float64 {
	c = c / 100
	g = g / 100

	l := g*(1.0-c) + 0.5*c
	var s float64 = 0

	if l > 0.0 && l < 0.5 {
		s = c / (2 * l)
	} else if l >= 0.5 && l < 1.0 {
		s = c / (2 * (1 - l))
	}

	var result [3]float64
	result[0] = math.Round(h)
	result[1] = math.Round(s * 100)
	result[2] = math.Round(l * 100)

	return result
}

func Hcg2Hwb(h float64, c float64, g float64) [3]float64 {
	c = c / 100
	g = g / 100

	v := c + g*(1.0-c)

	var result [3]float64
	result[0] = math.Round(h)
	result[1] = math.Round((v - c) * 100)
	result[2] = math.Round((1 - v) * 100)

	return result
}
