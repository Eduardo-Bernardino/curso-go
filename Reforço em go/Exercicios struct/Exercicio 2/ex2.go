package main

import "fmt"

/*Utilizando a solução anterior, coloque os valores do tipo "pessoa" num map, utilizando os sobrenomes como key.
Demonstre os valores do map utilizando range.
Os diferentes sabores devem ser demonstrados utilizando outro range, dentro do range anterior.*/

type pessoa struct {
	Nome      string
	Sobrenome string
	Sabores   []string
}

func main() {
	mapa := make(map[string]pessoa)

	mapa["Mello"] = pessoa{
		Nome:      "Eduardo",
		Sobrenome: "Mello",
		Sabores:   []string{"Morango", "Flocos"},
	}

	mapa["Doe"] = pessoa{
		Nome:      "Joe",
		Sobrenome: "Doe",
		Sabores:   []string{"Limao", "Milho", "Creme"},
	}

	mapa["Sampaio"] = pessoa{
		Nome:      "Rafael",
		Sobrenome: "Sampaio",
		Sabores:   []string{"Chocolate", "Pistache", "Queijo"},
	}

	for _, value := range mapa {
		fmt.Println(value)
	}
}
