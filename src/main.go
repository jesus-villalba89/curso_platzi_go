package main

import (
	"fmt"
)

func main() {
	// Declaración de constantes
	const pi float64 = 3.1415
	const pi2 = 3.141516
	fmt.Println("pi", pi)
	fmt.Println("pi2", pi2)

	//Declaración de variable enteras
	base := 12
	var altura int = 14
	var area int
	fmt.Println(base, altura, area)

	//Zero Values
	var a int
	var b float64
	var c string
	var d bool
	fmt.Println(a, b, c, d)

	//Area cuadrado
	const cuadradoBase = 10
	cuadradoArea := cuadradoBase * cuadradoBase
	fmt.Println("Area cuadrado es:", cuadradoArea)
}
