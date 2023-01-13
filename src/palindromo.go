package main

import (
	"fmt"
	"strings"
)

func obtenerSiEsPalindromo(palabra string) bool {
	palabra = strings.ToLower(palabra)
	var palabraInversa string
	for i := len(palabra) - 1; i >= 0; i-- {
		palabraInversa += string(palabra[i])
	}

	fmt.Println(palabra)
	fmt.Println(palabraInversa)
	return (palabraInversa == palabra)
}

func main() {
	var palabra string
	fmt.Scanln("%s", &palabra)
	fmt.Println(palabra)

	if obtenerSiEsPalindromo(palabra) {
		fmt.Println("Es Palindromo")
	} else {
		fmt.Println("No es Palindromo")
	}
}
