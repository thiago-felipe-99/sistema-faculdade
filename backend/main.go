package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"thiagofelipe.com.br/sistema-faculdade/data"
	"thiagofelipe.com.br/sistema-faculdade/data/database/mariadb"
)

type logFiles struct {
	pessoa         io.Writer
	curso          io.Writer
	aluno          io.Writer
	professor      io.Writer
	administrativo io.Writer
	matéria        io.Writer
	turma          io.Writer
}

func openLogFiles() *logFiles {
	defaultDir := "./logs/data/"

	pessoa, err := os.OpenFile(
		defaultDir+"pessoaLogs.txt",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0666)
	if err != nil {
		log.Fatal(err)
	}

	curso, err := os.OpenFile(
		defaultDir+"cursoLogs.txt",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0666)
	if err != nil {
		log.Fatal(err)
	}

	aluno, err := os.OpenFile(
		defaultDir+"alunoLogs.txt",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0666)
	if err != nil {
		log.Fatal(err)
	}

	professor, err := os.OpenFile(
		defaultDir+"professorLogs.txt",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0666)
	if err != nil {
		log.Fatal(err)
	}

	administrativo, err := os.OpenFile(
		defaultDir+"administrativoLogs.txt",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0666)
	if err != nil {
		log.Fatal(err)
	}

	matéria, err := os.OpenFile(
		defaultDir+"matériaLogs.txt",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0666)
	if err != nil {
		log.Fatal(err)
	}

	turma, err := os.OpenFile(
		defaultDir+"turmaLogs.txt",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0666)
	if err != nil {
		log.Fatal(err)
	}

	return &logFiles{
		pessoa:         pessoa,
		curso:          curso,
		aluno:          aluno,
		professor:      professor,
		administrativo: administrativo,
		matéria:        matéria,
		turma:          turma,
	}
}

func newData() *data.Data {
	logFiles := openLogFiles()

	config := mysql.Config{
		User:                 "root",
		Passwd:               "root",
		Net:                  "tcp",
		Addr:                 "localhost:9000",
		DBName:               "Administrativo",
		AllowNativePasswords: true,
		ParseTime:            true,
	}

	dns := config.FormatDSN()
	log.Println(dns)

	db, err := mariadb.NewDB(config.FormatDSN())
	if err != nil {
		log.Panicln(err.Message)
	}

	MariaDBPessoa := mariadb.PessoaDB{
		Connection: *mariadb.NewConnection(logFiles.pessoa, db),
		TableName:  "Pessoa",
	}

	return &data.Data{
		Pessoa: MariaDBPessoa,
	}
}

func prettyStruc(s ...interface{}) string {
	sJSON, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		log.Panicln(err.Error())
	}

	return string(sJSON)
}

func main() {
	r := gin.Default()

	Data := newData()

	r.GET("/ping", func(c *gin.Context) {
		novaPessoa := data.Pessoa{
			ID:               uuid.New(),
			Nome:             "Thiago Felipe",
			CPF:              "12345678910",
			DataDeNascimento: time.Date(1999, 12, 8, 1, 1, 1, 1, time.FixedZone("UTC-4", -4*60*60)),
			Senha:            "Senha",
		}

		err := Data.Pessoa.Insert(&novaPessoa)
		if err != nil {
			log.Panicln(err.Message)
		}

		resultGet, err := Data.Pessoa.Get(novaPessoa.ID)
		if err != nil {
			log.Panicln(err.Message)
		}

		log.Println(prettyStruc(resultGet))

		novaPessoa.Senha = "Passwd"
		novaPessoa.Nome = "Thiago Felipe Cruz E Souza"

		err = Data.Pessoa.Update(novaPessoa.ID, &novaPessoa)
		if err != nil {
			log.Panicln(err.Message)
		}

		resultGet, err = Data.Pessoa.Get(novaPessoa.ID)
		if err != nil {
			log.Panicln(err.Message)
		}

		log.Println(prettyStruc(resultGet))

		err = Data.Pessoa.Delete(novaPessoa.ID)
		if err != nil {
			log.Panicln(err.Message)
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/pong", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "ping",
		})
	})

	err := r.Run()
	if err != nil {
		panic(err)
	}
}
