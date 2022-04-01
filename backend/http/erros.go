package http

import (
	"github.com/gin-gonic/gin"
	"thiagofelipe.com.br/sistema-faculdade-backend/erros"
)

const (
	inesperado = iota
	requisiçãoSemBody
	decodificarJSON
	validarPessoa
	dataDeNascimentoInválida
	idNãoExisteContexto
	converterIDContexto
	pessoaNãoExisteContexto
	converterPessoaContexto
)

func enviarErro(c *gin.Context, código int, erros ...string) {
	c.JSON(código, gin.H{"erro": erros})
	c.Abort()
}

var criarErroPadrão = erros.NovoPadrãoFunc("HTTP") //nolint:gochecknoglobals

// Erros do servidor http.
var (
	ErroInesperado = criarErroPadrão(
		"Ocorreu um erro inesperado",
		inesperado,
	)
	ErroRequisiçãoSemBody = criarErroPadrão(
		"Não foi passado um body para a requisição",
		requisiçãoSemBody,
	)
	ErroDecodificarJSON = criarErroPadrão(
		"Erro ao decodificar o json",
		decodificarJSON,
	)
	ErroValidarPessoa = criarErroPadrão(
		"Erro ao validar pessoa",
		validarPessoa,
	)
	ErroDataDeNascimentoInválido = criarErroPadrão(
		"A DataDeNascimento precisa esta no formato \"DD/MM/AAAA\"",
		dataDeNascimentoInválida,
	)
	ErroIDNãoExisteContexto = criarErroPadrão(
		"Não foi encontrado ID nos contexto",
		idNãoExisteContexto,
	)
	ErroConverterIDContexto = criarErroPadrão(
		"Erro ao tentar converter o ID do contexto",
		converterIDContexto,
	)
	ErroPessoaNãoExisteContexto = criarErroPadrão(
		"Não foi encontrado body pessoa no contexto",
		pessoaNãoExisteContexto,
	)
	ErroConverterPessoaContexto = criarErroPadrão(
		"Erro ao tentar converter o body da pessoa no contexto",
		converterPessoaContexto,
	)
)
