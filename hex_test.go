package coco

import "testing"

// Test HEX conversions
func TestHex2Rgb(t *testing.T) {
	cases := []struct {
		hex string
		rgb [3]uint8
	}{
		{"4287F5", [3]uint8{66, 135, 245}},
		{"FF0000", [3]uint8{255, 0, 0}},
		{"00FF00", [3]uint8{0, 255, 0}},
		{"0000FF", [3]uint8{0, 0, 255}},
		{"FFFFFF", [3]uint8{255, 255, 255}},
		{"000000", [3]uint8{0, 0, 0}},
		{"808080", [3]uint8{128, 128, 128}},
	}

	for _, tc := range cases {
		result := Hex2Rgb(tc.hex)
		if result != tc.rgb {
			t.Errorf("Hex2Rgb(%s): expected %v, got %v", tc.hex, tc.rgb, result)
		}
	}

	// Test invalid hex
	result := Hex2Rgb("INVALID")
	expected := [3]uint8{0, 0, 0}
	if result != expected {
		t.Errorf("Hex2Rgb(invalid): expected %v, got %v", expected, result)
	}
}

// Benchmark test for Hex2Rgb
func BenchmarkHex2Rgb(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Hex2Rgb("8CC864")
	}
}
