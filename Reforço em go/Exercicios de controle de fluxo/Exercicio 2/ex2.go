package main

import (
	"fmt"
)

//Põe na tela: O unicode code point de todas as letras maiúsculas do alfabeto, três vezes cada.

func main() {
	for i := 65; i <= 90; i++ {
		fmt.Printf("%d \n", i)
		for j := 0; j < 3; j++ {
			fmt.Printf("%#U ", i)
		}
		fmt.Printf("\n")
	}
}
