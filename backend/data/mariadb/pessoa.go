package mariadb

import (
	"database/sql"

	. "thiagofelipe.com.br/sistema-faculdade-backend/data"
	"thiagofelipe.com.br/sistema-faculdade-backend/entidades"
	"thiagofelipe.com.br/sistema-faculdade-backend/erros"
)

// PessoaBD representa a conexão com o banco de dados MariaDB para fazer
// alterações na entidade Pessoa.
type PessoaBD struct {
	Conexão
	NomeDaTabela string
}

// Inserir é uma método que faz adiciona uma entidade Pessoa no banco de dados
// MariaDB.
func (bd PessoaBD) Inserir(pessoa *entidades.Pessoa) *erros.Aplicação {
	bd.Log.Informação("Inserindo Pessoa com o seguinte ID: " + pessoa.ID.String())

	query := "INSERT INTO " + bd.NomeDaTabela +
		" (ID, Nome, CPF, Data_De_Nascimento, Senha) VALUES (?, ?, ?, ?, ?)"

	_, erro := bd.BD.Exec(
		query,
		pessoa.ID,
		pessoa.Nome,
		pessoa.CPF,
		pessoa.DataDeNascimento,
		pessoa.Senha,
	)

	if erro != nil {
		bd.Log.Aviso(
			"Erro ao inserir a Pessoa com o seguinte ID: "+pessoa.ID.String(),
			"\n\t"+erros.ErroExterno(erro),
		)
		return erros.Novo(ErroInserirPessoa, nil, erro)
	}

	return nil
}

// Atualizar é uma método que faz a atualização de uma entidade Pessoa no banco
// de dados MariaDB.
func (bd PessoaBD) Atualizar(id entidades.ID, pessoa *entidades.Pessoa) *erros.Aplicação {
	bd.Log.Informação("Atualizando Pessoa com o seguinte ID: " + id.String())

	query := "UPDATE " + bd.NomeDaTabela +
		" SET Nome = ?, CPF = ?, Data_De_Nascimento = ?, Senha = ?" +
		" WHERE ID = ?"

	_, erro := bd.BD.Exec(
		query,
		pessoa.Nome,
		pessoa.CPF,
		pessoa.DataDeNascimento,
		pessoa.Senha,
		id,
	)

	if erro != nil {
		bd.Log.Aviso(
			"Erro ao atualizar a Pessoa com o seguinte ID: "+id.String(),
			"\n\t"+erros.ErroExterno(erro),
		)
		return erros.Novo(ErroAtualizarPessoa, nil, erro)
	}

	return nil
}

// Pegar é uma método que retorna uma entidade Pessoa no banco de dados MariaDB.
func (bd PessoaBD) Pegar(id entidades.ID) (*entidades.Pessoa, *erros.Aplicação) {
	bd.Log.Informação("Pegando Pessoa com o seguinte ID: " + id.String())

	var pessoa entidades.Pessoa

	query := "SELECT ID, Nome, CPF, Data_De_Nascimento, Senha FROM " +
		bd.NomeDaTabela + " WHERE ID = ?"

	row := bd.BD.QueryRow(query, id)

	erro := row.Scan(
		&pessoa.ID,
		&pessoa.Nome,
		&pessoa.CPF,
		&pessoa.DataDeNascimento,
		&pessoa.Senha,
	)

	if erro != nil {

		if erro == sql.ErrNoRows {
			bd.Log.Aviso(
				"Não foi encontrada nenhuma a pessoa com o seguinte ID: "+id.String(),
				"\n\t"+erros.ErroExterno(erro),
			)
			return nil, erros.Novo(ErroPessoaNãoEncontrada, nil, erro)
		}

		bd.Log.Aviso(
			"Erro ao tentar econtrar a pessoa com o seguinte ID: "+id.String(),
			"\n\t"+erros.ErroExterno(erro),
		)
		return nil, erros.Novo(ErroPegarPessoa, nil, erro)

	}

	return &pessoa, nil
}

// PegarPorCPF é uma método que retorna uma entidade Pessoa no banco de dados
// MariaDB.
func (bd PessoaBD) PegarPorCPF(cpf entidades.CPF) (*entidades.Pessoa, *erros.Aplicação) {
	bd.Log.Informação("Pegando Pessoa com o seguinte CPF: " + cpf)

	var pessoa entidades.Pessoa

	query := "SELECT ID, Nome, CPF, Data_De_Nascimento, Senha FROM " +
		bd.NomeDaTabela + " WHERE CPF = ?"

	row := bd.BD.QueryRow(query, cpf)

	erro := row.Scan(
		&pessoa.ID,
		&pessoa.Nome,
		&pessoa.CPF,
		&pessoa.DataDeNascimento,
		&pessoa.Senha,
	)

	if erro != nil {

		if erro == sql.ErrNoRows {
			bd.Log.Aviso(
				"Não foi encontrada nenhuma a pessoa com o seguinte CPF: "+cpf,
				"\n\t"+erros.ErroExterno(erro),
			)
			return nil, erros.Novo(ErroPessoaNãoEncontrada, nil, erro)
		}

		bd.Log.Aviso(
			"Erro ao tentar econtrar a pessoa com o seguinte CPF: "+cpf,
			"\n\t"+erros.ErroExterno(erro),
		)
		return nil, erros.Novo(ErroPegarPessoaPorCPF, nil, erro)

	}

	return &pessoa, nil

}

// Deletar é uma método que remove uma entidade Pessoa no banco de dados
// MariaDB.
func (bd PessoaBD) Deletar(id entidades.ID) *erros.Aplicação {
	bd.Log.Informação("Deletando Pessoa com o seguinte ID: " + id.String())

	query := "DELETE FROM " + bd.NomeDaTabela + " WHERE ID = ?"

	_, erro := bd.BD.Exec(query, id)

	if erro != nil {
		bd.Log.Aviso(
			"Erro ao tentar deletar a pessoa com o seguinte ID: "+id.String(),
			"\n\t"+erros.ErroExterno(erro),
		)

		return erros.Novo(ErroDeletarPessoa, nil, erro)
	}

	return nil
}
