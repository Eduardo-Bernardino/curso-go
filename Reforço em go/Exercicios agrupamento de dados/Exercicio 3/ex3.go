package main

import "fmt"

/*Utilizando como base o exercício anterior, utilize slicing para demonstrar os valores:
Do primeiro ao terceiro item do slice (incluindo o terceiro item!)
Do quinto ao último item do slice (incluindo o último item!)
Do segundo ao sétimo item do slice (incluindo o sétimo item!)
Do terceiro ao penúltimo item do slice (incluindo o penúltimo item!)
Desafio: obtenha o mesmo resultado acima utilizando a função len() para determinar o penúltimo item*/

func main() {
	slice := make([]int, 10, 10)

	for i := 0; i < 10; i++ {
		slice[i] = (i + 1) * 100
	}

	fmt.Println(slice)
	fmt.Println(slice[:3])
	fmt.Println(slice[4:])
	fmt.Println(slice[1:7])
	fmt.Println(slice[2 : len(slice)-1])

}
