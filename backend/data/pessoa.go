package data

import (
	"time"

	"thiagofelipe.com.br/sistema-faculdade/errors"
)

// CPF representa o documento CPF(Cadatro De Pessoa Física) no Brasil.
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

type pessoaToInsert struct {
	Nome             string
	CPF              CPF
	DataDeNascimento time.Time
	Senha            Senha
}

// PessoaData representa quais são as opereçãoes necessárias para salvar e
// alterar uma pessoa definitivamente.
type PessoaData interface {
	Insert(pessoa pessoaToInsert) errors.ApplicationError
	Update(id ID, pessoa pessoaToInsert) errors.ApplicationError
	Get(id ID) (Pessoa, errors.ApplicationError)
	Delete(id ID) errors.ApplicationError
}
