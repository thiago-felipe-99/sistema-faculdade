package mariadb

import (
	"database/sql"
	"log"
	"math/rand"
	"os"
	"regexp"
	"testing"
	"time"

	"github.com/go-sql-driver/mysql"
	"thiagofelipe.com.br/sistema-faculdade-backend/env"
	"thiagofelipe.com.br/sistema-faculdade-backend/logs"
)

const (
	MATÉRIAS_MÁXIMAS         = 20
	TAMANHO_MÁXIMO_PALAVRA   = 25
	TAMANHO_MÁXIMO_MATRÍCULA = 11
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

func criarConexão(m *testing.M) *sql.DB {

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

func criandoConexõesComAsTabelas(m *testing.M, bd *sql.DB) {

	logs := logs.AbrirArquivos("./logs/")

	pessoaBD = &PessoaBD{
		Conexão:      *NovaConexão(logs.Pessoa, bd),
		NomeDaTabela: "Pessoa",
	}

	pessoaBDInválido = &PessoaBD{
		Conexão:      *NovaConexão(logs.Pessoa, bd),
		NomeDaTabela: "PessoaErrada",
	}

	cursoBD = &CursoBD{
		Conexão:                *NovaConexão(logs.Curso, bd),
		NomeDaTabela:           "Curso",
		NomeDaTabelaSecundária: "CursoMatérias",
	}

	cursoBDInválido = &CursoBD{
		Conexão:                *NovaConexão(logs.Curso, bd),
		NomeDaTabela:           "CursoErrado",
		NomeDaTabelaSecundária: "CursoMatériasErrado",
	}

	cursoBDInválido2 = &CursoBD{
		Conexão:                *NovaConexão(logs.Curso, bd),
		NomeDaTabela:           "CursoErrado",
		NomeDaTabelaSecundária: "CursoMatérias",
	}

	alunoBD = &AlunoBD{
		Conexão:                *NovaConexão(logs.Aluno, bd),
		NomeDaTabela:           "Aluno",
		NomeDaTabelaSecundária: "AlunoTurma",
	}

	alunoBDInválido = &AlunoBD{
		Conexão:                *NovaConexão(logs.Aluno, bd),
		NomeDaTabela:           "AlunoErrado",
		NomeDaTabelaSecundária: "AlunoTurmaErrado",
	}

	alunoBDInválido2 = &AlunoBD{
		Conexão:                *NovaConexão(logs.Aluno, bd),
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

	rand.Seed(time.Now().UnixNano())

	bd := criarConexão(m)

	criandoConexõesComAsTabelas(m, bd)

	código := m.Run()

	deletarTabelas(bd)

	os.Exit(código)
}

// TestNovoBD verifica se a inicialização do banco de dados está okay.
//nolint: paralleltest
func TestNovoBD(t *testing.T) {
	var ambiente = env.PegandoVariáveisDeAmbiente()

	//nolint: exhaustivestruct
	config := mysql.Config{
		User:                 "Teste",
		Passwd:               "Teste",
		Net:                  "tcp",
		Addr:                 "localhost:" + ambiente.Portas.BDAdministrativo,
		DBName:               "Teste",
		AllowNativePasswords: true,
	}

	bd, erroAplicação := NovoBD(config.FormatDSN())
	if erroAplicação != nil {
		t.Fatalf("Erro ao configurar ao banco de dados: %v", erroAplicação)
	}

	erro := bd.Ping()
	if erro != nil {
		t.Fatalf("Erro ao conectar ao banco de dados: %v", erro)
	}
}

func TestNovoBD_EndereçoErrado(t *testing.T) {
	padrão, erroRegex := regexp.Compile(`invalid DSN`)
	if erroRegex != nil {
		t.Fatal("Erro ao compilar o regex")
	}

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
}
