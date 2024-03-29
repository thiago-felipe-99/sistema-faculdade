package logica

import (
	"context"
	"database/sql"
	"log"
	"os"
	"testing"
	"time"

	"github.com/go-sql-driver/mysql"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"thiagofelipe.com.br/sistema-faculdade-backend/data/mariadb"
	"thiagofelipe.com.br/sistema-faculdade-backend/data/mongodb"
	"thiagofelipe.com.br/sistema-faculdade-backend/data/padrao"
	"thiagofelipe.com.br/sistema-faculdade-backend/env"
	"thiagofelipe.com.br/sistema-faculdade-backend/logs"
)

const (
	tamanhoMáximoPalavra      = 25
	tamanhoMáximoCargaHorária = 10
	tamanhoMáximoCréditos     = 20
	tamanhoMáximoData         = 15
	tamanhoMáximoMatéria      = 10
)

//nolint: gochecknoglobals
var (
	logicaTeste         *Lógica
	pessoaBDInválido    *Pessoa
	pessoaDataInvalida  *Pessoa
	pessoaDataInvalida2 *Pessoa
	matériaBDTimeOut    *Matéria
	matériaBDInválido   *Matéria
	matériaBDInválido2  *Matéria
	cursoBDTimeOut      *Curso
	cursoBDInválido     *Curso
	ambiente            = env.PegandoVariáveisDeAmbiente()
)

func criarConexãoBDs() (*sql.DB, *mongo.Database) {
	//nolint:exhaustivestruct
	config := mysql.Config{
		User:                 "Teste",
		Passwd:               "Teste",
		Net:                  "tcp",
		Addr:                 "127.0.0.1:" + ambiente.Portas.BDAdministrativo,
		DBName:               "Teste",
		AllowNativePasswords: true,
		ParseTime:            true,
		MultiStatements:      true,
	}

	sqlConexão, erro := mariadb.NovoBD(config.FormatDSN())
	if erro != nil {
		log.Fatalf("Erro ao configurar o banco de dados MariaDB: %v", erro)
	}

	err := sqlConexão.Ping()
	if err != nil {
		log.Fatalf("Erro ao conectar o banco de dados MariaDB: %v", err)
	}

	uri := "mongodb://root:root@localhost:" + ambiente.Portas.BDMateria

	mongoConexão, erro := mongodb.NovoDB(context.Background(), uri, "Teste")
	if erro != nil {
		log.Fatalf("Erro ao configurar o banco de dados MongoDB: %v", erro)
	}

	err = mongoConexão.Client().Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatalf("Erro ao conectar o banco de dados MongoDB: %v", err)
	}

	return sqlConexão, mongoConexão
}

//nolint: funlen
func TestMain(m *testing.M) {
	sqlDB, mongoDB := criarConexãoBDs()
	logFiles := logs.AbrirArquivos("./logs/data/")
	log := logs.NovoLogEntidades(logFiles, logs.NívelDebug)

	Data := padrao.DataPadrão(log, sqlDB, mongoDB)

	logicaTeste = NovaLógica(Data)

	dataPessoaInválido := &mariadb.PessoaBD{
		Conexão:      *mariadb.NovaConexão(log.Pessoa, sqlDB),
		NomeDaTabela: "PessoaInválida",
	}
	pessoaBDInválido = &Pessoa{Data: dataPessoaInválido}

	pessoaDataInvalida = &Pessoa{
		&dataPessoaInvalida{logicaTeste.Pessoa.Data},
	}

	pessoaDataInvalida2 = &Pessoa{
		&dataPessoaInvalida2{logicaTeste.Pessoa.Data},
	}

	conexãoMatériaTimeOut := *mongodb.NovaConexão(
		context.Background(),
		log.Matéria,
		mongoDB,
	)
	conexãoMatériaTimeOut.Timeout = time.Nanosecond
	dataMatériaTimeOut := &mongodb.MatériaBD{
		Conexão:    conexãoMatériaTimeOut,
		Collection: conexãoMatériaTimeOut.BD.Collection("Matéria"),
	}
	matériaBDTimeOut = &Matéria{
		data: dataMatériaTimeOut,
	}

	matériaBDInválido = &Matéria{
		&dataMatériaInvalida{logicaTeste.Matéria.data},
	}

	matériaBDInválido2 = &Matéria{
		&dataMatériaInvalida2{logicaTeste.Matéria.data},
	}

	conexãoCursoTimeOut := *mongodb.NovaConexão(
		context.Background(),
		log.Matéria,
		mongoDB,
	)
	conexãoCursoTimeOut.Timeout = time.Nanosecond
	dataCursoTimeOut := &mongodb.CursoBD{
		Conexão:    conexãoMatériaTimeOut,
		Collection: conexãoMatériaTimeOut.BD.Collection("Curso"),
	}
	cursoBDTimeOut = &Curso{
		data:    dataCursoTimeOut,
		matéria: matériaBDTimeOut,
	}

	cursoBDInválido = &Curso{
		data:    &dataCursoInvalido{logicaTeste.Curso.data},
		matéria: matériaBDInválido2,
	}

	código := m.Run()

	os.Exit(código)
}
