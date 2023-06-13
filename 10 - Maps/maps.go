package main

import "fmt"

func main() {
	fmt.Println("Maps")

	//dentro é chave fora é os valores
	user := map[string]string{
		"name":     "Douglas",
		"lastName": "Dias",
	}
	fmt.Println(user)
	fmt.Println(user["lastName"])

	user2 := map[string]map[string]string{
		"name": {
			"firstnama": "Douglas",
			"lastname":  "Dias",
		},
		"course": {
			"name":   "Banco de dados",
			"campus": "Pitagoras",
		},
	}

	fmt.Println(user2)
	delete(user2, "name")
	fmt.Println(user2)

}
