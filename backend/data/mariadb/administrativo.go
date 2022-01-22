package mariadb

import (
	"thiagofelipe.com.br/sistema-faculdade-backend/entidades"
	"thiagofelipe.com.br/sistema-faculdade-backend/erros"
)

// AdministrativoBD representa a conexão com o banco de dados MariaDB para fazer
// alterações na entidade Administrativo.
type AdministrativoBD struct {
	Conexão
}

// Inserir é uma método que adiciona uma entidade Administrativo no banco
// de dados MariaDB.
func (bd AdministrativoBD) Inserir(*entidades.Administrativo) *erros.Aplicação {
	bd.Log.Informação("Inserindo Administrativo")

	return nil
}

// Atualizar é uma método que faz a atualização de uma entidade Administrativo
// no banco de dados MariaDB.
func (bd AdministrativoBD) Atualizar(entidades.ID, *entidades.Administrativo) *erros.Aplicação {
	bd.Log.Informação("Atualizando Administrativo")

	return nil
}

// Pegar é uma método que retorna uma entidade Administrativo no banco de dados
// MariaDB.
func (bd AdministrativoBD) Pegar(entidades.ID) (*entidades.Administrativo, *erros.Aplicação) {
	bd.Log.Informação("Pegando Administrativo")

	return nil, nil
}

// Deletar é uma método que remove uma entidade Administrativo no banco de dados
// MariaDB.
func (bd AdministrativoBD) Deletar(entidades.ID) *erros.Aplicação {
	bd.Log.Informação("Deletando Administrativo")

	return nil
}