package mariadb

import (
	"database/sql"

	// Driver para funcionar o mariadb.
	_ "github.com/go-sql-driver/mysql"
	"thiagofelipe.com.br/sistema-faculdade-backend/data"
	"thiagofelipe.com.br/sistema-faculdade-backend/entidades"
	"thiagofelipe.com.br/sistema-faculdade-backend/erros"
	"thiagofelipe.com.br/sistema-faculdade-backend/logs"
)

type (
	erro           = *erros.Aplicação
	id             = entidades.ID
	cpf            = entidades.CPF
	pessoa         = entidades.Pessoa
	curso          = entidades.Curso
	cursoMatéria   = entidades.CursoMatéria
	aluno          = entidades.Aluno
	turmaAluno     = entidades.TurmaAluno
	professor      = entidades.Professor
	administrativo = entidades.Administrativo
)

// Conexão representa a conexão com o banco de dados MariaDB.
type Conexão struct {
	ID  id
	Log *logs.Log
	BD  *sql.DB
}

// NovoBD cria um link com o banco de dados MariaDB.
func NovoBD(dsn string) (*sql.DB, erro) {
	bd, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, erros.Novo(data.ErroConfigurarBD, nil, err)
	}

	return bd, nil
}

// NovaConexão cria uma conexão com o banco de dados MariaDB.
func NovaConexão(log *logs.Log, bd *sql.DB) *Conexão {
	return &Conexão{
		ID:  entidades.NovoID(),
		Log: log,
		BD:  bd,
	}
}
