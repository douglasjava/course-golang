package main

import "fmt"

func main() {
	fmt.Println("Estrutura de controle")

	number := -50

	if number > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor ou igual a 15")
	}

	if otherNumber := number; otherNumber > 0 { // variavel só exist no scope atual
		fmt.Println("Numero é maior que zer")
	} else {
		fmt.Println("Número é menor que zxero")
	}

}
