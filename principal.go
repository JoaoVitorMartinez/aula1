package main

import (
	"aual1/atividade1"
	"aual1/atividade2"
	"aual1/atividade3"
	"aual1/atividade4"
	"aual1/atividade5"
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
		fmt.Println("4 - Atividade 4")
		fmt.Println("5 - Atividade 5")
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
			fmt.Println("Opção 3 foi escolhida")
			atividade3.Fatorial()

		case 4:
			atividade4.Calculadora()

		case 5:
			atividade5.MultiplosValores()
		case 0:
			fmt.Println("Saindo...")
			os.Exit(0)
		default:
			fmt.Println("Opção não reconhecida")
			os.Exit(-1)

		}

		fmt.Println("")
	}
}
