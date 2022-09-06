package main

import (
	"aual1/atividade1"
	"aual1/atividade2"
	"aual1/atividade3"
	"aual1/atividade4"
	"fmt"
	"os"
)

func main() {
	var (
		nome  string
		opcao int
	)

	for {
		fmt.Println("Olá ", nome)
		fmt.Println("1 - Atividade 1")
		fmt.Println("2 - Atividade 2")
		fmt.Println("3 - Atividade 3")
		fmt.Println("0 - Para sair")
		fmt.Scan(&opcao)
		switch opcao {
		case 1:
			fmt.Println("Opção 1 foi escolhida")
			atividade1.MaiorNumero()
		case 2:
			fmt.Println("Opção 2 foi escolhida")
			atividade2.CalcularEquacao()
		case 3:
			atividade3.Fatorial(2)

		case 4:
			atividade4.Calculadora()
		case 0:
			fmt.Println("Saindo...")
			os.Exit(0)
		}

		fmt.Println("\n\n")
	}
}
