package main

import "fmt"

func main() {
	fmt.Println("Anonimas")

	func() {
		fmt.Println("Hello world")
	}()

	func(texto string) {
		fmt.Println(texto)
	}("parametro")

	retorno := func(txt string) string { // como a função não faz nada apenas retorna, é necessario colocar ela dentro de uma variavel
		return fmt.Sprintf("Foi recebido: %s", txt)
	}("TESTE")

	fmt.Println(retorno)

}
