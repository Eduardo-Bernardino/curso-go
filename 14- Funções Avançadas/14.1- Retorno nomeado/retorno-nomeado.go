package main

import (
	"fmt"
)

func calculosMatematicos(n1, n2 int) (soma, subtração int) {
	// como ja esta no retorno nomeado não precisa declarar ":="
	soma = n1 + n2
	subtração = n1 + n2
	return
}

func main() {
	fmt.Println("funçoes nomeadas")
	soma, subtração := calculosMatematicos(10, 20)
	fmt.Println(soma, subtração)
}
