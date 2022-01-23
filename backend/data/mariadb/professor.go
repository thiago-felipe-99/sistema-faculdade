package mariadb

// ProfessorBD representa a conexão com o banco de dados MariaDB para fazer
// alterações na entidade Professor.
type ProfessorBD struct {
	Conexão
}

// Inserir é uma método que adiciona uma entidade Professor no banco de
// dados MariaDB.
func (bd ProfessorBD) Inserir(*professor) erro {
	bd.Log.Informação("Inserindo Professor")

	return nil
}

// Atualizar é uma método que faz a atualização de uma entidade Professor no
// banco de dados MariaDB.
func (bd ProfessorBD) Atualizar(id, *professor) erro {
	bd.Log.Informação("Atualizando Professor")

	return nil
}

// Pegar é uma método que retorna uma entidade Professor no banco de dados
// MariaDB.
func (bd ProfessorBD) Pegar(id) (*professor, erro) {
	bd.Log.Informação("Pegando Professor")

	return nil, nil
}

// Deletar é uma método que remove uma entidade Professor no banco de dados
// MariaDB.
func (bd ProfessorBD) Deletar(id) erro {
	bd.Log.Informação("Deletando Professor")

	return nil
}
