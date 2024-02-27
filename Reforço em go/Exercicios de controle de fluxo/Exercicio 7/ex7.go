package main

import (
	"fmt"
)

//Utilizando a solução anterior, adicione as opções else if e else.

func main() {
	num := 1

	if num > 1 {
		fmt.Println("Numero maior que 1")
	} else if num < 1 {
		fmt.Println("Numero menor que 1")
	} else {
		fmt.Println("Numero igual a:", num)
	}
}
