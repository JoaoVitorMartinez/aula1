package atividade5

import "fmt"

func MultiplosValores() {
	string1, string2 := ExibeMultiplosValores()

	fmt.Println(string1)
	fmt.Println(string2)

}

func ExibeMultiplosValores() (string, string) {
	return "multiplos", "valores"
}
