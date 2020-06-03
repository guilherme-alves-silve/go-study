package conta

import "github.com/guilherme.alves.silve/banco/cliente"

//Poupanca - Estrutura usada para armazenar dados da conta pouçança
type Poupanca struct {
	Titular  cliente.Titular
	Conta    int
	Agencia  int
	Operacao int
	saldo    float64
}

//Saca - Realizar saques na conta
func (c *Poupanca) Saca(saque float64) string {
	podeSacar := saque <= c.saldo && saque > 0
	if podeSacar {
		c.saldo -= saque
		return "Saque realizado com sucesso"
	} else {
		return "Não foi possível realizar o saque"
	}
}

//Deposita - Realizar depósitos na conta
func (c *Poupanca) Deposita(deposito float64) (string, bool) {
	if deposito > 0 {
		c.saldo += deposito
		return "Depósito realizado com sucesso", true
	} else {
		return "Não foi possível realizar o depósito", false
	}
}

//Saldo - Retorna o saldo da conta
func (c *Poupanca) Saldo() float64 {
	return c.saldo
}
