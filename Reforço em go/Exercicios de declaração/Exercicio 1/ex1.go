package main

import (
	"fmt"
)

/*
Utilizando o operador curto de declaração, atribua estes valores às variáveis com os identificadores "x", "y", e "z".
42
"James Bond"
true
Agora demonstre os valores nestas variáveis, com:
Uma única declaração print
Múltiplas declarações print
*/
func main() {
	x := 42
	y := "James Bond"
	z := true
	fmt.Printf("Valor de x: %d\nValor de y: %s\nValor de z: %v\n", x, y, z)
	fmt.Println("Valor de x:", x)
	fmt.Println("Valor de y:", y)
	fmt.Println("Valor de z:", z)
}
