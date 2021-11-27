package data

import (
	"time"

	"thiagofelipe.com.br/sistema-faculdade/errors"
)

// Administrativo representa a entidade Administrativo
type Administrativo struct {
	ID
	Pessoa
	Matrícula           string
	DataDeIngresso      time.Time
	DataDeSaída         time.Time
	Status              string
	Grau                string
	CargaHoráriaSemanal time.Duration
	HorárioDeAula       Horário
}

type AdministrativoToInsert struct {
	Pessoa              PessoaToInsert
	Matrícula           string
	DataDeIngresso      time.Time
	DataDeSaída         time.Time
	Status              string
	Grau                string
	Turmas              []TurmaID
	CargaHoráriaSemanal time.Duration
	HorárioDeAula       Horário
}

// AdministrativoData representa as opereçãoes que se possa fazer com a entidade
// Administrativo
type AdministrativoData interface {
	Insert(*AdministrativoToInsert) (*Administrativo, *errors.Application)
	Update(ID, *AdministrativoToInsert) (*Administrativo, *errors.Application)
	Get(ID) (*Administrativo, *errors.Application)
	Delete(ID) *errors.Application
}
