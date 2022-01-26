package padrao

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"testing"

	"github.com/go-sql-driver/mysql"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"thiagofelipe.com.br/sistema-faculdade-backend/data/mariadb"
	"thiagofelipe.com.br/sistema-faculdade-backend/data/mongodb"
	"thiagofelipe.com.br/sistema-faculdade-backend/env"
	"thiagofelipe.com.br/sistema-faculdade-backend/logs"
)

//nolint: gochecknoglobals
var ambiente = env.PegandoVariáveisDeAmbiente()

func criarConexãoDBs() (*sql.DB, *mongo.Database) {
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

func TestDataPadrão(t *testing.T) {
	t.Parallel()

	bdSQL, bdMongo := criarConexãoDBs()
	logFiles := logs.AbrirArquivos("./logs/")
	log := logs.NovoLogEntidades(logFiles, logs.NívelDebug)

	data := DataPadrão(log, bdSQL, bdMongo)

	tipos := map[string]struct{ quer, recebou string }{
		"Pessoa":         {"mariadb.PessoaBD", fmt.Sprintf("%T", data.Pessoa)},
		"Curso":          {"<nil>", fmt.Sprintf("%T", data.Curso)},
		"Aluno":          {"<nil>", fmt.Sprintf("%T", data.Aluno)},
		"Professor":      {"<nil>", fmt.Sprintf("%T", data.Professor)},
		"Administrativo": {"<nil>", fmt.Sprintf("%T", data.Administrativo)},
		"Matéria":        {"mongodb.MatériaBD", fmt.Sprintf("%T", data.Matéria)},
		"Turma":          {"mongodb.TurmaBD", fmt.Sprintf("%T", data.Turma)},
	}

	for chave, valor := range tipos {
		if valor.quer != valor.recebou {
			t.Fatalf(
				"A data da entidade %s quer: %s, porém recebeu: %s",
				chave, valor.quer, valor.recebou,
			)
		}
	}
}
