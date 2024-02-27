package main

import "fmt"

/*
Crie um tipo "pessoa" com tipo subjacente struct, que possa conter os seguintes campos:
Nome
Sobrenome
Sabores favoritos de sorvete
Crie dois valores do tipo "pessoa" e demonstre estes valores, utilizando range na slice que contem os sabores de sorvete.
*/

type pessoa struct {
	Nome      string
	Sobrenome string
	Sabores   []string
}

func main() {
	pessoa1 := pessoa{
		Nome:      "Eduardo",
		Sobrenome: "Mello",
		Sabores:   []string{"Morango", "Flocos"},
	}

	pessoa2 := pessoa{
		Nome:      "Joe",
		Sobrenome: "Doe",
		Sabores:   []string{"Limao", "Milho", "Creme"},
	}

	fmt.Println(pessoa1, pessoa2)
}
