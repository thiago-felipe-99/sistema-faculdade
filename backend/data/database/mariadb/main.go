package mariadb

import (
	"io"

	"github.com/google/uuid"
	"thiagofelipe.com.br/sistema-faculdade/data"
	"thiagofelipe.com.br/sistema-faculdade/logs"
)

type id = data.ID

type Connection struct {
	ID  id
	Log *logs.Log
}

func NewConnection(outlog io.Writer) *Connection {
	return &Connection{
		ID:  uuid.New(),
		Log: logs.NewLog(outlog),
	}
}
