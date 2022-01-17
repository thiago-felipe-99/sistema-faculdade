package logica

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/go-sql-driver/mysql"
	"thiagofelipe.com.br/sistema-faculdade-backend/data"
	"thiagofelipe.com.br/sistema-faculdade-backend/data/mariadb"
	"thiagofelipe.com.br/sistema-faculdade-backend/env"
	"thiagofelipe.com.br/sistema-faculdade-backend/logs"
)

const (
	tamanhoMáximoDaPalavra = 25
)

var (
	logicaTeste         *Lógica
	pessoaBDInválido    *Pessoa
	pessoaDataInvalida  *Pessoa
	pessoaDataInvalida2 *Pessoa
	ambiente            = env.PegandoVariáveisDeAmbiente()
)

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
	bd := criarConexãoMariaDB()
	logFiles := logs.AbrirArquivos("./logs/data/")
	log := logs.NovoLogEntidades(logFiles, logs.NívelDebug)

	Data := data.DataPadrão(log, bd)

	logicaTeste = NovaLógica(Data)

	dataPessoaInválido := &mariadb.PessoaBD{
		Conexão:      *mariadb.NovaConexão(log.Pessoa, bd),
		NomeDaTabela: "PessoaInválida",
	}

	pessoaBDInválido = &Pessoa{data: dataPessoaInválido}

	pessoaDataInvalida = &Pessoa{
		&dataPessoaInvalida{logicaTeste.Pessoa.data},
	}

	pessoaDataInvalida2 = &Pessoa{
		&dataPessoaInvalida2{logicaTeste.Pessoa.data},
	}

	código := m.Run()

	os.Exit(código)
}
