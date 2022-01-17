package mongodb

import (
	"thiagofelipe.com.br/sistema-faculdade-backend/entidades"
	"thiagofelipe.com.br/sistema-faculdade-backend/logs"
)

// Conexão representa a conexão com o banco de dados MongoDB.
type Connexão struct {
	ID  entidades.ID
	Log *logs.Log
}

// NovaConexão cria uma conexão com o banco de dados MongoDB.
func NovaConexão(log *logs.Log) *Connexão {
	return &Connexão{
		ID:  entidades.NovoID(),
		Log: log,
	}
}
