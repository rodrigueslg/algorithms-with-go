package module01

import (
	"fmt"
	"strconv"
	"strings"
)

// DecToBase will return a string representing
// the provided decimal number in the provided base.
// This is limited to bases 2-16 for simplicity.
//
// Eg:
//
//	DecToBase(14, 16) => "E"
//	DecToBase(14, 2) => "1110"
func DecToBase(dec, base int) string {

	if base > 16 {
		println("max base is 16")
		return ""
	}

	var result strings.Builder
	calc(dec, base, &result)

	// use the Reverse func when the calc() func does'nt append new digits to the left
	//return Reverse(result.String())

	return result.String()
}

func calc(dec, base int, sbResult *strings.Builder) int {
	if dec == 0 {
		return 0
	}

	reminder := dec % base
	dec = dec / base

	// save the current new base number string
	current := sbResult.String()

	// erase the builder
	sbResult.Reset()

	// append the new digit to the left
	sbResult.WriteString(fmt.Sprintf("%s%s", numberToLetter(reminder), current))

	return calc(dec, base, sbResult)
}

func numberToLetter(n int) string {
	switch n {
	case 10:
		return "A"
	case 11:
		return "B"
	case 12:
		return "C"
	case 13:
		return "D"
	case 14:
		return "E"
	case 15:
		return "F"
	default:
		return strconv.Itoa(n)
	}
}
