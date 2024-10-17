package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"checker/utils"
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

	// Read the instructions from the standard input
	reader := bufio.NewReader(os.Stdin)
	strInstructs, err := reader.ReadString(4)
	if err != io.EOF {
		fmt.Println("Error When we read standard input")
		return
	}
	// Convert the instructions to a stack
	for _, ins := range utils.ConvertToInstructs(strInstructs) {
		// Do the instructions
		if utils.DoInstruct(&a, &b, ins) != nil {
			fmt.Println("KO")
			return
		}
	}

	// Check if the stack a is sorted
	if utils.IsSorted(a) {
		fmt.Println("OK")
	} else {
		fmt.Println("KO")
	}
}
