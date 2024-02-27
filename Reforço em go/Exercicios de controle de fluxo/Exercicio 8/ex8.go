package main

import (
	"fmt"
)

//Crie um programa que utilize a declaração switch, sem switch statement especificado.

func main() {
	num := 233

	switch {

	case num == 0:
		fmt.Println("numero igual a:", num)
	case num == 1:
		fmt.Println("numero igual a:", num)
	case num == 2:
		fmt.Println("numero igual a:", num)
	default:
		fmt.Println("numero escolhido foi:", num)
	}

}
