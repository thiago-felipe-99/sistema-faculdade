package data

import (
	"time"

	"thiagofelipe.com.br/sistema-faculdade/errors"
)

// TurmaID representa as turmas do aluno
type TurmaID struct {
	ID     ID
	Status string
}

// Aluno representa a entidade Aluno
type Aluno struct {
	ID
	Pessoa
	Matrícula      string
	Curso          ID
	DataDeIngresso time.Time
	DataDeSaída    time.Time
	Período        string
	Status         string
	Turmas         []TurmaID
}

type AlunoToInsert struct {
	Pessoa         PessoaToInsert
	Matrícula      string
	Curso          ID
	DataDeIngresso time.Time
	DataDeSaída    time.Time
	Período        string
	Status         string
	Turmas         []TurmaID
}

// AlunoData representa as opereçãoes que se possa fazer com a entidade Aluno
type AlunoData interface {
	Insert(*AlunoToInsert) (*Aluno, *errors.Application)
	Update(ID, *AlunoToInsert) (*Aluno, *errors.Application)
	Get(ID) (*Aluno, *errors.Application)
	Delete(ID) *errors.Application
}
