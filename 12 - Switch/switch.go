package main

import "fmt"

func dayWeek(number int) string {
	switch number {
	case 1:
		return "Sunday"
	case 2:
		return "Monday"
	case 3:
		return "Tuesday"
	case 4:
		return "Wednesday"
	case 5:
		return "Thursday"
	case 6:
		return "Saturday"
	case 7:
		return "Sunday"
	default:
		return "Don't exist days of the week "

	}

}

func dayWeek2(number int) string {
	var result string
	switch number {
	case 1:
		result = "Sunday"
		//fallthrough -> joga para a outra opção sem validar, cairia automatico na proxima validação
	case 2:
		result = "Monday"
	case 3:
		result = "Tuesday"
	case 4:
		result = "Wednesday"
	case 5:
		result = "Thursday"
	case 6:
		result = "Saturday"
	case 7:
		result = "Sunday"
	default:
		return "Don't exist days of the week "

	}
	return result
}

func main() {
	fmt.Println("Estrutura de controler | Switch")

	day := dayWeek(20)
	fmt.Println(day)

	day2 := dayWeek2(20)
	fmt.Println(day2)

}
