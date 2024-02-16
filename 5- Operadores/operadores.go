package main

import "fmt"

func main() {
	soma := 1 + 1
	sub := 3 - 2
	mult := 1 * 1
	rest := 10 % 4

	fmt.Println(soma, sub, mult, rest)

	// Nota: aparentemente os numeros q vao ser divididos precisam estar em float para o resultado ser em
	var div1 float32 = 5 / 3
	var div2 float32 = 5.0 / 3.0
	var div3 float32 = float32(5) / float32(3)

	fmt.Println(div1, div2, div3)

	/*Algumas notas:
	não pode fazer nada com variaveis de tipos diferentes, por exemplo int8 * int 16, o go n deixa
	diferentes formas de atribuição, var pode mudar e const não pode */

	//var pode mudar o valor
	var variavel1 string = "str primeira"
	fmt.Println(variavel1)
	variavel1 = "str segunda"
	fmt.Println(variavel1)

	//atribuição apenas com o := aparentemente é igual var

	variavel2 := "str terceira"
	fmt.Println(variavel2)
	variavel2 = "str quarta"
	fmt.Println(variavel2)

	//const não muda de valor
	const variavel3 = "str quinta"
	fmt.Println(variavel3)
	// variavel3 = "teste" da o erro: cannot assign to variavel3 (neither addressable nor a map index expression)

	fmt.Println(variavel1, variavel2, variavel3)

	//operadores
	fmt.Println("operadores")
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 < 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 != 2)

	//operadores logicos
	verdadeiro, falso := true, false
	fmt.Println("operadores logicos")
	fmt.Println(verdadeiro && verdadeiro)
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || verdadeiro)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)
	fmt.Println(!falso)

	// operadores unarios
	fmt.Println("operadores unarios")
	num := 10
	fmt.Println("original: ", num)
	num++ // numero = numero + 1
	fmt.Println("++: ", num)
	num-- // numero = numero - 1
	fmt.Println("--: ", num)
	num += 10 // numero = numero + 10
	fmt.Println("+=: ", num)
	num -= 10 // numero = numero - 10
	fmt.Println("-=: ", num)
	num *= 3 // numero = numero * 3
	fmt.Println("*=: ", num)
	num /= 3 // numero = numero / 3
	fmt.Println("/=: ", num)
}
