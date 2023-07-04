package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	canal := multiplexar(write("Ola mundo"), write("Programando em GO!!"))

	for i := 0; i < 10; i++ {
		fmt.Println(<-canal)
	}

}

func multiplexar(canalEntrance1, canalEntrance2 <-chan string) <-chan string {
	canalExit := make(chan string)

	go func() {
		for {
			select {
			case message := <-canalEntrance1:
				canalExit <- message
			case message := <-canalEntrance2:
				canalExit <- message

			}
		}
	}()
	return canalExit
}

func write(text string) <-chan string {
	canal := make(chan string)

	go func() {
		for {
			canal <- fmt.Sprintf("Valor recebido: %s", text)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000)))
		}
	}()

	return canal
}
