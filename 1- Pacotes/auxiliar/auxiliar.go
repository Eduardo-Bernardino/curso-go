package auxiliar

import "fmt"

//Funções que são exportadas precisam de comentario descrevendo, no caso registra mensagem na tela
func Escrever() {
	fmt.Println("Escrevendo do auxiliar")
	escrever2() //nem precisa chamar o aux2, pois ta no mesmo pacote
}
