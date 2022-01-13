package mongodb

import (
	"thiagofelipe.com.br/sistema-faculdade-backend/entidades"
	"thiagofelipe.com.br/sistema-faculdade-backend/logs"
)

type Connexão struct {
	ID  entidades.ID
	Log *logs.Log
}

func NovaConexão(log *logs.Log) *Connexão {
	return &Connexão{
		ID:  entidades.NovoID(),
		Log: log,
	}
}
