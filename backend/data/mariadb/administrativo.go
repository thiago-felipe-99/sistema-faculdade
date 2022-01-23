package mariadb

// AdministrativoBD representa a conexão com o banco de dados MariaDB para fazer
// alterações na entidade Administrativo.
type AdministrativoBD struct {
	Conexão
}

// Inserir é uma método que adiciona uma entidade Administrativo no banco
// de dados MariaDB.
func (bd AdministrativoBD) Inserir(*administrativo) erro {
	bd.Log.Informação("Inserindo Administrativo")

	return nil
}

// Atualizar é uma método que faz a atualização de uma entidade Administrativo
// no banco de dados MariaDB.
func (bd AdministrativoBD) Atualizar(id, *administrativo) erro {
	bd.Log.Informação("Atualizando Administrativo")

	return nil
}

// Pegar é uma método que retorna uma entidade Administrativo no banco de dados
// MariaDB.
func (bd AdministrativoBD) Pegar(id) (*administrativo, erro) {
	bd.Log.Informação("Pegando Administrativo")

	return nil, nil
}

// Deletar é uma método que remove uma entidade Administrativo no banco de dados
// MariaDB.
func (bd AdministrativoBD) Deletar(id) erro {
	bd.Log.Informação("Deletando Administrativo")

	return nil
}
