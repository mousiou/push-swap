package utils

import "fmt"

// SortThree sorts the stack a
func SortThree(stack []int, name string) []int {
	// If the maximum value is at the first index, rotate up
	if BigIndex(stack) == 0 {
		fmt.Println("r" + name)
		Rup(&stack)
		// If the maximum value is at the second index, rotate down
	} else if BigIndex(stack) == 1 {
		fmt.Println("rr" + name)
		Rdown(&stack)
	}

	// If the first two values are in the wrong order, swap
	if CheckFirstTwo(stack) {
		fmt.Println("s" + name)
		Swap(&stack)
	}

	return stack
}

// Sort sorts the stack a
func Sort(stackA, stackB []int) []int {
	// Sort the stack a
	for len(stackA) > 3 {
		// If the minimum value is at the last index, rotate down
		if MinIndex(stackA) == len(stackA)-1 {
			Rdown(&stackA)
			fmt.Println("ra")
		} else if MinIndex(stackA) == 1 {
			// If the minimum value is at the first index, swap
			Swap(&stackA)
			fmt.Println("sa")
		} else if MinIndex(stackA) != 0 {
			// If the minimum value is not at the first index, rotate up
			for MinIndex(stackA) < len(stackA)/2 {
				Rup(&stackA)
				fmt.Println("ra")
			}
			// If the minimum value is at the last index, rotate down
			for MinIndex(stackA) >= len(stackA)/2 {
				Rdown(&stackA)
				fmt.Println("rra")
			}
		}
		// Push the first value of the stack a to the stack b
		fmt.Println("pb")
		PushTo(&stackA, &stackB)
	}

	// Sort the stack a
	stackA = SortThree(stackA, "a")

	// Push the first value of the stack b to the stack a
	for len(stackB) > 0 {
		PushTo(&stackB, &stackA)
		fmt.Println("pa")
	}

	return stackA
}
