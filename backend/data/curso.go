package data

import (
	"time"

	"thiagofelipe.com.br/sistema-faculdade/errors"
)

// Curso representa a entidade Curso
type Curso struct {
	ID                ID
	Nome              string
	DataDeInício      time.Time
	DataDeDesativação time.Time
	Matérias          []ID
}

type CursoToInsert struct {
	Nome              string
	DataDeInício      time.Time
	DataDeDesativação time.Time
	Matérias          []ID
}

// CursoData representa as operaçãoes para modificar um curso definitivamente
type CursoData interface {
	Insert(*CursoToInsert) (*Curso, *errors.Application)
	Update(ID, *CursoToInsert) (*Curso, *errors.Application)
	Get(ID) (*Curso, *errors.Application)
	Delete(ID) *errors.Application
}
