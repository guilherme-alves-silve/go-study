package route

import (
	"net/http"

	"github.com/guilherme.alves.silve/produto/controller"
)

//LoadRoutes - Utilizada para configurar as rotas do servidor HTTP
func LoadRoutes() {
	http.HandleFunc("/", controller.Index)
	http.HandleFunc("/new", controller.New)
	http.HandleFunc("/insert", controller.Insert)
	http.HandleFunc("/delete", controller.Delete)
	http.HandleFunc("/edit", controller.Edit)
	http.HandleFunc("/update", controller.Update)
}
