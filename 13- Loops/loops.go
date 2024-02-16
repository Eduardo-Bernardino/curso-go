package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Loops")

	i := 0

	for i < 10 {
		i++
		fmt.Println("incrementando i")
		time.Sleep(time.Second)
	}

	fmt.Println(i)

	for j := 0; j < 10; j++ { // variavel apenas declarada dentro do escopo do for
		time.Sleep(time.Second)
		fmt.Println("incrementando j: ", j)
	}

	nome := [3]string{"eduardo", "joe", "doe"}

	for index, value := range nome {
		fmt.Println((index + 1), value)
	}

	//Sequiser imprimir apenas um valor, não pode simplesmente omitir a variavel, é necessario colocar "_" para chegar no dado que voce quer
	for _, value := range nome {
		fmt.Println(value)
	}

	for index, value := range "TESTE" {
		fmt.Println(index+1, string(value))
	}

	user := map[string]string{
		"nome":      "Joe",
		"sobrenome": "Doe",
	}

	for index, value := range user {
		fmt.Println(index, value)
	}

	//Não tem como usar o range em struct

	//Para executar um loop infinito é só não colocar nem uma condição para o for

	for {
		fmt.Println("Infinito")
	}
}
