package logica_test

import (
	"database/sql"
	"log"
	"math/rand"
	"os"
	"testing"
	"time"

	"github.com/go-sql-driver/mysql"
	"thiagofelipe.com.br/sistema-faculdade-backend/data"
	"thiagofelipe.com.br/sistema-faculdade-backend/data/database/mariadb"
	"thiagofelipe.com.br/sistema-faculdade-backend/env"
	. "thiagofelipe.com.br/sistema-faculdade-backend/logica"
	"thiagofelipe.com.br/sistema-faculdade-backend/logs"
)

const (
	//nolint:deadcode,varcheck,unused
	tamanhoMáximoDaPalavra = 25
)

//nolint:gochecknoglobals
var logica *Lógica

//nolint:gochecknoglobals
var ambiente = env.PegandoVariáveisDeAmbiente()

func criarConexãoMariaDB() *sql.DB {
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

	conexão, erro := mariadb.NovoBD(config.FormatDSN())
	if erro != nil {
		log.Fatalf("Erro ao configurar o banco de dados: %s", erro)
	}

	erroPing := conexão.Ping()
	if erroPing != nil {
		log.Fatalf("Erro ao conectar o banco de dados: %s", erroPing)
	}

	return conexão
}

func TestMain(m *testing.M) {
	rand.Seed(time.Now().UnixNano())

	bd := criarConexãoMariaDB()
	logData := logs.AbrirArquivos("./logs/data/")

	Data := data.DataPadrão(logs.NovoLogEntidades(logData, logs.NívelDebug), bd)

	logica = NovaLógica(Data)

	código := m.Run()

	os.Exit(código)
}