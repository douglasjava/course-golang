package main

import (
	"errors"
	"fmt"
)

func main() {

	var number1 rune = 12345
	fmt.Println(number1)

	var number2 byte = 123
	fmt.Println(number2)

	var number3 float32 = 123.45
	fmt.Println(number3)

	var number4 float64 = 124152454545454.45
	fmt.Println(number4)

	number5 := 12345.5
	fmt.Println(number5)

	number6 := "Texto1"
	fmt.Println(number6)

	char := 'B'
	fmt.Println(char) //print number of table ASCII

	var text string
	fmt.Println(text) //all type of data has a value 0

	boo1 := false
	fmt.Println(boo1)

	var bool2 bool
	fmt.Println(bool2)

	var erro error
	fmt.Println(erro)

	var erro1 error = errors.New("error internal")
	fmt.Println(erro1)

}
