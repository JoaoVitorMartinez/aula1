package atividade2

import (
	"fmt"
	"math"
)

func CalcularEquacao() {
	fmt.Println("Dado uma equação 2º grau informe os valores a seguir")
	var (
		a  int
		b  int
		c  int
		x1 float64
		x2 float64
	)

	fmt.Println("Informe o valor de a")
	fmt.Scan(&a)
	fmt.Println("Informe o valor de b")
	fmt.Scan(&b)
	fmt.Println("Informe o valor de c")
	fmt.Scan(&c)
	delta := math.Pow(float64(b), 2) - float64(4*a*c)
	if delta < 0 {
		fmt.Println("Não é possível calcular a raíz")
		return
	}
	x1 = (float64(-b) + math.Sqrt(delta)) / float64(2.0*a)
	x2 = (float64(-b) - math.Sqrt(delta)) / float64(2.0*a)
	fmt.Printf("As Raízes são x'=%f, x''=%f\n", x1, x2)
}
