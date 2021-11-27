package mariadb

import (
	"thiagofelipe.com.br/sistema-faculdade/data"
	"thiagofelipe.com.br/sistema-faculdade/errors"
)

type pessoaToInsert = data.PessoaToInsert

// Pessoa representa a conexão com o banco de dados MariaDB para fazer alterações
// na entidade Pessoa.
type Pessoa struct {
	Connection
}

// Insert é uma função que faz inserção de uma Pessoa no banco de dados MariaDB.
func (pessoa Pessoa) Insert(*pessoaToInsert) (*data.Pessoa, *errors.Application) {
	pessoa.Log.Info.Println("Inserindo Pessoa")

	return nil, nil
}

// Update é uma função que faz a atualização de Pessoa no banco de dados MariaDB.
func (pessoa Pessoa) Update(id, *pessoaToInsert) (*data.Pessoa, *errors.Application) {
	pessoa.Log.Info.Println("Atualizando Pessoa")

	return nil, nil
}

// Get é uma função que retorna uma Pessoa do banco de dados MariaDB.
func (pessoa Pessoa) Get(id) (*data.Pessoa, *errors.Application) {
	pessoa.Log.Info.Println("Pegando Pessoa")

	return nil, nil
}

// Delete é uma função que remove uma Pessoa do banco de dados MariaDB.
func (pessoa Pessoa) Delete(id) *errors.Application {
	pessoa.Log.Info.Print("Deletando Pessoa")

	return nil
}
