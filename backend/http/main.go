package http

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/es"
	"github.com/go-playground/locales/pt_BR"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_tradução "github.com/go-playground/validator/v10/translations/en"
	es_tradução "github.com/go-playground/validator/v10/translations/es"
	pt_tradução "github.com/go-playground/validator/v10/translations/pt_BR"
	"thiagofelipe.com.br/sistema-faculdade-backend/entidades"
	"thiagofelipe.com.br/sistema-faculdade-backend/erros"
	"thiagofelipe.com.br/sistema-faculdade-backend/logica"
)

type (
	id   = entidades.ID
	erro = *erros.Aplicação
)

var uni *ut.UniversalTranslator //nolint: gochecknoglobals

const dataFormatato = "2/1/2006"

func pegarTradutor(c *gin.Context) ut.Translator {
	trans, existe := uni.GetTranslator(c.Request.Header.Get("Accept-Language"))
	if !existe {
		trans, _ = uni.GetTranslator("pt_BR")
	}

	return trans
}

func Rotas(url string, lógica *logica.Lógica) {
	roteamento := gin.Default()

	validate := validator.New()

	uni = ut.New(pt_BR.New(), en.New(), es.New())

	trans, _ := uni.GetTranslator("pt_BR")

	err := pt_tradução.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		panic(err)
	}

	trans, _ = uni.GetTranslator("en")

	err = en_tradução.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		panic(err)
	}

	trans, _ = uni.GetTranslator("es")

	err = es_tradução.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		panic(err)
	}

	pessoaTeste := CriarPessoa(lógica.Pessoa, validate)
	PessoaRotas(roteamento.Group("/pessoa"), pessoaTeste)

	err = roteamento.Run(url)
	if err != nil {
		panic(err)
	}
}
