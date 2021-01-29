package main

import (
	"fmt"
)

func main() {
	sayMessage("Hola Go!")
	greeting := "hello"
	name := "Sebastian"
	sayGreeting(&greeting, &name) // estoy pasando los punteros de las varables para que la funcion pueda modificar el valor
	fmt.Println(name)

	sum(1, 2, 3, 4, 5)
	s := sum_return(1, 2, 3, 4, 5)
	fmt.Println("The sum with return is", s)

	d, err := divide(5.0, 2.5)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(d)

	func() {
		fmt.Println("Hello Go!")
	}() // cuando llamo una funcion sin nombre se llaman anonymous function y se ejecuta inmediatamente con los parentesis al final

	// defino una variavle como type greeter
	g := greeter{
		greeting: "Hello",
		name:     "Go",
	}
	g.greet()
}

func sayMessage(msg string) {
	fmt.Println(msg)
}

func sayGreeting(greeting, name *string) {
	fmt.Println(*greeting, *name)
	*name = "Nicole"
	fmt.Println(*name)
}

func sum(values ...int) { //los tres puntos hacen referencia a un slice
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}
	fmt.Println("The sum is ", result)
}

func sum_return(values ...int) (result int) { //los tres puntos hacen referencia a un slice y con el int final defino la variable de retorno result
	fmt.Println(values)
	for _, v := range values {
		result += v
	}
	return // uso return para devolver el valor de result
}

func divide(a, b float64) (float64, error) {
	if b == 0.0 {
		return 0.0, fmt.Errorf("Cannot divide by zero")
	}
	return a / b, nil
}

// metodos como funciones.  Son funciones definidas con tipos de datos

type greeter struct {
	greeting string
	name     string
}

func (g greeter) greet() { // se define (g greeter) como receiver, si trabajo con punteros en *greeter se puede manipular el valor excepto con slices y maps que siempre se manipulan
	fmt.Println(g.greeting, g.name)
}
