package coco

import (
	"strconv"
)

func Hex2Rgb(input string) [3]uint8 {
	value, err := strconv.ParseUint(input, 16, 32)

	if err != nil {
		return [3]uint8{0, 0, 0}
	}

	r := uint8(value >> 16)
	g := uint8((value >> 8) & 0xFF)
	b := uint8(value & 0xFF)

	var result [3]uint8
	result[0] = r
	result[1] = g
	result[2] = b

	return result
}
