package mariadb

import (
	"thiagofelipe.com.br/sistema-faculdade/data"
	"thiagofelipe.com.br/sistema-faculdade/errors"
)

type professorToInsert = data.Professor

// Professor representa a conexão com o banco de dados MariaDB para fazer
// alterações na entidade Professor.
type Professor struct {
	Conexão
}

// Insert é uma função que faz inserção de uma Professor no banco de dados MariaDB.
func (professor Professor) Insert(*professorToInsert) (*data.Professor, *errors.Application) {
	professor.Log.Info.Println("Inserindo Professor")

	return nil, nil
}

// Update é uma função que faz a atualização de Professor no banco de dados
// MariaDB.
func (professor Professor) Update(id, *professorToInsert) (*data.Professor, *errors.Application) {
	professor.Log.Info.Println("Atualizando Professor")

	return nil, nil
}

// Get é uma função que retorna uma Professor do banco de dados MariaDB.
func (professor Professor) Get(id) (*data.Professor, *errors.Application) {
	professor.Log.Info.Println("Pegando Professor")

	return nil, nil
}

// Delete é uma função que remove uma Professor do banco de dados MariaDB.
func (professor Professor) Delete(id) *errors.Application {
	professor.Log.Info.Print("Deletando Professor")

	return nil
}
