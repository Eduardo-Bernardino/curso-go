package main

import (
	"fmt"
)

/*Crie um programa que utilize a declaração switch, onde o switch statement seja uma variável do tipo string com identificador "esporteFavorito".*/

func main() {
	str := "futebol"

	switch {
	case str == "futebol":
		fmt.Println("neymar")
	case str == "volei":
		fmt.Println("serginho escadinha")
	case str == "luta":
		fmt.Println("anderson silva")
	default:
		fmt.Println("sedentario?")
	}

}
