package main

import "fmt"

func studentApproved(n1, n2 float64) bool {
	defer recoverRunningTime()

	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	panic("A MÉDIA É EXATAMENTE 6!")

}

func recoverRunningTime() {
	fmt.Println("Tentativa de recuperar a execução")

	if r := recover(); r != nil {
		fmt.Println("Execução recuperada com sucesso!!")
	}

}

func main() {
	fmt.Println(studentApproved(6, 6))
	fmt.Println("Pós execution")
}
