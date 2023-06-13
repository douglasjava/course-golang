package main

import "fmt"

type user struct {
	name    string
	age     uint8
	address address
}

type address struct {
	name string
}

func main() {
	fmt.Println("Arquivo Structs")

	var u user
	u.name = "Douglas"
	u.age = 32
	fmt.Println(u)

	var address = address{"Rua San Marino"}

	user1 := user{"Débora", 26, address}
	fmt.Println(user1)

	user2 := user{age: 26}
	fmt.Println(user2)

}
