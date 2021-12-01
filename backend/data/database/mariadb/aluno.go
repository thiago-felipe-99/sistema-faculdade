package mariadb

import (
	"thiagofelipe.com.br/sistema-faculdade/entidades"
	"thiagofelipe.com.br/sistema-faculdade/errors"
)

// AlunoBD representa a conexão com o banco de dados MariaDB para fazer alterações
// na entidade AlunoBD.
type AlunoBD struct {
	Conexão
}

// Inserir é uma função que faz inserção de uma Aluno no banco de dados MariaDB.
func (bd AlunoBD) Inserir(*entidades.Aluno) *errors.Aplicação {
	bd.Log.Informação.Println("Inserindo Aluno")

	return nil
}

// Atualizar é uma função que faz a atualização de Aluno no banco de dados MariaDB.
func (bd AlunoBD) Atualizar(id, *entidades.Aluno) *errors.Aplicação {
	bd.Log.Informação.Println("Atualizando Aluno")

	return nil
}

// Pegar é uma função que retorna uma Aluno do banco de dados MariaDB.
func (bd AlunoBD) Pegar(id) (*entidades.Aluno, *errors.Aplicação) {
	bd.Log.Informação.Println("Pegando Aluno")

	return nil, nil
}

// Deletar é uma função que remove uma Aluno do banco de dados MariaDB.
func (bd AlunoBD) Deletar(id) *errors.Aplicação {
	bd.Log.Informação.Print("Deletando Aluno")

	return nil
}
