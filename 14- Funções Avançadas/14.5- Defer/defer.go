package main

import "fmt"

func funcao1() {
	fmt.Println("executando função 1")
}

func funcao2() {
	fmt.Println("executando função 2")
}

func funcao3() {
	fmt.Println("executando função 3")
}

func alunoestaaprovado(nota1, nota2 float32) bool {
	defer fmt.Println("media calculada, retornando resultado") /* mesmo que esteja na primeira linha da função, ele vai executar
	a função toda antes de printar na tela, antes apenas do return*/
	fmt.Println("executando função para ver se o aluno foi aprovado")

	media := (nota1 + nota2) / 2

	if media > 6 {
		return true
	} else {
		return false
	}

}

func main() {
	fmt.Println("Defer")

	defer funcao1() // defer é adiar, ou seja, adiar até ser o ultimo possivel, vai executar tudo antes
	funcao2()
	funcao3()
	funcao2()
	funcao2()
	funcao3()

	fmt.Println(alunoestaaprovado(7, 8))
	fmt.Println(alunoestaaprovado(5, 4))
}
