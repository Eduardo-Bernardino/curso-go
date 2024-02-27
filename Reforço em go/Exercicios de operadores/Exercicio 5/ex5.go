package main

import (
	"fmt"
)

/*Utilizando iota, crie 4 constantes cujos valores sejam os próximos 4 anos.
Demonstre estes valores.*/

const (
	x = iota + 2024
	y
	z
	t
)

func main() {
	fmt.Println(x, y, z, t)
}
