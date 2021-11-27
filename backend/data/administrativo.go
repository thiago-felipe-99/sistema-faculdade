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

type administrativoToInsert struct {
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

// AdministrativoData representa as opereçãoes que se possa fazer com a entidade
// Administrativo
type AdministrativoData interface {
	Insert(administrativoToInsert) (Administrativo, errors.ApplicationError)
	Update(ID, administrativoToInsert) (Administrativo, errors.ApplicationError)
	Get(ID) (Administrativo, errors.ApplicationError)
	Delete(ID) errors.ApplicationError
}
