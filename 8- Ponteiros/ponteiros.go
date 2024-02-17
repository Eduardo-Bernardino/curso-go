package main

import "fmt"

func main() {
	fmt.Println("Ponteiros")

	var variavel1 int = 10
	var variavel2 int = variavel1

	fmt.Println(variavel1, variavel2)

	variavel1++

	fmt.Println(variavel1, variavel2)

	//ponteiro é uma referencia de memoria

	fmt.Println("--------------------------")

	var variavel3 int
	var ponteiro *int

	variavel3 = 100

	fmt.Println(variavel3, ponteiro)

	ponteiro = &variavel3

	fmt.Println(variavel3, ponteiro) //imprime o lugar da memoria onde esta a variavel 3, pois foi passado &varivael3, ou seja, o endereço da variavel 3

	variavel3++

	fmt.Println(variavel3, *ponteiro)
	/*quando coloca o *ponteiro, quer dizer q com o "*"
	eu quero saber o que esta dentro do endereço q ele recebeu como parametro
	desta forma imprimindo o conteudo da variavel3*/

	var exponteiro int = *&variavel3
	/*desta forma o que eu estou passando é o conteudo presente dentro do endereço de variavel 3,ou seja,
	ele imprime o conteudo presente na variavel 3, porem de forma fixa(só uma vez, que é na atribuição da variavel)
	assim se o valor de varivel3 for alterado, ele n muda junto, afinal ele não é um ponteiro, apenas uma variavel
	que recebeu o conteudo presente no endereço da variavel 3*/
	fmt.Println(variavel3, *ponteiro, exponteiro)

	variavel3++
	fmt.Println(variavel3, *ponteiro, exponteiro)
}
