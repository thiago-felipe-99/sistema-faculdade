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
	"thiagofelipe.com.br/sistema-faculdade/env"
	"thiagofelipe.com.br/sistema-faculdade/logs"
)

var (
	pessoaBD         *PessoaBD
	pessoaBDInválido *PessoaBD
	cursoBD          *CursoBD
	cursoBDInválido  *CursoBD
	cursoBDInválido2 *CursoBD
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

func deletarTabelas(bd *sql.DB) {
	query := "DELETE FROM CursoMatérias"
	bd.Exec(query)

	query = "DELETE FROM Curso"
	bd.Exec(query)

	query = "DELETE FROM Pessoa"
	bd.Exec(query)
}

func TestMain(m *testing.M) {

	rand.Seed(time.Now().UnixNano())

	bd := criarConexão(m)

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
