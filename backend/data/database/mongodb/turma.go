package mongodb

import (
	"thiagofelipe.com.br/sistema-faculdade/data"
	"thiagofelipe.com.br/sistema-faculdade/errors"
)

type turmaToInsert = data.TurmaToInsert

// Turma representa a conexão com o banco de dados MariaDB para fazer alterações
// na entidade Turma.
type Turma struct {
	Connection
}

// Insert é uma função que faz inserção de uma Turma no banco de dados MariaDB.
func (turma Turma) Insert(*turmaToInsert) (*data.Turma, *errors.Application) {
	turma.Log.Info.Println("Inserindo Turma")

	return nil, nil
}

// Update é uma função que faz a atualização de Turma no banco de dados MariaDB.
func (turma Turma) Update(id, *turmaToInsert) (*data.Turma, *errors.Application) {
	turma.Log.Info.Println("Atualizando Turma")

	return nil, nil
}

// Get é uma função que retorna uma Turma do banco de dados MariaDB.
func (turma Turma) Get(id) (*data.Turma, *errors.Application) {
	turma.Log.Info.Println("Pegando Turma")

	return nil, nil
}

// Delete é uma função que remove uma Turma do banco de dados MariaDB.
func (turma Turma) Delete(id) *errors.Application {
	turma.Log.Info.Print("Deletando Turma")

	return nil
}
