package data

import (
	"time"

	"thiagofelipe.com.br/sistema-faculdade/errors"
)

// CursoMatéria representa as matérias que um curso pode ter
type CursoMatéria struct {
	ID_Curso   ID
	ID_Matéria ID
	Período    string
	Tipo       string
	Status     string
	Observação string
}

// Curso representa a entidade Curso
type Curso struct {
	ID                ID
	Nome              string
	DataDeInício      time.Time
	DataDeDesativação time.Time
	Matérias          []CursoMatéria
}

type CursoToInsert struct {
	Nome              string
	DataDeInício      time.Time
	DataDeDesativação time.Time
	Matérias          []CursoMatéria
}

// CursoData representa as operaçãoes para modificar um curso definitivamente
type CursoData interface {
	InserirMatérias(*[]CursoMatéria) *errors.Application
	Inserir(*Curso) *errors.Application
	AtualizarMatérias(*[]CursoMatéria) *errors.Application
	Atualizar(ID, *Curso) *errors.Application
	PegarMatérias(ID) (*[]CursoMatéria, *errors.Application)
	Pegar(ID) (*Curso, *errors.Application)
	DeletarMatérias(ID) *errors.Application
	Deletar(ID) *errors.Application
}
