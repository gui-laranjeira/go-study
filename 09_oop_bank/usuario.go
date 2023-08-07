package main

type Usuario struct {
	Id        int
	Nome      string
	Sobrenome string
	CPF       string
	Email     string
	senha     string
	Endereco  string
	DataNasc  string
}

func (u *Usuario) Autenticar(senha string) bool {
	return u.senha == senha
}
