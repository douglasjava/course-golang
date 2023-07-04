package main

import "fmt"

func main() {
	canal := make(chan string, 2)
	canal <- "Olá Mundo"

	message := <-canal
	fmt.Println(message)

}
