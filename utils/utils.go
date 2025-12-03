package utils

import "strconv"

func Assert(expected, actual int) {
	if expected != actual {
		panic("Assertion failed: expected " + strconv.Itoa(expected) + " but got " + strconv.Itoa(actual))
	}
}

func AssertBool(expected, actual bool) {
	if expected != actual {
		panic("Assertion failed: expected " + strconv.FormatBool(expected) + " but got " + strconv.FormatBool(actual))
	}
}
