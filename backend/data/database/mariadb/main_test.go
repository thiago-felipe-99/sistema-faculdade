package mariadb_test

import (
	"database/sql"
	"testing"

	"github.com/go-sql-driver/mysql"
)

var config = mysql.Config{
	User:                 "Teste",
	Passwd:               "Teste",
	Net:                  "tcp",
	Addr:                 "localhost:9000",
	DBName:               "Teste",
	AllowNativePasswords: true,
}

// TestNewDB verifica se a inicialização do banco de dados está okay.
func TestNewDB(t *testing.T) {
	db, err := sql.Open("mysql", config.FormatDSN())
	if err != nil {
		t.Fatalf("Erro ao configurar ao banco de dados: %v", err)
	}

	err = db.Ping()
	if err != nil {
		t.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}
}
