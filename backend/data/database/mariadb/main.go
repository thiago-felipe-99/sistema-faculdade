package mariadb

import (
	"database/sql"
	"io"

	// Driver para funcionar o mariadb.
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"thiagofelipe.com.br/sistema-faculdade/data"
	"thiagofelipe.com.br/sistema-faculdade/errors"
	"thiagofelipe.com.br/sistema-faculdade/logs"
)

type id = data.ID

type Connection struct {
	ID  id
	Log *logs.Log
	DB  *sql.DB
}

func NewDB(dsn string) (*sql.DB, *errors.Application) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, errors.New(errors.ConfigurarBD, nil, err)
	}

	return db, nil
}

func NewConnection(outlog io.Writer, db *sql.DB) *Connection {
	return &Connection{
		ID:  uuid.New(),
		Log: logs.NewLog(outlog),
		DB:  db,
	}
}
