package contas

import "github.com/luizgnaguiar/banco/clientes"

type ContaCorrente struct {
	Titular                    clientes.Titular
	NumeroAgencia, NumeroConta int
	saldo                      float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso"
	} else {
		return "saldo insuficiente"
	}
}
func (c *ContaCorrente) Depositar(valorDeDeposito float64) (string, float64) {
	if valorDeDeposito > 0 {
		c.saldo += valorDeDeposito
		return "Deposito realizado com sucesso", c.saldo
	} else {
		return "O valor do deposito e menor que 0", c.saldo
	}
}
func (c *ContaCorrente) Transferir(valorDeTransferencia float64, contaDestino *ContaCorrente) bool {
	if valorDeTransferencia < c.saldo && valorDeTransferencia > 0 {
		c.saldo -= valorDeTransferencia
		contaDestino.Depositar(valorDeTransferencia)
		return true
	} else {
		return false
	}
}
func (c *ContaCorrente) ObterSaldo() float64 {
	return c.saldo
}
