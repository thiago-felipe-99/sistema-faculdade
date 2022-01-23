package mongodb

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"thiagofelipe.com.br/sistema-faculdade-backend/data"
	"thiagofelipe.com.br/sistema-faculdade-backend/entidades"
	"thiagofelipe.com.br/sistema-faculdade-backend/erros"
	"thiagofelipe.com.br/sistema-faculdade-backend/logs"
)

type erro = *erros.Aplicação

// Conexão representa a conexão com o banco de dados MongoDB.
type Conexão struct {
	ID      entidades.ID
	Log     *logs.Log
	BD      *mongo.Database
	Timeout time.Duration
	ctx     context.Context
}

// NovoDB cria um link com o banco de dados MongoDB.
func NovoDB(ctx context.Context, uri string, nomeDB string) (
	*mongo.Database,
	*erros.Aplicação,
) {
	bd, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, erros.Novo(data.ErroConfigurarBD, nil, err)
	}

	return bd.Database(nomeDB), nil
}

// NovaConexão cria uma conexão com o banco de dados MongoDB.
func NovaConexão(ctx context.Context, log *logs.Log, bd *mongo.Database) *Conexão {
	const quantidade = 1

	return &Conexão{
		ID:      entidades.NovoID(),
		Log:     log,
		BD:      bd,
		Timeout: time.Second * quantidade,
		ctx:     ctx,
	}
}
