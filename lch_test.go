package coco

import "testing"

// Test LCH conversions
func TestLch2Lab(t *testing.T) {
	cases := []struct {
		lch [3]float64
		lab [3]float64
	}{
		{[3]float64{30, 79, 50}, [3]float64{30, 51, 61}},
		{[3]float64{100, 0, 0}, [3]float64{100, 0, 0}},
	}

	for _, tc := range cases {
		result := Lch2Lab(tc.lch[0], tc.lch[1], tc.lch[2])
		assertFloat3Equal(t, "Lch2Lab", "test", tc.lab, result, 15.0)
	}
}
