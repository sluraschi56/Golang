package main

import (
	"fmt"
)

func main() {
	arr := [...][]int{
		{20, 1, -1, 2, -2, 3, 3, 5, 5, 1, 2, 4, 20, 4, -1, -2, 5},
		{1, 1, 2, -2, 5, 2, 4, 4, -1, -2, 5},
		{20, 1, 1, 2, 2, 3, 3, 5, 5, 4, 20, 4, 5},
		{10},
		{1, 1, 1, 1, 1, 1, 10, 1, 1, 1, 1},
		{5, 4, 3, 2, 1, 5, 4, 3, 2, 10, 10},
	}
	for i := 0; i < len(arr); i++ {
		FindOdd(arr[i])
	}
}

func FindOdd(seq []int) int {
	length := len(seq)
	var odd_number int = seq[0]
	for i := 1; i < length; i++ {
		// Perform XOR (^) operator to all elements in array
		odd_number = odd_number ^ seq[i]
	}
	fmt.Printf("\nOdd number occurance is: %d", odd_number)
	return odd_number
}
