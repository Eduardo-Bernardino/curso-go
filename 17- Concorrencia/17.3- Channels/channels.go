package main

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan string)

	go escrever("Ola mundo", canal)

	fmt.Println("Depois da exec do escrever")

	for mensagem := range canal {
		fmt.Println(mensagem)
	}

	fmt.Println("GGWP")
}

func escrever(texto string, canal chan string) {
	for i := 0; i < 3; i++ {
		canal <- texto
		time.Sleep(time.Second)
	}
	close(canal)
}
