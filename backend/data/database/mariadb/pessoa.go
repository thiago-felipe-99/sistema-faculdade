package mariadb

import (
	"database/sql"

	"thiagofelipe.com.br/sistema-faculdade/data"
	"thiagofelipe.com.br/sistema-faculdade/errors"
)

// PessoaDB representa a conexão com o banco de dados MariaDB para fazer
// alterações na entidade PessoaDB.
type PessoaDB struct {
	Connection
	TableName string
}

// Insert é uma função que faz inserção de uma Pessoa no banco de dados MariaDB.
func (db PessoaDB) Insert(pessoa *data.Pessoa) *errors.Application {
	db.Log.Info.Println("Inserindo Pessoa com o seguinte ID: " + pessoa.ID.String())

	query := "INSERT INTO " + db.TableName +
		" (ID, Nome, CPF, Data_De_Nascimento, Senha) VALUES (?, ?, ?, ?, ?)"

	_, err := db.DB.Exec(
		query,
		pessoa.ID,
		pessoa.Nome,
		pessoa.CPF,
		pessoa.DataDeNascimento,
		pessoa.Senha,
	)

	if err != nil {
		db.Log.Warning.Println(
			"Erro ao inserir a Pessoa com o seguinte ID: "+pessoa.ID.String(),
			"\nErro: "+err.Error(),
		)
		return errors.New(errors.InserirPessoa, nil, err)
	}

	return nil
}

// Update é uma função que faz a atualização de Pessoa no banco de dados MariaDB.
func (db PessoaDB) Update(id id, pessoa *data.Pessoa) *errors.Application {
	db.Log.Info.Println("Atualizando Pessoa com o seguinte ID: " + id.String())

	query := "UPDATE " + db.TableName +
		" SET Nome = ?, CPF = ?, Data_De_Nascimento = ?, Senha = ?" +
		" WHERE ID = ?"

	_, err := db.DB.Exec(
		query,
		pessoa.Nome,
		pessoa.CPF,
		pessoa.DataDeNascimento,
		pessoa.Senha,
		id,
	)

	if err != nil {
		db.Log.Warning.Println(
			"Erro ao atualizar a Pessoa com o seguinte ID: "+id.String(),
			"\nErro: "+err.Error(),
		)
		return errors.New(errors.AtualizarPessoa, nil, err)
	}
	return nil
}

// Get é uma função que retorna uma Pessoa do banco de dados MariaDB.
func (db PessoaDB) Get(id id) (*data.Pessoa, *errors.Application) {
	db.Log.Info.Println("Pegando Pessoa com o seguinte ID: " + id.String())

	var pessoa data.Pessoa

	query := "SELECT ID, Nome, CPF, Data_De_Nascimento, Senha FROM " + db.TableName +
		" WHERE ID = ?"

	row := db.DB.QueryRow(query, id)

	err := row.Scan(
		&pessoa.ID,
		&pessoa.Nome,
		&pessoa.CPF,
		&pessoa.DataDeNascimento,
		&pessoa.Senha,
	)

	if err != nil {

		if err == sql.ErrNoRows {
			db.Log.Warning.Println(
				"Não foi encontrada nenhuma a pessoa com o seguinte ID: "+id.String(),
				"\nErro: "+err.Error(),
			)
			return nil, errors.New(errors.PessoaNãoEncontrada, nil, err)
		}

		db.Log.Warning.Println(
			"Erro ao tentar econtrar a pessoa com o seguinte ID: "+id.String(),
			"\nErro: "+err.Error(),
		)
		return nil, errors.New(errors.PegarPessoa, nil, err)

	}

	return &pessoa, nil
}

// Delete é uma função que remove uma Pessoa do banco de dados MariaDB.
func (db PessoaDB) Delete(id id) *errors.Application {
	db.Log.Info.Print("Deletando Pessoa com o seguinte ID: " + id.String())

	query := "DELETE FROM " + db.TableName + " WHERE ID = ?"

	_, err := db.DB.Exec(query, id)

	if err != nil {
		db.Log.Warning.Println(
			"Erro ao tentar deletar a pessoa com o seguinte ID: "+id.String(),
			"\nErro: "+err.Error(),
		)

		return errors.New(errors.DeletarPessoa, nil, err)
	}

	return nil
}
