package main

import "fmt"

func diaDaSemana(numero int) string {
	switch numero {
	case 1:
		return "domingo"
	case 2:
		return "segunda"
	case 3:
		return "terça"
	case 4:
		return "quarta"
	case 5:
		return "quinta"
	case 6:
		return "sexta"
	case 7:
		return "sabado"
	default:
		return "numero errado"
	}
}

func diaDaSemana2(numero int) string {
	switch {
	case numero == 1:
		return "domingo"
	case numero == 2:
		return "segunda"
	case numero == 3:
		return "terça"
	case numero == 4:
		return "quarta"
	case numero == 5:
		return "quinta"
	case numero == 6:
		return "sexta"
	case numero == 7:
		return "sabado"
	default:
		return "numero errado"
	}
}

func diaDaSemana3(numero int) string {
	var retorno string
	switch {
	case numero == 1:
		retorno = "domingo"
	case numero == 2:
		retorno = "segunda"
	case numero == 3:
		retorno = "terça"
	case numero == 4:
		retorno = "quarta"
	case numero == 5:
		retorno = "quinta"
	case numero == 6:
		retorno = "sexta"
	case numero == 7:
		retorno = "sabado"
	default:
		retorno = "numero errado"
	}

	return retorno
}

func main() {
	fmt.Println("Switch")
	var retorno string
	retorno = diaDaSemana(4)
	fmt.Println(retorno)
	fmt.Println("segunda função")
	retorno = diaDaSemana2(3)
	fmt.Println(retorno)
	fmt.Println("terceira função")
	retorno = diaDaSemana3(7)
	fmt.Println(retorno)
}
