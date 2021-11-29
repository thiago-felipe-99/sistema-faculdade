package mongodb

import (
	"io"

	"github.com/google/uuid"
	"thiagofelipe.com.br/sistema-faculdade/data"
	"thiagofelipe.com.br/sistema-faculdade/logs"
)

type id = data.ID

type Connexão struct {
	ID  id
	Log *logs.Log
}

func NovaConexão(arquivolog io.Writer) *Connexão {
	return &Connexão{
		ID:  uuid.New(),
		Log: logs.NovoLog(arquivolog),
	}
}
