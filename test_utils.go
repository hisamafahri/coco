package coco

import (
	"math"
	"testing"
)

// Test data structure for color conversions
type testColorData struct {
	name  string
	rgb   [3]float64
	hsl   [3]float64
	hsv   [3]float64
	hwb   [3]float64
	cmyk  [4]float64
	lab   [3]float64
	xyz   [3]float64
	lch   [3]float64
	hex   string
	hcg   [3]float64
	apple [3]float64
	gray  [1]float64
}

// Common test colors with expected values
var testColors = []testColorData{
	{
		name:  "Red",
		rgb:   [3]float64{255, 0, 0},
		hsl:   [3]float64{0, 100, 50},
		hsv:   [3]float64{0, 100, 100},
		hwb:   [3]float64{0, 0, 0},
		cmyk:  [4]float64{0, 100, 100, 0},
		lab:   [3]float64{53, 80, 67},
		xyz:   [3]float64{41, 21, 2},
		hex:   "FF0000",
		hcg:   [3]float64{0, 100, 0},
		apple: [3]float64{65535, 0, 0},
		gray:  [1]float64{33},
	},
	{
		name:  "Green",
		rgb:   [3]float64{0, 255, 0},
		hsl:   [3]float64{120, 100, 50},
		hsv:   [3]float64{120, 100, 100},
		hwb:   [3]float64{120, 0, 0},
		cmyk:  [4]float64{100, 0, 100, 0},
		lab:   [3]float64{88, -86, 83},
		xyz:   [3]float64{36, 72, 12},
		hex:   "FF00",
		hcg:   [3]float64{120, 100, 0},
		apple: [3]float64{0, 65535, 0},
		gray:  [1]float64{33},
	},
	{
		name:  "Blue",
		rgb:   [3]float64{0, 0, 255},
		hsl:   [3]float64{240, 100, 50},
		hsv:   [3]float64{240, 100, 100},
		hwb:   [3]float64{240, 0, 0},
		cmyk:  [4]float64{100, 100, 0, 0},
		lab:   [3]float64{32, 79, -108},
		xyz:   [3]float64{18, 7, 95},
		hex:   "FF",
		hcg:   [3]float64{240, 100, 0},
		apple: [3]float64{0, 0, 65535},
		gray:  [1]float64{33},
	},
	{
		name:  "White",
		rgb:   [3]float64{255, 255, 255},
		hsl:   [3]float64{0, 0, 100},
		hsv:   [3]float64{0, 0, 100},
		hwb:   [3]float64{0, 100, 0},
		cmyk:  [4]float64{0, 0, 0, 0},
		lab:   [3]float64{100, 0, 0},
		xyz:   [3]float64{95, 100, 109},
		hex:   "FFFFFF",
		hcg:   [3]float64{0, 0, 100},
		apple: [3]float64{65535, 65535, 65535},
		gray:  [1]float64{100},
	},
	{
		name:  "Black",
		rgb:   [3]float64{0, 0, 0},
		hsl:   [3]float64{0, 0, 0},
		hsv:   [3]float64{0, 0, 0},
		hwb:   [3]float64{0, 0, 100},
		cmyk:  [4]float64{0, 0, 0, 100},
		lab:   [3]float64{0, 0, 0},
		xyz:   [3]float64{0, 0, 0},
		hex:   "0",
		hcg:   [3]float64{0, 0, 0},
		apple: [3]float64{0, 0, 0},
		gray:  [1]float64{0},
	},
	{
		name:  "Gray",
		rgb:   [3]float64{128, 128, 128},
		hsl:   [3]float64{0, 0, 50},
		hsv:   [3]float64{0, 0, 50},
		hwb:   [3]float64{0, 50, 50},
		cmyk:  [4]float64{0, 0, 0, 50},
		lab:   [3]float64{54, 0, 0},
		xyz:   [3]float64{21, 22, 24},
		hex:   "808080",
		hcg:   [3]float64{0, 0, 50},
		apple: [3]float64{32896, 32896, 32896},
		gray:  [1]float64{50},
	},
}

// Helper function to compare float arrays with tolerance
func floatArraysEqual(a, b []float64, tolerance float64) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if math.Abs(a[i]-b[i]) > tolerance {
			return false
		}
	}
	return true
}

// Helper function to compare 3-element float arrays with tolerance
func float3ArraysEqual(a, b [3]float64, tolerance float64) bool {
	for i := range a {
		if math.Abs(a[i]-b[i]) > tolerance {
			return false
		}
	}
	return true
}

// Helper function to compare 4-element float arrays with tolerance
func float4ArraysEqual(a, b [4]float64, tolerance float64) bool {
	for i := range a {
		if math.Abs(a[i]-b[i]) > tolerance {
			return false
		}
	}
	return true
}

// Helper function to compare 1-element float arrays with tolerance
func float1ArraysEqual(a, b [1]float64, tolerance float64) bool {
	return math.Abs(a[0]-b[0]) <= tolerance
}

// Helper function to assert test results
func assertFloat3Equal(t *testing.T, funcName, testName string, expected, actual [3]float64, tolerance float64) {
	if !float3ArraysEqual(actual, expected, tolerance) {
		t.Errorf("%s(%s): expected %v, got %v", funcName, testName, expected, actual)
	}
}

func assertFloat4Equal(t *testing.T, funcName, testName string, expected, actual [4]float64, tolerance float64) {
	if !float4ArraysEqual(actual, expected, tolerance) {
		t.Errorf("%s(%s): expected %v, got %v", funcName, testName, expected, actual)
	}
}

func assertFloat1Equal(t *testing.T, funcName, testName string, expected, actual [1]float64, tolerance float64) {
	if !float1ArraysEqual(actual, expected, tolerance) {
		t.Errorf("%s(%s): expected %v, got %v", funcName, testName, expected, actual)
	}
}

func assertStringEqual(t *testing.T, funcName, testName string, expected, actual string) {
	if actual != expected {
		t.Errorf("%s(%s): expected %s, got %s", funcName, testName, expected, actual)
	}
}
