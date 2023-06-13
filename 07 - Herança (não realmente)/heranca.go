package main

import (
	"fmt"
)

type person struct {
	name     string
	lastName string
	age      uint8
	hight    uint8
}

type student struct {
	person
	course  string
	college string
}

func main() {

	fmt.Println("Herança")

	p1 := person{"Douglas", "Dias", 26, 160}
	fmt.Println(p1)

	e1 := student{p1, "Analyse", "Pitagorae"}
	fmt.Println(e1)

	fmt.Println(p1.name)

}
