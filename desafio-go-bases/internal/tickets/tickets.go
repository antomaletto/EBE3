package tickets

import (
	"errors"
	"fmt"
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
func (r *Reservas) GetTotalTickets(destination string, tickets []Ticket) (int, error) {
	var count int = 0
	for _, ticket := range tickets {
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
func GetMornings(time string, tickets []Ticket) (int, error) {

}

// funcion 3
func AverageDestination(destination string, total int) (int, error) {}
