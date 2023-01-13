package main

import "fmt"

type computadora struct {
	memoriaRam int
	disco      int
	marca      string
}

func (miComputadora computadora) ping() {
	fmt.Println(miComputadora.marca, "Pong")
}

func (miComputadora *computadora) duplicarRam() {
	miComputadora.memoriaRam *= 2
}

//Impresión personalizada
//fmt.Println(nombreClase)
func (miComputadora computadora) String() string {
	return fmt.Sprintf("Tengo %d de GB RAM, %d GB de Disco y es un(a) %s", miComputadora.memoriaRam, miComputadora.disco, miComputadora.marca)
}

func main() {
	a := 50
	//acceder dirección de memoría
	b := &a

	fmt.Println(b)
	//acceder valor de memoria
	fmt.Println(*b)

	//modificar el valor del puntero
	*b = 100
	fmt.Println(a)

	miComputadora := computadora{memoriaRam: 8, disco: 512, marca: "HP"}
	miComputadora.ping()

	fmt.Println(miComputadora)
	miComputadora.duplicarRam()
	fmt.Println(miComputadora)

}
