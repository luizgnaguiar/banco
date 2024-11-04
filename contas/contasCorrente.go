package contas

type ContaCorrente struct {
	Titular       string
	NumeroAgencia int
	NumeroConta   int
	Saldo         float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) (string, float64) {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.Saldo
	if podeSacar {
		c.Saldo -= valorDoSaque
		return "Saque realizado com sucesso", c.Saldo
	} else {
		return "Saldo insuficiente", c.Saldo
	}
}
func (c *ContaCorrente) Depositar(valorDeDeposito float64) (string, float64) {
	if valorDeDeposito > 0 {
		c.Saldo += valorDeDeposito
		return "Deposito realizado com sucesso", c.Saldo
	} else {
		return "O valor do deposito e menor que 0", c.Saldo
	}
}
func (c *ContaCorrente) Transferir(valorDeTransferencia float64, contaDestino *ContaCorrente) bool {
	if valorDeTransferencia < c.Saldo && valorDeTransferencia > 0 {
		c.Saldo -= valorDeTransferencia
		contaDestino.Depositar(valorDeTransferencia)
		return true
	} else {
		return false
	}
}
