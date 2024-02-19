package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitGroup sync.WaitGroup // criação de um grupo de espera para go routines, traduzindo do literal

	waitGroup.Add(2) // com essa função dizemos a quantidade de go routines que vai ser executado ao mesmo tempo (quantas vao estar na fila)

	go func() {
		escrever("salve")
		waitGroup.Done() // quando se utiliza o done ele diminui 1 do contador do "Add", diminuindo a quantidade de routines que devem ser executadas
	}()

	go func() {
		escrever("teste")
		waitGroup.Done()
	}()

	waitGroup.Wait() //Forma de o programa não só finalizar, essa função diz para o programa esperar não ter nem uma go routine executando para poder finalizar.

}

func escrever(texto string) {
	for i := 0; i < 3; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
