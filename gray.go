package coco

import "math"

func Gray2Rgb(gray float64) [3]float64 {
	var result [3]float64
	result[0] = math.Round(gray / 100 * 255)
	result[1] = math.Round(gray / 100 * 255)
	result[2] = math.Round(gray / 100 * 255)

	return result
}

func Gray2Hsl(gray float64) [3]float64 {
	var result [3]float64
	result[0] = 0
	result[1] = 0
	result[2] = math.Round(gray)

	return result
}

func Gray2Hsv(gray float64) [3]float64 {
	var result [3]float64
	result[0] = 0
	result[1] = 0
	result[2] = math.Round(gray)

	return result
}

func Gray2Hwb(gray float64) [3]float64 {
	var result [3]float64
	result[0] = 0
	result[1] = 100
	result[2] = math.Round(gray)

	return result
}
