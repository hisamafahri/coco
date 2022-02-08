package coco

import (
	"math"
	"strconv"
	"strings"
)

func Rgb2Cmyk(r float64, g float64, b float64) [4]float64 {
	r = r / 255
	g = g / 255
	b = b / 255

	k := math.Min(math.Min(1-r, 1-g), 1-b)

	c := (1 - r - k) / (1 - k)
	if math.IsNaN(c) {
		c = 0
	}

	m := (1 - g - k) / (1 - k)
	if math.IsNaN(m) {
		m = 0
	}

	y := (1 - b - k) / (1 - k)
	if math.IsNaN(y) {
		y = 0
	}

	var result [4]float64
	result[0] = math.Round(c * 100)
	result[1] = math.Round(m * 100)
	result[2] = math.Round(y * 100)
	result[3] = math.Round(k * 100)

	return result
}

func Rgb2Hsl(r float64, g float64, b float64) [3]float64 {
	r = r / 255
	g = g / 255
	b = b / 255

	min := math.Min(math.Min(r, g), b)

	max := math.Max(math.Max(r, g), b)

	delta := max - min

	var h float64
	var s float64

	if max == min {
		h = 0
	} else if r == max {
		h = (g - b) / delta
	} else if g == max {
		h = 2 + (b-r)/delta
	} else if b == max {
		h = 4 + (r-g)/delta
	}

	h = math.Min(h*60, 360)

	if h < 0 {
		h += 360
	}

	l := (min + max) / 2

	if max == min {
		s = 0
	} else if l <= 0.5 {
		s = delta / (max + min)
	} else {
		s = delta / (2 - max - min)
	}

	var result [3]float64
	result[0] = math.Round(h)
	result[1] = math.Round(s * 100)
	result[2] = math.Round(l * 100)

	return result
}

func Rgb2Hsv(r float64, g float64, b float64) [3]float64 {
	var rdif float64
	var gdif float64
	var bdif float64
	var h float64
	var s float64

	r = r / 255
	g = g / 255
	b = b / 255

	min := math.Min(math.Min(r, g), b)

	v := math.Max(math.Max(r, g), b)

	diff := v - min

	diffc := func(c float64) float64 {
		return (v-c)/6.0/diff + (1.0 / 2.0)
	}

	if diff == 0 {
		h = 0
		s = 0
	} else {
		s = diff / v
		rdif = diffc(r)
		gdif = diffc(g)
		bdif = diffc(b)

		if r == v {
			h = bdif - gdif
		} else if g == v {
			h = (1.0 / 3.0) + rdif - bdif
		} else if b == v {
			h = (2.0 / 3.0) + gdif - rdif
		}

		if h < 0 {
			h += 1
		} else if h > 1 {
			h -= 1
		}
	}

	var result [3]float64
	result[0] = math.Round(h * 360)
	result[1] = math.Round(s * 100)
	result[2] = math.Round(v * 100)

	return result
}

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

func Rgb2Xyz(r float64, g float64, b float64) [3]float64 {
	r = r / 255
	g = g / 255
	b = b / 255

	// Assume sRGB
	if r > 0.04045 {
		r = math.Pow(((r + 0.055) / 1.055), 2.4)
	} else {
		r = (r / 12.92)
	}

	if g > 0.04045 {
		g = math.Pow(((g + 0.055) / 1.055), 2.4)
	} else {
		g = (g / 12.92)
	}

	if b > 0.04045 {
		b = math.Pow(((b + 0.055) / 1.055), 2.4)
	} else {
		b = (b / 12.92)
	}

	x := (r * 0.4124564) + (g * 0.3575761) + (b * 0.1804375)
	y := (r * 0.2126729) + (g * 0.7151522) + (b * 0.072175)
	z := (r * 0.0193339) + (g * 0.119192) + (b * 0.9503041)

	var result [3]float64
	result[0] = math.Round(x * 100)
	result[1] = math.Round(y * 100)
	result[2] = math.Round(z * 100)

	return result
}

func Rgb2Hex(r float64, g float64, b float64) string {
	integer := ((int64(math.Round(r)) & 0xFF) << 16) + ((int64(math.Round(g)) & 0xFF) << 8) + (int64(math.Round(b)) & 0xFF)

	myString := strings.ToUpper(strconv.FormatInt(integer, 16))
	return myString
}

// func Rgb2Ansi16(r float64, g float64, b float64) float64 {
// 	value := Rgb2Hsv(r, g, b)[2] // Hsv -> ansi16 optimization

// 	/*
// 		// should check the saturation value
// 		var value float64

// 		if saturation == nil {
// 			value = Rgb2Hsv(r, g, b)[2]
// 		} else {
// 			value = saturation
// 		}
// 	*/

// 	value = math.Round(value / 50)

// 	if value == 0 {
// 		return 30
// 	}

// 	rBit := math.Round(r / 255)
// 	gBit := math.Round(g / 255)
// 	bBit := math.Round(b / 255)

// 	ansi := (30 + (int64(bBit) << 2) | (int64(gBit) << 1) | int64(rBit))

// 	fmt.Println(ansi)

// 	if value == 2 {
// 		ansi += 60
// 	}

// 	return float64(ansi)
// }
