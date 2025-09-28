package coco

import "testing"

// Test XYZ conversions
func TestXyz2Rgb(t *testing.T) {
	cases := []struct {
		xyz [3]float64
		rgb [3]float64
	}{
		{[3]float64{21, 18, 5}, [3]float64{166, 103, 46}},
		{[3]float64{41, 21, 2}, [3]float64{255, 0, 0}},
		{[3]float64{95, 100, 109}, [3]float64{255, 255, 255}},
	}

	for _, tc := range cases {
		result := Xyz2Rgb(tc.xyz[0], tc.xyz[1], tc.xyz[2])
		assertFloat3Equal(t, "Xyz2Rgb", "test", tc.rgb, result, 15.0)
	}
}

func TestXyz2Lab(t *testing.T) {
	cases := []struct {
		xyz [3]float64
		lab [3]float64
	}{
		{[3]float64{21, 18, 5}, [3]float64{49, 20, 41}},
		{[3]float64{95, 100, 109}, [3]float64{100, 0, 0}},
		{[3]float64{0, 0, 0}, [3]float64{0, 0, 0}},
	}

	for _, tc := range cases {
		result := Xyz2Lab(tc.xyz[0], tc.xyz[1], tc.xyz[2])
		assertFloat3Equal(t, "Xyz2Lab", "test", tc.lab, result, 15.0)
	}
}

// Test round-trip XYZ conversions
func TestXyzRoundTripConversions(t *testing.T) {
	t.Run("XYZ -> LAB -> XYZ", func(t *testing.T) {
		original := [3]float64{21, 18, 5}
		lab := Xyz2Lab(original[0], original[1], original[2])
		result := Lab2Xyz(lab[0], lab[1], lab[2])

		assertFloat3Equal(t, "XYZ->LAB->XYZ round-trip", "test", original, result, 5.0)
	})
}
