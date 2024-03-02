package main

import (
	"github.com/bootcamp-go/desafio-go-bases/internal/tickets"
	// "encoding/csv"
	"fmt"
	"time"
)

func main() {

	go func(pais string) {
		total, err := tickets.GetTotalTickets(pais)
		if err != nil {
			fmt.Printf("Error en GetTotalTickets para %s: %v", pais, err)
			return
		}
		fmt.Printf("Total tickets for %s: %d\n",pais,total)
	}("Brazil")
	

	filePath := "tickets.csv"
	times := []string{"Madrugada", "Mañana", "Tarde", "Noche"}
	for _, time := range times {
		go func(t string) {
			count, err := tickets.GetCountByPeriod(filePath, t)
			if err != nil {
				fmt.Printf("Error en GetCountByPeriod para %s: %v\n", t, err)
				return
			}
			fmt.Printf("Personas que viajan en %s: %d\n", t, count)
		}(time)
	}
	go func(paisPorcentaje string) {
        total, err := tickets.PercentageDestination(filePath,paisPorcentaje)
        if err != nil {
             fmt.Printf("Error en PercentageDestination para %s: %v", paisPorcentaje, err)
            return
        }
         fmt.Printf("El porcentaje de personas que viajaron a %s es %.2f\n", paisPorcentaje, total)
    }("China")



	time.Sleep(2 * time.Second)

}

	



