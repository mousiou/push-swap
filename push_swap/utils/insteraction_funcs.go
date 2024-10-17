package utils

import "errors"

// PushTo pushes the first value of the stack a to the stack b
func PushTo(a, b *[]int) error {
	if len(*a) < 1 {
		return errors.New("stack length is less than one")
	}

	*b = append([]int{(*a)[0]}, (*b)...)
	*a = (*a)[1:]

	return nil
}

// Swap swaps the first two values in the stack
func Swap(stack *[]int) error {
	if len(*stack) < 2 {
		return errors.New("stack length is less than 2")
	}

	temp := (*stack)[0]
	(*stack)[0] = (*stack)[1]
	(*stack)[1] = temp

	return nil
}

// Rup rotates the stack up
func Rup(stack *[]int) error {
	if len(*stack) < 2 {
		return errors.New("stack length is less than 2")
	}

	temp := (*stack)[0]
	(*stack) = append((*stack), temp)

	*stack = (*stack)[1:]

	return nil
}

// Rdown rotates the stack down
func Rdown(stack *[]int) error {
	if len(*stack) < 2 {
		return errors.New("stack length is less than 2")
	}

	temp := (*stack)[len(*stack)-1]
	(*stack) = append([]int{temp}, (*stack)[:len(*stack)-1]...)

	return nil
}
