package main

import (
	"context"
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"thiagofelipe.com.br/sistema-faculdade-backend/data/mariadb"
	"thiagofelipe.com.br/sistema-faculdade-backend/data/mongodb"
	"thiagofelipe.com.br/sistema-faculdade-backend/data/padrao"
	"thiagofelipe.com.br/sistema-faculdade-backend/env"
	"thiagofelipe.com.br/sistema-faculdade-backend/http"
	"thiagofelipe.com.br/sistema-faculdade-backend/logica"
	"thiagofelipe.com.br/sistema-faculdade-backend/logs"
)

func criarConexãoBDs(ambiente env.VariáveisDeAmbiente) (*sql.DB, *mongo.Database) {
	//nolint:exhaustivestruct
	config := mysql.Config{
		User:                 "root",
		Passwd:               "root",
		Net:                  "tcp",
		Addr:                 "127.0.0.1:" + ambiente.Portas.BDAdministrativo,
		DBName:               "Administrativo",
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

	mongoConexão, erro := mongodb.NovoDB(context.Background(), uri, "Matéria")
	if erro != nil {
		log.Fatalf("Erro ao configurar o banco de dados MongoDB: %v", erro)
	}

	err = mongoConexão.Client().Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatalf("Erro ao conectar o banco de dados MongoDB: %v", err)
	}

	return sqlConexão, mongoConexão
}

func main() {
	ambiente := env.PegandoVariáveisDeAmbiente()
	sqlDB, mongoDB := criarConexãoBDs(ambiente)
	logFiles := logs.AbrirArquivos("./logs/data/")
	log := logs.NovoLogEntidades(logFiles, logs.NívelDebug)
	Data := padrao.DataPadrão(log, sqlDB, mongoDB)
	logica := logica.NovaLógica(Data)

	http.Rotas("127.0.0.1:8080", logica)
}
