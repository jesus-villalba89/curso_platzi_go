package main

import "fmt"

func main() {
	var arreglo [4]int
	arreglo[0] = 1
	arreglo[1] = 2
	//cap => capacidad maxima del arreglo
	fmt.Println(arreglo, len(arreglo), cap(arreglo))

	//slice rebanadas
	slice := []int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println(slice, len(slice), cap(slice))

	//Métodos en el slice
	fmt.Println(slice[0])
	fmt.Println(slice[:3])
	fmt.Println(slice[2:4])
	fmt.Println(slice[4:])

	//append
	slice = append(slice, 7)
	fmt.Println(slice)
	//append con una nueva lista
	newSlice := []int{8, 9, 10}
	slice = append(slice, newSlice...)
	fmt.Println(slice)
}
