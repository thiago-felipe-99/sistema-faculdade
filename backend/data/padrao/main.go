package padrao

import (
	"database/sql"

	"thiagofelipe.com.br/sistema-faculdade-backend/data"
	"thiagofelipe.com.br/sistema-faculdade-backend/data/mariadb"
	"thiagofelipe.com.br/sistema-faculdade-backend/data/mongodb"
	"thiagofelipe.com.br/sistema-faculdade-backend/logs"
)

// DataPadrão cria um Data que pode ser utilizado na aplicação.
func DataPadrão(log *logs.Entidades, bdSQL *sql.DB) *data.Data {
	MariaDBPessoa := mariadb.PessoaBD{
		Conexão:      *mariadb.NovaConexão(log.Pessoa, bdSQL),
		NomeDaTabela: "Pessoa",
	}

	MariaDBCurso := mariadb.CursoBD{
		Conexão:                *mariadb.NovaConexão(log.Curso, bdSQL),
		NomeDaTabela:           "Curso",
		NomeDaTabelaSecundária: "CursoMatérias",
	}

	MariaDBAluno := mariadb.AlunoBD{
		Conexão:                *mariadb.NovaConexão(log.Aluno, bdSQL),
		NomeDaTabela:           "Aluno",
		NomeDaTabelaSecundária: "AlunoTurma",
	}

	MariaDBProfessor := mariadb.ProfessorBD{
		Conexão: *mariadb.NovaConexão(log.Professor, bdSQL),
	}

	MariaDBAdministrativo := mariadb.AdministrativoBD{
		Conexão: *mariadb.NovaConexão(log.Administrativo, bdSQL),
	}

	MariaDBMatéria := mongodb.MatériaBD{
		Connexão: *mongodb.NovaConexão(log.Matéria),
	}

	MariaDBTurma := mongodb.TurmaBD{
		Connexão: *mongodb.NovaConexão(log.Turma),
	}

	return &data.Data{
		Pessoa:         MariaDBPessoa,
		Curso:          MariaDBCurso,
		Aluno:          MariaDBAluno,
		Professor:      MariaDBProfessor,
		Administrativo: MariaDBAdministrativo,
		Matéria:        MariaDBMatéria,
		Turma:          MariaDBTurma,
	}
}
