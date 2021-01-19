package main

import (
	"fmt"
)

func main() {
	//tres maneras de definir variables:
	//	var i int
	//	var j int = 27
	k := 99.
	fmt.Printf("%v, %T", k, k)
}

//no se pueden redeclarar variables dentro del programa, pero se pueden hacer shadow
//todas las variables deben ser usadas sino tira error
//en minuscula la primer letra es usada a nivel package
//en mayuscula la primer letra es usada a nivel de exportar
//no hay variables privadas
//type conversion con el package strconv
