package tickets

import (
	"encoding/csv"
	"os"
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
func GetMornings(time string) (int, error) {

}

// Requerimiento 3
func AverageDestination(destination string, total int) (int, error) {

}
