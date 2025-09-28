package coco

import "testing"

// Test CMYK conversions
func TestCmyk2Rgb(t *testing.T) {
	cases := []struct {
		cmyk [4]float64
		rgb  [3]float64
	}{
		{[4]float64{70, 0, 51, 34}, [3]float64{50, 168, 82}},
		{[4]float64{0, 100, 100, 0}, [3]float64{255, 0, 0}},
		{[4]float64{0, 0, 0, 0}, [3]float64{255, 255, 255}},
		{[4]float64{0, 0, 0, 100}, [3]float64{0, 0, 0}},
	}

	for _, tc := range cases {
		result := Cmyk2Rgb(tc.cmyk[0], tc.cmyk[1], tc.cmyk[2], tc.cmyk[3])
		assertFloat3Equal(t, "Cmyk2Rgb", "test", tc.rgb, result, 10.0)
	}
}
