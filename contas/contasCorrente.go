package contas

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) (string, float64) {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso", c.saldo
	} else {
		return "Saldo insuficiente", c.saldo
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
