package main

import (
	"fmt"
	"net/http"

	"github.com/guilherme.alves.silve/produto/route"
)

func main() {

	fmt.Println("Iniciando servidor...")
	route.LoadRoutes()
	http.ListenAndServe(":8000", nil)
}
