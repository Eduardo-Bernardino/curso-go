package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
	/* salvar func(), isso não é um metodo, só estou dizendo que o meu usuario tem um campo salvar, e esse camo é do tipo função
	   desta forma ele não pode ser invocado por exempo usuario.func...*/
}

//isso quer dizer que todos os usuarios, eles tem um metodo chamado salvar
func (u usuario) salvar() { // o "u" é uma variavel utilizado para referenciar outros campos que chamou o metodo em questão
	fmt.Printf("Salvando o usuario: %s\n", u.nome)
}

func (u usuario) maiorIdade() bool {
	fmt.Printf("Usuario: %s, possui %d de idade\n", u.nome, u.idade)
	return u.idade >= 18
}

func (u *usuario) fazerAniversario() { /*para alterar o valor de idade do usuario é necessario passar o valor por referenciação,
	desta forma ela mantem a alteração em todo o codigo*/
	u.idade++
}

func main() {
	fmt.Println("Metodos")

	usuario1 := usuario{
		nome:  "joe",
		idade: 21,
	}

	fmt.Println(usuario1)

	usuario2 := usuario{"doe", 40}

	fmt.Println(usuario2)

	//desta forma após o metodo ter sido criado ele pode ser invocado da seguinte forma:
	usuario1.salvar()
	usuario2.salvar()

	usuario3 := usuario{"eduardo", 17}

	fmt.Println(usuario3.maiorIdade())
	usuario3.fazerAniversario()
	fmt.Println(usuario3.maiorIdade())

}
