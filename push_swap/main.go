package main

import (
	"fmt"
	"os"
	"push-swap/utils"
)

func main() {
	// Get the arguments from the command line
	args := os.Args[1:]
	// Check arguments
	if len(args) == 0 {
		return
	} else if len(args) != 1 || !utils.Checker(args[0]) {
		fmt.Println("Error")
		return
	}

	// Convert the input string to a stack
	a := utils.ConvertToStack(args[0])
	// Initialize the stack b
	b := []int{}

	// If the stack a is already sorted, return
	if utils.IsSorted(a) {
		return
	}

	// Sort the stack a
	a = utils.Sort(a, b)

	// Print the stack a and b
	fmt.Println(a, b)
}
