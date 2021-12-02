package mongodb

import (
	"io"

	"thiagofelipe.com.br/sistema-faculdade/data"
	"thiagofelipe.com.br/sistema-faculdade/entidades"
	"thiagofelipe.com.br/sistema-faculdade/logs"
)

type id = data.ID

type Connexão struct {
	ID  id
	Log *logs.Log
}

func NovaConexão(arquivolog io.Writer) *Connexão {
	return &Connexão{
		ID:  entidades.NovoID(),
		Log: logs.NovoLog(arquivolog),
	}
}
