package main

import "fmt"

type atividade interface {
	fazerAtividade()
}

type pessoa struct {
	Nome      string
	Idade     uint8
	Atividade atividade
}

type profissao struct {
	Profissao string
	Formacao  string
}

func (prof profissao) fazerAtividade() {
	fmt.Printf("trampando como: %s\n", prof.Profissao)
}

type passeio struct {
	Cidade string
	Como   string
}

func (pas passeio) fazerAtividade() {
	fmt.Printf("passeando na cidade: %s\n", pas.Cidade)
}

func main() {
	fmt.Println("reforço")

	dev := profissao{"dev", "ufg"}
	enf := profissao{"enfermeiro", "ufcat"}
	cat := passeio{"Catalao", "barco"}
	udi := passeio{"Uberlandia", "aviao"}

	eduardo := pessoa{"Eduardo", 25, dev}
	joe := pessoa{"Joe", 21, enf}
	doe := pessoa{"Doe", 50, cat}
	rafael := pessoa{"Rafael", 21, udi}

	eduardo.Atividade.fazerAtividade()
	fmt.Println(eduardo)
	joe.Atividade.fazerAtividade()
	doe.Atividade.fazerAtividade()
	rafael.Atividade.fazerAtividade()

	eduardo = pessoa{"Eduardo", 25, cat}

	eduardo.Atividade.fazerAtividade()

}
