package data

import (
	"time"

	"thiagofelipe.com.br/sistema-faculdade/errors"
)

// CPF representa o documento CPF(Cadatro De Pessoa Física) do Brasil.
type CPF = string

// Senha representa uma senha na aplicação.
type Senha = string

// Pessoa representa a entidade Pessoa.
type Pessoa struct {
	ID               ID
	Nome             string
	CPF              CPF
	DataDeNascimento time.Time
	Senha            Senha
}

// PessoaData representa quais são as opereçãoes necessárias para salvar e
// alterar uma pessoa definitivamente.
type PessoaData interface {
	Inserir(*Pessoa) *errors.Aplicação
	Atualizar(ID, *Pessoa) *errors.Aplicação
	Pegar(ID) (*Pessoa, *errors.Aplicação)
	Deletar(ID) *errors.Aplicação
}
