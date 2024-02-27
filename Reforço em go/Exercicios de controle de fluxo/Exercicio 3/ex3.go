package main

import (
	"fmt"
	"time"
)

/*Crie um loop utilizando a sintaxe: for condition {}
Utilize-o para demonstrar os anos desde que você nasceu.*/

func main() {
	for i := 1999; i <= time.Now().Year(); i++ {
		fmt.Println(i)
	}
}
