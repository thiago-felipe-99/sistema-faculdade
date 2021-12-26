package main

import (
	"encoding/json"
	"log"
	"net/http"
	"reflect"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"thiagofelipe.com.br/sistema-faculdade-backend/data"
	"thiagofelipe.com.br/sistema-faculdade-backend/data/database/mariadb"
	"thiagofelipe.com.br/sistema-faculdade-backend/logica"
	"thiagofelipe.com.br/sistema-faculdade-backend/logs"
)

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

	return data.DataPadrão(logFiles, bd)
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

	data := newData()
	lógica := logica.NovaLógica(*data)

	r.GET("/ping", func(c *gin.Context) {
		pessoaCriar, erro := lógica.Pessoa.Criar(
			"Thiago Felipe",
			"00000000000",
			time.Date(1999, 12, 8, 0, 0, 0, 0, time.UTC),
			"senha",
		)
		if erro != nil {
			log.Println(erro.Traçado())

			return
		}

		pessoaPegar, erro := lógica.Pessoa.Pegar(pessoaCriar.ID)
		if erro != nil {
			log.Println(erro.Traçado())

			return
		}

		if !reflect.DeepEqual(pessoaCriar, pessoaPegar) {
			log.Printf(
				"Devia chegar %s chegou %s\n",
				prettyStruct(pessoaCriar),
				prettyStruct(pessoaPegar),
			)

			return
		}

		log.Println(prettyStruct(pessoaPegar))

		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
			"pessoa":  pessoaPegar,
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
