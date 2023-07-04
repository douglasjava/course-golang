package main

import "fmt"

/*
*
montando a interface aceita qualqer tipo
*/
func generica(interf interface{}) {
	fmt.Println(interf)
}

func main() {
	generica("String")
	generica(1)
	generica(true)
}
