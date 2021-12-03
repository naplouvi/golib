package golib

import "strconv"

func toInteger(str string) int {
	intValue, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}

	return intValue
}
