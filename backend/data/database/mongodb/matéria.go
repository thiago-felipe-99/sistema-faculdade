package mongodb

import (
	"thiagofelipe.com.br/sistema-faculdade-backend/entidades"
	"thiagofelipe.com.br/sistema-faculdade-backend/erros"
)

// MatériaBD representa a conexão com o banco de dados MongoDB para fazer
// alterações na entidade Matéria.
type MatériaBD struct {
	Connexão
}

// Inserir é uma método que adiciona uma entidade Matéria no banco de
// dados MongoDB.
func (bd MatériaBD) Inserir(*entidades.Matéria) *erros.Aplicação {
	bd.Log.Informação("Inserindo Matéria")

	return nil
}

// Atualizar é uma método que faz a atualização de uma entidade Matéria no banco
// de dados MongoDB.
func (bd MatériaBD) Atualizar(entidades.ID, *entidades.Matéria) *erros.Aplicação {
	bd.Log.Informação("Atualizando Matéria")

	return nil
}

// Pegar é uma método que retorna uma entidade Matéria no banco de dados MongoDB.
func (bd MatériaBD) Pegar(entidades.ID) (*entidades.Matéria, *erros.Aplicação) {
	bd.Log.Informação("Pegando Matéria")

	return nil, nil
}

// Deletar é uma método que remove uma entidade Matéria no banco de dados MongoDB.
func (bd MatériaBD) Deletar(entidades.ID) *erros.Aplicação {
	bd.Log.Informação("Deletando Matéria")

	return nil
}
