package main

import (
	"fmt"
)

const (
	a = iota
	b = iota
	c = iota
)

func main() {
	//const myConst int = 42 //con la M mayuscula es norma para exportar y con minuscula es para uso interno.
	//fmt.Printf("%v, %T\n", myConst, myConst)
	fmt.Printf("%v\n", a)
	fmt.Printf("%v\n", b)
	fmt.Printf("%v\n", c)
}
