package main

import (
	"fmt"
)

/*Crie uma variável de tipo string utilizando uma raw string literal.
Demonstre-a.*/

func main() {
	x := `testando
	como
		printar em
		r	
		aw,teste`
	fmt.Println(x)
}
