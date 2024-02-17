package main

import "fmt"

func inverterSinal(numero int) int {
	return numero * -1 // para inverter o sinal é só multiplicar por menos 1
}

func inverterSinalComPonteiro(numero *int) { /*Para mudar o valor onde foi chamado a função deve se receber o endereço da variavel,
	desta forma onde esta chegando o valor desta variavel deve ser um ponteiro, para poder receber e modificar os valores da variavel
	onde foi invocada, e não apenas fazer uma copia*/
	*numero = *numero * -1
}

func main() {
	numero := 20

	//Obs: Passando o parametro por copia
	numeroInvertido := inverterSinal(numero) /* utilizando essa forma de chamada, quer dizer apenas que a função onde chega o valor
	da variavel "numero" está criando uma copia do valor q foi enviado, desta forma a variavel original permanece inalterada*/
	fmt.Println(numeroInvertido)
	fmt.Println(numero)

	//Obs: Passando o parametro por referencia
	novoNumero := 40
	fmt.Println(novoNumero)
	inverterSinalComPonteiro(&novoNumero) /*como a variavel que recebe este valor que esta sendo passado como parametro é um ponteiro
	deve-se enviar um endereço para ela, desta forma utilizando o "&" para isso */
	fmt.Println(novoNumero)

}
