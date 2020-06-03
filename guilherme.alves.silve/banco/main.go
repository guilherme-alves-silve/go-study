package main

import (
	"fmt"

	"github.com/guilherme.alves.silve/banco/cliente"
	"github.com/guilherme.alves.silve/banco/conta"
)

type iverficadorConta interface {
	Saca(saque float64) string
}

func pagaBoleto(valorBoleto float64, conta iverficadorConta) {
	fmt.Println(conta.Saca(valorBoleto))
}

func main() {

	conta1 := conta.Corrente{
		Titular: cliente.Titular{
			Nome:      "Guilherme",
			Documento: "123.123.123-12",
			Profissao: "Desenvolvedor",
		},
		Conta:   123,
		Agencia: 1234567,
	}
	conta1.Deposita(800.45)

	conta2 := conta.Corrente{
		Titular: cliente.Titular{
			Nome:      "João",
			Documento: "333.333.333-33",
			Profissao: "Consultor Externo",
		},
		Conta:   555,
		Agencia: 555666777,
	}
	conta2.Deposita(8500.75)

	var conta3 *conta.Corrente
	conta3 = new(conta.Corrente)
	conta3.Titular = cliente.Titular{
		Nome:      "Maria",
		Documento: "777.777.777-77",
		Profissao: "Contadora",
	}
	conta3.Conta = 745
	conta3.Agencia = 651364365
	conta3.Deposita(3500.22)

	conta4 := conta.Poupanca{
		Titular: cliente.Titular{
			Nome:      "João",
			Documento: "333.333.333-33",
			Profissao: "Consultor Externo",
		},
		Conta:    555,
		Operacao: 13,
		Agencia:  555666777,
	}
	conta4.Deposita(5000)

	fmt.Println(conta1)
	fmt.Println(conta2)
	fmt.Println(conta3)
	fmt.Println(*conta3)
	fmt.Println(conta4)

	fmt.Println("******************")

	fmt.Println(conta1, conta2)
	fmt.Println(conta1.Saca(100))
	fmt.Println(conta1.Saca(1000))
	fmt.Println(conta2.Saca(-100))
	fmt.Println(conta2.Saca(1000))

	fmt.Println("******************")

	fmt.Println(conta1, conta2)
	fmt.Println(conta1.Deposita(100))
	fmt.Println(conta1.Deposita(1000))
	fmt.Println(conta2.Deposita(-100))
	fmt.Println(conta2.Deposita(1000))

	fmt.Println("******************")

	fmt.Println(conta1, conta2)
	fmt.Println(conta1.Transfere(100000, &conta2))
	fmt.Println(conta1.Transfere(1000, &conta2))
	fmt.Println(conta1.Transfere(800, &conta1))
	fmt.Println(conta2.Transfere(-100, &conta1))
	fmt.Println(conta2.Transfere(1000, &conta1))

	fmt.Println("******************")

	pagaBoleto(1000, &conta2)
	pagaBoleto(1000, conta3)
	pagaBoleto(6000, &conta4)
	pagaBoleto(50, &conta4)
}
