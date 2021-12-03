package mariadb

import (
	"database/sql"

	"thiagofelipe.com.br/sistema-faculdade/entidades"
	"thiagofelipe.com.br/sistema-faculdade/erros"
)

// AlunoBD representa a conexão com o banco de dados MariaDB para fazer alterações
// na entidade AlunoBD.
type AlunoBD struct {
	Conexão
	NomeDaTabela           string
	NomeDaTabelaSecundária string
}

// InserirTurmas é um método que faz a inserção das turmas de um Aluno no banco
// de dados MariaDB.
func (bd AlunoBD) InserirTurmas(turmas *[]entidades.TurmaAluno) *erros.Aplicação {
	bd.Log.Informação.Println("Inserindo turmas de um Aluno")

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
		bd.Log.Aviso.Println(
			"Erro ao inserir as turmas de um Aluno\n" + erros.ErroExterno(erroTx),
		)
		return erros.Novo(erros.InserirAlunoTurma, nil, erroTx)
	}

	defer tx.Rollback()

	_, erro := tx.Exec(query, params...)
	if erro != nil {
		bd.Log.Aviso.Println(
			"Erro ao inserir as turmas de um Aluno\n" + erros.ErroExterno(erro),
		)
		return erros.Novo(erros.InserirAlunoTurma, nil, erro)
	}

	if erroTx = tx.Commit(); erroTx != nil {
		bd.Log.Aviso.Println(
			"Erro ao inserir as turmas de um Aluno\nErro: " + erros.ErroExterno(erroTx),
		)
		return erros.Novo(erros.InserirAlunoTurma, nil, erroTx)
	}

	return nil
}

// Inserir é um método que faz inserção de uma Aluno no banco de dados MariaDB.
func (bd AlunoBD) Inserir(aluno *entidades.Aluno) *erros.Aplicação {
	bd.Log.Informação.Println("Inserindo Aluno com o seguinte ID: " + aluno.ID.String())

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
		bd.Log.Aviso.Println(
			"Erro ao inserir o Aluno com o seguinte ID: "+aluno.ID.String(),
			"\n"+erros.ErroExterno(erro),
		)
		return erros.Novo(erros.InserirAluno, nil, erro)
	}

	return bd.InserirTurmas(&aluno.Turmas)
}

// Atualizar é um método que faz a atualização de Aluno no banco de dados MariaDB.
func (bd AlunoBD) Atualizar(id, *entidades.Aluno) *erros.Aplicação {
	bd.Log.Informação.Println("Atualizando Aluno")

	return nil
}

// PegarTurmas é um método que pega as turmas de um Aluno no banco de dados
// MariaDB.
func (bd AlunoBD) PegarTurmas(idAluno id) (*[]entidades.TurmaAluno, *erros.Aplicação) {
	bd.Log.Informação.Println("Pegando as turmas do aluno com seguinte ID: " + idAluno.String())

	var turmas []entidades.TurmaAluno

	query := "SELECT ID_Aluno, ID_Turma, Status FROM " + bd.NomeDaTabelaSecundária +
		" WHERE ID_Aluno = ?"

	linhas, erro := bd.BD.Query(query, idAluno)
	if erro != nil {
		bd.Log.Aviso.Println(
			"Erro ao pegar as turmas do Aluno com o seguinte ID: "+idAluno.String(),
			"\n"+erros.ErroExterno(erro),
		)

		return nil, erros.Novo(erros.PegarAlunoTurma, nil, erro)
	}
	defer linhas.Close()

	for linhas.Next() {
		var turma entidades.TurmaAluno

		erro := linhas.Scan(&turma.IDAluno, &turma.IDTurma, &turma.Status)
		if erro != nil {
			bd.Log.Aviso.Println(
				"Erro ao pegar as turmas do aluno com o seguinte ID: "+idAluno.String(),
				"\n"+erros.ErroExterno(erro),
			)

			return nil, erros.Novo(erros.PegarAlunoTurma, nil, erro)
		}

		turmas = append(turmas, turma)
	}
	if erro = linhas.Err(); erro != nil {
		bd.Log.Aviso.Println(
			"Erro ao pegar as turmas do Curso com o seguinte ID: "+idAluno.String(),
			"\n"+erros.ErroExterno(erro),
		)

		return nil, erros.Novo(erros.PegarAlunoTurma, nil, erro)
	}

	return &turmas, nil
}

// Pegar é uma função que retorna uma Aluno do banco de dados MariaDB.
func (bd AlunoBD) Pegar(id id) (*entidades.Aluno, *erros.Aplicação) {
	bd.Log.Informação.Println("Pegando Aluno com o seguinte ID: " + id.String())

	turmas, erroAplicação := bd.PegarTurmas(id)
	if erroAplicação != nil &&
		!erroAplicação.ÉPadrão(erros.AlunoTurmaNãoEncontrado) {
		bd.Log.Aviso.Println(
			"Erro ao pegar as turmas do Aluno com o seguinte ID: "+id.String(),
			"\n"+erroAplicação.Error(),
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
			bd.Log.Aviso.Println(
				"Não foi encontrada nenhum Aluno com o seguinte ID: "+id.String(),
				"\n"+erros.ErroExterno(erro),
			)
			return nil, erros.Novo(erros.AlunoNãoEncontrado, nil, erro)
		}

		bd.Log.Aviso.Println(
			"Erro ao pegar o Aluno com o seguinte ID: "+id.String(),
			"\n"+erros.ErroExterno(erro),
		)
		return nil, erros.Novo(erros.PegarAluno, nil, erro)
	}

	return &aluno, nil
}

// Deletar é uma função que remove uma Aluno do banco de dados MariaDB.
func (bd AlunoBD) Deletar(id) *erros.Aplicação {
	bd.Log.Informação.Print("Deletando Aluno")

	return nil
}
