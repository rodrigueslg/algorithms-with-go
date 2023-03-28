package module01

import (
	"math"
	"strconv"
)

// BaseToDec takes in a number and the base it is currently
// in and returns the decimal equivalent as an integer.
//
// Eg:
//
//	BaseToDec("E", 16) => 14
//	BaseToDec("1110", 2) => 14
func BaseToDec(value string, base int) int {
	if base > 16 {
		println("max base is 16")
		return 0
	}

	var sum int
	for i, r := range value {
		num := letterToNumber(r)
		pow := math.Pow(float64(base), float64(len(value)-i-1))
		sum += num * int(pow)
	}
	return sum
}

func letterToNumber(r rune) int {
	switch r {
	case 'A':
		return 10
	case 'B':
		return 11
	case 'C':
		return 12
	case 'D':
		return 13
	case 'E':
		return 14
	case 'F':
		return 15
	default:
		num, _ := strconv.Atoi(string(r))
		return num
	}
}
