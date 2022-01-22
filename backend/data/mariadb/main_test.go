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
	matériasMáximas        = 20
	tamanhoMáximoPalavra   = 25
	tamanhoMáximoMatrícula = 11
)

var (
	pessoaBD         *PessoaBD
	pessoaBDInválido *PessoaBD
	cursoBD          *CursoBD
	cursoBDInválido  *CursoBD
	cursoBDInválido2 *CursoBD
	alunoBD          *AlunoBD
	alunoBDInválido  *AlunoBD
	alunoBDInválido2 *AlunoBD
)

var ambiente = env.PegandoVariáveisDeAmbiente()

func criarConexão() *sql.DB {
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

	erroPing := conexão.Ping()
	if erroPing != nil {
		log.Fatalf("Erro ao conectar o banco de dados: %s", erroPing)
	}

	return conexão
}

func criandoConexõesComAsTabelas(bd *sql.DB) {
	arquivos := logs.AbrirArquivos("./logs/")

	logPessoa := logs.NovoLog(arquivos.Pessoa, logs.NívelDebug)

	logCurso := logs.NovoLog(arquivos.Curso, logs.NívelDebug)

	logAluno := logs.NovoLog(arquivos.Aluno, logs.NívelDebug)

	pessoaBD = &PessoaBD{
		Conexão:      *NovaConexão(logPessoa, bd),
		NomeDaTabela: "Pessoa",
	}

	pessoaBDInválido = &PessoaBD{
		Conexão:      *NovaConexão(logPessoa, bd),
		NomeDaTabela: "PessoaErrada",
	}

	cursoBD = &CursoBD{
		Conexão:                *NovaConexão(logCurso, bd),
		NomeDaTabela:           "Curso",
		NomeDaTabelaSecundária: "CursoMatérias",
	}

	cursoBDInválido = &CursoBD{
		Conexão:                *NovaConexão(logCurso, bd),
		NomeDaTabela:           "CursoErrado",
		NomeDaTabelaSecundária: "CursoMatériasErrado",
	}

	cursoBDInválido2 = &CursoBD{
		Conexão:                *NovaConexão(logCurso, bd),
		NomeDaTabela:           "CursoErrado",
		NomeDaTabelaSecundária: "CursoMatérias",
	}

	alunoBD = &AlunoBD{
		Conexão:                *NovaConexão(logAluno, bd),
		NomeDaTabela:           "Aluno",
		NomeDaTabelaSecundária: "AlunoTurma",
	}

	alunoBDInválido = &AlunoBD{
		Conexão:                *NovaConexão(logAluno, bd),
		NomeDaTabela:           "AlunoErrado",
		NomeDaTabelaSecundária: "AlunoTurmaErrado",
	}

	alunoBDInválido2 = &AlunoBD{
		Conexão:                *NovaConexão(logAluno, bd),
		NomeDaTabela:           "AlunoErrado",
		NomeDaTabelaSecundária: "AlunoTurma",
	}
}

func deletarTabelas(bd *sql.DB) {
	query := ""

	query += "DELETE FROM AlunoTurma;"
	query += "DELETE FROM Aluno;"
	query += "DELETE FROM CursoMatérias;"
	query += "DELETE FROM Curso;"
	query += "DELETE FROM Pessoa;"

	_, erro := bd.Exec(query)
	if erro != nil {
		log.Fatalf("Erro ao deletar os valores das tabelas: %s", erro.Error())
	}
}

func TestMain(m *testing.M) {
	bd := criarConexão()

	criandoConexõesComAsTabelas(bd)

	código := m.Run()

	deletarTabelas(bd)

	os.Exit(código)
}

func TestNovoBD(t *testing.T) {
	t.Run("OKAY", func(t *testing.T) {
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
		padrão := regexp.MustCompile(`invalid DSN`)

		_, erro := NovoBD("endereço inválido")
		if erro == nil {
			t.Fatalf("Devia dar um erro na configuração")
		}

		if erro.ErroExterno == nil {
			t.Fatalf("Devia da um erro externo")
		}

		if !padrão.MatchString(erro.ErroExterno.Error()) {
			t.Fatalf("Esperava por um erro de configuração no DSN, chegou: %v", erro.ErroExterno.Error())
		}
	})
}
