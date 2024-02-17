package main

import "fmt"

func fibonacci(posicao uint) uint {
	if posicao <= 1 {
		return posicao
	}

	return fibonacci(posicao-2) + fibonacci(posicao-1)
}

func main() {
	fmt.Println("Recursivas")

	posicao := uint(10)             // para não ser declarado como int especifica como uint
	fmt.Println(fibonacci(posicao)) // e assim passa ele por parametro para a funcao

	for i := uint(1); i <= posicao; i++ {
		fmt.Println(fibonacci(i))
	}

}
