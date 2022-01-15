package mariadb

import (
	"database/sql"

	. "thiagofelipe.com.br/sistema-faculdade-backend/data/erros"
	"thiagofelipe.com.br/sistema-faculdade-backend/entidades"
	"thiagofelipe.com.br/sistema-faculdade-backend/erros"
)

// AlunoBD representa a conexão com o banco de dados MariaDB para fazer alterações
// na entidade Aluno.
type AlunoBD struct {
	Conexão
	NomeDaTabela           string
	NomeDaTabelaSecundária string
}

// InserirTurmas é um método que adiciona turmas de entidades Aluno
// no banco de dados MariaDB.
func (bd AlunoBD) InserirTurmas(turmas *[]entidades.TurmaAluno) *erros.Aplicação {
	bd.Log.Informação("Inserindo turmas")

	if len(*turmas) <= 0 {
		return nil
	}

	query := "INSERT INTO " + bd.NomeDaTabelaSecundária +
		"(ID_Aluno, ID_Turma, Status) VALUES"

	var params []interface{}
	for _, turma := range *turmas {
		query += " (?, ?, ?),"
		params = append(
			params,
			turma.IDAluno,
			turma.IDTurma,
			turma.Status,
		)
	}

	query = query[:len(query)-1]

	tx, erroTx := bd.BD.Begin()
	if erroTx != nil {
		bd.Log.Aviso(
			"Erro ao inserir as turmas\n\t" + erros.ErroExterno(erroTx),
		)
		return erros.Novo(ErroInserirAluno, nil, erroTx)
	}

	defer tx.Rollback()

	_, erro := tx.Exec(query, params...)
	if erro != nil {
		bd.Log.Aviso(
			"Erro ao inserir as turmas\n\t" + erros.ErroExterno(erro),
		)
		return erros.Novo(ErroInserirAlunoTurma, nil, erro)
	}

	if erroTx = tx.Commit(); erroTx != nil {
		bd.Log.Aviso(
			"Erro ao inserir as turmas\n\tErro: " + erros.ErroExterno(erroTx),
		)
		return erros.Novo(ErroInserirAlunoTurma, nil, erroTx)
	}

	return nil
}

// Inserir é um método que adiciona uma entidade Aluno no banco de dados
// MariaDB.
func (bd AlunoBD) Inserir(aluno *entidades.Aluno) *erros.Aplicação {
	bd.Log.Informação("Inserindo Aluno com o seguinte ID: " + aluno.ID.String())

	query := "INSERT INTO " + bd.NomeDaTabela +
		"(ID, ID_Pessoa, ID_Curso, Matrícula, Data_De_Ingresso, Data_De_Saída," +
		" Período, Status) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"

	_, erro := bd.BD.Exec(
		query,
		aluno.ID,
		aluno.Pessoa,
		aluno.Curso,
		aluno.Matrícula,
		aluno.DataDeIngresso,
		aluno.DataDeSaída,
		aluno.Período,
		aluno.Status,
	)
	if erro != nil {
		bd.Log.Aviso(
			"Erro ao inserir o Aluno com o seguinte ID: "+aluno.ID.String(),
			"\n\t"+erros.ErroExterno(erro),
		)
		return erros.Novo(ErroInserirAluno, nil, erro)
	}

	return bd.InserirTurmas(&aluno.Turmas)
}

// Atualizar é um método que faz a atualização de uma entidade Aluno no banco de
// dados MariaDB.
func (bd AlunoBD) Atualizar(entidades.ID, *entidades.Aluno) *erros.Aplicação {
	bd.Log.Informação("Atualizando Aluno")

	return nil
}

// PegarTurmas é um método que pega as turmas de um entidade Aluno no banco de
// dados MariaDB.
func (bd AlunoBD) PegarTurmas(idAluno entidades.ID) (*[]entidades.TurmaAluno, *erros.Aplicação) {
	bd.Log.Informação("Pegando as turmas do aluno com seguinte ID: " + idAluno.String())

	var turmas []entidades.TurmaAluno

	query := "SELECT ID_Aluno, ID_Turma, Status FROM " + bd.NomeDaTabelaSecundária +
		" WHERE ID_Aluno = ?"

	linhas, erro := bd.BD.Query(query, idAluno)
	if erro != nil {
		bd.Log.Aviso(
			"Erro ao pegar as turmas do Aluno com o seguinte ID: "+idAluno.String(),
			"\n\t"+erros.ErroExterno(erro),
		)

		return nil, erros.Novo(ErroPegarAlunoTurma, nil, erro)
	}
	defer linhas.Close()

	for linhas.Next() {
		var turma entidades.TurmaAluno

		erro := linhas.Scan(&turma.IDAluno, &turma.IDTurma, &turma.Status)
		if erro != nil {
			bd.Log.Aviso(
				"Erro ao pegar as turmas do aluno com o seguinte ID: "+idAluno.String(),
				"\n\t"+erros.ErroExterno(erro),
			)

			return nil, erros.Novo(ErroPegarAlunoTurma, nil, erro)
		}

		turmas = append(turmas, turma)
	}
	if erro = linhas.Err(); erro != nil {
		bd.Log.Aviso(
			"Erro ao pegar as turmas do Curso com o seguinte ID: "+idAluno.String(),
			"\n\t"+erros.ErroExterno(erro),
		)

		return nil, erros.Novo(ErroPegarAlunoTurma, nil, erro)
	}

	return &turmas, nil
}

// Pegar é uma método que retorna uma entidade Aluno no banco de dados MariaDB.
func (bd AlunoBD) Pegar(id entidades.ID) (*entidades.Aluno, *erros.Aplicação) {
	bd.Log.Informação("Pegando Aluno com o seguinte ID: " + id.String())

	turmas, erroAplicação := bd.PegarTurmas(id)
	if erroAplicação != nil &&
		!erroAplicação.ÉPadrão(ErroAlunoTurmaNãoEncontrado) {
		bd.Log.Aviso(
			"Erro ao pegar as turmas do Aluno com o seguinte ID: "+id.String(),
			"\n\t"+erroAplicação.Error(),
		)
	}

	var aluno entidades.Aluno
	aluno.Turmas = *turmas

	query := "SELECT ID, ID_Pessoa, ID_Curso, Matrícula, Data_De_Ingresso," +
		" Data_De_Saída, Período, Status FROM " + bd.NomeDaTabela + " WHERE ID = ?"

	linha := bd.BD.QueryRow(query, id)

	erro := linha.Scan(
		&aluno.ID,
		&aluno.Pessoa,
		&aluno.Curso,
		&aluno.Matrícula,
		&aluno.DataDeIngresso,
		&aluno.DataDeSaída,
		&aluno.Período,
		&aluno.Status,
	)

	if erro != nil {
		if erro == sql.ErrNoRows {
			bd.Log.Aviso(
				"Não foi encontrada nenhum Aluno com o seguinte ID: "+id.String(),
				"\n\t"+erros.ErroExterno(erro),
			)
			return nil, erros.Novo(ErroAlunoNãoEncontrado, nil, erro)
		}

		bd.Log.Aviso(
			"Erro ao pegar o Aluno com o seguinte ID: "+id.String(),
			"\n\t"+erros.ErroExterno(erro),
		)
		return nil, erros.Novo(ErroPegarAluno, nil, erro)
	}

	return &aluno, nil
}

// Deletar é uma método que remove uma entidade Aluno no banco de dados MariaDB.
func (bd AlunoBD) Deletar(entidades.ID) *erros.Aplicação {
	bd.Log.Informação("Deletando Aluno")

	return nil
}
