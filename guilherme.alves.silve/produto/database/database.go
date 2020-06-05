package database

import (
	"database/sql"

	//Pacote utilizado para realizar a coneção com o Postgres
	_ "github.com/lib/pq"
)

//Conecta - Utilizado para conectar no banco de dados do sistema de produtos
func Conecta() *sql.DB {
	conexao := "user=postgres dbname=alura password=admin host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error())
	}

	return db
}
