package main

func main() {
	canal := make(chan string, 2)
	canal <- "Teste"

	msg := <-canal
	println(msg)
}
