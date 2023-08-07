package main

import (
	"fmt"
)

type conta interface {
	Sacar(valor float64) error
	Depositar(valor float64) error
	GetSaldo() float64
}

func main() {
	//#region Conta Corrente
	contaDaCris := new(ContaCorrente)
	contaDaCris.Titular = Usuario{
		Nome:      "Cris",
		CPF:       "123.123.123-12",
		Sobrenome: "Silva",
		Email:     "cris.silva@teteteste.com",
		senha:     "123456",
		Endereco:  "Rua ABC, 123",
		DataNasc:  "01/01/2000",
	}
	contaDaCris.NumeroAgencia = 589
	contaDaCris.NumeroConta = 123456
	contaDaCris.saldo = 125.50
	//#endregion

	//#region Conta Poupança
	contaDoGui := new(ContaPoupanca)
	contaDoGui.Titular = Usuario{
		Nome:      "Gui",
		CPF:       "123.123.123-12",
		Sobrenome: "Silva",
		Email:     "gui.silva@teteteste.com",
		senha:     "123456",
		Endereco:  "Rua ABC, 123",
		DataNasc:  "01/01/2000",
	}
	contaDoGui.NumeroAgencia = 589
	contaDoGui.NumeroConta = 123456
	contaDoGui.saldo = 125.50
	//#endregion

	err := contaDaCris.Sacar(100)
	if err != nil {
		fmt.Printf("Erro: %v\n", err)
	}

	err = contaDaCris.Depositar(300)
	if err != nil {
		fmt.Printf("Erro: %v\n", err)
	}

	err = contaDaCris.Transferir(300, "1234", contaDoGui)
	if err != nil {
		fmt.Printf("Erro: %v\n", err)
	}

	fmt.Printf("Cris: %v\n", contaDaCris.GetSaldo())
	fmt.Printf("Gui: %v\n", contaDoGui.GetSaldo())

	err = contaDaCris.Transferir(300, "123456", contaDoGui)
	if err != nil {
		fmt.Printf("Erro: %v\n", err)
	}

	fmt.Printf("Cris: %v\n", contaDaCris.GetSaldo())
	fmt.Printf("Gui: %v\n", contaDoGui.GetSaldo())
	contaDaCris.Depositar(1000)

	PagarBoleto(contaDaCris, 60)
	PagarBoleto(contaDoGui, 34)

	fmt.Println("\n\nBOLETOS:")
	fmt.Printf("Cris: %v\n", contaDaCris.GetSaldo())
	fmt.Printf("Gui: %v\n", contaDoGui.GetSaldo())
}

func PagarBoleto(conta conta, valorBoleto float64) {
	conta.Sacar(valorBoleto)
}
