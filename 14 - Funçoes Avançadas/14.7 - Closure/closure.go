package main

import (
	"fmt"
)

func closure() func() {
	text := "Dentro da função closure"

	function := func() {
		fmt.Println(text)
	}

	return function
}

func main() {
	text := "Dentro da função main"
	fmt.Println(text)

	functionNew := closure()
	functionNew()

}
