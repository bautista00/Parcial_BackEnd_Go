package tickets

import (
	"encoding/csv"
	"os"
	"strconv"
	"strings"
)


type Ticket struct {
	
	ID  int
	Name string
	Email string
	Destination string
	Time string
	Price float64

}

// Requerimiento 1
// Una función que calcule cuántas personas viajan a un país determinado.

func GetTotalTickets(destination string) (int, error) {
	file, err := os.Open("tickets.csv")
	
	if err != nil {
		return 0, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	rows, err := reader.ReadAll()
	if err != nil {
		return 0, err
	}

	total := 0
	for _, row := range rows {
		if row[3] == destination {
			total++
		}
	}

	return total, nil
}

// Requerimiento 2
// Una o varias funciones que calculen cuántas personas 
// viajan en madrugada (0 → 6),
// mañana (7 → 12), tarde (13 → 19), y noche (20 → 23).

func GetCountByPeriod(filePath string, time string) (int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	rows, err := reader.ReadAll()
	if err != nil {
		return 0, err
	}

	tickets := make([]Ticket, len(rows))
	for i, row := range rows {
		id, _ := strconv.Atoi(row[0])
		price, _ := strconv.ParseFloat(row[5], 64)

		tickets[i] = Ticket{
			ID:          id,
			Name:        row[1],
			Email:       row[2],
			Destination: row[3],
			Time:        row[4],
			Price:       price,
		}
	}

	count := 0
	for _, ticket := range tickets {
		hour, err := strconv.Atoi(strings.Split(ticket.Time, ":")[0])
		if err != nil {
			return 0, err
		}

		switch time {
		case "Madrugada":
			if hour >= 0 && hour < 6 {
				count++
			}
		case "Mañana":
			if hour >= 7 && hour < 12 {
				count++
			}
		case "Tarde":
			if hour >= 13 && hour < 19 {
				count++
			}
		case "Noche":
			if hour >= 20 && hour <= 23 {
				count++
			}
		}
	}

	return count, nil
}

// Requerimiento 3
// Calcular el porcentaje de personas que viajan a 
// un país determinado en un día.

func PercentageDestination(filePath string, destination string) (float64, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	rows, err := reader.ReadAll()
	if err != nil {
		return 0, err
	}
	count := 0
	for _, row := range rows {
		if row[3] == destination{
			count++
		}
	}
	percentage := float64(count) / float64(1000) * 100

	return percentage, nil
}
