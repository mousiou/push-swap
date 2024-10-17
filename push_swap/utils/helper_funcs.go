package utils

import (
	"reflect"
	"sort"
	"strconv"
	"strings"
)

// BigIndex finds the index of the maximum value in the stack
func BigIndex(stack []int) int {
	max := 0
	maxIndex := 0

	for i, num := range stack {
		if num > max {
			max = num
			maxIndex = i
		}
	}

	return maxIndex
}

// MinIndex finds the index of the minimum value in the stack
func MinIndex(stack []int) int {
	min := stack[0]
	minIndex := 0

	for i, num := range stack {
		if num < min {
			min = num
			minIndex = i
		}
	}

	return minIndex
}

// CheckFirstTwo checks if the first two values are in the correct order
func CheckFirstTwo(stack []int) bool {
	return stack[0]-stack[1] > 0
}

// ConvertToStack converts a string to a stack
func ConvertToStack(str string) []int {
	res := []int{}

	for _, item := range strings.Split(strings.TrimSpace(str), " ") {
		num, _ := strconv.Atoi(item)
		res = append(res, num)
	}

	return res
}

// IsSorted checks if the stack is sorted
func IsSorted(stack []int) bool {
	temp := make([]int, len(stack))

	copy(temp, stack)
	sort.Ints(temp)

	return reflect.DeepEqual(temp, stack)
}
