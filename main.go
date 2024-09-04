package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	tools "push-swap/Tools"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Error in Argument")
		return
	}
	stack := BuildStack(os.Args[1])
	ThreeElem(&stack, "a")
	// stack := []int{111, 10, 1, 90, 3, 2}
	// i := 0
	// for i < len(stack) {
	// 	if stack[i]-1 > len(stack) {
	// 		tindex := len(stack) - 1
	// 		x := stack[i]
	// 		stack[i] = stack[tindex]
	// 		stack[tindex] = x
	// 	} else if stack[i]-1 != i && stack[i]-1 <= len(stack) {
	// 		tindex := stack[i] - 1
	// 		x := stack[i]
	// 		stack[i] = stack[tindex]
	// 		stack[tindex] = x
	// 	} else {
	// 		i++
	// 	}
	// }
	fmt.Println(stack)
}

// 2 2
// 1 3
// 3 1
func ThreeElem(s *[]int, name string) {
	if (*s)[len(*s)-1] == 0 { 2 == 0
		fmt.Printf("r%s\n", name)
		tools.RevRotate(s)
	}else if (*s)[len(*s)-1] == 2 {
		fmt.Printf("rr%s\n", name)
		tools.Rotate(s)
	}
	if (*s)[0] > (*s)[1] {
		fmt.Printf("s%s\n", name)
		tools.SwapFirstTwo(s)
	}
}

func BuildStack(str string) []int {
	st := strings.Fields(str)
	var stack []int

	for _, j := range st {
		nbr, err := strconv.Atoi(j)
		if err != nil {
			fmt.Printf("Value is not an integer: %v\n", j)
			return nil
		}
		stack = append(stack, nbr)
	}
	return stack
}
