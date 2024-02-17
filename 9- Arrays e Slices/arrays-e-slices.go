package main

import "fmt"

func main() {

	/*Arrays possuem uma estrutura muito rigida e fixa, sendo sempre necessario colocar
	tamanho e tipo na sua tribuição e desta forma não é muito utilizado na linguagem*/
	fmt.Println("Arrays")

	var array1 [5]int //como não foi atribuido valores no array ele coloca tudo como 0
	fmt.Println(array1)

	//Populando o array
	array1[0] = 25
	fmt.Println(array1)

	//Obs: sempre tem q expecificar o tamanho do array, pois se n expecificar ele vira um slice

	array2 := [5]string{} /*outra forma de declarar um array, declarando tamanho e tipo,
	porem tem q especificar o que vai nele. No caso de array de strings, se não declarar nada
	ele atribui apenas preenche com string vazia*/
	fmt.Println(array2)

	array3 := [5]string{"p1", "p2", "p3", "p4", "p5"} //o array pode ser inicializado na sua declaração
	fmt.Println(array3)

	array4 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array4)
	/*outra forma de declarar um array, é com o "..." como tamanho, desta forma ele fixa o tamanho do array
	de acordo com o tamanho que eu inicializei ele na atribuição*/

	/*Sliced permitem uma flexibilidade no tamanho, porem ainda possui atribuição fixa de tipo*/
	fmt.Println("Slices")

	slice1 := []int{1, 2, 3, 4, 5}
	fmt.Println(slice1)

	/*Slices não são como arrays na sua atribuição, o slice aponta para uma parte na memoria
	onde está os dados em forma de array,*/

	slice1 = append(slice1, 6) /*Como o slice não é fixo, a função append apenas adiciona o novo dado no final do "array",
	então ele copia todos os dados e retorna um novo "array" com o novo item adicionado na devida posição, ou seja,
	ele cria um outro "array" em outro lugar copia tudo o que tinha no antigo, adiciona o novo item e te devolve o novo*/
	fmt.Println(slice1)

	/*slice pode ser meio q traduzido do ingles, como um pedaço, levando meio q o sentido da coisa, da para pegar partes de array
	e colocar para ele*/

	slice2 := array3[1:4] /* o segundo parametro é onde ele vai pagar, no caso no indice 4. O slide tbm é como se fosse um ponteiro
	apontando para essas posicoes do array*/
	fmt.Println(slice2)

	array3[2] = "p11"
	array3[3] = "p99"
	fmt.Println(array3)
	fmt.Println(slice2)

	//Arrays internos
	fmt.Println("Arrays internos")
	slice3 := make([]float32, 10, 11)
	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))

	//cuiriosidade sobre o slice, toda vez q ele estoura o tamanho ele dobra o tamanho, desta forma nunca fica sem "espaço"

	slice3 = append(slice3, 8)
	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))

	fmt.Println("estourando o slice")
	slice3 = append(slice3, 9)
	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3)) //demonstra o slice estourando e dobrando o tamanho, de 12 para 24

}
