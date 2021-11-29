package main

import (
	"encoding/json"
	"io"
	"log"
	"math/rand"
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
		User:                 "Administrativo",
		Passwd:               "Administrativo",
		Net:                  "tcp",
		Addr:                 "localhost:9000",
		DBName:               "Administrativo",
		AllowNativePasswords: true,
		ParseTime:            true,
	}

	dns := config.FormatDSN()
	log.Println(dns)

	db, err := mariadb.NovoBD(config.FormatDSN())
	if err != nil {
		log.Panicln(err.Message)
	}

	MariaDBCurso := mariadb.CursoBD{
		Conexão:                *mariadb.NovaConexão(logFiles.pessoa, db),
		NomeDaTabela:           "Curso",
		NomeDaTabelaSecundária: "CursoMatérias",
	}

	return &data.Data{
		Curso: MariaDBCurso,
	}
}

func prettyStruct(s ...interface{}) string {
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
		rand.Seed(time.Now().UnixNano())
		id := uuid.New()

		materias := &[]data.CursoMatéria{
			{
				ID_Curso:   id,
				ID_Matéria: uuid.New(),
				Período:    "Teste",
				Tipo:       "Não sei",
				Status:     "Testando",
				Observação: "Okay",
			},
			{
				ID_Curso:   id,
				ID_Matéria: uuid.New(),
				Período:    "Teste",
				Tipo:       "Não sei",
				Status:     "Testando",
				Observação: "Okay",
			},
		}

		curso := &data.Curso{
			ID:                id,
			Nome:              "Curso novo",
			DataDeInício:      time.Now(),
			DataDeDesativação: time.Now(),
			Matérias:          *materias,
		}

		err := Data.Curso.Inserir(curso)
		if err != nil {
			log.Panicln(err.Error())
			if err.SystemError != nil {
				log.Println(err.SystemError.Error())
			}
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
