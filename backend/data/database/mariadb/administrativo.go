package mariadb

import (
	"thiagofelipe.com.br/sistema-faculdade/data"
	"thiagofelipe.com.br/sistema-faculdade/errors"
)

type administrativoToInsert = data.Administrativo

// Administrativo representa a conexão com o banco de dados MariaDB para fazer
// alterações na entidade Administrativo.
type Administrativo struct {
	Conexão
}

// Insert é uma função que faz inserção de uma Administrativo no banco de dados
// MariaDB.
func (administrativo Administrativo) Insert(*administrativoToInsert) (*data.Administrativo, *errors.Application) {
	administrativo.Log.Info.Println("Inserindo Administrativo")

	return nil, nil
}

// Update é uma função que faz a atualização de Administrativo no banco de dados
// MariaDB.
func (administrativo Administrativo) Update(id, *administrativoToInsert) (*data.Administrativo, *errors.Application) {
	administrativo.Log.Info.Println("Atualizando Administrativo")

	return nil, nil
}

// Get é uma função que retorna uma Administrativo do banco de dados MariaDB.
func (administrativo Administrativo) Get(id) (*data.Administrativo, *errors.Application) {
	administrativo.Log.Info.Println("Pegando Administrativo")

	return nil, nil
}

// Delete é uma função que remove uma Administrativo do banco de dados MariaDB.
func (administrativo Administrativo) Delete(id) *errors.Application {
	administrativo.Log.Info.Print("Deletando Administrativo")

	return nil
}
