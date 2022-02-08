package coco

import "math"

func Gray2Rgb(gray float64) [3]float64 {
	var result [3]float64
	result[0] = math.Round(gray / 100 * 255)
	result[1] = math.Round(gray / 100 * 255)
	result[2] = math.Round(gray / 100 * 255)

	return result
}
