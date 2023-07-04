package main

import (
	"fmt"
	"os/exec"
)

func execute(cmd string) {

	out, err := exec.Command("/bin/sh", "-c", "./execute.sh").Output()

	if err != nil {
		fmt.Printf("Falha ao executar o comando %s falha -> %s", cmd, err)
	} else {
		fmt.Println("Command Successfully Executed")
	}

	output := string(out[:])
	fmt.Println(output)
}

func main() {
	execute("peer version")
}
