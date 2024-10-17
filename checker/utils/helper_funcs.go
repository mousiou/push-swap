package utils

import (
	"reflect"
	"sort"
	"strconv"
	"strings"
)

// ConvertToStack converts a string to a stack
func ConvertToStack(str string) []int {
	res := []int{}

	for _, item := range strings.Split(strings.TrimSpace(str), " ") {
		num, _ := strconv.Atoi(item)
		res = append(res, num)
	}

	return res
}

// ConvertToInstructs converts a string to a stack of instructions
func ConvertToInstructs(instr string) []string {
	res := []string{}

	for _, item := range strings.Split(instr, "\n") {
		res = append(res, strings.TrimSpace(item))
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

// DoInstruct does the instructions
func DoInstruct(a, b *[]int, instruction string) error {
	switch instruction {
	// Swap the first two values in the stack a
	case "sa":
		if err := Swap(a); err != nil {
			return err
		}
	// Swap the first two values in the stack b
	case "sb":
		if err := Swap(b); err != nil {
			return err
		}
	// Swap the first two values in the stack a and b
	case "ss":
		if err := Swap(a); err != nil {
			return err
		}
		if err := Swap(b); err != nil {
			return err
		}
	// Rotate the stack a up
	case "ra":
		if err := Rup(a); err != nil {
			return err
		}
	// Rotate the stack b up
	case "rb":
		if err := Rup(b); err != nil {
			return err
		}
	// Rotate the stack a and b up
	case "rr":
		if err := Rup(a); err != nil {
			return err
		}
		if err := Rup(b); err != nil {
			return err
		}
	// Rotate the stack a down
	case "rra":
		if err := Rdown(a); err != nil {
			return err
		}
	// Rotate the stack b down
	case "rrb":
		if err := Rdown(b); err != nil {
			return err
		}
	// Rotate the stack a and b down
	case "rrr":
		if err := Rdown(a); err != nil {
			return err
		}
		if err := Rdown(b); err != nil {
			return err
		}
	// Push the first value of the stack a to the stack b
	case "pb":
		if err := PushTo(a, b); err != nil {
			return err
		}
	// Push the first value of the stack b to the stack a
	case "pa":
		if err := PushTo(b, a); err != nil {
			return err
		}
	}

	return nil
}
