package main

import "fmt"

func sum(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func main() {

	/*
		calc := sum(25, 30)
		fmt.Println(calc)

		var f = func() {
			fmt.Println("Function Y")
		}

		f()

		var f1 = func(txt string) {
			fmt.Println(txt)
		}

		f1("Function two with parameters")

		var f2 = func(txt string) string {
			fmt.Println(txt)
			return txt
		}

		result := f2("Function two with parameters new")
		fmt.Println(result)
	*/

	sum, subtraction := calcMath(15, 60)
	fmt.Println(sum, subtraction)

	sum2, _ := calcMath(15, 60)
	fmt.Println(sum2)

}

func calcMath(n1, n2 int8) (int8, int8) {
	sum := n1 + n2
	subtraction := n1 - n2
	return sum, subtraction
}
