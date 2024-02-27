package main

import (
	"fmt"
)

/*Crie um programa que:
Atribua um valor int a uma variável
Demonstre este valor em decimal, binário e hexadecimal
Desloque os bits dessa variável 1 para a esquerda, e atribua o resultado a outra variável
Demonstre esta outra variável em decimal, binário e hexadecimal*/

var x int = 20

func main() {
	fmt.Printf("%d\n%b\n%x\n", x, x, x)
}
