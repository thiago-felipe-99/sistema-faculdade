package mongodb

import (
	"context"
	"log"
	"os"
	"regexp"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"thiagofelipe.com.br/sistema-faculdade-backend/env"
	"thiagofelipe.com.br/sistema-faculdade-backend/logs"
)

//nolint: gochecknoglobals
var (
	matériaBD         *MatériaBD
	matériaBDInválido *MatériaBD
	ambiente          = env.PegandoVariáveisDeAmbiente()
)

const (
	tamanhoMáximoPréRequisito = 10
	tamanhoMáximoPalavra      = 20
	cargaHoráriaMáxima        = 10
)

func criarConexão(ctx context.Context) *mongo.Database {
	uri := "mongodb://root:root@localhost:" + ambiente.Portas.BDMateria

	db, erro := NovoDB(ctx, uri, "Teste")
	if erro != nil {
		log.Fatalf("Erro ao configurar o banco de dados: %s", erro)
	}

	err := db.Client().Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatalf("Erro ao conectar o banco de dados: %s", err)
	}

	return db
}

func criandoConexõesComAsColeções(bd *mongo.Database) {
	arquivos := logs.AbrirArquivos("./logs/")

	logMatéria := logs.NovoLog(arquivos.Matéria, logs.NívelDebug)

	conexãoMatéria := *NovaConexão(context.Background(), logMatéria, bd)

	matériaBD = &MatériaBD{
		Conexão:    conexãoMatéria,
		Collection: conexãoMatéria.BD.Collection("Matéria"),
	}

	matériaBDInválido = &MatériaBD{
		Conexão:    conexãoMatéria,
		Collection: conexãoMatéria.BD.Collection("MatériaInválida"),
	}

	matériaBDInválido.Timeout = time.Microsecond
}

func TestMain(m *testing.M) {
	bd := criarConexão(context.Background())

	criandoConexõesComAsColeções(bd)

	código := m.Run()

	// deletarColeções(bd)

	os.Exit(código)
}

func TestNovoDB(t *testing.T) {
	t.Parallel()

	t.Run("OKAY", func(t *testing.T) {
		t.Parallel()

		uri := "mongodb://root:root@localhost:" + ambiente.Portas.BDMateria

		bd, erro := NovoDB(context.Background(), uri, "Teste")
		if erro != nil {
			t.Fatalf("Erro ao configurar ao banco de dados: %v", erro)
		}

		err := bd.Client().Ping(context.Background(), readpref.Primary())
		if err != nil {
			t.Fatalf("Erro ao conectar ao banco de dados: %v", err)
		}
	})

	t.Run("EndereçoInválido", func(t *testing.T) {
		t.Parallel()

		padrão := regexp.MustCompile(`error parsing uri`)

		_, erro := NovoDB(context.Background(), "endereço inválido", "Teste")
		if erro == nil {
			t.Fatalf("Devia dar um erro na configuração")
		}

		if erro.ErroExterno == nil {
			t.Fatalf("Devia da um erro externo")
		}

		if !padrão.MatchString(erro.ErroExterno.Error()) {
			t.Fatalf(
				"Esperava por um erro de configuração no URI, chegou: %v",
				erro.ErroExterno.Error(),
			)
		}
	})
}