package main

import "fmt"

func main() {
	//	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { w.Write([]byte("Hola Mundo!")) })
	//	err := http.ListenAndServe(":8080", nil)
	//	if err != nil {
	//		panic(err.Error())
	//	}

	fmt.Println("start")
	panic("something bad happened") //fmt.Println("middle")
	fmt.Println("end")
}
