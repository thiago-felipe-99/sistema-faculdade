package mariadb

import (
	"database/sql"

	. "thiagofelipe.com.br/sistema-faculdade-backend/data/erros"
	"thiagofelipe.com.br/sistema-faculdade-backend/entidades"
	"thiagofelipe.com.br/sistema-faculdade-backend/erros"
)

// CursoBD representa a conexão com o banco de dados MariaDB para fazer
// alterações na entidade Curso.
type CursoBD struct {
	Conexão
	NomeDaTabela           string
	NomeDaTabelaSecundária string
}

// InserirMatérias inseres as matérias de um curso no banco de dados MariaDB.
func (bd CursoBD) InserirMatérias(matérias *[]entidades.CursoMatéria) *erros.Aplicação {
	bd.Log.Informação.Println("Inserindo matérias no Curso")

	if len(*matérias) <= 0 {
		bd.Log.Aviso.Println("Erro ao inserir matérias do curso, não tem a quantidade mínima de matérias")
		return erros.Novo(ErroInserirCursoMatériasTamanhoMínimo, nil, nil)
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

	tx, erroTx := bd.BD.Begin()
	if erroTx != nil {
		bd.Log.Aviso.Println(
			"Erro ao inserir as matérias do curso\n" + erros.ErroExterno(erroTx),
		)
		return erros.Novo(ErroInserirCursoMatérias, nil, erroTx)
	}

	defer tx.Rollback()

	_, erro := tx.Exec(query, params...)
	if erro != nil {
		bd.Log.Aviso.Println(
			"Erro ao inserir as matérias do curso\n" + erros.ErroExterno(erro),
		)
		return erros.Novo(ErroInserirCursoMatérias, nil, erro)
	}

	if erroTx = tx.Commit(); erroTx != nil {
		bd.Log.Aviso.Println(
			"Erro ao inserir as matérias do curso\nErro: " + erros.ErroExterno(erroTx),
		)
		return erros.Novo(ErroInserirCursoMatérias, nil, erroTx)
	}

	return nil
}

// Inserir é uma função que faz inserção de uma Curso no banco de dados MariaDB.
func (bd CursoBD) Inserir(curso *entidades.Curso) *erros.Aplicação {
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
			"\n"+erros.ErroExterno(erro),
		)
		return erros.Novo(ErroInserirCurso, nil, erro)
	}

	return bd.InserirMatérias(&curso.Matérias)
}

// AtualizarMatérias é uma função que atualiza as matérias dos cursos no banco
// de dados MariaDB
func (bd CursoBD) AtualizarMatérias(matérias *[]entidades.CursoMatéria) *erros.Aplicação {
	bd.Log.Informação.Println("Atualizando matérias no Curso")

	query := "UPDATE " + bd.NomeDaTabelaSecundária +
		" SET Período = ?, Tipo = ?, Status = ?, Observação = ? " +
		" WHERE ID_Curso = ? AND ID_Matéria = ?"

	tx, erroTx := bd.BD.Begin()
	if erroTx != nil {
		bd.Log.Aviso.Println(
			"Erro ao atualizar as matérias do curso\n" + erros.ErroExterno(erroTx),
		)
		return erros.Novo(ErroInserirCurso, nil, erroTx)
	}

	defer tx.Rollback()

	for _, matéria := range *matérias {
		_, erro := tx.Exec(
			query,
			matéria.Período,
			matéria.Tipo,
			matéria.Status,
			matéria.Observação,
			matéria.IDCurso,
			matéria.IDMatéria,
		)
		if erro != nil {
			bd.Log.Aviso.Println(
				"Erro ao atualizar as matérias do curso\n" + erros.ErroExterno(erro),
			)
			return erros.Novo(ErroInserirCursoMatérias, nil, erro)
		}
	}

	if erroTx = tx.Commit(); erroTx != nil {
		bd.Log.Aviso.Println(
			"Erro ao atualizar as matérias do curso\n" + erros.ErroExterno(erroTx),
		)
		return erros.Novo(ErroInserirCurso, nil, erroTx)
	}

	return nil
}

// Atualizar é uma função que faz a atualização de Curso no banco de dados
// MariaDB.
func (bd CursoBD) Atualizar(id entidades.ID, curso *entidades.Curso) *erros.Aplicação {
	bd.Log.Informação.Println("Atualizando Curso com o seguinte ID: " + id.String())

	erro := bd.AtualizarMatérias(&curso.Matérias)
	if erro != nil {
		bd.Log.Aviso.Println(
			"Erro ao atualizar o curso com o seguinte ID: "+id.String(),
			"\n"+erro.Error(),
		)
		return erros.Novo(ErroAtualizarCurso, erro, nil)
	}

	query := "UPDATE " + bd.NomeDaTabela +
		" SET Nome = ?, Data_De_Início = ?, Data_De_Desativação = ? WHERE ID = ?"

	_, erroBD := bd.BD.Exec(
		query,
		curso.Nome,
		curso.DataDeInício,
		curso.DataDeDesativação,
		id,
	)

	if erroBD != nil {
		bd.Log.Aviso.Println(
			"Erro ao atualizar o curso com o seguinte ID: "+id.String(),
			"\n"+erros.ErroExterno(erroBD),
		)
		return erros.Novo(ErroAtualizarCurso, nil, erroBD)
	}

	return nil
}

// PegarMatérias é uma função que retonar as matérias de um Curso que está salvo
// no banco de dados MariaDB.
func (bd CursoBD) PegarMatérias(idCurso entidades.ID) (*[]entidades.CursoMatéria, *erros.Aplicação) {
	bd.Log.Informação.Println("Pegando as matérias do Curso com o seguinte ID: " + idCurso.String())

	var matérias []entidades.CursoMatéria

	query := "SELECT ID_Curso, ID_Matéria, Período, Tipo, Status, Observação FROM " +
		bd.NomeDaTabelaSecundária + " WHERE ID_Curso = ?"

	linhas, erro := bd.BD.Query(query, idCurso)
	if erro != nil {
		bd.Log.Aviso.Println(
			"Erro ao pegar as matérias do Curso com o seguinte ID: "+idCurso.String(),
			"\n"+erros.ErroExterno(erro),
		)

		return nil, erros.Novo(ErroPegarCursoMatérias, nil, erro)
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
				"Erro ao pegar as matérias do Curso com o seguinte ID: "+idCurso.String(),
				"\n"+erros.ErroExterno(erro),
			)

			return nil, erros.Novo(ErroPegarCursoMatérias, nil, erro)
		}

		matérias = append(matérias, matéria)

	}

	if erro = linhas.Err(); erro != nil {
		bd.Log.Aviso.Println(
			"Erro ao pegar as matérias do Curso com o seguinte ID: "+idCurso.String(),
			"\n"+erros.ErroExterno(erro),
		)

		return nil, erros.Novo(ErroPegarCursoMatérias, nil, erro)
	}

	return &matérias, nil
}

// Pegar é uma função que retorna uma Curso do banco de dados MariaDB.
func (bd CursoBD) Pegar(id entidades.ID) (*entidades.Curso, *erros.Aplicação) {
	bd.Log.Informação.Println("Pegando Curso com o seguinte ID: " + id.String())

	matérias, erroAplicação := bd.PegarMatérias(id)
	if erroAplicação != nil &&
		!erroAplicação.ÉPadrão(ErroCursoMatériasNãoEncontrado) {

		bd.Log.Aviso.Println(
			"Erro ao pegar as matérias do Curso com o seguinte ID: "+id.String(),
			"\n"+erroAplicação.Error(),
		)

		return nil, erros.Novo(ErroPegarCurso, erroAplicação, nil)
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
		if erro == sql.ErrNoRows {
			bd.Log.Aviso.Println(
				"Não foi encontrada nenhum curso com o seguinte ID: "+id.String(),
				"\n"+erros.ErroExterno(erro),
			)
			return nil, erros.Novo(ErroCursoNãoEncontrado, nil, erro)
		}

		bd.Log.Aviso.Println(
			"Erro ao pegar o curso com o seguinte ID: "+id.String(),
			"\n"+erros.ErroExterno(erro),
		)
		return nil, erros.Novo(ErroPegarCurso, nil, erro)
	}

	return &curso, nil
}

// DeletarMatérias é uma função que deleta as matérias de um Curso que está salvo
// no banco de dados MariaDB.
func (bd CursoBD) DeletarMatérias(idCurso entidades.ID) *erros.Aplicação {
	bd.Log.Informação.Print("Deletando as matérias do Curso com o seguinte ID: " + idCurso.String())

	query := "DELETE FROM " + bd.NomeDaTabelaSecundária + " WHERE ID_Curso = ?"

	_, erro := bd.BD.Exec(query, idCurso)

	if erro != nil {
		bd.Log.Aviso.Println(
			"Erro ao tentar deletar as matérias do curso com o seguinte ID: "+idCurso.String(),
			"\n"+erros.ErroExterno(erro),
		)

		return erros.Novo(ErroDeletarCursoMatérias, nil, erro)
	}

	return nil
}

// Deletar é uma função que remove uma Curso do banco de dados MariaDB.
func (bd CursoBD) Deletar(id entidades.ID) *erros.Aplicação {
	bd.Log.Informação.Print("Deletando Curso com o seguinte ID: " + id.String())

	erro := bd.DeletarMatérias(id)
	if erro != nil {
		bd.Log.Aviso.Println(
			"Erro ao tentar deletar curso com o seguinte ID: "+id.String(),
			"\n"+erro.Error(),
		)
		return erros.Novo(ErroDeletarCurso, erro, nil)
	}

	query := "DELETE FROM " + bd.NomeDaTabela + " WHERE ID = ?"

	_, erroBD := bd.BD.Exec(query, id)

	if erroBD != nil {
		bd.Log.Aviso.Println(
			"Erro ao tentar deletar curso com o seguinte ID: "+id.String(),
			"\n"+erros.ErroExterno(erroBD),
		)

		return erros.Novo(ErroDeletarCurso, nil, erroBD)
	}

	return nil
}
