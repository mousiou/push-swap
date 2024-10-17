package utils

import (
	"strconv"
	"strings"
)

// Checker checks if the input string is valid
func Checker(arg string) bool {
	if isInvalidValues(arg) {
		return false
	}
	stack := ConvertToStack(arg)

	return !isDuplicate(stack)
}

// isInvalidValues checks if the input string contains invalid values
func isInvalidValues(arg string) bool {

	for _, item := range strings.Fields(arg) {
		_, err := strconv.Atoi(item)
		if err != nil {
			return true
		}
	}

	return false
}

// isDuplicate checks if the stack contains duplicate values
func isDuplicate(stack []int) bool {
	mapChecker := map[int]bool{}

	for _, num := range stack {
		if mapChecker[num] {
			return true
		} else {
			mapChecker[num] = true
		}
	}

	return false
}
