package main

import (
	"github.com/bootcamp-go/desafio-go-bases/internal/tickets"
	// "encoding/csv"
	"fmt"
)

func main() {
	
	go func(){
		
		total, err := tickets.GetTotalTickets("Brazil")
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("Total tickets for Brazil: %d\n", total)
		}()
}
