package helper

import "math"

func ComparativeDistance(x []float64, y []float64) float64 {
	/*
		See https://en.m.wikipedia.org/wiki/Euclidean_distance#Squared_Euclidean_distance
	*/
	return math.Pow((x[0]-y[0]), 2) + math.Pow((x[1]-y[1]), 2) + math.Pow((x[2]-y[2]), 2)
}
