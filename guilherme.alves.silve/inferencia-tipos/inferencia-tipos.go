package main

import (
	"fmt"
	"reflect"
)

func main() {

	var idade = 15
	var nome = "Guilherme"
	var versao = 1.1

	fmt.Println("Seu nome é sr.", nome, "e sua idade é", idade, "anos")
	fmt.Println("Esse programa se encontra na versão", versao)

	idade1 := 33
	nome1 := "João Pedro"
	versao1 := 2.5

	fmt.Println("\nSeu nome é sr.", nome1, "e sua idade é", idade1, "anos,", "sendo a variável nome do tipo", reflect.TypeOf(nome1), "e a variável idade do tipo", reflect.TypeOf(idade1))
	fmt.Println("Esse programa se encontra na versão", versao1, "com a variável do tipo", reflect.TypeOf(versao1))
}
