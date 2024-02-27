package main

import (
	"fmt"
)

/*
Usando uma literal composta:
Crie um array que suporte 5 valores to tipo int
Atribua valores aos seus índices
Utilize range e demonstre os valores do array.
Utilizando format printing, demonstre o tipo do array.*/

func main() {
	array := [5]int{10, 20, 30, 40, 50}
	for _, valor := range array {
		fmt.Println(valor)
	}
}
