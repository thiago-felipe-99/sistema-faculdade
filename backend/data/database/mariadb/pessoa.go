package mariadb

import (
	"database/sql"

	"thiagofelipe.com.br/sistema-faculdade/entidades"
	"thiagofelipe.com.br/sistema-faculdade/errors"
)

// PessoaBD representa a conexão com o banco de dados MariaDB para fazer
// alterações na entidade PessoaBD.
type PessoaBD struct {
	Conexão
	NomeDaTabela string
}

// Inserir é uma função que faz inserção de uma Pessoa no banco de dados MariaDB.
func (bd PessoaBD) Inserir(pessoa *entidades.Pessoa) *errors.Aplicação {
	bd.Log.Informação.Println("Inserindo Pessoa com o seguinte ID: " + pessoa.ID.String())

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
		bd.Log.Aviso.Println(
			"Erro ao inserir a Pessoa com o seguinte ID: "+pessoa.ID.String(),
			"\nErro: "+err.Error(),
		)
		return errors.New(errors.InserirPessoa, nil, err)
	}

	return nil
}

// Atualizar é uma função que faz a atualização de Pessoa no banco de dados MariaDB.
func (bd PessoaBD) Atualizar(id id, pessoa *entidades.Pessoa) *errors.Aplicação {
	bd.Log.Informação.Println("Atualizando Pessoa com o seguinte ID: " + id.String())

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
		bd.Log.Aviso.Println(
			"Erro ao atualizar a Pessoa com o seguinte ID: "+id.String(),
			"\nErro: "+err.Error(),
		)
		return errors.New(errors.AtualizarPessoa, nil, err)
	}
	return nil
}

// Pegar é uma função que retorna uma Pessoa do banco de dados MariaDB.
func (bd PessoaBD) Pegar(id id) (*entidades.Pessoa, *errors.Aplicação) {
	bd.Log.Informação.Println("Pegando Pessoa com o seguinte ID: " + id.String())

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

		if err == sql.ErrNoRows {
			bd.Log.Aviso.Println(
				"Não foi encontrada nenhuma a pessoa com o seguinte ID: "+id.String(),
				"\nErro: "+err.Error(),
			)
			return nil, errors.New(errors.PessoaNãoEncontrada, nil, err)
		}

		bd.Log.Aviso.Println(
			"Erro ao tentar econtrar a pessoa com o seguinte ID: "+id.String(),
			"\nErro: "+err.Error(),
		)
		return nil, errors.New(errors.PegarPessoa, nil, err)

	}

	return &pessoa, nil
}

// Deletar é uma função que remove uma Pessoa do banco de dados MariaDB.
func (bd PessoaBD) Deletar(id id) *errors.Aplicação {
	bd.Log.Informação.Print("Deletando Pessoa com o seguinte ID: " + id.String())

	query := "DELETE FROM " + bd.NomeDaTabela + " WHERE ID = ?"

	_, err := bd.BD.Exec(query, id)

	if err != nil {
		bd.Log.Aviso.Println(
			"Erro ao tentar deletar a pessoa com o seguinte ID: "+id.String(),
			"\nErro: "+err.Error(),
		)

		return errors.New(errors.DeletarPessoa, nil, err)
	}

	return nil
}
