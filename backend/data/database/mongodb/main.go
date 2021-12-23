package mongodb

import (
	"io"

	"thiagofelipe.com.br/sistema-faculdade-backend/entidades"
	"thiagofelipe.com.br/sistema-faculdade-backend/logs"
)

type Connexão struct {
	ID  entidades.ID
	Log *logs.Log
}

func NovaConexão(arquivolog io.Writer) *Connexão {
	return &Connexão{
		ID:  entidades.NovoID(),
		Log: logs.NovoLog(arquivolog),
	}
}
