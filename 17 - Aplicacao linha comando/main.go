package main

import (
	"fmt"
	"linha-de-comando/app"
	"log"
	"os"
)

func main() {
	fmt.Println("Ponto de partida")

	application := app.Gerar()
	if errorReturn := application.Run(os.Args); errorReturn != nil {
		log.Fatal(errorReturn)
	}

}
