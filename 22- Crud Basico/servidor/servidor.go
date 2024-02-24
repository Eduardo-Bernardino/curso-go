package servidor

import (
	"crud/banco"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type user struct {
	ID    uint32 `json:id`
	Nome  string `json:nome`
	Email string `json:email`
}

func CriateUser(w http.ResponseWriter, r *http.Request) {
	corpo, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		w.Write([]byte("Falha ao ler o corpo da requisição"))
		return
	}

	var usuario user

	if erro = json.Unmarshal(corpo, &usuario); erro != nil {
		w.Write([]byte("Erro ao converter o usuario em json"))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar com o BD"))
		return
	}
	defer db.Close()

	//Prepare statement

	statement, erro := db.Prepare("insert into usuarios (nome, email) values (?,?)")
	if erro != nil {
		w.Write([]byte("Erro ao executar o statement"))
		return
	}
	defer statement.Close()

	insercao, erro := statement.Exec(usuario.Nome, usuario.Email)
	if erro != nil {
		w.Write([]byte("Erro ao executar o statement"))
		return
	}

	idInserido, erro := insercao.LastInsertId()
	if erro != nil {
		w.Write([]byte("Erro ao obter o id inserido"))
		return
	}

	//Status code
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("User criado com sucesso! ID: %d", idInserido)))
}
