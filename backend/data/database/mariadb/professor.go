package mariadb

import (
	"thiagofelipe.com.br/sistema-faculdade/data"
	"thiagofelipe.com.br/sistema-faculdade/errors"
)

// ProfessorBD representa a conexão com o banco de dados MariaDB para fazer
// alterações na entidade ProfessorBD.
type ProfessorBD struct {
	Conexão
}

// Inserir é uma função que faz inserção de uma Professor no banco de dados MariaDB.
func (bd ProfessorBD) Inserir(*data.Professor) *errors.Aplicação {
	bd.Log.Informação.Println("Inserindo Professor")

	return nil
}

// Atualizar é uma função que faz a atualização de Professor no banco de dados
// MariaDB.
func (bd ProfessorBD) Atualizar(id, *data.Professor) *errors.Aplicação {
	bd.Log.Informação.Println("Atualizando Professor")

	return nil
}

// Pegar é uma função que retorna uma Professor do banco de dados MariaDB.
func (bd ProfessorBD) Pegar(id) (*data.Professor, *errors.Aplicação) {
	bd.Log.Informação.Println("Pegando Professor")

	return nil, nil
}

// Deletar é uma função que remove uma Professor do banco de dados MariaDB.
func (bd ProfessorBD) Deletar(id) *errors.Aplicação {
	bd.Log.Informação.Print("Deletando Professor")

	return nil
}
