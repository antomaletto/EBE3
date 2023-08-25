package main

import (
	"os"
	"strings"

	"github.com/antomaletto/EBE3/desafio-go-bases/internal/tickets/internal/tickets"
)

const (
	filename = "./tickets.csv"
)

func main() {

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
