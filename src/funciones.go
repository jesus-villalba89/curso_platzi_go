package main

import "fmt"

func imprimirMensaje(mensaje string) {
	fmt.Println(mensaje)
}

func multipleArgumentos(a, b int, c string) {
	fmt.Println(a, b, c)
}

func retornarValor(a int) int {
	return (a * 2)
}

func retornarMasUnValor(a int) (c, b int) {
	return a, a * 2
}

func main() {
	imprimirMensaje("Hola Mundo")
	multipleArgumentos(1, 2, "Hola Mundo")
	fmt.Println(retornarValor(2))
	valor1, valor2 := retornarMasUnValor(2)
	fmt.Println(valor1, valor2)
	valor1, _ = retornarMasUnValor(3)
	fmt.Println(valor1)
}
