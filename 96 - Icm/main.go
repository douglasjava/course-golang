package main

import (
	"fmt"
	"icm/calculos"
)

func main() {

	fmt.Println("------------- BAIRRO DA GRAÇA ---------------------")
	response1 := calculos.CalcAverageFrequency("./calculos/frequency-bairro-graca.json")
	fmt.Println(response1)

	fmt.Println("------------- CARLOS PRATES ---------------------")
	response2 := calculos.CalcAverageFrequency("./calculos/frequency-carlos-prates.json")
	fmt.Println(response2)

	fmt.Println("-------------  CENTRAL ---------------------")
	response3 := calculos.CalcAverageFrequency("./calculos/frequency-central.json")
	fmt.Println(response3)

	fmt.Println("------------- CIDADE NOVA ---------------------")
	response4 := calculos.CalcAverageFrequency("./calculos/frequency-cidade-nova.json")
	fmt.Println(response4)

	fmt.Println("------------- COLORADO ---------------------")
	response5 := calculos.CalcAverageFrequency("./calculos/frequency-colorado.json")
	fmt.Println(response5)

	fmt.Println("------------- NACIONAL ---------------------")
	response6 := calculos.CalcAverageFrequency("./calculos/frequency-nacional.json")
	fmt.Println(response6)

	fmt.Println("------------- PALMARES ---------------------")
	response7 := calculos.CalcAverageFrequency("./calculos/frequency-palmares.json")
	fmt.Println(response7)

	fmt.Println("------------- UNIÃO ---------------------")
	response8 := calculos.CalcAverageFrequency("./calculos/frequency-uniao.json")
	fmt.Println(response8)

	fmt.Println("------------- ESTRELA DALVA ---------------------")
	response9 := calculos.CalcAverageFrequency("./calculos/frequency-estrela-dalva.json")
	fmt.Println(response9)

}
