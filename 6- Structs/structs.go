package main

import "fmt"

type user struct {
	nome  string
	idade uint8
}

type endereco struct {
	rua    string
	numero uint8
}

type usercompleto struct {
	nome      string
	idade     uint8
	endereco  endereco
	documento string
}

func main() {
	fmt.Println("arquivos structs")

	var usuario1 user
	fmt.Println(usuario1)

	usuario1.nome = "joe"
	fmt.Println(usuario1)

	usuario1.idade = 21
	fmt.Println(usuario1)

	usuario2 := user{"doe", 31}
	fmt.Println(usuario2)

	usuario3 := user{nome: "joe doe"}
	fmt.Println(usuario3)

	usuario4 := user{idade: 40}
	fmt.Println(usuario4)

	usuario5 := usercompleto{
		nome:      "eduardo",
		idade:     25,
		documento: "cpf: 1892738917239"}

	fmt.Println(usuario5)

	enderecoexemplo := endereco{"sabia", 123}

	usuario6 := usercompleto{
		nome:      "eduardo",
		idade:     25,
		endereco:  enderecoexemplo,
		documento: "cpf: 1892738917239"}

	fmt.Println(usuario6)

	usuario7 := usercompleto{
		nome:      "eduardo",
		idade:     25,
		endereco:  endereco{"bobos", 0},
		documento: "cpf: 1892738917239"}

	fmt.Println(usuario7)
}
