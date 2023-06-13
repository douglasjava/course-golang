package main

import (
	"fmt"
)

func main() {

	fmt.Println("Ponteiros")

	var variable1 int = 10
	var variable2 int = variable1

	variable1++
	fmt.Println(variable1, variable2)

	//PONTEIRO É UMA REFERENCIA DE MEMORIA
	var variable3 int
	var ponteiro *int

	variable3 = 100
	ponteiro = &variable3 //TODO apontando o endereçode memoria [&]

	fmt.Println(variable3, *ponteiro) //com [*] recuperamoso valor do 08 - Ponteiros

}
