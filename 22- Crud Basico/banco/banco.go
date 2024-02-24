package banco

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" //Driver para mysql
)

// Abre a conexão com o BD
func Conectar() (*sql.DB, error) {
	stringConexao := "golang:123321@/golang?charset=utf8&parseTime=True&loc=Local"

	db, erro := sql.Open("mysql", stringConexao)

	if erro != nil {
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		return nil, erro
	}

	return db, nil
}
