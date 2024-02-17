package main

import "fmt"

var n int

func main() {
	fmt.Println("Função main sendo executada")
	fmt.Println(n)
}

// Afunção init sempre é executada primeiro, até mesmo do main, e até mesmo sendo declarada depois.
func init() {
	fmt.Println("Executando a função init")
	n = 10
}

//Obs: muito utilizada para inicializar algo (setar algo) "antes da execução do programa em si"
