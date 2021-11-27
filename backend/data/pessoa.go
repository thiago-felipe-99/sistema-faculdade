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

// PessoaToInsert contém somente os campos requeridos para modificar uma pessoa.
type PessoaToInsert struct {
	Nome             string
	CPF              CPF
	DataDeNascimento time.Time
	Senha            Senha
}

// PessoaData representa quais são as opereçãoes necessárias para salvar e
// alterar uma pessoa definitivamente.
type PessoaData interface {
	Insert(*PessoaToInsert) (*Pessoa, *errors.Application)
	Update(ID, *PessoaToInsert) (*Pessoa, *errors.Application)
	Get(ID) (*Pessoa, *errors.Application)
	Delete(ID) *errors.Application
}
