package main

type ContaCorrente struct {
	Titular       Usuario
	NumeroAgencia int
	NumeroConta   int
	saldo         float64
}

func (c *ContaCorrente) Sacar(valor float64) error {
	if valor <= 0 {
		return ErroValorInvalido()
	}

	if valor > c.saldo {
		return ErroSaldoInsuficiente()
	}

	c.saldo -= valor
	return nil
}

func (c *ContaCorrente) Depositar(valor float64) error {
	if valor <= 0 {
		return ErroValorInvalido()
	}

	c.saldo += valor
	return nil
}

func (c *ContaCorrente) Transferir(valor float64, senha string, contaDestino conta) error {
	if valor <= 0 {
		return ErroValorInvalido()
	}

	if valor > c.saldo {
		return ErroSaldoInsuficiente()
	}

	a := c.Titular.Autenticar(senha)

	if !a {
		return ErroSenhaIncorreta()
	}

	c.saldo -= valor
	contaDestino.Depositar(valor)
	return nil
}

func (c *ContaCorrente) GetSaldo() float64 {
	return c.saldo
}
