package mongodb

import (
	"thiagofelipe.com.br/sistema-faculdade/data"
	"thiagofelipe.com.br/sistema-faculdade/errors"
)

type matériaToInsert = data.MatériaToInsert

// Matéria representa a conexão com o banco de dados MariaDB para fazer alterações
// na entidade Matéria.
type Matéria struct {
	Connection
}

// Insert é uma função que faz inserção de uma Matéria no banco de dados MariaDB.
func (matéria Matéria) Insert(*matériaToInsert) (*data.Matéria, *errors.Application) {
	matéria.Log.Info.Println("Inserindo Matéria")

	return nil, nil
}

// Update é uma função que faz a atualização de Matéria no banco de dados MariaDB.
func (matéria Matéria) Update(id, *matériaToInsert) (*data.Matéria, *errors.Application) {
	matéria.Log.Info.Println("Atualizando Matéria")

	return nil, nil
}

// Get é uma função que retorna uma Matéria do banco de dados MariaDB.
func (matéria Matéria) Get(id) (*data.Matéria, *errors.Application) {
	matéria.Log.Info.Println("Pegando Matéria")

	return nil, nil
}

// Delete é uma função que remove uma Matéria do banco de dados MariaDB.
func (matéria Matéria) Delete(id) *errors.Application {
	matéria.Log.Info.Print("Deletando Matéria")

	return nil
}
