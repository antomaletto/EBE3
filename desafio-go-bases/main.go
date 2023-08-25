package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/antomaletto/EBE3/desafio-go-bases/internal/tickets/internal/tickets"
)

const (
	filename = "./tickets.csv"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Fatal(err)
		}
	}()
	var entrada string

	fmt.Print("Ingrese destino a buscar: ")

	_, err := fmt.Scanln(&entrada)

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	reservas := tickets.Reservas{
		Tickets: readFile(filename),
	}
	// Crear canales para comunicarnos con las go rutinas
	canalTickets := make(chan tickets.Ticket)
	defer close(canalTickets)
	canalErr := make(chan error)
	defer close(canalErr)

	go func(chan tickets.Ticket, chan error) {

		ticket, err := reservas.GetTotalTickets(entrada)
		if err != nil {
			canalErr <- err
			return
		}

		canalTickets <- ticket

	}(canalTickets, canalErr)

	select {
	case pr := <-canalTickets:
		fmt.Println(pr)
	case err := <-canalErr:
		log.Fatal(err)
		os.Exit(1)
	}

}
func readFile(filename string) []tickets.Ticket {
	file, err := os.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	data := strings.Split(string(file), "\n")

	var resultado []tickets.Ticket
	for i := 0; i < len(data); i++ {

		if len(data[i]) > 0 {
			file := strings.Split(string(data[i]), ",")
			ticket := tickets.Ticket{
				Id:          file[0],
				Nombre:      file[1],
				Email:       file[2],
				PaisDestino: file[3],
				HoraVuelo:   file[4],
				Precio:      file[5],
			}
			resultado = append(resultado, ticket)
		}

	}

	return resultado

}
