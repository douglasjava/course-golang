package main

import (
	"fmt"
)

func sum(numbers ...int) int {
	fmt.Println(numbers)

	total := 0
	for _, number := range numbers {
		total += number
	}

	return total

}

func write(text string, numbers ...int) {

	for _, number := range numbers {
		fmt.Println(text, number)
	}

}

func main() {

	valueSum := sum(1, 2, 6, 8, 9, 55, 58, 44, 472, 12)
	fmt.Println(valueSum)

	write("Hello Go!!", 459, 121, 1, 5, 6, 97, 2)

}
