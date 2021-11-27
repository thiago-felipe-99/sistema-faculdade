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

// ProfessorData representa as opereçãoes que se possa fazer com a entidade
// Professor
type ProfessorData interface {
	Insert(*Professor) (*Professor, *errors.Application)
	Update(ID, *Professor) (*Professor, *errors.Application)
	Get(ID) (*Professor, *errors.Application)
	Delete(ID) *errors.Application
}