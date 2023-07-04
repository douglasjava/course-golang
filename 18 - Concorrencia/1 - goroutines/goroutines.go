package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Concorrencia")

	go write("Olá Mundo") // executa esse método e não espera o retorno
	write("Progrmando em Go!")

}

func write(text string) {
	for {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}
