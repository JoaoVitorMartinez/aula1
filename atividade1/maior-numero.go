package atividade1

import "fmt"

func MaiorNumero() {
	var (
		notas [5]float32
		maior float32
	)
	for i := 0; i < 5; i++ {
		fmt.Printf("Informe a %d nota", i)
		fmt.Scan(&notas[i])
	}
	for i := 0; i < 5; i++ {
		if (i == 0) || (maior < notas[i]) {
			maior = notas[i]
		}
	}
	fmt.Println("Dado as notas: ", notas)
	fmt.Println("A maior nota é: ", maior)

}
