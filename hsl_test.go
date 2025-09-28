package coco

import "testing"

// Test HSL conversions
func TestHsl2Rgb(t *testing.T) {
	cases := []struct {
		name string
		hsl  [3]float64
		rgb  [3]float64
	}{
		{"Red", [3]float64{0, 100, 50}, [3]float64{255, 0, 0}},
		{"Green", [3]float64{120, 100, 50}, [3]float64{0, 255, 0}},
		{"Blue", [3]float64{240, 100, 50}, [3]float64{0, 0, 255}},
		{"White", [3]float64{0, 0, 100}, [3]float64{255, 255, 255}},
		{"Black", [3]float64{0, 0, 0}, [3]float64{0, 0, 0}},
		{"Gray", [3]float64{0, 0, 50}, [3]float64{128, 128, 128}},
	}

	for _, tc := range cases {
		result := Hsl2Rgb(tc.hsl[0], tc.hsl[1], tc.hsl[2])
		assertFloat3Equal(t, "Hsl2Rgb", tc.name, tc.rgb, result, 10.0)
	}
}

func TestHsl2Hsv(t *testing.T) {
	cases := []struct {
		hsl [3]float64
		hsv [3]float64
	}{
		{[3]float64{136, 54, 43}, [3]float64{136, 70, 66}},
		{[3]float64{0, 100, 50}, [3]float64{0, 100, 100}},
		{[3]float64{0, 0, 0}, [3]float64{0, 0, 0}},
	}

	for _, tc := range cases {
		result := Hsl2Hsv(tc.hsl[0], tc.hsl[1], tc.hsl[2])
		assertFloat3Equal(t, "Hsl2Hsv", "test", tc.hsv, result, 10.0)
	}
}

func TestHsl2Hcg(t *testing.T) {
	cases := []struct {
		hsl [3]float64
		hcg [3]float64
	}{
		{[3]float64{136, 54, 43}, [3]float64{136, 46, 37}},
		{[3]float64{0, 100, 50}, [3]float64{0, 100, 0}},
	}

	for _, tc := range cases {
		result := Hsl2Hcg(tc.hsl[0], tc.hsl[1], tc.hsl[2])
		assertFloat3Equal(t, "Hsl2Hcg", "test", tc.hcg, result, 10.0)
	}
}
