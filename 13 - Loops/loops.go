package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Lopps")

	i := 0

	for i < 10 {
		i++
		//fmt.Println("Incrementando i")
		time.Sleep(time.Second)
	}

	//fmt.Println(i)

	fmt.Println("----------///------------")

	for j := 0; j < 10; j++ {
		//fmt.Println("Incrementando i", j)
		time.Sleep(time.Second)
	}

	fmt.Println("----------///------------")

	names := [3]string{"Douglas", "Débora", "Leticia"}

	for indice, name := range names {
		fmt.Println(indice, name)
	}

	fmt.Println("----------///------------")

	for indice, name := range "PALAVRA" {
		fmt.Println(indice, name)
	}

	fmt.Println("----------///------------")

	for indice, name := range "PALAVRA" {
		fmt.Println(indice, string(name))
	}

	fmt.Println("MAP ----------///------------")

	user := map[string]string{
		"name":     "Douglas",
		"lastname": "Dias",
	}

	for key, value := range user {
		fmt.Println(key, value)
	}

	fmt.Println("MAP ----------///------------")

	for {
		fmt.Println("Executando infinitamente")
	}

}
