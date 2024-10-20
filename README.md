# Push-Swap

## Description

The **Push-Swap** project consists of implementing a sorting algorithm using two programs: `push-swap` and `checker`. 

- `push-swap`: This program calculates the smallest possible set of instructions to sort a stack of integers (`a` stack) in ascending order using another stack (`b` stack) and a specific set of instructions.
  
- `checker`: This program verifies if the instructions produced by `push-swap` actually sort the stack `a`. It reads the instructions, applies them, and checks if `a` is sorted while `b` is empty.

The objective is to sort the stack `a` with the least number of operations.

## Features

- **Two Programs**:
  - `push-swap`: Generates sorting instructions for the stack.
  - `checker`: Verifies the correctness of the sorting.
  
- **Sorting Instructions**:
  - `pa`: Push top element from stack `b` to stack `a`.
  - `pb`: Push top element from stack `a` to stack `b`.
  - `sa`: Swap the first two elements in stack `a`.
  - `sb`: Swap the first two elements in stack `b`.
  - `ss`: Perform `sa` and `sb` simultaneously.
  - `ra`: Rotate stack `a` (first element becomes the last).
  - `rb`: Rotate stack `b`.
  - `rr`: Perform `ra` and `rb` simultaneously.
  - `rra`: Reverse rotate stack `a` (last element becomes the first).
  - `rrb`: Reverse rotate stack `b`.
  - `rrr`: Perform `rra` and `rrb` simultaneously.

## Usage

### 1. Push-Swap Program

To sort a stack using the `push-swap` program:

```bash
$ ./push-swap "4 67 3 87 23"
pb
pb
ra
sa
rrr
pa
pa
```

In case of invalid inputs (e.g., non-integers or duplicates):

```bash
$ ./push-swap "0 one 2 3"
Error
```

### 2. Checker Program

To check if the instructions generated by `push-swap` correctly sort the stack:

```bash
$ ARG="4 67 3 87 23"; ./push-swap "$ARG" | ./checker "$ARG"
OK
```

If the instructions do not correctly sort the stack:

```bash
$ ./checker "3 2 1 0"
KO
```

### Error Handling

Both programs handle various errors, including:
- Non-integer values in the input.
- Duplicate values.
- Invalid or non-existing instructions.

In such cases, an error message is displayed:

```bash
$ ./push-swap "3 2 one"
Error
```

## Installation

To compile the programs, simply use the Go compiler:

```bash
go build -o push-swap push-swap.go
go build -o checker checker.go
```

Ensure that both `push-swap` and `checker` executables are available for use.
