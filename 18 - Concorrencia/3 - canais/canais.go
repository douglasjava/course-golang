package main

import (
	"fmt"
	"time"
)

func main() {

	canal := make(chan string)
	go write("Olá mundo", canal)

	/*
		for {
			message, open := <-canal
			if !open { //validarndo o status do canal
				break
			}
			fmt.Println(message)
		}
	*/

	// mesma coisa do código acima, ela já valida se o canal ta ligado
	for message := range canal {
		fmt.Println(message)
	}

	fmt.Println("Fim do programa!!")
}

func write(text string, canal chan string) {
	for i := 0; i < 5; i++ {
		canal <- text
		time.Sleep(time.Second)
	}
	close(canal) // vai fechar o canal
}
