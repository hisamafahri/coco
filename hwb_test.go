package coco

import "testing"

// Test HWB conversions
func TestHwb2Rgb(t *testing.T) {
	cases := []struct {
		hwb [3]float64
		rgb [3]float64
	}{
		{[3]float64{136, 0, 43}, [3]float64{0, 145, 107}},
		{[3]float64{0, 0, 0}, [3]float64{255, 255, 0}},
		{[3]float64{0, 100, 0}, [3]float64{255, 255, 255}},
	}

	for _, tc := range cases {
		result := Hwb2Rgb(tc.hwb[0], tc.hwb[1], tc.hwb[2])
		assertFloat3Equal(t, "Hwb2Rgb", "test", tc.rgb, result, 10.0)
	}
}

func TestHwb2Hcg(t *testing.T) {
	cases := []struct {
		hwb [3]float64
		hcg [3]float64
	}{
		{[3]float64{136, 0, 43}, [3]float64{136, 57, 0}},
		{[3]float64{0, 50, 50}, [3]float64{0, 0, 50}},
	}

	for _, tc := range cases {
		result := Hwb2Hcg(tc.hwb[0], tc.hwb[1], tc.hwb[2])
		assertFloat3Equal(t, "Hwb2Hcg", "test", tc.hcg, result, 10.0)
	}
}
