package main

import "fmt"

func main() {
	//defer
	defer fmt.Println("Hola")
	fmt.Println("Mundo")

	//continue break
	for i := 0; i < 10; i++ {
		esPar := i % 2
		if esPar == 0 {
			break
			continue
		}
		fmt.Println(i)
	}
}
