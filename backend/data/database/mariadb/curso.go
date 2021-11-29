package mariadb

import (
	"thiagofelipe.com.br/sistema-faculdade/data"
	"thiagofelipe.com.br/sistema-faculdade/errors"
)

// CursoBD representa a conexão com o banco de dados MariaDB para fazer
// alterações na entidade Curso.
type CursoBD struct {
	Conexão
	NomeDaTabela           string
	NomeDaTabelaSecundária string
}

// InserirMatérias inseres as matérias de um curso no banco de dados MariaDB.
func (bd CursoBD) InserirMatérias(matérias *[]data.CursoMatéria) *errors.Aplicação {
	bd.Log.Informação.Println("Inserindo matérias no Curso")

	if len(*matérias) <= 0 {
		return errors.New(errors.InserirCursoMatériasTamanhoMínimo, nil, nil)
	}

	query := "INSERT INTO " + bd.NomeDaTabelaSecundária +
		"(ID_Curso, ID_Matéria, Período, Tipo, Status, Observação) VALUES"

	var params []interface{}
	for _, matéria := range *matérias {
		query += " (?, ?, ?, ?, ?, ?),"
		params = append(
			params,
			matéria.ID_Curso,
			matéria.ID_Matéria,
			matéria.Período,
			matéria.Tipo,
			matéria.Status,
			matéria.Observação,
		)
	}

	query = query[:len(query)-1]

	_, err := bd.BD.Exec(query, params...)
	if err != nil {
		bd.Log.Aviso.Println(
			"Erro ao inserir as matérias do curso\nErro: " + err.Error(),
		)
		return errors.New(errors.InserirCursoMatérias, nil, err)
	}

	return nil
}

// Inserir é uma função que faz inserção de uma Curso no banco de dados MariaDB.
func (bd CursoBD) Inserir(curso *data.Curso) *errors.Aplicação {
	bd.Log.Informação.Println("Inserindo Curso com o seguinte ID: " + curso.ID.String())

	query := "INSERT INTO " + bd.NomeDaTabela +
		"(ID, Nome, Data_De_Início, Data_De_Desativação) VALUES(?, ?, ?, ?)"

	_, err := bd.BD.Exec(
		query,
		curso.ID,
		curso.Nome,
		curso.DataDeInício,
		curso.DataDeDesativação,
	)
	if err != nil {
		bd.Log.Aviso.Println(
			"Erro ao inserir o Curso com o seguinte ID: "+curso.ID.String(),
			"\nErro: "+err.Error(),
		)
		return errors.New(errors.InserirCurso, nil, err)
	}

	return bd.InserirMatérias(&curso.Matérias)
}

// AtualizarMatérias é uma função que atualiza as matérias dos cursos no banco
// de dados MariaDB
func (bd CursoBD) AtualizarMatérias(matérias *[]data.CursoMatéria) *errors.Aplicação {
	bd.Log.Informação.Println("Atualizando matérias no Curso")

	return nil
}

// Atualizar é uma função que faz a atualização de Curso no banco de dados
// MariaDB.
func (bd CursoBD) Atualizar(id id, curso *data.Curso) *errors.Aplicação {
	bd.Log.Informação.Println("Atualizando Curso com o seguinte ID: " + id.String())

	return nil
}

// PegarMatérias é uma função que retonar as matérias de um Curso que está salvo
// no banco de dados MariaDB.
func (bd CursoBD) PegarMatérias(id id) (*[]data.CursoMatéria, *errors.Aplicação) {
	return nil, nil
}

// Pegar é uma função que retorna uma Curso do banco de dados MariaDB.
func (bd CursoBD) Pegar(id id) (*data.Curso, *errors.Aplicação) {
	bd.Log.Informação.Println("Pegando Curso com o seguinte ID: " + id.String())

	return nil, nil
}

// DeletarMatérias é uma função que deleta as matérias de um Curso que está salvo
// no banco de dados MariaDB.
func (bd CursoBD) DeletarMatérias(id id) *errors.Aplicação {
	return nil
}

// Deletar é uma função que remove uma Curso do banco de dados MariaDB.
func (bd CursoBD) Deletar(id id) *errors.Aplicação {
	bd.Log.Informação.Print("Deletando Curso com o seguinte ID: " + id.String())

	return nil
}
