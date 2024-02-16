package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint8
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {
	fmt.Println("herança")

	pessoa1 := pessoa{"joe", "doe", 30, 123}
	fmt.Println(pessoa1)

	estudante1 := estudante{pessoa1, "enfermagem", "unb"}
	fmt.Println(estudante1)

	fmt.Println(estudante1.altura)

	//só para mostar que pode colocar um struct direto, e chamar por exemplo, estudante.nome ao invez de estudante.pessoa.nome

}
