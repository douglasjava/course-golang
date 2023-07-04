package main

import (
	"fmt"
)

func funcao1() {
	fmt.Println("Executando a função 1")
}

func funcao2() {
	fmt.Println("Executando a função 2")
}

func studentApproved(n1, n2 float32) bool {
	fmt.Println("Entrando na função para verificar se o aluno está aprovado")

	media := (n1 + n2) / 2
	if media >= 6 {
		return true
	}
	return false

}

func main() {
	defer funcao1()
	funcao2()

	fmt.Println(studentApproved(7, 8))
}
