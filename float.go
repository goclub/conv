package xconv

import (
	"math"
	"strconv"
)

func Float64String(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}
func Float32String(f float32) string {
	return strconv.FormatFloat(float64(f), 'f', -1, 32)
}

func Float64Round(f float64, precision uint) float64 {
	p := math.Pow10(int(precision))
	return math.Round(f*p) / p
}
