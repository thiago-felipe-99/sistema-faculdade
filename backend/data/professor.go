package data

import (
	"time"

	"thiagofelipe.com.br/sistema-faculdade/errors"
)

// Professor representa a entidade Professor
type Professor struct {
	ID
	Pessoa
	Matrícula           string
	DataDeIngresso      time.Time
	DataDeSaída         time.Time
	Status              string
	Grau                string
	Turmas              []TurmaID
	CargaHoráriaSemanal time.Duration
	HorárioDeAula       Horário
}

type professorToInsert struct {
	Pessoa              pessoaToInsert
	Matrícula           string
	DataDeIngresso      time.Time
	DataDeSaída         time.Time
	Status              string
	Grau                string
	Turmas              []TurmaID
	CargaHoráriaSemanal time.Duration
	HorárioDeAula       Horário
}

// ProfessorData representa as opereçãoes que se possa fazer com a entidade Professor
type ProfessorData interface {
	Insert(professor professorToInsert) (Professor, errors.ApplicationError)
	Update(id, ID, professor professorToInsert) (Professor, errors.ApplicationError)
	Get(id ID) (Professor, errors.ApplicationError)
	Delete(id ID) errors.ApplicationError
}
