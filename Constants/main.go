package main

import (
	"fmt"
)

func main() {
	const myConst int = 42 //con la M mayuscula es norma para exportar y con minuscula es para uso interno.
	fmt.Printf("%v, %T\n", myConst, myConst)
}
