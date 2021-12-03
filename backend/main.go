package main

import (
	"encoding/json"
	"fmt"
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
		Conexão:                *mariadb.NovaConexão(logFiles.Aluno, bd),
		NomeDaTabela:           "Aluno",
		NomeDaTabelaSecundária: "AlunoTurma",
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

//nolint: funlen
func main() {
	r := gin.Default()

	Data := newData()

	r.GET("/ping", func(c *gin.Context) {
		rand.Seed(time.Now().UnixNano())
		pessoa := &entidades.Pessoa{
			ID:               entidades.NovoID(),
			Nome:             "Nome da PEssoa",
			CPF:              fmt.Sprintf("%011d", rand.Intn(99999999999)), //nolint:gosec,gomnd,lll
			DataDeNascimento: time.Now(),
			Senha:            "Senha errada",
		}

		err := Data.Pessoa.Inserir(pessoa)
		if err != nil {
			log.Panicln(err.Error())
		}

		idCurso := entidades.NovoID()

		materias := &[]entidades.CursoMatéria{
			{
				IDCurso:    idCurso,
				IDMatéria:  entidades.NovoID(),
				Período:    "Teste",
				Tipo:       "Não sei",
				Status:     "Testando",
				Observação: "Okay",
			},
			{
				IDCurso:    idCurso,
				IDMatéria:  entidades.NovoID(),
				Período:    "Teste",
				Tipo:       "Não sei",
				Status:     "Testando",
				Observação: "Okay",
			},
		}

		curso := &entidades.Curso{
			ID:                idCurso,
			Nome:              "Curso novo",
			DataDeInício:      time.Now(),
			DataDeDesativação: time.Now(),
			Matérias:          *materias,
		}

		err = Data.Curso.Inserir(curso)
		if err != nil {
			log.Println(err.Error())
		}

		cursoSalvo, err := Data.Curso.Pegar(curso.ID)
		if err != nil {
			log.Println(err.Error())
		}

		log.Println(prettyStruct(cursoSalvo))

		idAluno := entidades.NovoID()

		turmas := &[]entidades.TurmaAluno{
			{
				IDAluno: idAluno,
				IDTurma: entidades.NovoID(),
				Status:  "Testando",
			},
			{
				IDAluno: idAluno,
				IDTurma: entidades.NovoID(),
				Status:  "Testando",
			},
		}

		aluno := &entidades.Aluno{
			ID:             idAluno,
			Pessoa:         pessoa.ID,
			DataDeIngresso: time.Now(),
			DataDeSaída:    time.Now(),
			Turmas:         *turmas,
			Matrícula:      fmt.Sprintf("%011d", rand.Intn(99999999999)), //nolint:gosec,gomnd,lll
			Curso:          curso.ID,
			Período:        "2022",
			Status:         "Okay",
		}

		err = Data.Aluno.Inserir(aluno)
		if err != nil {
			log.Panicln(err.Error())
		}

		alunoSalvo, err := Data.Aluno.Pegar(aluno.ID)
		if err != nil {
			log.Panicln(err.Error())
		}

		log.Println(prettyStruct(alunoSalvo))

		// aluno.Nome = "Nome Velho"
		// aluno.Matérias[0].Observação = "Observação1 "
		// aluno.Matérias[1].Observação = "Observação 2"
		// aluno.Matérias[0].Status = "Status NOveo "
		// aluno.Matérias[1].Status = "Novo status"

		// err = Data.Curso.Atualizar(aluno.ID, aluno)
		// if err != nil {
		// 	log.Println(err.Error())
		// }

		// cursoSalvo, err = Data.Curso.Pegar(aluno.ID)
		// if err != nil {
		// 	log.Println(err.Error())
		// }

		// log.Println(prettyStruct(cursoSalvo))

		// err = Data.Curso.Deletar(aluno.ID)
		// if err != nil {
		// 	log.Println(err.Error())
		// }

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
