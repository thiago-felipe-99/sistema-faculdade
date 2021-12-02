package mariadb

import (
	"database/sql"
	"io"

	// Driver para funcionar o mariadb.
	_ "github.com/go-sql-driver/mysql"
	"thiagofelipe.com.br/sistema-faculdade/data"
	"thiagofelipe.com.br/sistema-faculdade/entidades"
	"thiagofelipe.com.br/sistema-faculdade/erros"
	"thiagofelipe.com.br/sistema-faculdade/logs"
)

type id = data.ID

type Conexão struct {
	ID  id
	Log *logs.Log
	BD  *sql.DB
}

func NovoBD(dsn string) (*sql.DB, *erros.Aplicação) {
	bd, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, erros.Novo(erros.ConfigurarBD, nil, err)
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
