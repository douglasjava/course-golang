package main

import (
	"fmt"
	"github.com/badoux/checkmail"
	"modulo/auxiliar"
)

func main() {

	auxiliar.Write()

	fmt.Println("Escrevendo do arquivo main")

	error := checkmail.ValidateFormat("douglas.marques@montreal.com.br")
	fmt.Println(error)

}
