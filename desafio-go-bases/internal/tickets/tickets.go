package tickets

import (
	"errors"
	"fmt"
	"log"
)

// estructura
type Ticket struct {
	Id          string
	Nombre      string
	Email       string
	PaisDestino string
	HoraVuelo   string
	Precio      string
}
type Reservas struct {
	Tickets []Ticket
}

// metodo
func (r *Reservas) PrintInfo() {
	fmt.Printf("%v+", r.Tickets)
}

// funcion 1
func (r *Reservas) GetTotalTickets(destination string) (int, error) {
	var count int = 0
	for _, ticket := range r.Tickets {
		if ticket.PaisDestino == destination {
			count++
		}
	}
	if count == 0 {
		return count, errors.New("No se encontraron tickets para ese país")
	}
	return count, nil
}

// funcion 2

func (r *Reservas) GetCountByPeriod(time string) (string, error) {
	var count int = 0
	for _, ticket := range r.Tickets {

		switch time {
		case "madrugada":
			if ticket.HoraVuelo >= "00" && ticket.HoraVuelo <= "06" {
				count++
			}
		case "maniana":
			if ticket.HoraVuelo >= "07" && ticket.HoraVuelo <= "12" {
				count++
			}
		case "tarde":
			if ticket.HoraVuelo >= "13" && ticket.HoraVuelo <= "19" {
				count++
			}
		case "noche":
			if ticket.HoraVuelo >= "20" && ticket.HoraVuelo <= "23" {
			}

		}

	}
	message := fmt.Sprintf("Total de pasajeros que viajan por la %s: %d", time, count)
	return message, nil
}

// funcion 3
func (r *Reservas) PercentageDestination(destination string) (float64, error) {
	total, err := GetTotalTickets(destination)
	if err != nil {
		log.Fatal(err)
	}

	return float64(total) / float64(len(r.Tickets)) * 100, nil
}
