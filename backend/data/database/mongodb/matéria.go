package mongodb

import (
	"thiagofelipe.com.br/sistema-faculdade/entidades"
	"thiagofelipe.com.br/sistema-faculdade/errors"
)

// MatériaBD representa a conexão com o banco de dados MariaDB para fazer alterações
// na entidade MatériaBD.
type MatériaBD struct {
	Connexão
}

// Inserir é uma função que faz inserção de uma Matéria no banco de dados MariaDB.
func (bd MatériaBD) Inserir(*entidades.Matéria) *errors.Aplicação {
	bd.Log.Informação.Println("Inserindo Matéria")

	return nil
}

// Atualizar é uma função que faz a atualização de Matéria no banco de dados MariaDB.
func (bd MatériaBD) Atualizar(id, *entidades.Matéria) *errors.Aplicação {
	bd.Log.Informação.Println("Atualizando Matéria")

	return nil
}

// Pegar é uma função que retorna uma Matéria do banco de dados MariaDB.
func (bd MatériaBD) Pegar(id) (*entidades.Matéria, *errors.Aplicação) {
	bd.Log.Informação.Println("Pegando Matéria")

	return nil, nil
}

// Deletar é uma função que remove uma Matéria do banco de dados MariaDB.
func (bd MatériaBD) Deletar(id) *errors.Aplicação {
	bd.Log.Informação.Print("Deletando Matéria")

	return nil
}
