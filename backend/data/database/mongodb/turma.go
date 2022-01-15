package mongodb

import (
	"thiagofelipe.com.br/sistema-faculdade-backend/entidades"
	"thiagofelipe.com.br/sistema-faculdade-backend/erros"
)

// TurmaBD representa a conexão com o banco de dados MongoDB para fazer alterações
// na entidade Turma.
type TurmaBD struct {
	Connexão
}

// Inserir é uma método que adiciona uma entidade Turma no banco de dados
// MongoDB.
func (bd TurmaBD) Inserir(*entidades.Turma) *erros.Aplicação {
	bd.Log.Informação("Inserindo Turma")

	return nil
}

// Atualizar é uma método que faz a atualização de entidade Turma no banco de
// dados MongoDB.
func (bd TurmaBD) Atualizar(entidades.ID, *entidades.Turma) *erros.Aplicação {
	bd.Log.Informação("Atualizando Turma")

	return nil
}

// Pegar é uma método que retorna uma entidade Turma no banco de dados MongoDB.
func (bd TurmaBD) Pegar(entidades.ID) (*entidades.Turma, *erros.Aplicação) {
	bd.Log.Informação("Pegando Turma")

	return nil, nil
}

// Deletar é uma método que remove uma entidade Turma no banco de dados MongoDB.
func (bd TurmaBD) Deletar(entidades.ID) *erros.Aplicação {
	bd.Log.Informação("Deletando Turma")

	return nil
}
