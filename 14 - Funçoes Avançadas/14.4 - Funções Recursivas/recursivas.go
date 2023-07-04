package main

import "fmt"

/**
fibonacci
*/

func fibonacci(position uint) uint {
	if position <= 1 {
		return position
	}
	return fibonacci(position-2) + fibonacci(position-1)
}

func main() {
	fmt.Println("Funções Recursivas")

	position := uint(10)

	for i := uint(0); i < position; i++ {
		fmt.Println(fibonacci(i))
	}

	//fmt.Println(fibonacci(position))

}
