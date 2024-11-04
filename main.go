package main

import (
	"fmt"

	c "github.com/luizgnaguiar/banco/contas"
)

func main() {
	contaDoLuiz := c.ContaCorrente{Titular: "Luiz", NumeroAgencia: 222, NumeroConta: 11111, Saldo: 200}
	contaDaMari := c.ContaCorrente{Titular: "Mari", NumeroAgencia: 222, NumeroConta: 11112, Saldo: 1000}
	status, valor := contaDoLuiz.Depositar(2500)
	fmt.Println(status, valor)
	contaDaMari.Transferir(500, &contaDoLuiz)
	fmt.Println(contaDoLuiz)
	fmt.Println(contaDaMari)

}
