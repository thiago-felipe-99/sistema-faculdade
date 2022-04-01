package http

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"thiagofelipe.com.br/sistema-faculdade-backend/entidades"
	"thiagofelipe.com.br/sistema-faculdade-backend/erros"
	"thiagofelipe.com.br/sistema-faculdade-backend/logica"
)

type Pessoa struct {
	logica    *logica.Pessoa
	validator *validator.Validate
}

func CriarPessoa(lógica *logica.Pessoa, validador *validator.Validate) *Pessoa {
	return &Pessoa{
		logica:    lógica,
		validator: validador,
	}
}

func (p *Pessoa) enviar(c *gin.Context, codigo int) {
	c.Status(codigo)
	c.Abort()
}

func (p *Pessoa) enviarPessoa(c *gin.Context, codigo int, pessoa *entidades.Pessoa) {
	enviar := struct {
		ID               id        `json:"id"`
		Nome             string    `json:"nome"`
		CPF              string    `json:"cpf"`
		DataDeNascimento time.Time `json:"dataDeNascimento"`
	}{
		ID:               pessoa.ID,
		Nome:             pessoa.Nome,
		CPF:              pessoa.CPF,
		DataDeNascimento: pessoa.DataDeNascimento,
	}

	c.JSON(codigo, enviar)
	c.Abort()
}

func (p *Pessoa) enviarErro(c *gin.Context, erro erro) {
	status := 0
	mensagem := ""

	switch {
	case erro.ÉPadrão(logica.ErroCPFInválido):
		status = http.StatusBadRequest
		mensagem = logica.ErroCPFInválido.Mensagem

	case erro.ÉPadrão(logica.ErroCPFExiste):
		status = http.StatusConflict
		mensagem = logica.ErroCPFExiste.Mensagem

	case erro.ÉPadrão(logica.ErroDataDeNascimentoInválido):
		status = http.StatusBadRequest
		mensagem = logica.ErroDataDeNascimentoInválido.Mensagem

	case erro.ÉPadrão(logica.ErroSenhaInválida):
		status = http.StatusBadRequest
		mensagem = logica.ErroSenhaInválida.Mensagem

	case erro.ÉPadrão(logica.ErroPessoaNãoEncontrada):
		status = http.StatusNotFound
		mensagem = logica.ErroPessoaNãoEncontrada.Mensagem

	case erro.ÉPadrão(ErroDecodificarJSON):
		status = http.StatusBadRequest
		mensagem = ErroDecodificarJSON.Mensagem

	case erro.ÉPadrão(ErroRequisiçãoSemBody):
		status = http.StatusBadRequest
		mensagem = ErroRequisiçãoSemBody.Mensagem

	case erro.ÉPadrão(ErroDataDeNascimentoInválido):
		status = http.StatusBadRequest
		mensagem = ErroDataDeNascimentoInválido.Mensagem

	default:
		status = http.StatusInternalServerError
		mensagem = ErroInesperado.Mensagem
	}

	c.Abort()
	enviarErro(c, status, mensagem)
}

func (p *Pessoa) pegarID(c *gin.Context) {
	id, erro := entidades.ParseID(c.Params.ByName("id"))
	if erro != nil {
		p.enviarErro(c, erro)

		return
	}

	c.Set("id", id)
	c.Next()
}

func (p *Pessoa) pegarIDContexto(c *gin.Context) (*id, erro) {
	IDGet, existe := c.Get("id")
	if !existe {
		return nil, erros.Novo(ErroIDNãoExisteContexto, nil, nil)
	}

	id, okay := IDGet.(*id)
	if !okay {
		return nil, erros.Novo(ErroConverterIDContexto, nil, nil)
	}

	return id, nil
}

func (p *Pessoa) pegarBody(c *gin.Context) {
	decodificador := json.NewDecoder(c.Request.Body)
	pessoaString := struct {
		Nome             string `json:"nome" validate:"required"`
		CPF              string `json:"cpf" validate:"required"`
		DataDeNascimento string `json:"dataDeNascimento" validate:"required"`
		Senha            string `json:"senha" validate:"required"`
	}{}

	err := decodificador.Decode(&pessoaString)
	if err != nil {
		if errors.Is(err, io.EOF) {
			p.enviarErro(c, erros.Novo(ErroRequisiçãoSemBody, nil, err))

			return
		}

		p.enviarErro(c, erros.Novo(ErroDecodificarJSON, nil, err))

		return
	}

	err = p.validator.Struct(pessoaString)
	if err != nil {
		if erros, ok := err.(validator.ValidationErrors); ok { //nolint: errorlint
			mensagens := []string{}
			for _, erro := range erros.Translate(pegarTradutor(c)) {
				mensagens = append(mensagens, erro)
			}

			enviarErro(c, http.StatusBadRequest, mensagens...)

			return
		}

		p.enviarErro(c, erros.Novo(ErroValidarPessoa, nil, nil))

		return
	}

	data, err := time.Parse(dataFormatato, pessoaString.DataDeNascimento)
	if err != nil {
		p.enviarErro(c, erros.Novo(ErroDataDeNascimentoInválido, nil, nil))

		return
	}

	pessoa := &entidades.Pessoa{ //nolint: exhaustivestruct
		Nome:             pessoaString.Nome,
		CPF:              pessoaString.CPF,
		DataDeNascimento: data,
		Senha:            pessoaString.Senha,
	}

	c.Set("pessoa", pessoa)
	c.Next()
}

func (p *Pessoa) pegarPessoaContexto(c *gin.Context) (*entidades.Pessoa, erro) {
	pessoaGet, existe := c.Get("pessoa")
	if !existe {
		return nil, erros.Novo(ErroPessoaNãoExisteContexto, nil, nil)
	}

	pessoa, okay := pessoaGet.(*entidades.Pessoa)
	if !okay {
		return nil, erros.Novo(ErroConverterPessoaContexto, nil, nil)
	}

	return pessoa, nil
}

func (p *Pessoa) Criar(c *gin.Context) {
	body, erro := p.pegarPessoaContexto(c)
	if erro != nil {
		p.enviarErro(c, erro)

		return
	}

	pessoa, erro := p.logica.
		Criar(body.Nome, body.CPF, body.DataDeNascimento, body.Senha)
	if erro != nil {
		p.enviarErro(c, erro)

		return
	}

	p.enviarPessoa(c, http.StatusCreated, pessoa)
}

func (p *Pessoa) Atualizar(c *gin.Context) {
	id, erro := p.pegarIDContexto(c)
	if erro != nil {
		p.enviarErro(c, erro)

		return
	}

	body, erro := p.pegarPessoaContexto(c)
	if erro != nil {
		p.enviarErro(c, erro)

		return
	}

	pessoa, erro := p.logica.
		Atualizar(*id, body.Nome, body.CPF, body.DataDeNascimento, body.Senha)
	if erro != nil {
		p.enviarErro(c, erro)

		return
	}

	p.enviarPessoa(c, http.StatusOK, pessoa)
}

func (p *Pessoa) Pegar(c *gin.Context) {
	id, erro := p.pegarIDContexto(c)
	if erro != nil {
		p.enviarErro(c, erro)

		return
	}

	pessoa, erro := p.logica.Pegar(*id)
	if erro != nil {
		p.enviarErro(c, erro)

		return
	}

	p.enviarPessoa(c, http.StatusOK, pessoa)
}

func (p *Pessoa) Deletar(c *gin.Context) {
	id, erro := p.pegarIDContexto(c)
	if erro != nil {
		p.enviarErro(c, erro)

		return
	}

	erro = p.logica.Deletar(*id)
	if erro != nil {
		p.enviarErro(c, erro)

		return
	}

	p.enviar(c, http.StatusNoContent)
}

func PessoaRotas(roteamento *gin.RouterGroup, pessoa *Pessoa) {
	roteamento.POST("", pessoa.pegarBody, pessoa.Criar)
	roteamento.PUT("/:id", pessoa.pegarID, pessoa.pegarBody, pessoa.Atualizar)
	roteamento.GET("/:id", pessoa.pegarID, pessoa.Pegar)
	roteamento.DELETE("/:id", pessoa.pegarID, pessoa.Deletar)
}
