package conta

import "github.com/guilherme.alves.silve/banco/cliente"

//Corrente - Estrutura usada para armazenar dados da conta corrente
type Corrente struct {
	Titular cliente.Titular
	Conta   int
	Agencia int
	saldo   float64
}

//Saca - Realizar saques na conta
func (c *Corrente) Saca(saque float64) string {
	podeSacar := saque <= c.saldo && saque > 0
	if podeSacar {
		c.saldo -= saque
		return "Saque realizado com sucesso"
	} else {
		return "Não foi possível realizar o saque"
	}
}

//Deposita - Realizar depósitos na conta
func (c *Corrente) Deposita(deposito float64) (string, bool) {
	if deposito > 0 {
		c.saldo += deposito
		return "Depósito realizado com sucesso", true
	} else {
		return "Não foi possível realizar o depósito", false
	}
}

//Transfere - Realizar transferência dessa conta para a conta do parâmetro contaParaTransferir
func (c *Corrente) Transfere(valorTransferencia float64, contaParaTransferir *Corrente) (string, float64) {
	if valorTransferencia > 0 &&
		c.saldo >= valorTransferencia &&
		c != contaParaTransferir {
		c.saldo -= valorTransferencia
		contaParaTransferir.saldo += valorTransferencia
		return "Transferência realizada com sucesso", c.saldo
	} else {
		return "Não foi possível realizar a transferência", c.saldo
	}
}

//Saldo - Retorna o saldo da conta
func (c *Corrente) Saldo() float64 {
	return c.saldo
}
