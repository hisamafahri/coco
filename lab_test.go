package coco

import "testing"

// Test LAB conversions
func TestLab2Xyz(t *testing.T) {
	cases := []struct {
		lab [3]float64
		xyz [3]float64
	}{
		{[3]float64{50, 21, 38}, [3]float64{22, 18, 6}},
		{[3]float64{100, 0, 0}, [3]float64{95, 100, 109}},
		{[3]float64{0, 0, 0}, [3]float64{0, 0, 0}},
	}

	for _, tc := range cases {
		result := Lab2Xyz(tc.lab[0], tc.lab[1], tc.lab[2])
		assertFloat3Equal(t, "Lab2Xyz", "test", tc.xyz, result, 15.0)
	}
}

func TestLab2Lch(t *testing.T) {
	cases := []struct {
		lab [3]float64
		lch [3]float64
	}{
		{[3]float64{50, 21, 38}, [3]float64{50, 43, 61}},
		{[3]float64{100, 0, 0}, [3]float64{100, 0, 0}},
	}

	for _, tc := range cases {
		result := Lab2Lch(tc.lab[0], tc.lab[1], tc.lab[2])
		assertFloat3Equal(t, "Lab2Lch", "test", tc.lch, result, 15.0)
	}
}

// Test round-trip LAB conversions
func TestLabRoundTripConversions(t *testing.T) {
	t.Run("LAB -> LCH -> LAB", func(t *testing.T) {
		original := [3]float64{50, 21, 38}
		lch := Lab2Lch(original[0], original[1], original[2])
		result := Lch2Lab(lch[0], lch[1], lch[2])

		assertFloat3Equal(t, "LAB->LCH->LAB round-trip", "test", original, result, 5.0)
	})
}
