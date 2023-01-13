package main

import "fmt"

func main() {
	diccionario := make(map[string]int)

	diccionario["jesus"] = 23
	diccionario["david"] = 23
	fmt.Println(diccionario)

	//recorrer diccionario
	for indice, valor := range diccionario {
		fmt.Println(indice, valor)
	}

	//encontrar valor
	valor, existeValor := diccionario["jesus"]
	fmt.Println(valor, existeValor)
}
