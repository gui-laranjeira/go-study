package main

type ContaPoupanca struct {
	Titular       Usuario
	NumeroAgencia int
	NumeroConta   int
	Operacao      int
	saldo         float64
}

func (c *ContaPoupanca) Sacar(valor float64) error {
	if valor <= 0 {
		return ErroValorInvalido()
	}

	if valor > c.saldo {
		return ErroSaldoInsuficiente()
	}

	c.saldo -= valor
	return nil
}

func (c *ContaPoupanca) Depositar(valor float64) error {
	if valor <= 0 {
		return ErroValorInvalido()
	}

	c.saldo += valor
	return nil
}

func (c *ContaPoupanca) GetSaldo() float64 {
	return c.saldo
}
