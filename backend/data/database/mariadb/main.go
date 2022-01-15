package mariadb

import (
	"database/sql"

	// Driver para funcionar o mariadb.
	_ "github.com/go-sql-driver/mysql"

	//nolint:revive,stylecheck
	. "thiagofelipe.com.br/sistema-faculdade-backend/data/erros"
	"thiagofelipe.com.br/sistema-faculdade-backend/entidades"
	"thiagofelipe.com.br/sistema-faculdade-backend/erros"
	"thiagofelipe.com.br/sistema-faculdade-backend/logs"
)

// Conexão representa a conexão com o banco de dados MariaDB.
type Conexão struct {
	ID  entidades.ID
	Log *logs.Log
	BD  *sql.DB
}

// NovoBD cria um link com o banco de dados MariaDB.
func NovoBD(dsn string) (*sql.DB, *erros.Aplicação) {
	bd, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, erros.Novo(ErroConfigurarBD, nil, err)
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
