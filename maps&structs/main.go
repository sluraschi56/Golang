package main

import (
	"fmt"
	"reflect"
)

type Doctor struct {
	number     int
	actorName  string
	companions []string
}

//En Golang no hay herencia pero se puede hacer los siguiente, haciendo una composicion en Bird con las caracteristicas de Animal
type Animal struct {
	Name   string `required max:"100"`
	Origin string
}
type Bird struct {
	Animal
	SpeedKPH float32
	CanFly   bool
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

	b := Bird{}
	b.Name = "Emu"
	b.Origin = "Australia"
	b.SpeedKPH = 48
	b.CanFly = false
	fmt.Println(b)
	//uso de tags en los campos de un struct
	t := reflect.TypeOf(Animal{})
	field, _ := t.FieldByName("Name")
	fmt.Println(field.Tag)
}
