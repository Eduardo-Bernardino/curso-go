package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type pet struct {
	Nome  string `json: "nome"`
	Raca  string `json: "raca"`
	Idade uint   `json: "idade"`
}

func main() {
	cachorroEmJSON := `{"nome":"Thor","raca":"Golden","idade":1}`

	var cachorro pet

	if erro := json.Unmarshal([]byte(cachorroEmJSON), &cachorro); erro != nil {
		log.Fatal(erro)
	}
	fmt.Println(cachorro)

	cachorro2EmJSON := `{"nome":"bolt","raca":"york"}`

	cachorro2 := make(map[string]string)

	if erro := json.Unmarshal([]byte(cachorro2EmJSON), &cachorro2); erro != nil {
		log.Fatal(erro)
	}
	fmt.Println(cachorro2)
}
