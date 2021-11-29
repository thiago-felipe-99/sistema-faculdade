package data

import (
	"time"

	"thiagofelipe.com.br/sistema-faculdade/errors"
)

type cursosOfertado struct {
	ID
	Vagas   int
	Período string
}

type nota struct {
	Aluno  ID
	Nota   float32
	Status string
}

// Turma representa a entidade Turma
type Turma struct {
	ID
	Matéria
	Professores        []ID
	Alunos             []ID
	CursosResponsáveis []ID
	CursosOfertados    []cursosOfertado
	HorárioDasAulas    []Horário
	Notas              []nota
	DataDeInício       time.Time
	DataDeTérmino      time.Time
}

// TurmaData representa as opereções que se possa fazer com a entidade
// Turma
type TurmaData interface {
	Inserir(*Turma) *errors.Aplicação
	Atualizar(ID, *Turma) *errors.Aplicação
	Pegar(ID) (*Turma, *errors.Aplicação)
	Deletar(ID) *errors.Aplicação
}
