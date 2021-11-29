package mariadb_test

import (
	"database/sql"
	"testing"

	"github.com/go-sql-driver/mysql"
	"thiagofelipe.com.br/sistema-faculdade/env"
)

// TestNovoBD verifica se a inicialização do banco de dados está okay.
func TestNovoBD(t *testing.T) {
	ambiente := env.PegandoVariáveisDeAmbiente()

	var config = mysql.Config{
		User:                 "Teste",
		Passwd:               "Teste",
		Net:                  "tcp",
		Addr:                 "localhost:" + ambiente.Portas.BDAdministrativo,
		DBName:               "Teste",
		AllowNativePasswords: true,
	}

	bd, err := sql.Open("mysql", config.FormatDSN())
	if err != nil {
		t.Fatalf("Erro ao configurar ao banco de dados: %v", err)
	}

	err = bd.Ping()
	if err != nil {
		t.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}
}
