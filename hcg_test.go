package coco

import "testing"

// Test HCG conversions
func TestHcg2Rgb(t *testing.T) {
	cases := []struct {
		hcg [3]float64
		rgb [3]float64
	}{
		{[3]float64{44, 98, 50}, [3]float64{252, 186, 3}},
		{[3]float64{0, 100, 0}, [3]float64{255, 0, 0}},
		{[3]float64{0, 0, 50}, [3]float64{128, 128, 128}},
	}

	for _, tc := range cases {
		result := Hcg2Rgb(tc.hcg[0], tc.hcg[1], tc.hcg[2])
		assertFloat3Equal(t, "Hcg2Rgb", "test", tc.rgb, result, 10.0)
	}
}

func TestHcg2Hsv(t *testing.T) {
	cases := []struct {
		hcg [3]float64
		hsv [3]float64
	}{
		{[3]float64{44, 98, 50}, [3]float64{44, 99, 99}},
		{[3]float64{0, 100, 0}, [3]float64{0, 100, 100}},
	}

	for _, tc := range cases {
		result := Hcg2Hsv(tc.hcg[0], tc.hcg[1], tc.hcg[2])
		assertFloat3Equal(t, "Hcg2Hsv", "test", tc.hsv, result, 10.0)
	}
}

func TestHcg2Hsl(t *testing.T) {
	cases := []struct {
		hcg [3]float64
		hsl [3]float64
	}{
		{[3]float64{44, 98, 50}, [3]float64{44, 98, 50}},
		{[3]float64{0, 100, 0}, [3]float64{0, 100, 50}},
	}

	for _, tc := range cases {
		result := Hcg2Hsl(tc.hcg[0], tc.hcg[1], tc.hcg[2])
		assertFloat3Equal(t, "Hcg2Hsl", "test", tc.hsl, result, 10.0)
	}
}

func TestHcg2Hwb(t *testing.T) {
	cases := []struct {
		hcg [3]float64
		hwb [3]float64
	}{
		{[3]float64{44, 98, 50}, [3]float64{44, 1, 1}},
		{[3]float64{0, 0, 50}, [3]float64{0, 50, 50}},
	}

	for _, tc := range cases {
		result := Hcg2Hwb(tc.hcg[0], tc.hcg[1], tc.hcg[2])
		assertFloat3Equal(t, "Hcg2Hwb", "test", tc.hwb, result, 10.0)
	}
}
