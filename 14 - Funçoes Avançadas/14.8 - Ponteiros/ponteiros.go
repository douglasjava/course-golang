package main

import "fmt"

func invertSignal(number int) int {
	return number * -1
}

func main() {
	number := 20
	numeroInvertido := invertSignal(number)

	fmt.Println(numeroInvertido)
	fmt.Println(number)
}
