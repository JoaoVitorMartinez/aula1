package atividade4

import (
	"fmt"
	"strings"
)

func Calculadora() {
	var (
		opcao            string
		numero1, numero2 float64
	)

menu:
	for {
		fmt.Println("Informe o primeiro número")
		fmt.Scan(&numero1)
		fmt.Println("Informe o segundo número")
		fmt.Scan(&numero2)

		fmt.Println("Escolha uma operação: * / + -")
		fmt.Scan(&opcao)

		switch opcao {
		case "*":
			fmt.Println(numero1 * numero2)

		case "/":
			fmt.Println(numero1 / numero2)

		case "-":
			fmt.Println(numero1 - numero2)

		case "+":
			fmt.Println(numero1 + numero2)
		}

		fmt.Println("Deseja sair ? N/S")
		fmt.Scan(&opcao)

		if "S" == strings.ToUpper(opcao) {
			break menu
		}

	}

}
