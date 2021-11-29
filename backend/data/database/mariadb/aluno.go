package mariadb

import (
	"thiagofelipe.com.br/sistema-faculdade/data"
	"thiagofelipe.com.br/sistema-faculdade/errors"
)

type alunoToInsert = data.Aluno

// Aluno representa a conexão com o banco de dados MariaDB para fazer alterações
// na entidade Aluno.
type Aluno struct {
	Conexão
}

// Insert é uma função que faz inserção de uma Aluno no banco de dados MariaDB.
func (aluno Aluno) Insert(*alunoToInsert) (*data.Aluno, *errors.Application) {
	aluno.Log.Info.Println("Inserindo Aluno")

	return nil, nil
}

// Update é uma função que faz a atualização de Aluno no banco de dados MariaDB.
func (aluno Aluno) Update(id, *alunoToInsert) (*data.Aluno, *errors.Application) {
	aluno.Log.Info.Println("Atualizando Aluno")

	return nil, nil
}

// Get é uma função que retorna uma Aluno do banco de dados MariaDB.
func (aluno Aluno) Get(id) (*data.Aluno, *errors.Application) {
	aluno.Log.Info.Println("Pegando Aluno")

	return nil, nil
}

// Delete é uma função que remove uma Aluno do banco de dados MariaDB.
func (aluno Aluno) Delete(id) *errors.Application {
	aluno.Log.Info.Print("Deletando Aluno")

	return nil
}
