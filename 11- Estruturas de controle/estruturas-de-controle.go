package main

import "fmt"

func main() {
	fmt.Println("Estruturas de controle")

	numero := 10

	if numero > 15 { // não precisa de parentes
		fmt.Println("numero maior que 15")
	} else {
		fmt.Println("numero menor que 15")
	}

	if numero2 := numero; numero2 > 0 { // Obs: se utilizar o if init, voce só usa dentro do escopo, não tem como utilizar fora
		fmt.Println("Numero maior que 0")
	} else if numero2 == 0 {
		fmt.Println("numero igual que 0")
	} else {
		fmt.Println("numero menor que 0")
	}

}
