package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"

	"github.com/guilherme.alves.silve/produto/models"
	"github.com/guilherme.alves.silve/produto/repository"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

//Index - Retorna a página principal que contém todos os produtos
func Index(w http.ResponseWriter, r *http.Request) {

	produtos := repository.GetAllProdutos()
	templates.ExecuteTemplate(w, "Index", produtos)
}

//New - Retorna a página de cadastro de produto
func New(w http.ResponseWriter, r *http.Request) {

	templates.ExecuteTemplate(w, "New", nil)
}

//Insert - Utilizado para cadastrar um novo produto
func Insert(w http.ResponseWriter, r *http.Request) {
	if "POST" == r.Method {

		produto := getProdutoDoForm(r)

		repository.SaveProduto(produto)

		http.Redirect(w, r, "/", http.StatusMovedPermanently)
	}
}

//Delete - Utilizado para deletar um produto
func Delete(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	repository.DeleteProduto(id)
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

//Edit - Retorna a tela do produto que vai ser editado
func Edit(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	produto := repository.FindProduto(id)
	templates.ExecuteTemplate(w, "Edit", produto)
}

//Update - Atualiza os dados do produto
func Update(w http.ResponseWriter, r *http.Request) {
	if "POST" == r.Method {
		produto := getProdutoDoForm(r)

		id, err := strconv.Atoi(r.FormValue("id"))
		if err != nil {
			fmt.Println("Não foi possível converter o ID")
		}

		produto.ID = id
		repository.UpdateProduto(produto)
		http.Redirect(w, r, "/", http.StatusMovedPermanently)
	}
}

func getProdutoDoForm(r *http.Request) models.Produto {
	nome := r.FormValue("nome")
	descricao := r.FormValue("descricao")
	preco, errPreco := strconv.ParseFloat(r.FormValue("preco"), 64)
	if errPreco != nil {
		fmt.Println("Não foi possível converter o preço")
	}

	quantidade, errQuantidade := strconv.Atoi(r.FormValue("quantidade"))
	if errQuantidade != nil {
		fmt.Println("Não foi possível converter a quantidade")
	}

	return models.Produto{
		Nome:       nome,
		Descricao:  descricao,
		Preco:      preco,
		Quantidade: quantidade,
	}
}
