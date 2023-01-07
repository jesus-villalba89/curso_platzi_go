package main

import "fmt"

func main() {
	valor1 := 1
	if valor1 == 1 {
		fmt.Println("Es igual a 1")
	} else {
		fmt.Println("No es igual a 1")
	}

	//switch
	switch modulo := 5 % 2; modulo {
	case 0:
		fmt.Println("Es par")
	default:
		fmt.Println("Es impar")
	}

	//sin condición
	valor := 50
	switch {
	case valor > 100:
		fmt.Println("Es mayor a 100")
	case valor < 0:
		fmt.Println("Es menor a 0")
	default:
		fmt.Println("No condición")
	}
}
