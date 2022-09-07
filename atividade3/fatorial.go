package atividade3

import (
	"fmt"
)

func Fatorial() {
	var numero int

	fmt.Println("Informe um numero:")
	fmt.Scan(&numero)

	fmt.Println("O fatorial é", CalculaFatorial(numero))
}

func CalculaFatorial(numero int) int {

	if numero == 1 {
		return 1
	} else {
		return numero * CalculaFatorial(numero-1)

	}

}
