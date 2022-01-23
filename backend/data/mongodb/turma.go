package mongodb

// TurmaBD representa a conexão com o banco de dados MongoDB para fazer alterações
// na entidade Turma.
type TurmaBD struct {
	Conexão
}

// Inserir é uma método que adiciona uma entidade Turma no banco de dados
// MongoDB.
func (bd TurmaBD) Inserir(*turma) erro {
	bd.Log.Informação("Inserindo Turma")

	return nil
}

// Atualizar é uma método que faz a atualização de entidade Turma no banco de
// dados MongoDB.
func (bd TurmaBD) Atualizar(id, *turma) erro {
	bd.Log.Informação("Atualizando Turma")

	return nil
}

// Pegar é uma método que retorna uma entidade Turma no banco de dados MongoDB.
func (bd TurmaBD) Pegar(id) (*turma, erro) {
	bd.Log.Informação("Pegando Turma")

	return nil, nil
}

// Deletar é uma método que remove uma entidade Turma no banco de dados MongoDB.
func (bd TurmaBD) Deletar(id) erro {
	bd.Log.Informação("Deletando Turma")

	return nil
}
