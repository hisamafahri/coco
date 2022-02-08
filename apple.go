package coco

import "math"

func Apple2Rgb(r16 float64, g16 float64, b16 float64) [3]float64 {
	var result [3]float64
	result[0] = math.Round((r16 / 65535) * 255)
	result[1] = math.Round((g16 / 65535) * 255)
	result[2] = math.Round((b16 / 65535) * 255)

	return result
}
