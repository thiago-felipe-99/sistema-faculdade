package mariadb

import (
	"database/sql"
	"log"
	"os"
	"regexp"
	"testing"

	"github.com/go-sql-driver/mysql"
	"thiagofelipe.com.br/sistema-faculdade-backend/env"
	"thiagofelipe.com.br/sistema-faculdade-backend/logs"
)

const (
	tamanhoMáximoPalavra = 25
)

//nolint: gochecknoglobals
var (
	pessoaBD         *PessoaBD
	pessoaBDInválido *PessoaBD
	ambiente         = env.PegandoVariáveisDeAmbiente()
)

func criarConexão() *sql.DB {
	//nolint: exhaustivestruct
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

	conexão, erro := NovoBD(config.FormatDSN())
	if erro != nil {
		log.Fatalf("Erro ao configurar o banco de dados: %s", erro)
	}

	if erroPing := conexão.Ping(); erroPing != nil {
		log.Fatalf("Erro ao conectar o banco de dados: %s", erroPing)
	}

	return conexão
}

func criandoConexõesComAsTabelas(banco *sql.DB) {
	arquivos := logs.AbrirArquivos("./logs/")

	logPessoa := logs.NovoLog(arquivos.Pessoa, logs.NívelDebug)

	pessoaBD = &PessoaBD{
		Conexão:      *NovaConexão(logPessoa, banco),
		NomeDaTabela: "Pessoa",
	}

	pessoaBDInválido = &PessoaBD{
		Conexão:      *NovaConexão(logPessoa, banco),
		NomeDaTabela: "PessoaErrada",
	}
}

func deletarTabelas(banco *sql.DB) {
	query := ""

	query += "DELETE FROM AlunoTurma;"
	query += "DELETE FROM Aluno;"
	query += "DELETE FROM CursoMaterias;"
	query += "DELETE FROM Curso;"
	query += "DELETE FROM Pessoa;"

	if _, erro := banco.Exec(query); erro != nil {
		log.Fatalf("Erro ao deletar os valores das tabelas: %s", erro.Error())
	}
}

func TestMain(m *testing.M) {
	banco := criarConexão()

	criandoConexõesComAsTabelas(banco)

	código := m.Run()

	deletarTabelas(banco)

	os.Exit(código)
}

func TestNovoBD(t *testing.T) {
	t.Parallel()

	t.Run("OKAY", func(t *testing.T) {
		t.Parallel()

		//nolint: exhaustivestruct
		config := mysql.Config{
			User:                 "Teste",
			Passwd:               "Teste",
			Net:                  "tcp",
			Addr:                 "localhost:" + ambiente.Portas.BDAdministrativo,
			DBName:               "Teste",
			AllowNativePasswords: true,
		}

		bd, erro := NovoBD(config.FormatDSN())
		if erro != nil {
			t.Fatalf("Erro ao configurar ao banco de dados: %v", erro)
		}

		err := bd.Ping()
		if err != nil {
			t.Fatalf("Erro ao conectar ao banco de dados: %v", err)
		}
	})

	t.Run("EndereçoInválido", func(t *testing.T) {
		t.Parallel()

		padrão := regexp.MustCompile(`invalid DSN`)

		_, erro := NovoBD("endereço inválido")
		if erro == nil {
			t.Fatalf("Devia dar um erro na configuração")
		}

		if erro.ErroExterno == nil {
			t.Fatalf("Devia da um erro externo")
		}

		if !padrão.MatchString(erro.ErroExterno.Error()) {
			t.Fatalf(
				"Esperava por um erro de configuração no DSN, chegou: %v",
				erro.ErroExterno.Error(),
			)
		}
	})
}
