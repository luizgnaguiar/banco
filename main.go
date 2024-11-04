package main

import (
	"fmt"

	"github.com/luizgnaguiar/banco/clientes"
	"github.com/luizgnaguiar/banco/contas"
)

type verificarConta interface {
	Sacar(valor float64) string
}

func main() {

	clienteLuiz := clientes.Titular{Nome: "Luiz", CPF: "111.111.111.11", Profissao: "Desenvolvedor"}
	contaDoLuiz := contas.ContaCorrente{Titular: clienteLuiz, NumeroAgencia: 1111, NumeroConta: 11111}

	contaDaMari := contas.ContaCorrente{}
	contaDaMari.Depositar(1000)
	contaDaMari.Transferir(500, &contaDoLuiz)

	PagarBoleto(&contaDaMari, 500)

	fmt.Println(contaDoLuiz)
	fmt.Println(contaDaMari)

}
func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}
