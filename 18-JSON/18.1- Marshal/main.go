package main

import (
	"bytes"
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
	cachorro := pet{"Thor", "Golden", 1}
	fmt.Println(cachorro)

	cachorroEmJSON, erro := json.Marshal(cachorro)
	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(cachorroEmJSON)

	fmt.Println(bytes.NewBuffer(cachorroEmJSON))

	cachorro2 := map[string]string{
		"nome": "bolt",
		"raca": "york",
	}

	cachorro2EmJSON, erro := json.Marshal(cachorro2)
	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(cachorro2EmJSON)

	fmt.Println(bytes.NewBuffer(cachorro2EmJSON))
}
