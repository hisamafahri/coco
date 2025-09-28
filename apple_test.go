package coco

import "testing"

// Test Apple conversions
func TestApple2Rgb(t *testing.T) {
	cases := []struct {
		apple [3]float64
		rgb   [3]float64
	}{
		{[3]float64{64764, 47802, 771}, [3]float64{252, 186, 3}},
		{[3]float64{65535, 0, 0}, [3]float64{255, 0, 0}},
		{[3]float64{0, 65535, 0}, [3]float64{0, 255, 0}},
		{[3]float64{0, 0, 65535}, [3]float64{0, 0, 255}},
		{[3]float64{65535, 65535, 65535}, [3]float64{255, 255, 255}},
		{[3]float64{0, 0, 0}, [3]float64{0, 0, 0}},
	}

	for _, tc := range cases {
		result := Apple2Rgb(tc.apple[0], tc.apple[1], tc.apple[2])
		assertFloat3Equal(t, "Apple2Rgb", "test", tc.rgb, result, 1.0)
	}
}
