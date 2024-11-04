package main

import (
	"fmt"
)

func main() {
	contaDoLuiz := ContaCorrente{"Luiz", 222, 11111, 200}
	contaDaMari := ContaCorrente{"Mari", 222, 11112, 1000}
	status, valor := contaDoLuiz.Depositar(2500)
	fmt.Println(status, valor)
	contaDaMari.Transferir(500, &contaDoLuiz)
	fmt.Println(contaDoLuiz)
	fmt.Println(contaDaMari)

}
