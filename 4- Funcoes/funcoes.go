package main

import "fmt"

func somar(num1 int8, num2 int8) int8 { //o retorno da função vem depois dos parenteses
	return num1 + num2
}

func testesub(num3, num4 int8) int8 { // se os parametros forem do mesmo tipo podem ser declarados desta maneira, colocando o tipo só dpois
	return num3 - num4
}

func calculosmat(num5, num6 int8) (int8, int8, float32, int16) { // voce pode dizer que a função retorna 2 parametros só colocar entre parenteses e expecificar quais são em ordem
	soma := num5 + num6
	subtr := num5 - num6
	div := float32(num5) / float32(num6)
	mult := int16(num5) * int16(num6)
	return soma, subtr, div, mult
}

func main() {
	soma := somar(20, 30) //declara a variavel soma que recebe o valor retornado da função somar
	fmt.Println(soma)

	var f = func(txt string) string {
		fmt.Println("funcao f")
		return txt
	}

	f("texto da func")
	result := f("texto do texto do texto")
	fmt.Println(result)

	resultsub := testesub(12, 5)
	fmt.Println(resultsub)

	resultcalcsoma1, resultcalsubtr1, resultadodiv1, resultmult1 := calculosmat(18, 10)
	fmt.Println(resultcalcsoma1, resultcalsubtr1, resultadodiv1, resultmult1)

	// caso vc n queria o valor retornado é só colocar um "_"
	resultcalcsoma2, _, _, _ := calculosmat(30, 15)
	fmt.Println(resultcalcsoma2)

	_, resultcalsubtr2, _, _ := calculosmat(30, 15)
	fmt.Println(resultcalsubtr2)

	_, _, resultadodiv2, _ := calculosmat(90, 3)
	fmt.Println(resultadodiv2)

	_, _, _, resultmult2 := calculosmat(99, 98)
	fmt.Println(resultmult2)

}
