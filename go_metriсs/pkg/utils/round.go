package utils

import "math"

func Round(value float64, precision int) float64 {
	p := math.Pow(10, float64(precision))
	value = math.Round(value*p) / p
	return value
}
