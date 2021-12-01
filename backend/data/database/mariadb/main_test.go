package mariadb

import (
	"log"
	"math/rand"
	"os"
	"regexp"
	"testing"
	"time"

	"github.com/go-sql-driver/mysql"
	"thiagofelipe.com.br/sistema-faculdade/env"
)

var (
	pessoaBD         *PessoaBD
	pessoaBDInválido *PessoaBD
	cursoBD          *CursoBD
	cursoBDInválido  *CursoBD
)

var ambiente = env.PegandoVariáveisDeAmbiente()

//nolint:funlen
func TestMain(m *testing.M) {

	rand.Seed(time.Now().UnixNano())

	config := mysql.Config{
		User:                 "Teste",
		Passwd:               "Teste",
		Net:                  "tcp",
		Addr:                 "127.0.0.1:" + ambiente.Portas.BDAdministrativo,
		DBName:               "Teste",
		AllowNativePasswords: true,
		ParseTime:            true,
	}

	connexão, erro := NovoBD(config.FormatDSN())
	if erro != nil {
		log.Fatalf("Erro ao configurar o banco de dados: %s", erro)
	}

	erroPing := connexão.Ping()
	if erroPing != nil {
		log.Fatalf("Erro ao conectar o banco de dados: %s", erroPing)
	}

	pessoaBD = &PessoaBD{
		Conexão:      *NovaConexão(os.Stderr, connexão),
		NomeDaTabela: "Pessoa",
	}

	pessoaBDInválido = &PessoaBD{
		Conexão:      *NovaConexão(os.Stderr, connexão),
		NomeDaTabela: "PessoaErrada",
	}

	cursoBD = &CursoBD{
		Conexão:                *NovaConexão(os.Stderr, connexão),
		NomeDaTabela:           "Curso",
		NomeDaTabelaSecundária: "CursoMatérias",
	}

	cursoBDInválido = &CursoBD{
		Conexão:                *NovaConexão(os.Stderr, connexão),
		NomeDaTabela:           "CursoMatérias",
		NomeDaTabelaSecundária: "CursoMatériasErrado",
	}

	código := m.Run()

	query := "DELETE FROM Curso"
	connexão.Exec(query)

	query = "DELETE FROM Pessoa"
	connexão.Exec(query)

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
