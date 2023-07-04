package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("Concorrencia")

	var waitgroup sync.WaitGroup
	waitgroup.Add(2)

	go func() {
		write("Olá Mundo")
		waitgroup.Done() // -1 line 13
	}()

	go func() {
		write("Progrmando em Go!")
		waitgroup.Done() // -2 line 13
	}()

	waitgroup.Wait() // espera as duas, sincronizar a execuão dos métodos
}

func write(text string) {
	for i := 0; i < 5; i++ {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}
