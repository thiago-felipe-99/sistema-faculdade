package mariadb

import (
	"thiagofelipe.com.br/sistema-faculdade/entidades"
	"thiagofelipe.com.br/sistema-faculdade/errors"
)

// AdministrativoBD representa a conexão com o banco de dados MariaDB para fazer
// alterações na entidade AdministrativoBD.
type AdministrativoBD struct {
	Conexão
}

// Inserir é uma função que faz inserção de uma Administrativo no banco de dados
// MariaDB.
func (bd AdministrativoBD) Inserir(*entidades.Administrativo) *errors.Aplicação {
	bd.Log.Informação.Println("Inserindo Administrativo")

	return nil
}

// Atualizar é uma função que faz a atualização de Administrativo no banco de dados
// MariaDB.
func (bd AdministrativoBD) Atualizar(id, *entidades.Administrativo) *errors.Aplicação {
	bd.Log.Informação.Println("Atualizando Administrativo")

	return nil
}

// Pegar é uma função que retorna uma Administrativo do banco de dados MariaDB.
func (bd AdministrativoBD) Pegar(id) (*entidades.Administrativo, *errors.Aplicação) {
	bd.Log.Informação.Println("Pegando Administrativo")

	return nil, nil
}

// Deletar é uma função que remove uma Administrativo do banco de dados MariaDB.
func (bd AdministrativoBD) Deletar(id) *errors.Aplicação {
	bd.Log.Informação.Print("Deletando Administrativo")

	return nil
}
