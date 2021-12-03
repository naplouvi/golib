package golib

import "strconv"

// ToInteger takes a string an returns an integer
func ToInteger(str string) int {
	intValue, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}

	return intValue
}
