package calculos

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"time"
)

type Item struct {
	RN            string `json:"rn"`
	CodSequencial string `json:"cod_sequencial"`
	DatRegistroF  string `json:"dat_registro_f"`
	DatOrder      string `json:"dat_order"`
	CodIgreja     string `json:"cod_igreja"`
	IdtEbd        string `json:"idt_ebd"`
	QtdMembro     string `json:"qtd_membro"`
	QtdVisitante  string `json:"qtd_visitante"`
	TxtObservacao string `json:"txt_observacao"`
	Total         string `json:"total"`
}

type Data struct {
	Total       int    `json:"total"`
	PerPage     int    `json:"perPage"`
	CurrentPage int    `json:"currentPage"`
	LastPage    int    `json:"lastPage"`
	From        int    `json:"from"`
	To          int    `json:"to"`
	Items       []Item `json:"items"`
}

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    Data   `json:"data"`
}

func CalcAverageFrequency(filepath string) string {

	// Ler o arquivo JSON
	file, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatal("Erro ao ler o arquivo:", err)
	}

	// Decodificar o JSON
	var response Response
	err = json.Unmarshal(file, &response)
	if err != nil {
		log.Fatal("Erro ao decodificar o JSON:", err)
	}

	// Calcular a média total
	sum := 0
	count := 0

	// Calcular a média quarta-feira
	sumWednesday := 0
	countWednesday := 0

	// Calcular a média quinta-feira
	sumThursday := 0
	countThursday := 0

	// Calcular a média quinta-feira
	sumSaturday := 0
	countSaturday := 0

	// Calcular a média para domingos com idt_ebd "sim" e "não"
	sumSim := 0
	countSim := 0
	sumNao := 0
	countNao := 0

	for _, item := range response.Data.Items {
		// Verificar se o registro é do ano de 2023
		date, err := time.Parse("02/01/2006", item.DatRegistroF)
		if err != nil {
			log.Fatal("Erro ao converter a data:", err)
		}

		if date.Year() == 2023 {
			// Converter o campo "total" para inteiro
			total, err := strconv.Atoi(strings.TrimSpace(item.Total))
			if err != nil {
				//log.Println("Erro ao converter data: ", total)
			}

			sum += total
			count++

			if date.Weekday() == time.Wednesday {
				sumWednesday += total
				countWednesday++
			}

			if date.Weekday() == time.Thursday {
				sumThursday += total
				countThursday++
			}

			if date.Weekday() == time.Saturday {
				sumSaturday += total
				countSaturday++
			}

			if date.Weekday() == time.Sunday {

				if item.IdtEbd == "Sim" {
					sumSim += total
					countSim++
				} else if item.IdtEbd == "Não" {
					sumNao += total
					countNao++
				}

			}

		}

	}

	// Calcular a média
	average := float64(sum) / float64(count)
	averageWednesday := float64(sumWednesday) / float64(countWednesday)
	averageThursday := float64(sumThursday) / float64(countThursday)
	averageSaturday := float64(sumSaturday) / float64(countSaturday)
	averageSundayNight := float64(sumNao) / float64(countNao)
	averageSundayEBD := float64(sumSim) / float64(countSim)

	fmt.Printf("Média total de frequencia de 2023: %.2f\n", average)
	fmt.Printf("Média dos cultos de quarta-feira em 2023: %.2f\n", averageWednesday)
	fmt.Printf("Média dos cultos de quinta-feira em 2023: %.2f\n", averageThursday)
	fmt.Printf("Média dos cultos de Sabado em 2023: %.2f\n", averageSaturday)
	fmt.Printf("Média dos cultos de Domingo a noite em 2023: %.2f\n", averageSundayNight)
	fmt.Printf("Média dos cultos de Domingo(EBD) em 2023: %.2f\n", averageSundayEBD)

	return "Finish!!"
}
