package main

import (
	"fmt"
)

/*
Escreva expressões utilizando os seguintes operadores, e atribua seus valores a variáveis.
==
!=
<=
<
=

Demonstre estes valores.
*/
func main() {
	a := (10 == 100)
	fmt.Println(a)
	a2 := (10 == 10)
	fmt.Println(a2)
	b := (10 != 100)
	fmt.Println(b)
	b2 := (10 != 10)
	fmt.Println(b2)
	c := (10 <= 100)
	fmt.Println(c)
	c2 := (10 <= 10)
	fmt.Println(c2)
	d := (10 < 100)
	fmt.Println(d)
	d2 := (10 < 9)
	fmt.Println(d2)
	e := (10 >= 100)
	fmt.Println(e)
	e2 := (10 >= 10)
	fmt.Println(e2)
	f := (10 > 100)
	fmt.Println(f)
	f2 := (10 > 9)
	fmt.Println(f2)
}
