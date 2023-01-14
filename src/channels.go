package main

import "fmt"

func say(texto string, c chan<- string) {
	c <- texto
}

func main() {
	c := make(chan string)
	fmt.Println("Hola")

	go say("Bye", c)
	fmt.Println(<-c)

}
