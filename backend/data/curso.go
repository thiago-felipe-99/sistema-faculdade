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

type CursoObrigatório struct {
	Nome              string
	DataDeInício      time.Time
	DataDeDesativação time.Time
	Matérias          []CursoMatéria
}

// CursoData representa as operaçãoes para modificar um curso definitivamente
type CursoData interface {
	InserirMatérias(*[]CursoMatéria) *errors.Aplicação
	Inserir(*Curso) *errors.Aplicação
	AtualizarMatérias(*[]CursoMatéria) *errors.Aplicação
	Atualizar(ID, *Curso) *errors.Aplicação
	PegarMatérias(ID) (*[]CursoMatéria, *errors.Aplicação)
	Pegar(ID) (*Curso, *errors.Aplicação)
	DeletarMatérias(ID) *errors.Aplicação
	Deletar(ID) *errors.Aplicação
}
