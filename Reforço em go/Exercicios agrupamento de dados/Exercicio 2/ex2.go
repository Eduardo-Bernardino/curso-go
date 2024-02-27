package main

import "fmt"

/*Usando uma literal composta:
Crie uma slice de tipo int
Atribua 10 valores a ela
Utilize range para demonstrar todos estes valores.
E utilize format printing para demonstrar seu tipo.*/

func main() {
	slice := make([]int, 10, 10)

	for i := 0; i < 10; i++ {
		slice[i] = (i + 1) * 100
	}

	for _, value := range slice {
		fmt.Println(value)
	}

	fmt.Printf("%T\n", slice)
}
