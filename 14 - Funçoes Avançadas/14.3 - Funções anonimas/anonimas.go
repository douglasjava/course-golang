package main

import (
	"fmt"
)

func main() {

	retorno := func(text string) string {
		return fmt.Sprintf("Recebido -> %s - %d", text, 15)
	}("Parameter")

	fmt.Println(retorno)
}
