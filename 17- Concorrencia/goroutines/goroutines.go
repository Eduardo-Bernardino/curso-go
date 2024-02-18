package main

import (
	"fmt"
	"time"
)

func main() {
	go escrever("Olá Mundo!") /* goroutine, ela quer dizer para o programa que não há necessidade de esperar essa chamada e função
	terminar de executar para executar as outras partes do programa. Se colocar go routine em todas as linhas, ele simplesmente
	não executa nada e termina o programa, então tem q tomar cuidado onde vai ser utilizado */
	escrever("Programando em Go!")
}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
