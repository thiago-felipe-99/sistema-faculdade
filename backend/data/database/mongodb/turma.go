package mongodb

import (
	"thiagofelipe.com.br/sistema-faculdade/data"
	"thiagofelipe.com.br/sistema-faculdade/errors"
)

// TurmaBD representa a conexão com o banco de dados MariaDB para fazer alterações
// na entidade TurmaBD.
type TurmaBD struct {
	Connexão
}

// Inserir é uma função que faz inserção de uma Turma no banco de dados MariaDB.
func (bd TurmaBD) Inserir(*data.Turma) *errors.Aplicação {
	bd.Log.Informação.Println("Inserindo Turma")

	return nil
}

// Atualizar é uma função que faz a atualização de Turma no banco de dados MariaDB.
func (bd TurmaBD) Atualizar(id, *data.Turma) *errors.Aplicação {
	bd.Log.Informação.Println("Atualizando Turma")

	return nil
}

// Pegar é uma função que retorna uma Turma do banco de dados MariaDB.
func (bd TurmaBD) Pegar(id) (*data.Turma, *errors.Aplicação) {
	bd.Log.Informação.Println("Pegando Turma")

	return nil, nil
}

// Deletar é uma função que remove uma Turma do banco de dados MariaDB.
func (bd TurmaBD) Deletar(id) *errors.Aplicação {
	bd.Log.Informação.Print("Deletando Turma")

	return nil
}
