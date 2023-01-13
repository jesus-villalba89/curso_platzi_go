package main

import "fmt"

type carro struct {
	marca  string
	modelo int
}

func main() {
	miCarro := carro{marca: "Renault", modelo: 2020}
	fmt.Println(miCarro)

	//otra manera de instancias
	var miOtroCarro carro
	miOtroCarro.marca = "Ferrari"
	miOtroCarro.modelo = 2022
	fmt.Println(miOtroCarro)
}
