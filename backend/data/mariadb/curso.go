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

// InserirMatérias inseres as matérias da entidade Curso no banco de dados
// MariaDB.
func (bd CursoBD) InserirMatérias(matérias *[]entidades.CursoMatéria) *erros.Aplicação {
	bd.Log.Informação("Inserindo matérias no Curso")

	if len(*matérias) <= 0 {
		bd.Log.Aviso("Erro ao inserir matérias do curso, não tem a quantidade mínima de matérias")
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
		bd.Log.Aviso(
			"Erro ao inserir as matérias do curso\n\t" + erros.ErroExterno(erroTx),
		)
		return erros.Novo(ErroInserirCursoMatérias, nil, erroTx)
	}

	defer tx.Rollback()

	_, erro := tx.Exec(query, params...)
	if erro != nil {
		bd.Log.Aviso(
			"Erro ao inserir as matérias do curso\n\t" + erros.ErroExterno(erro),
		)
		return erros.Novo(ErroInserirCursoMatérias, nil, erro)
	}

	if erroTx = tx.Commit(); erroTx != nil {
		bd.Log.Aviso(
			"Erro ao inserir as matérias do curso\n\tErro: " + erros.ErroExterno(erroTx),
		)
		return erros.Novo(ErroInserirCursoMatérias, nil, erroTx)
	}

	return nil
}

// Inserir é uma método que adiciona uma entidade Curso no banco de dados
// MariaDB.
func (bd CursoBD) Inserir(curso *entidades.Curso) *erros.Aplicação {
	bd.Log.Informação("Inserindo Curso com o seguinte ID: " + curso.ID.String())

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
		bd.Log.Aviso(
			"Erro ao inserir o Curso com o seguinte ID: "+curso.ID.String(),
			"\n\t"+erros.ErroExterno(erro),
		)
		return erros.Novo(ErroInserirCurso, nil, erro)
	}

	return bd.InserirMatérias(&curso.Matérias)
}

// AtualizarMatérias é uma método que atualiza as matérias de entidades Cursos
// no banco de dados MariaDB.
func (bd CursoBD) AtualizarMatérias(matérias *[]entidades.CursoMatéria) *erros.Aplicação {
	bd.Log.Informação("Atualizando matérias no Curso")

	query := "UPDATE " + bd.NomeDaTabelaSecundária +
		" SET Período = ?, Tipo = ?, Status = ?, Observação = ? " +
		" WHERE ID_Curso = ? AND ID_Matéria = ?"

	tx, erroTx := bd.BD.Begin()
	if erroTx != nil {
		bd.Log.Aviso(
			"Erro ao atualizar as matérias do curso\n\t" + erros.ErroExterno(erroTx),
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
			bd.Log.Aviso(
				"Erro ao atualizar as matérias do curso\n\t" + erros.ErroExterno(erro),
			)
			return erros.Novo(ErroInserirCursoMatérias, nil, erro)
		}
	}

	if erroTx = tx.Commit(); erroTx != nil {
		bd.Log.Aviso(
			"Erro ao atualizar as matérias do curso\n\t" + erros.ErroExterno(erroTx),
		)
		return erros.Novo(ErroInserirCurso, nil, erroTx)
	}

	return nil
}

// Atualizar é uma método que faz a atualização de uma entidade Curso no banco
// de dados MariaDB.
func (bd CursoBD) Atualizar(id entidades.ID, curso *entidades.Curso) *erros.Aplicação {
	bd.Log.Informação("Atualizando Curso com o seguinte ID: " + id.String())

	erro := bd.AtualizarMatérias(&curso.Matérias)
	if erro != nil {
		bd.Log.Aviso(
			"Erro ao atualizar o curso com o seguinte ID: "+id.String(),
			"\n\t"+erro.Error(),
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
		bd.Log.Aviso(
			"Erro ao atualizar o curso com o seguinte ID: "+id.String(),
			"\n\t"+erros.ErroExterno(erroBD),
		)
		return erros.Novo(ErroAtualizarCurso, nil, erroBD)
	}

	return nil
}

// PegarMatérias é uma método que retonar as matérias de uma entidade Curso
// no banco de dados MariaDB.
func (bd CursoBD) PegarMatérias(idCurso entidades.ID) (*[]entidades.CursoMatéria, *erros.Aplicação) {
	bd.Log.Informação("Pegando as matérias do Curso com o seguinte ID: " + idCurso.String())

	var matérias []entidades.CursoMatéria

	query := "SELECT ID_Curso, ID_Matéria, Período, Tipo, Status, Observação FROM " +
		bd.NomeDaTabelaSecundária + " WHERE ID_Curso = ?"

	linhas, erro := bd.BD.Query(query, idCurso)
	if erro != nil {
		bd.Log.Aviso(
			"Erro ao pegar as matérias do Curso com o seguinte ID: "+idCurso.String(),
			"\n\t"+erros.ErroExterno(erro),
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
			bd.Log.Aviso(
				"Erro ao pegar as matérias do Curso com o seguinte ID: "+idCurso.String(),
				"\n\t"+erros.ErroExterno(erro),
			)

			return nil, erros.Novo(ErroPegarCursoMatérias, nil, erro)
		}

		matérias = append(matérias, matéria)

	}

	if erro = linhas.Err(); erro != nil {
		bd.Log.Aviso(
			"Erro ao pegar as matérias do Curso com o seguinte ID: "+idCurso.String(),
			"\n\t"+erros.ErroExterno(erro),
		)

		return nil, erros.Novo(ErroPegarCursoMatérias, nil, erro)
	}

	return &matérias, nil
}

// Pegar é uma método que retorna uma entidade Curso no banco de dados MariaDB.
func (bd CursoBD) Pegar(id entidades.ID) (*entidades.Curso, *erros.Aplicação) {
	bd.Log.Informação("Pegando Curso com o seguinte ID: " + id.String())

	matérias, erroAplicação := bd.PegarMatérias(id)
	if erroAplicação != nil &&
		!erroAplicação.ÉPadrão(ErroCursoMatériasNãoEncontrado) {

		bd.Log.Aviso(
			"Erro ao pegar as matérias do Curso com o seguinte ID: "+id.String(),
			"\n\t"+erroAplicação.Error(),
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
			bd.Log.Aviso(
				"Não foi encontrada nenhum curso com o seguinte ID: "+id.String(),
				"\n\t"+erros.ErroExterno(erro),
			)
			return nil, erros.Novo(ErroCursoNãoEncontrado, nil, erro)
		}

		bd.Log.Aviso(
			"Erro ao pegar o curso com o seguinte ID: "+id.String(),
			"\n\t"+erros.ErroExterno(erro),
		)
		return nil, erros.Novo(ErroPegarCurso, nil, erro)
	}

	return &curso, nil
}

// DeletarMatérias é uma método que remove as matérias de uma entidade Curso no
// banco de dados MariaDB.
func (bd CursoBD) DeletarMatérias(idCurso entidades.ID) *erros.Aplicação {
	bd.Log.Informação("Deletando as matérias do Curso com o seguinte ID: " + idCurso.String())

	query := "DELETE FROM " + bd.NomeDaTabelaSecundária + " WHERE ID_Curso = ?"

	_, erro := bd.BD.Exec(query, idCurso)

	if erro != nil {
		bd.Log.Aviso(
			"Erro ao tentar deletar as matérias do curso com o seguinte ID: "+idCurso.String(),
			"\n\t"+erros.ErroExterno(erro),
		)

		return erros.Novo(ErroDeletarCursoMatérias, nil, erro)
	}

	return nil
}

// Deletar é uma método que remove uma entidade Curso no banco de dados MariaDB.
func (bd CursoBD) Deletar(id entidades.ID) *erros.Aplicação {
	bd.Log.Informação("Deletando Curso com o seguinte ID: " + id.String())

	erro := bd.DeletarMatérias(id)
	if erro != nil {
		bd.Log.Aviso(
			"Erro ao tentar deletar curso com o seguinte ID: "+id.String(),
			"\n\t"+erro.Error(),
		)
		return erros.Novo(ErroDeletarCurso, erro, nil)
	}

	query := "DELETE FROM " + bd.NomeDaTabela + " WHERE ID = ?"

	_, erroBD := bd.BD.Exec(query, id)

	if erroBD != nil {
		bd.Log.Aviso(
			"Erro ao tentar deletar curso com o seguinte ID: "+id.String(),
			"\n\t"+erros.ErroExterno(erroBD),
		)

		return erros.Novo(ErroDeletarCurso, nil, erroBD)
	}

	return nil
}
