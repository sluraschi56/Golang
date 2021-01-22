package main

import (
	"fmt"
)

func main() {
	var students [3]string
	grades := [...]int{97, 85, 93} //los 3 puntos quiere decir que inicialice el array tan grande como se carguen los datos
	fmt.Printf("Grades: %v\n", grades)
	students[0] = "Lisa"
	students[1] = "Momo"
	students[2] = "Neni"
	fmt.Printf("Students #1: %v\n", students[1])
	fmt.Printf("Numero de estudiantes: %v\n", len(students))
	var MatrizIndentidad [3][3]int
	MatrizIndentidad[0] = [3]int{1, 0, 0}
	MatrizIndentidad[1] = [3]int{0, 1, 0}
	MatrizIndentidad[2] = [3]int{0, 0, 1}
	fmt.Println(MatrizIndentidad)
	copiaGrades := &grades //el & se usa para copiar el puntero de un array
	fmt.Println(copiaGrades)
	//algunas funciones utiles son make(), append(), cap(), len()
	slice := make([]int, 10, 100) //crea un slice con un total de 10(len) elementos y capacidad de 100(cap)
	fmt.Println(slice)
}
