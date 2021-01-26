package main

import (
	"fmt"
)

func main() {
	if true {
		fmt.Println("The test is true")
	}

	statePopulations := make(map[string]int)
	statePopulations = map[string]int{
		"California": 123,
		"Texas":      456,
		"Florida":    890,
	}
	if pop, ok := statePopulations["Florida"]; ok {
		fmt.Println(pop)
	}

	number := 50
	guess := 50
	if guess < 1 || guess > 100 {
		fmt.Println("The guest must be between 1 and 100!")
	} else {
		if guess < number {
			fmt.Println("Too Low")
		}
		if guess > number {
			fmt.Println("Too High")
		}
		if guess == number {
			fmt.Println("You Got it!")
		}
		fmt.Println(number <= guess, number >= guess, number != guess)
	}
	switch 3 {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	default:
		fmt.Println("Not one or two")
	}
}

func returnTrue() bool {
	fmt.Println("Returning true")
	return true
}
