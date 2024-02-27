package main

import (
	"fmt"
	"time"
)

/*Crie um loop utilizando a sintaxe: for {}
Utilize-o para demonstrar os anos desde que você nasceu.*/

func main() {
	i := 1999
	for {
		fmt.Println(i)
		i++
		if i > time.Now().Year() {
			break
		}
	}
}
