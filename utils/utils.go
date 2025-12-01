package utils

import "strconv"

func Assert(expected, actual int) {
	if expected != actual {
		panic("Assertion failed: expected " + strconv.Itoa(expected) + " but got " + strconv.Itoa(actual))
	}
}
