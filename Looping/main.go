package main

import "fmt"

func main() {
	//for i := 0; i < 5; i++ {
	//	fmt.Println(i)
	//}
	// incrementar con dos variables
	//for i, j := 0, 0; i < 5; i, j = i+1, j+1 {
	//		fmt.Println(i, j)
	//	}
	//el while se hace con for
	//for i := 0; i < 10; i++ { //solo imprime los valores impares y se usa continue para saltar e ir a la proxima iteracion
	//	if i%2 == 0 {
	//		continue
	//	}
	//	fmt.Println(i)
	//}
	//i := 0
	//for {
	//	fmt.Println(i)
	//	i++
	//	if i == 5 {
	//		break // se usa break para salir del for mas cercano
	//	}
	//}

Loop:
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Println(i * j)
			if i*j >= 3 {
				break Loop
			}
		}
	}

	// recorrer array con for
	s := []int{1, 2, 3}
	for k, v := range s {
		fmt.Println(k, v)
	}
}
