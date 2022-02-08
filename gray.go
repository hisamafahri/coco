package coco

import (
	"math"
	"strconv"
	"strings"
)

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

func Gray2Cmyk(gray float64) [4]float64 {
	var result [4]float64
	result[0] = 0
	result[1] = 0
	result[2] = 0
	result[3] = math.Round(gray)

	return result
}

func Gray2Lab(gray float64) [3]float64 {
	var result [3]float64
	result[0] = math.Round(gray)
	result[1] = 0
	result[2] = 0

	return result
}

func Gray2Hex(gray float64) string {
	val := int64(math.Round(gray/100*255)) & 0xFF
	integer := (val << 16) + (val << 8) + val

	myString := strings.ToUpper(strconv.FormatInt(integer, 16))
	return myString
}
