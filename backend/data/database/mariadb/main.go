package mariadb

import (
	"database/sql"
	"io"

	// Driver para funcionar o mariadb.
	_ "github.com/go-sql-driver/mysql"

	//nolint:revive,stylecheck
	. "thiagofelipe.com.br/sistema-faculdade-backend/data/erros"
	"thiagofelipe.com.br/sistema-faculdade-backend/entidades"
	"thiagofelipe.com.br/sistema-faculdade-backend/erros"
	"thiagofelipe.com.br/sistema-faculdade-backend/logs"
)

type Conexão struct {
	ID  entidades.ID
	Log *logs.Log
	BD  *sql.DB
}

func NovoBD(dsn string) (*sql.DB, *erros.Aplicação) {
	bd, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, erros.Novo(ErroConfigurarBD, nil, err)
	}

	return bd, nil
}

func NovaConexão(arquivolog io.Writer, bd *sql.DB) *Conexão {
	return &Conexão{
		ID:  entidades.NovoID(),
		Log: logs.NovoLog(arquivolog),
		BD:  bd,
	}
}
