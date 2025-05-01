package lexical_order

import (
	"fmt"
	"math"
)

// RoundToNearestEven 반올림하되 .5에서 짝수로 반올림
func RoundToNearestEven(x float64) int {
	fraction := x - math.Floor(x)
	d := int(x - fraction)

	if fraction < 0.5 {
		return d
	}
	if fraction > 0.5 {
		return d + 1
	}
	if d%2 == 0 {
		return d
	}
	return d + 1
}

// LogBase 계산: log_b(n) = log(n) / log(b)
func LogBase(n, base float64) float64 {
	return math.Log(n) / math.Log(base)
}

// Pow 계산: a^b
func Pow(a, b float64) float64 {
	return math.Pow(a, b)
}

func CharAt(s string, index int) rune {
	runes := []rune(s)
	if index < 0 || index >= len(runes) {
		panic(fmt.Sprintf("유효하지 않은 인덱스입니다. %s %d", s, index))
	}
	return runes[index]
}

type KeySpace struct {
	keys       []string
	keyToIndex map[rune]int
}
