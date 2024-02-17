package main

import "fmt"

/* Para se utilizar a função recover é necessario armazerar o resultado da função, ou seja, sempre criar uma variavel para receber ela. Se estiver
utilizando a funçãor recover sem uma função panic o recover vai ser nil e por isso não executa, desta forma é importante sempre comprar ela.*/
func recuperarexecucao() {
	if r := recover(); r != nil {
		fmt.Println("execução recuperada com sucesso")
	}
}

func alunoestaaprovado(nota1, nota2 float32) bool {
	defer recuperarexecucao()
	media := (nota1 + nota2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	panic("a media é exatamente 6") /*a função panic para de executar tudo do programa, só vai chamar a função que tem defer
	e foram chamadas anteriormente a execução da função panie e encerra*/
	// se a função for executada depois do panic, mesmo com defer não funciona -> "defer recuperarexecucao()""
}

func main() {
	fmt.Println("Panic e recover")

	fmt.Println(alunoestaaprovado(8, 6))
	fmt.Println("pós execução")
}
