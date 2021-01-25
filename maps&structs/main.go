package main

import (
	"fmt"
)

type Doctor struct {
	number     int
	actorName  string
	companions []string
}

func main() {
	statePopulations := make(map[string]int)
	statePopulations = map[string]int{
		"California": 123,
		"Texas":      456,
		"Florida":    890,
	}
	statePopulations["Georgia"] = 2222
	fmt.Println(statePopulations)
	fmt.Println(statePopulations["Florida"])
	delete(statePopulations, "Georgia")
	fmt.Println(statePopulations)
	pop, ok := statePopulations["California"]
	fmt.Println(pop, ok)
	println(len(statePopulations))

	aDoctor := Doctor{
		number:    3,
		actorName: "Jon Pertwee",
		companions: []string{
			"Liz Grant",
			"Jo Grant",
			"Sarah Jane Smith",
		},
	}
	fmt.Println(aDoctor.companions[2])
}
