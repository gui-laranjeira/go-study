package main

type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}

func ErroSaldoInsuficiente() error {
	return &errorString{"Saldo insuficiente"}
}

func ErroValorInvalido() error {
	return &errorString{"Valor inválido, toda movimentação deve ser maior que zero"}
}

func ErroSenhaIncorreta() error {
	return &errorString{"Senha incorreta"}
}
