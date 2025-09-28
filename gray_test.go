package coco

import "testing"

// Test Gray conversions
func TestGray2Rgb(t *testing.T) {
	cases := []struct {
		gray float64
		rgb  [3]float64
	}{
		{46, [3]float64{117, 117, 117}},
		{0, [3]float64{0, 0, 0}},
		{100, [3]float64{255, 255, 255}},
		{50, [3]float64{128, 128, 128}},
	}

	for _, tc := range cases {
		result := Gray2Rgb(tc.gray)
		assertFloat3Equal(t, "Gray2Rgb", "test", tc.rgb, result, 1.0)
	}
}

func TestGray2Hsl(t *testing.T) {
	cases := []struct {
		gray float64
		hsl  [3]float64
	}{
		{46, [3]float64{0, 0, 46}},
		{0, [3]float64{0, 0, 0}},
		{100, [3]float64{0, 0, 100}},
	}

	for _, tc := range cases {
		result := Gray2Hsl(tc.gray)
		assertFloat3Equal(t, "Gray2Hsl", "test", tc.hsl, result, 1.0)
	}
}

func TestGray2Hsv(t *testing.T) {
	cases := []struct {
		gray float64
		hsv  [3]float64
	}{
		{46, [3]float64{0, 0, 46}},
		{0, [3]float64{0, 0, 0}},
		{100, [3]float64{0, 0, 100}},
	}

	for _, tc := range cases {
		result := Gray2Hsv(tc.gray)
		assertFloat3Equal(t, "Gray2Hsv", "test", tc.hsv, result, 1.0)
	}
}

func TestGray2Hwb(t *testing.T) {
	cases := []struct {
		gray float64
		hwb  [3]float64
	}{
		{46, [3]float64{0, 100, 46}},
		{0, [3]float64{0, 100, 0}},
		{100, [3]float64{0, 100, 100}},
	}

	for _, tc := range cases {
		result := Gray2Hwb(tc.gray)
		assertFloat3Equal(t, "Gray2Hwb", "test", tc.hwb, result, 1.0)
	}
}

func TestGray2Cmyk(t *testing.T) {
	cases := []struct {
		gray float64
		cmyk [4]float64
	}{
		{46, [4]float64{0, 0, 0, 46}},
		{0, [4]float64{0, 0, 0, 0}},
		{100, [4]float64{0, 0, 0, 100}},
	}

	for _, tc := range cases {
		result := Gray2Cmyk(tc.gray)
		assertFloat4Equal(t, "Gray2Cmyk", "test", tc.cmyk, result, 1.0)
	}
}

func TestGray2Lab(t *testing.T) {
	cases := []struct {
		gray float64
		lab  [3]float64
	}{
		{46, [3]float64{46, 0, 0}},
		{0, [3]float64{0, 0, 0}},
		{100, [3]float64{100, 0, 0}},
	}

	for _, tc := range cases {
		result := Gray2Lab(tc.gray)
		assertFloat3Equal(t, "Gray2Lab", "test", tc.lab, result, 1.0)
	}
}

func TestGray2Hex(t *testing.T) {
	cases := []struct {
		gray     float64
		expected string
	}{
		{46, "757575"},
		{0, "0"},
		{100, "FFFFFF"},
		{50, "808080"},
	}

	for _, tc := range cases {
		result := Gray2Hex(tc.gray)
		assertStringEqual(t, "Gray2Hex", "test", tc.expected, result)
	}
}
