package mariadb

import (
	"regexp"
	"testing"

	"github.com/go-sql-driver/mysql"
	"thiagofelipe.com.br/sistema-faculdade/env"
)

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
