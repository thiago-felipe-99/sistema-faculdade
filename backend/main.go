package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"thiagofelipe.com.br/sistema-faculdade/data"
	"thiagofelipe.com.br/sistema-faculdade/data/database/mariadb"
	"thiagofelipe.com.br/sistema-faculdade/data/database/mongodb"
	"thiagofelipe.com.br/sistema-faculdade/entidades"
	"thiagofelipe.com.br/sistema-faculdade/logs"
)

type id = entidades.ID

//nolint:funlen
func newData() *data.Data {
	logFiles := logs.AbrirArquivos("./logs/data/")

	//nolint:exhaustivestruct
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

	bd, err := mariadb.NovoBD(config.FormatDSN())
	if err != nil {
		log.Panicln(err.Mensagem)
	}

	MariaDBPessoa := mariadb.PessoaBD{
		Conexão:      *mariadb.NovaConexão(logFiles.Pessoa, bd),
		NomeDaTabela: "Pessoa",
	}

	MariaDBCurso := mariadb.CursoBD{
		Conexão:                *mariadb.NovaConexão(logFiles.Curso, bd),
		NomeDaTabela:           "Curso",
		NomeDaTabelaSecundária: "CursoMatérias",
	}

	MariaDBAluno := mariadb.AlunoBD{
		Conexão: *mariadb.NovaConexão(logFiles.Aluno, bd),
	}

	MariaDBProfessor := mariadb.ProfessorBD{
		Conexão: *mariadb.NovaConexão(logFiles.Professor, bd),
	}

	MariaDBAdministrativo := mariadb.AdministrativoBD{
		Conexão: *mariadb.NovaConexão(logFiles.Administrativo, bd),
	}

	MariaDBMatéria := mongodb.MatériaBD{
		Connexão: *mongodb.NovaConexão(logFiles.Matéria),
	}

	MariaDBTurma := mongodb.TurmaBD{
		Connexão: *mongodb.NovaConexão(logFiles.Turma),
	}

	return &data.Data{
		Pessoa:         MariaDBPessoa,
		Curso:          MariaDBCurso,
		Aluno:          MariaDBAluno,
		Professor:      MariaDBProfessor,
		Administrativo: MariaDBAdministrativo,
		Matéria:        MariaDBMatéria,
		Turma:          MariaDBTurma,
	}
}

func prettyStruct(s ...interface{}) string {
	sJSON, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		log.Panicln(err.Error())
	}

	return string(sJSON)
}

//nolint: funlen, cyclop
func main() {
	r := gin.Default()

	Data := newData()

	r.GET("/ping", func(c *gin.Context) {
		rand.Seed(time.Now().UnixNano())
		id := entidades.NovoID()

		materias := &[]entidades.CursoMatéria{
			{
				IDCurso:    id,
				IDMatéria:  entidades.NovoID(),
				Período:    "Teste",
				Tipo:       "Não sei",
				Status:     "Testando",
				Observação: "Okay",
			},
			{
				IDCurso:    id,
				IDMatéria:  entidades.NovoID(),
				Período:    "Teste",
				Tipo:       "Não sei",
				Status:     "Testando",
				Observação: "Okay",
			},
		}

		curso := &entidades.Curso{
			ID:                id,
			Nome:              "Curso novo",
			DataDeInício:      time.Now(),
			DataDeDesativação: time.Now(),
			Matérias:          *materias,
		}

		err := Data.Curso.Inserir(curso)
		if err != nil {
			log.Println(err.Error())
			if err.ErroExterno != nil {
				log.Println(err.ErroExterno.Error())
			}
		}

		cursoSalvo, err := Data.Curso.Pegar(curso.ID)
		if err != nil {
			log.Println(err.Error())
			if err.ErroExterno != nil {
				log.Panicln(err.ErroExterno.Error())
			}
		}

		log.Println(prettyStruct(cursoSalvo))

		curso.Nome = "Nome Velho"
		curso.Matérias[0].Observação = "Observação1 "
		curso.Matérias[1].Observação = "Observação 2"
		curso.Matérias[0].Status = "Status NOveo "
		curso.Matérias[1].Status = "Novo status"

		err = Data.Curso.Atualizar(curso.ID, curso)
		if err != nil {
			log.Println(err.Error())
			if err.ErroExterno != nil {
				log.Panicln(err.ErroExterno.Error())
			}
		}

		cursoSalvo, err = Data.Curso.Pegar(curso.ID)
		if err != nil {
			log.Println(err.Error())
			if err.ErroExterno != nil {
				log.Panicln(err.ErroExterno.Error())
			}
		}

		log.Println(prettyStruct(cursoSalvo))

		err = Data.Curso.Deletar(curso.ID)
		if err != nil {
			log.Println(err.Error())
			if err.ErroExterno != nil {
				log.Panicln(err.ErroExterno.Error())
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
