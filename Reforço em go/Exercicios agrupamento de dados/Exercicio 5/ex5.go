package main

import (
	"fmt"
)

/*Comece com essa slice:
x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
Utilizando slicing e append, crie uma slice y que contenha os valores:
[42, 43, 44, 48, 49, 50, 51]*/

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	y := []int{}

	y = append(y, x[:3]...)
	y = append(y, x[6:]...)

	fmt.Println(y)
}
