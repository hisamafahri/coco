package coco

import "testing"

// Test RGB to other color space conversions
func TestRgb2Cmyk(t *testing.T) {
	for _, color := range testColors {
		result := Rgb2Cmyk(color.rgb[0], color.rgb[1], color.rgb[2])
		assertFloat4Equal(t, "Rgb2Cmyk", color.name, color.cmyk, result, 5.0)
	}

	// Test edge cases
	result := Rgb2Cmyk(0, 0, 0) // Black should result in K=100
	expected := [4]float64{0, 0, 0, 100}
	assertFloat4Equal(t, "Rgb2Cmyk", "black", expected, result, 1.0)
}

func TestRgb2Hsl(t *testing.T) {
	for _, color := range testColors {
		result := Rgb2Hsl(color.rgb[0], color.rgb[1], color.rgb[2])
		// More lenient tolerance for HSL conversions due to rounding
		assertFloat3Equal(t, "Rgb2Hsl", color.name, color.hsl, result, 10.0)
	}
}

func TestRgb2Hsv(t *testing.T) {
	for _, color := range testColors {
		result := Rgb2Hsv(color.rgb[0], color.rgb[1], color.rgb[2])
		assertFloat3Equal(t, "Rgb2Hsv", color.name, color.hsv, result, 10.0)
	}
}

func TestRgb2Hwb(t *testing.T) {
	for _, color := range testColors {
		result := Rgb2Hwb(color.rgb[0], color.rgb[1], color.rgb[2])
		assertFloat3Equal(t, "Rgb2Hwb", color.name, color.hwb, result, 10.0)
	}
}

func TestRgb2Lab(t *testing.T) {
	for _, color := range testColors {
		result := Rgb2Lab(color.rgb[0], color.rgb[1], color.rgb[2])
		// Lab conversions can have larger tolerances due to complex calculations
		assertFloat3Equal(t, "Rgb2Lab", color.name, color.lab, result, 15.0)
	}
}

func TestRgb2Xyz(t *testing.T) {
	for _, color := range testColors {
		result := Rgb2Xyz(color.rgb[0], color.rgb[1], color.rgb[2])
		assertFloat3Equal(t, "Rgb2Xyz", color.name, color.xyz, result, 15.0)
	}
}

func TestRgb2Hex(t *testing.T) {
	for _, color := range testColors {
		result := Rgb2Hex(color.rgb[0], color.rgb[1], color.rgb[2])
		assertStringEqual(t, "Rgb2Hex", color.name, color.hex, result)
	}

	// Test specific hex conversions
	cases := []struct {
		r, g, b  float64
		expected string
	}{
		{140, 200, 100, "8CC864"},
		{66, 135, 245, "4287F5"},
		{252, 186, 3, "FCBA03"},
	}

	for _, tc := range cases {
		result := Rgb2Hex(tc.r, tc.g, tc.b)
		if result != tc.expected {
			t.Errorf("Rgb2Hex(%v, %v, %v): expected %s, got %s", tc.r, tc.g, tc.b, tc.expected, result)
		}
	}
}

func TestRgb2Hcg(t *testing.T) {
	for _, color := range testColors {
		result := Rgb2Hcg(color.rgb[0], color.rgb[1], color.rgb[2])
		assertFloat3Equal(t, "Rgb2Hcg", color.name, color.hcg, result, 10.0)
	}
}

func TestRgb2Apple(t *testing.T) {
	for _, color := range testColors {
		result := Rgb2Apple(color.rgb[0], color.rgb[1], color.rgb[2])
		assertFloat3Equal(t, "Rgb2Apple", color.name, color.apple, result, 1.0)
	}
}

func TestRgb2Gray(t *testing.T) {
	for _, color := range testColors {
		result := Rgb2Gray(color.rgb[0], color.rgb[1], color.rgb[2])
		assertFloat1Equal(t, "Rgb2Gray", color.name, color.gray, result, 10.0)
	}
}

// Test edge cases and boundary conditions for RGB functions
func TestRgbEdgeCases(t *testing.T) {
	t.Run("RGB edge cases", func(t *testing.T) {
		// Test with maximum values
		result := Rgb2Hsl(255, 255, 255)
		expected := [3]float64{0, 0, 100}
		assertFloat3Equal(t, "Rgb2Hsl", "max values", expected, result, 1.0)

		// Test with minimum values
		result = Rgb2Hsl(0, 0, 0)
		expected = [3]float64{0, 0, 0}
		assertFloat3Equal(t, "Rgb2Hsl", "min values", expected, result, 1.0)
	})

	t.Run("Division by zero handling", func(t *testing.T) {
		// Test CMYK conversion with black (k=1)
		result := Rgb2Cmyk(0, 0, 0)
		// Should handle division by zero in CMYK calculation
		for i := 0; i < 3; i++ {
			if result[i] != result[i] { // Check for NaN
				t.Errorf("Rgb2Cmyk(0,0,0) produced NaN at index %d: %v", i, result)
			}
		}
	})

	t.Run("Hue wraparound", func(t *testing.T) {
		// Test hue values that should wrap around 360 degrees
		result := Rgb2Hsl(255, 0, 128) // Should produce hue around 330
		if result[0] < 0 || result[0] >= 360 {
			t.Errorf("Hue should be in range [0, 360), got %v", result[0])
		}
	})
}

// Test round-trip conversions to ensure consistency
func TestRgbRoundTripConversions(t *testing.T) {
	t.Run("RGB -> HSL -> RGB", func(t *testing.T) {
		original := [3]float64{140, 200, 100}
		hsl := Rgb2Hsl(original[0], original[1], original[2])
		result := Hsl2Rgb(hsl[0], hsl[1], hsl[2])

		assertFloat3Equal(t, "RGB->HSL->RGB round-trip", "test", original, result, 2.0)
	})

	t.Run("RGB -> HSV -> RGB", func(t *testing.T) {
		original := [3]float64{140, 200, 100}
		hsv := Rgb2Hsv(original[0], original[1], original[2])
		result := Hsv2Rgb(hsv[0], hsv[1], hsv[2])

		assertFloat3Equal(t, "RGB->HSV->RGB round-trip", "test", original, result, 2.0)
	})

	t.Run("RGB -> HEX -> RGB", func(t *testing.T) {
		original := [3]float64{140, 200, 100}
		hex := Rgb2Hex(original[0], original[1], original[2])
		hexResult := Hex2Rgb(hex)
		result := [3]float64{float64(hexResult[0]), float64(hexResult[1]), float64(hexResult[2])}

		assertFloat3Equal(t, "RGB->HEX->RGB round-trip", "test", original, result, 1.0)
	})
}

// Benchmark tests for RGB functions
func BenchmarkRgb2Hsl(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Rgb2Hsl(140, 200, 100)
	}
}

func BenchmarkRgb2Hex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Rgb2Hex(140, 200, 100)
	}
}

func BenchmarkRgb2Lab(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Rgb2Lab(140, 200, 100)
	}
}
