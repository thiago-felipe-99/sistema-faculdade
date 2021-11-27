package mariadb

import (
	"thiagofelipe.com.br/sistema-faculdade/data"
	"thiagofelipe.com.br/sistema-faculdade/errors"
)

type cursoToInsert = data.CursoToInsert

// Curso representa a conexão com o banco de dados MariaDB para fazer alterações
// na entidade Curso.
type Curso struct {
	Connection
}

// Insert é uma função que faz inserção de uma Curso no banco de dados MariaDB.
func (curso Curso) Insert(*cursoToInsert) (*data.Curso, *errors.Application) {
	curso.Log.Info.Println("Inserindo Curso")

	return nil, nil
}

// Update é uma função que faz a atualização de Curso no banco de dados MariaDB.
func (curso Curso) Update(id, *cursoToInsert) (*data.Curso, *errors.Application) {
	curso.Log.Info.Println("Atualizando Curso")

	return nil, nil
}

// Get é uma função que retorna uma Curso do banco de dados MariaDB.
func (curso Curso) Get(id) (*data.Curso, *errors.Application) {
	curso.Log.Info.Println("Pegando Curso")

	return nil, nil
}

// Delete é uma função que remove uma Curso do banco de dados MariaDB.
func (curso Curso) Delete(id) *errors.Application {
	curso.Log.Info.Print("Deletando Curso")

	return nil
}
