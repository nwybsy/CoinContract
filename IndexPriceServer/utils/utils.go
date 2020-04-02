package utils

import (
	"strconv"

	"github.com/shopspring/decimal"
)

func Int32ToString(n int32) string {
	r := strconv.FormatInt(int64(n), 10)
	return r
}

func Int64ToString(n int64) string {
	return strconv.FormatInt(n, 10)
}

func StringToInt32(n string) int32 {
	s, _ := strconv.ParseInt(n, 10, 64)
	return int32(s)
}

func StringToInt64(n string) int64 {
	s, _ := strconv.ParseInt(n, 10, 64)
	return s
}

func Float64ToString(n float64) string {
	return decimal.NewFromFloat(n).String()
}

// string -> float64
func StringToFloat64(n string) float64 {
	float, _ := strconv.ParseFloat(n, 64)
	return float
}
