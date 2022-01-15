package logs

import (
	"bytes"
	"fmt"
	"os"
	"reflect"
	"regexp"
	"testing"

	"thiagofelipe.com.br/sistema-faculdade-backend/erros"
)

func TestNovoLog(t *testing.T) {
	t.Run("OKAY", func(t *testing.T) {
		mensagem := "teste"
		flags := ` - [0-9]{4}/[0-9]{2}/[0-9]{2} [0-9]{2}:[0-9]{2}:[0-9]{2} main.go:[0-9]+: ` + mensagem + `\n$`

		outs := map[string]string{
			"Panic":      "PANIC" + flags,
			"Erro":       "ERRO" + flags,
			"Aviso":      "AVISO" + flags,
			"Informação": "INFORMAÇÃO" + flags,
			"Debug":      "DEBUG" + flags,
		}

		testes := []struct {
			nível uint
			outs  map[string]bool
		}{
			{
				NívelPanic,
				map[string]bool{
					"Erro":       false,
					"Aviso":      false,
					"Informação": false,
					"Debug":      false,
				},
			},
			{
				NívelErro,
				map[string]bool{
					"Erro":       true,
					"Aviso":      false,
					"Informação": false,
					"Debug":      false,
				},
			},
			{
				NívelAviso,
				map[string]bool{
					"Erro":       true,
					"Aviso":      true,
					"Informação": false,
					"Debug":      false,
				},
			},
			{
				NívelInfo,
				map[string]bool{
					"Erro":       true,
					"Aviso":      true,
					"Informação": true,
					"Debug":      false,
				},
			},
			{
				NívelDebug,
				map[string]bool{
					"Erro":       true,
					"Aviso":      true,
					"Informação": true,
					"Debug":      true,
				},
			},
		}

		mensagemInput := []reflect.Value{reflect.ValueOf(mensagem)}

		for _, teste := range testes {
			t.Run(fmt.Sprint(teste.nível), func(t *testing.T) {
				var buffer bytes.Buffer
				log := NovoLog(&buffer, teste.nível)

				for função, imprimir := range teste.outs {
					buffer.Reset()
					reflect.ValueOf(log).MethodByName(função).Call(mensagemInput)

					if imprimir {
						padrão := regexp.MustCompile(outs[função])

						if !padrão.MatchString(buffer.String()) {
							t.Fatalf(
								"Esperava a mensagem com o padrão `%s`\nChegou: %s",
								outs[função], buffer.String(),
							)
						}
					} else {
						if buffer.String() != "" {
							t.Fatalf("Não esperava mensagem porém chegou: %s", buffer.String())
						}
					}
				}

				defer func() {
					r := recover()

					rValue := fmt.Sprintf("%v", r)
					mensagemInputValue := fmt.Sprintf("%v", mensagemInput)

					if rValue != mensagemInputValue {
						t.Fatalf("Esperava: %s, chegou: %v", mensagemInput, r)
					}

					padrão := regexp.MustCompile(outs["Panic"])

					if !padrão.MatchString(buffer.String()) {
						t.Fatalf(
							"Esperava a mensagem com o padrão `%s`\nChegou: %s",
							outs["Panic"], buffer.String(),
						)
					}
				}()

				buffer.Reset()
				log.Panic(mensagem)

			})
		}
	})

	t.Run("NívelErrado", func(t *testing.T) {
		var buffer bytes.Buffer

		defer func() {
			r := recover()
			erroEsperado := erros.Novo(ErroNívelInválido, nil, nil)

			if !reflect.DeepEqual(r, erroEsperado) {
				t.Fatalf("Esperava: %v\nChegou: %v", erroEsperado, r)
			}
		}()

		NovoLog(&buffer, NívelDebug+1)
	})
}

func TestAbrirArquivos(t *testing.T) {
	pasta := "./logs/"
	entidades := []string{"Pessoa", "Curso", "Aluno", "Professor", "Administrativo", "Matéria", "Turma"}

	t.Run("OKAY", func(t *testing.T) {
		arquivos := AbrirArquivos(pasta)
		arquivosRefletidos := reflect.ValueOf(*arquivos)

		for _, entidade := range entidades {

			t.Run(entidade, func(t *testing.T) {
				arquivo := arquivosRefletidos.FieldByName(entidade).Elem()
				tipo := arquivo.Type().String()

				if tipo != "*os.File" {
					t.Fatalf("Espera um arquivo do tipo *os.File, chegou: %s", tipo)
				}

				caminhoArquivo := arquivo.MethodByName("Name").Call(nil)[0].String()
				caminhoEsperado := pasta + entidade + "Logs.txt"
				if caminhoArquivo != caminhoEsperado {
					t.Fatalf("Esperava: %s, chegou: %s", caminhoEsperado, caminhoArquivo)
				}
			})
		}
	})

	t.Run("PermisãoErrada", func(t *testing.T) {
		for _, entidade := range entidades {

			t.Run(entidade, func(t *testing.T) {
				const flags = os.O_CREATE

				const mode os.FileMode = 0644

				caminhoArquivo := pasta + entidade + "Logs.txt"

				arquivo, erro := os.OpenFile(caminhoArquivo, flags, mode)
				if erro != nil {
					t.Fatalf("Um erro inesperado aconteceu: %v", erro)
				}

				defer arquivo.Chmod(mode)
				arquivo.Chmod(0000)

				defer func() {
					r := recover()
					pathError := os.PathError{
						Op:   "open",
						Path: caminhoArquivo,
						Err:  os.ErrPermission,
					}
					erroEsperado := erros.ErroExterno(&pathError)
					if r != erroEsperado {
						t.Fatalf("Esperava: %v\nChegou: %v", erroEsperado, r)
					}
				}()

				AbrirArquivos(pasta)
			})
		}
	})
}

func TestNovoLogEntidades(t *testing.T) {
	NovoLogEntidades(AbrirArquivos("./logs"), NívelDebug)
}