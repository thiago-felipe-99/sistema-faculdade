package mariadb

import (
	"thiagofelipe.com.br/sistema-faculdade/entidades"
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
func (bd CursoBD) InserirMatérias(matérias *[]entidades.CursoMatéria) *errors.Aplicação {
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
			matéria.IDCurso,
			matéria.IDMatéria,
			matéria.Período,
			matéria.Tipo,
			matéria.Status,
			matéria.Observação,
		)
	}

	query = query[:len(query)-1]

	_, erro := bd.BD.Exec(query, params...)
	if erro != nil {
		bd.Log.Aviso.Println(
			"Erro ao inserir as matérias do curso\nErro: " + erro.Error(),
		)
		return errors.New(errors.InserirCursoMatérias, nil, erro)
	}

	return nil
}

// Inserir é uma função que faz inserção de uma Curso no banco de dados MariaDB.
func (bd CursoBD) Inserir(curso *entidades.Curso) *errors.Aplicação {
	bd.Log.Informação.Println("Inserindo Curso com o seguinte ID: " + curso.ID.String())

	query := "INSERT INTO " + bd.NomeDaTabela +
		"(ID, Nome, Data_De_Início, Data_De_Desativação) VALUES(?, ?, ?, ?)"

	_, erro := bd.BD.Exec(
		query,
		curso.ID,
		curso.Nome,
		curso.DataDeInício,
		curso.DataDeDesativação,
	)
	if erro != nil {
		bd.Log.Aviso.Println(
			"Erro ao inserir o Curso com o seguinte ID: "+curso.ID.String(),
			"\nErro: "+erro.Error(),
		)
		return errors.New(errors.InserirCurso, nil, erro)
	}

	return bd.InserirMatérias(&curso.Matérias)
}

// AtualizarMatérias é uma função que atualiza as matérias dos cursos no banco
// de dados MariaDB
func (bd CursoBD) AtualizarMatérias(matérias *[]entidades.CursoMatéria) *errors.Aplicação {
	bd.Log.Informação.Println("Atualizando matérias no Curso")

	return nil
}

// Atualizar é uma função que faz a atualização de Curso no banco de dados
// MariaDB.
func (bd CursoBD) Atualizar(id id, curso *entidades.Curso) *errors.Aplicação {
	bd.Log.Informação.Println("Atualizando Curso com o seguinte ID: " + id.String())

	return nil
}

// PegarMatérias é uma função que retonar as matérias de um Curso que está salvo
// no banco de dados MariaDB.
func (bd CursoBD) PegarMatérias(id id) (*[]entidades.CursoMatéria, *errors.Aplicação) {

	var matérias []entidades.CursoMatéria

	query := "SELECT ID_Curso, ID_Matéria, Período, Tipo, Status, Observação FROM " +
		bd.NomeDaTabelaSecundária + " WHERE ID_Curso = ?"

	linhas, erro := bd.BD.Query(query, id)
	if erro != nil {
		bd.Log.Aviso.Println(
			"Erro ao pegar as matérias do Curso com o seguinte ID: "+id.String(),
			"\nErro: "+erro.Error(),
		)

		return nil, errors.New(errors.PegarCursoMatérias, nil, erro)
	}
	defer linhas.Close()

	for linhas.Next() {
		var matéria entidades.CursoMatéria

		erro := linhas.Scan(
			&matéria.IDCurso,
			&matéria.IDMatéria,
			&matéria.Período,
			&matéria.Tipo,
			&matéria.Status,
			&matéria.Observação,
		)

		if erro != nil {
			bd.Log.Aviso.Println(
				"Erro ao pegar as matérias do Curso com o seguinte ID: "+id.String(),
				"\nErro: "+erro.Error(),
			)

			return nil, errors.New(errors.PegarCursoMatérias, nil, erro)
		}

		matérias = append(matérias, matéria)

	}

	if erro = linhas.Err(); erro != nil {
		bd.Log.Aviso.Println(
			"Erro ao pegar as matérias do Curso com o seguinte ID: "+id.String(),
			"\nErro: "+erro.Error(),
		)

		return nil, errors.New(errors.PegarCursoMatérias, nil, erro)
	}

	return &matérias, nil
}

// Pegar é uma função que retorna uma Curso do banco de dados MariaDB.
func (bd CursoBD) Pegar(id id) (*entidades.Curso, *errors.Aplicação) {
	bd.Log.Informação.Println("Pegando Curso com o seguinte ID: " + id.String())

	matérias, erroAplicação := bd.PegarMatérias(id)
	if erroAplicação != nil {
		bd.Log.Aviso.Println(
			"Erro ao pegar as matérias do Curso com o seguinte ID: "+id.String(),
			"\nErro: "+erroAplicação.Error(),
		)

		return nil, errors.New(errors.PegarCurso, erroAplicação, nil)
	}

	var curso entidades.Curso
	curso.Matérias = *matérias

	query := "SELECT ID, Nome, Data_De_Início, Data_De_Desativação FROM " +
		bd.NomeDaTabela + " WHERE ID = ?"

	linha := bd.BD.QueryRow(query, id)

	erro := linha.Scan(
		&curso.ID,
		&curso.Nome,
		&curso.DataDeInício,
		&curso.DataDeDesativação,
	)

	if erro != nil {
		return nil, errors.New(errors.PegarCurso, nil, erro)
	}

	return &curso, nil
}

// DeletarMatérias é uma função que deleta as matérias de um Curso que está salvo
// no banco de dados MariaDB.
func (bd CursoBD) DeletarMatérias(id id) *errors.Aplicação {
	bd.Log.Informação.Print("Deletando as matérias do Curso com o seguinte ID: " + id.String())

	return nil
}

// Deletar é uma função que remove uma Curso do banco de dados MariaDB.
func (bd CursoBD) Deletar(id id) *errors.Aplicação {
	bd.Log.Informação.Print("Deletando Curso com o seguinte ID: " + id.String())

	return bd.DeletarMatérias(id)
}
