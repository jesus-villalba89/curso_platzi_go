package main

import (
	"curso_platzi_go/src/mypackage"
	"fmt"
)

func main() {
	var carro mypackage.CarroPublico
	carro.Marca = "Ferrari"
	fmt.Println(carro)
}
