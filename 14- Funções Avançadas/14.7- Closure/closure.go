package main

import "fmt"

func closure() func() {
	texto := "Dentro da função closure" // sempre vai se referir a esse texto, no caso onde ela foi declarada...

	funcao := func() {
		fmt.Println(texto)
	}

	return funcao
}

func main() {
	texto := "Dentro da função main"
	fmt.Println(texto)

	funcaoNova := closure()
	funcaoNova()
}

// Obs: função clousure "tem uma memoria de onde ela veio", ou seja mesmo q possua uma variavel com o mesmo nome no escopo onde foi chamado ele printa o de dentro
