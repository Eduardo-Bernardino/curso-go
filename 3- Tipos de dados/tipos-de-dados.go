package main

import (
	"errors"
	"fmt"
)

func main() {
	/*tipos de inteiros,o que muda é o tamanho (quantidade de digitos)
	int8, int16, int32, int64
	*/

	var num int16 = 100
	fmt.Println(num)

	// se colocar apenas "int" ele vai usar o padrão do pc se for 32 bits vai criar int 32, se for 64 vai criar int64...

	num2 := 100000000000

	fmt.Println(num2)

	//tbm existe o "uint" ou seja inteiro sem sinal (não aceita negativo)

	var num3 uint32 = 1000000
	fmt.Println(num3)

	//alguns tipos tbm tem apelidos chamados "alias"
	//int32 = rune
	var num4 rune = 1928381
	fmt.Println(num4)

	//int8 = byte (byte tem 8 bits...)
	var num5 byte = 123
	fmt.Println(num5)

	var numr1 float32 = 123.333332 // sempre ponto "."
	fmt.Println(numr1)

	var numr2 float64 = 1230102100.333332
	fmt.Println(numr2)

	// mesma coisa para o int, vai da arquitetura padrão

	numr3 := 12873912873.12212
	fmt.Println(numr3)

	var str string = "texto do tamanho que voce quiser" // sempre aspas suplas, aspas simples é char
	fmt.Println(str)

	str2 := "texto 2"
	fmt.Println(str2)

	char := 'b'
	fmt.Println(char) // retorna um numero, a posição deste caracter que voce inseriu na "tabela asc", desta forma sempre vai printar o int...

	var booleano1 bool // se n explicitar é sempre false
	fmt.Println(booleano1)

	var booleano2 bool = true
	fmt.Println(booleano2)

	// Os erros tambem são um tipo em go

	var erro1 error
	fmt.Println(erro1) // printa o valor 0, ou seja o padrão q é <nil>

	// Para criar um erro voce precisa de um novo pacote do go chamado errors

	var erro2 error = errors.New("erro interno")
	fmt.Println(erro2)

}
