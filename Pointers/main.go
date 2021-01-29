package main

import (
	"fmt"
)

func main() {
	var a int = 42
	var b *int = &a    //declaro la variable b como puntero a un entero
	fmt.Println(&a, b) //veo ambos punteros, para referenciar se usa &
	fmt.Println(a, b)  //veo el valor de a y el puntero b
	fmt.Println(a, *b) //veo los valores de a y el valor que apunta b

	var ms *myStruct
	ms = new(myStruct)
	ms.foo = 42
	fmt.Println(ms.foo)

	//cuando trabajo con slices o maps se trabaja con punteros
}

type myStruct struct {
	foo int
}
