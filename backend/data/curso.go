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

type cursoToInsert struct {
	Nome              string
	DataDeInício      time.Time
	DataDeDesativação time.Time
	Matérias          []ID
}

// CursoData representa as operaçãoes para modificar um curso definitivamente
type CursoData interface {
	Insert(curso cursoToInsert) (Curso, errors.ApplicationError)
	Update(id ID, curso cursoToInsert) (Curso, errors.ApplicationError)
	Get(id ID) (Curso, errors.ApplicationError)
	Delete(id ID) errors.ApplicationError
}
