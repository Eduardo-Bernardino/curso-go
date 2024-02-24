package servidor

import (
	"crud/banco"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type user struct {
	ID    uint32 `json:id`
	Nome  string `json:nome`
	Email string `json:email`
}

// Cria e insere o usuario no banco
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

// Busca todos os usuarios do banco
func FindUsers(w http.ResponseWriter, r *http.Request) {
	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar com o banco"))
		return
	}
	defer db.Close()

	linhas, erro := db.Query("select * from usuarios")
	if erro != nil {
		w.Write([]byte("Erro ao buscar usuarios no banco"))
		return
	}
	defer linhas.Close()

	var users []user

	for linhas.Next() {
		var user user

		if erro := linhas.Scan(&user.ID, &user.Nome, &user.Email); erro != nil {
			w.Write([]byte("Erro ao escanear o usuario no banco"))
			return
		}
		users = append(users, user)
	}

	w.WriteHeader(http.StatusOK)
	if erro := json.NewEncoder(w).Encode(users); erro != nil {
		w.Write([]byte("Erro converter usuarios para json"))
		return
	}
}

// Busca apenas 1 usuario expecifico no banco
func FindUser(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	ID, erro := strconv.ParseUint(parametros["id"], 10, 64)
	if erro != nil {
		w.Write([]byte("Erro ao converter o parametro para int"))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar com o banco"))
		return
	}
	defer db.Close()

	linha, erro := db.Query("select * from usuarios where id = ?", ID)
	if erro != nil {
		w.Write([]byte("Erro ao encontrar o usuario no banco"))
		return
	}
	defer linha.Close()

	var user user
	if linha.Next() {
		if erro := linha.Scan(&user.ID, &user.Nome, &user.Email); erro != nil {
			w.Write([]byte("Erro ao escanear o usuario no banco"))
			return
		}
	}

	if user.ID == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Usuario com o id informado não foi encontrado no banco"))
		return
	}

	w.WriteHeader(http.StatusOK)
	if erro := json.NewEncoder(w).Encode(user); erro != nil {
		w.Write([]byte("Erro converter o usuario para json"))
		return
	}

}
