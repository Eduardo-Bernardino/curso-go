package main

import "fmt"

func soma(numeros ...int) int { //isso quer dizer que a função vai receber de 1 a n ints
	total := 0

	for _, numero := range numeros {
		total += numero
	}

	return total
}

func escrever(texto string, numeros ...int) { // só pode utilizar 1 variatico por função e ele tem obrigatoriamente que ser o ultimo, se n da erro
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}

func main() {
	fmt.Println("Variaticas")
	resultSoma := soma(10, 20, 30, 40, 50, 60, 70)
	fmt.Println(resultSoma)

	escrever("Salve", 1, 2, 3, 4, 5, 6, 7)
}
