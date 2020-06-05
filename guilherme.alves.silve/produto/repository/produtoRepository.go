package repository

import (
	"github.com/guilherme.alves.silve/produto/database"
	"github.com/guilherme.alves.silve/produto/models"
)

//GetAllProdutos - Retorna todos os produtos cadastrados no banco de dados
func GetAllProdutos() []models.Produto {
	db := database.Conecta()

	produtosSelect, err := db.Query("SELECT * FROM produtos;")
	if err != nil {
		panic(err.Error())
	}

	produtos := []models.Produto{}
	for produtosSelect.Next() {
		produto := models.Produto{}
		produtosSelect.Scan(
			&produto.ID,
			&produto.Nome,
			&produto.Descricao,
			&produto.Preco,
			&produto.Quantidade,
		)
		produtos = append(produtos, produto)
	}

	defer db.Close()

	return produtos
}

//SaveProduto - Salva o produto no banco de dados
func SaveProduto(produto models.Produto) {

	db := database.Conecta()

	stmt, err := db.Prepare("INSERT INTO produtos(nome, descricao, preco, quantidade) VALUES ($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}

	_, errResult := stmt.Exec(produto.Nome, produto.Descricao, produto.Preco, produto.Quantidade)
	if errResult != nil {
		panic(err.Error())
	}

	defer db.Close()
}

//DeleteProduto - Deleta o produto do banco de dados
func DeleteProduto(id string) {

	db := database.Conecta()

	stmt, err := db.Prepare("DELETE FROM produtos WHERE id = $1")
	if err != nil {
		panic(err.Error())
	}

	_, errDelete := stmt.Exec(id)
	if errDelete != nil {
		panic(errDelete.Error())
	}

	defer db.Close()
}

//FindProduto - Retorna o produto através do id informado
func FindProduto(id string) models.Produto {

	db := database.Conecta()

	produto := models.Produto{}

	rows, err := db.Query("SELECT * FROM produtos WHERE id = $1", id)
	if err != nil {
		panic(err.Error())
	}

	for rows.Next() {
		rows.Scan(
			&produto.ID,
			&produto.Nome,
			&produto.Descricao,
			&produto.Preco,
			&produto.Quantidade,
		)
	}

	defer db.Close()

	return produto
}

//UpdateProduto - Atualiza os dados do produto
func UpdateProduto(produto models.Produto) {

	db := database.Conecta()

	stmt, err := db.Prepare(`UPDATE produtos SET 
							 nome = $1,
							 descricao = $2, 
							 preco = $3, 
							 quantidade = $4
	 						 WHERE id = $5`)
	if err != nil {
		panic(err.Error())
	}

	_, errUpdate := stmt.Exec(
		produto.Nome,
		produto.Descricao,
		produto.Preco,
		produto.Quantidade,
		produto.ID,
	)

	if errUpdate != nil {
		panic(errUpdate.Error())
	}

	defer db.Close()
}
