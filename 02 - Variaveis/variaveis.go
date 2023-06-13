package main

import "fmt"

func main() {
	var variable1 string = "variable1"
	variable2 := "variable2"
	var (
		variable3 string = "variable3"
		variable4 string = "variable4"
	)

	variable5, variable6 := "variable5", "variable6"

	fmt.Println(variable1)
	fmt.Println(variable2)
	fmt.Println(variable3)
	fmt.Println(variable4)
	fmt.Println(variable5)
	fmt.Println(variable6)

	variable5, variable6 = variable6, variable5

	fmt.Println(variable5)
	fmt.Println(variable6)

}
