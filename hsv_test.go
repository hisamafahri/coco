package coco

import "testing"

// Test HSV conversions
func TestHsv2Rgb(t *testing.T) {
	cases := []struct {
		hsv [3]float64
		rgb [3]float64
	}{
		{[3]float64{4, 78, 92}, [3]float64{235, 64, 52}},
		{[3]float64{0, 100, 100}, [3]float64{255, 0, 0}},
		{[3]float64{0, 0, 0}, [3]float64{0, 0, 0}},
	}

	for _, tc := range cases {
		result := Hsv2Rgb(tc.hsv[0], tc.hsv[1], tc.hsv[2])
		assertFloat3Equal(t, "Hsv2Rgb", "test", tc.rgb, result, 10.0)
	}
}

func TestHsv2Hsl(t *testing.T) {
	cases := []struct {
		hsv [3]float64
		hsl [3]float64
	}{
		{[3]float64{4, 78, 92}, [3]float64{4, 82, 56}},
		{[3]float64{0, 100, 100}, [3]float64{0, 100, 50}},
	}

	for _, tc := range cases {
		result := Hsv2Hsl(tc.hsv[0], tc.hsv[1], tc.hsv[2])
		assertFloat3Equal(t, "Hsv2Hsl", "test", tc.hsl, result, 10.0)
	}
}

func TestHsv2Hcg(t *testing.T) {
	cases := []struct {
		hsv [3]float64
		hcg [3]float64
	}{
		{[3]float64{4, 78, 92}, [3]float64{4, 72, 72}},
		{[3]float64{0, 100, 100}, [3]float64{0, 100, 0}},
	}

	for _, tc := range cases {
		result := Hsv2Hcg(tc.hsv[0], tc.hsv[1], tc.hsv[2])
		assertFloat3Equal(t, "Hsv2Hcg", "test", tc.hcg, result, 10.0)
	}
}
