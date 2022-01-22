package mariadb

import (
	"database/sql"
	"errors"

	"thiagofelipe.com.br/sistema-faculdade-backend/data"
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

	_, err := bd.BD.Exec(
		query,
		pessoa.ID,
		pessoa.Nome,
		pessoa.CPF,
		pessoa.DataDeNascimento,
		pessoa.Senha,
	)
	if err != nil {
		bd.Log.Aviso(
			"Erro ao inserir a Pessoa com o seguinte ID: "+pessoa.ID.String(),
			"\n\t"+erros.ErroExterno(err),
		)

		return erros.Novo(data.ErroInserirPessoa, nil, err)
	}

	return nil
}

// Atualizar é uma método que faz a atualização de uma entidade Pessoa no banco
// de dados MariaDB.
func (bd PessoaBD) Atualizar(
	id entidades.ID,
	pessoa *entidades.Pessoa,
) *erros.Aplicação {
	bd.Log.Informação("Atualizando Pessoa com o seguinte ID: " + id.String())

	query := "UPDATE " + bd.NomeDaTabela +
		" SET Nome = ?, CPF = ?, Data_De_Nascimento = ?, Senha = ?" +
		" WHERE ID = ?"

	_, err := bd.BD.Exec(
		query,
		pessoa.Nome,
		pessoa.CPF,
		pessoa.DataDeNascimento,
		pessoa.Senha,
		id,
	)
	if err != nil {
		bd.Log.Aviso(
			"Erro ao atualizar a Pessoa com o seguinte ID: "+id.String(),
			"\n\t"+erros.ErroExterno(err),
		)

		return erros.Novo(data.ErroAtualizarPessoa, nil, err)
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

	err := row.Scan(
		&pessoa.ID,
		&pessoa.Nome,
		&pessoa.CPF,
		&pessoa.DataDeNascimento,
		&pessoa.Senha,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			bd.Log.Aviso(
				"Não foi encontrada nenhuma a pessoa com o seguinte ID: "+id.String(),
				"\n\t"+erros.ErroExterno(err),
			)

			return nil, erros.Novo(data.ErroPessoaNãoEncontrada, nil, err)
		}

		bd.Log.Aviso(
			"Erro ao tentar econtrar a pessoa com o seguinte ID: "+id.String(),
			"\n\t"+erros.ErroExterno(err),
		)

		return nil, erros.Novo(data.ErroPegarPessoa, nil, err)
	}

	return &pessoa, nil
}

// PegarPorCPF é uma método que retorna uma entidade Pessoa no banco de dados
// MariaDB.
func (bd PessoaBD) PegarPorCPF(cpf entidades.CPF) (
	*entidades.Pessoa,
	*erros.Aplicação,
) {
	bd.Log.Informação("Pegando Pessoa com o seguinte CPF: " + cpf)

	var pessoa entidades.Pessoa

	query := "SELECT ID, Nome, CPF, Data_De_Nascimento, Senha FROM " +
		bd.NomeDaTabela + " WHERE CPF = ?"

	row := bd.BD.QueryRow(query, cpf)

	err := row.Scan(
		&pessoa.ID,
		&pessoa.Nome,
		&pessoa.CPF,
		&pessoa.DataDeNascimento,
		&pessoa.Senha,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			bd.Log.Aviso(
				"Não foi encontrada nenhuma a pessoa com o seguinte CPF: "+cpf,
				"\n\t"+erros.ErroExterno(err),
			)

			return nil, erros.Novo(data.ErroPessoaNãoEncontrada, nil, err)
		}

		bd.Log.Aviso(
			"Erro ao tentar econtrar a pessoa com o seguinte CPF: "+cpf,
			"\n\t"+erros.ErroExterno(err),
		)

		return nil, erros.Novo(data.ErroPegarPessoaPorCPF, nil, err)
	}

	return &pessoa, nil
}

// Deletar é uma método que remove uma entidade Pessoa no banco de dados
// MariaDB.
func (bd PessoaBD) Deletar(id entidades.ID) *erros.Aplicação {
	bd.Log.Informação("Deletando Pessoa com o seguinte ID: " + id.String())

	query := "DELETE FROM " + bd.NomeDaTabela + " WHERE ID = ?"

	_, err := bd.BD.Exec(query, id)
	if err != nil {
		bd.Log.Aviso(
			"Erro ao tentar deletar a pessoa com o seguinte ID: "+id.String(),
			"\n\t"+erros.ErroExterno(err),
		)

		return erros.Novo(data.ErroDeletarPessoa, nil, err)
	}

	return nil
}
