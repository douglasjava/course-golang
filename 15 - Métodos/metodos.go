package main

import "fmt"

type user struct {
	name  string
	idade uint8
}

func (u user) salvar() {
	fmt.Println("Dentro do método salvar")
}

func (u user) maiorDeIdade() bool {
	return u.idade >= 18
}

func (u *user) fazerAniversaio() {
	u.idade++
}

func main() {
	user1 := user{"Usuário 1", 20}
	fmt.Println(user1)
	user1.salvar()
	user1.maiorDeIdade()
	user1.fazerAniversaio()

	user2 := user{"Usuário 2", 32}
	fmt.Println(user1)
	user2.salvar()

	user2.fazerAniversaio()
	fmt.Println(user2.idade)

}
