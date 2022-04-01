//nolint: dupl
package mongodb

import (
	"reflect"
	"testing"

	"go.mongodb.org/mongo-driver/mongo"
	"thiagofelipe.com.br/sistema-faculdade-backend/aleatorio"
	"thiagofelipe.com.br/sistema-faculdade-backend/data"
	"thiagofelipe.com.br/sistema-faculdade-backend/entidades"
)

func criarCursoAleatória() *curso {
	matérias := []cursoMatéria{}
	índice := uint(0)

	for ; índice <= aleatorio.Número(tamanhoMáximoMatérias); índice++ {
		matérias = append(matérias, cursoMatéria{
			Matéria:    entidades.NovoID(),
			Período:    aleatorio.Palavra(aleatorio.Número(tamanhoMáximoPalavra) + 1),
			Tipo:       aleatorio.Palavra(aleatorio.Número(tamanhoMáximoPalavra) + 1),
			Status:     aleatorio.Palavra(aleatorio.Número(tamanhoMáximoPalavra) + 1),
			Observação: aleatorio.Palavra(tamanhoMáximoPalavra),
		})
	}

	dataAgora := entidades.DataAtual()

	return &curso{
		ID:                entidades.NovoID(),
		Nome:              aleatorio.Palavra(aleatorio.Número(tamanhoMáximoPalavra) + 1),
		DataDeInício:      dataAgora,
		DataDeDesativação: dataAgora.AddDate(int(aleatorio.Número(50)), 0, 0),
		Matérias:          matérias,
	}
}

func adicionarCurso(t *testing.T, curso *curso) id {
	t.Helper()

	if erro := cursoBD.Inserir(curso); erro != nil {
		t.Fatalf("Erro ao inserir a curso no banco de dados: %v", erro)
	}

	cursoSalva, erro := cursoBD.Pegar(curso.ID)
	if erro != nil {
		t.Fatalf("Erro ao pegar a curso no banco de dados: %v", erro)
	}

	if !reflect.DeepEqual(curso, cursoSalva) {
		t.Fatalf(
			"Erro ao salvar a curso no banco de dados, queria: %v, chegou: %v",
			curso,
			cursoSalva,
		)
	}

	t.Cleanup(func() {
		removerCurso(t, curso.ID)
	})

	return curso.ID
}

func removerCurso(t *testing.T, id id) {
	t.Helper()

	if erro := cursoBD.Deletar(id); erro != nil {
		t.Fatalf("Erro ao tentar deletar a curso teste: %v", erro.Error())
	}
}

func TestInserirCurso(t *testing.T) {
	t.Parallel()

	t.Run("OKAY", func(t *testing.T) {
		t.Parallel()
		adicionarCurso(t, criarCursoAleatória())
	})

	t.Run("Duplicado/ID", func(t *testing.T) {
		t.Parallel()

		cursoTeste := criarCursoAleatória()

		adicionarCurso(t, cursoTeste)

		erro := cursoBD.Inserir(cursoTeste)
		if erro == nil || erro.ErroExterno == nil {
			t.Fatalf("Não foi enviado erro do sistema")
		}

		if !mongo.IsDuplicateKeyError(erro.ErroExterno) {
			t.Fatalf(
				"Erro ao inserir a curso queria: %v, chegou: %v",
				"Erro de id duplicado",
				erro.ErroExterno,
			)
		}
	})

	t.Run("TimeOut", func(t *testing.T) {
		t.Parallel()
		erro := cursoBDInválido.Inserir(criarCursoAleatória())
		if erro == nil || !mongo.IsTimeout(erro.ErroExterno) {
			t.Fatalf("Esperava um erro de Timeout, chegou: %v", erro)
		}
	})
}

func TestAtualizarCurso(t *testing.T) {
	t.Parallel()

	id := adicionarCurso(t, criarCursoAleatória())
	novoCurso := criarCursoAleatória()
	novoCurso.ID = id

	t.Run("OKAY", func(t *testing.T) {
		t.Parallel()

		erro := cursoBD.Atualizar(id, novoCurso)
		if erro != nil {
			t.Fatalf("Não deveria ter erro ao atualizar curso: %v", erro)
		}

		cursoSalvo, erro := cursoBD.Pegar(id)
		if erro != nil {
			t.Fatalf("Erro ao pegar a curso no banco de dados: %v", erro)
		}

		if !reflect.DeepEqual(novoCurso, cursoSalvo) {
			t.Fatalf(
				"Erro ao salvar a curso no banco de dados, queria: %v, chegou: %v",
				novoCurso,
				cursoSalvo,
			)
		}
	})

	t.Run("TimeOut", func(t *testing.T) {
		t.Parallel()
		erro := cursoBDInválido.Atualizar(id, novoCurso)
		if erro == nil || !mongo.IsTimeout(erro.ErroExterno) {
			t.Fatalf("Esperava um erro de Timeout, chegou: %v", erro)
		}
	})
}

func TestPegarCurso(t *testing.T) {
	t.Parallel()

	t.Run("OKAY", func(t *testing.T) {
		t.Parallel()
		adicionarCurso(t, criarCursoAleatória())
	})

	t.Run("CursoNãoEcontrado", func(t *testing.T) {
		t.Parallel()
		_, erro := cursoBD.Pegar(entidades.NovoID())
		if erro == nil || !erro.ÉPadrão(data.ErroCursoNãoEncontrado) {
			t.Fatalf(
				"Erro ao pegar a curso, queria: %v, chegou: %v",
				data.ErroCursoNãoEncontrado,
				erro,
			)
		}
	})

	t.Run("TimeOut", func(t *testing.T) {
		t.Parallel()
		_, erro := cursoBDInválido.Pegar(entidades.NovoID())
		if erro == nil || !mongo.IsTimeout(erro.ErroExterno) {
			t.Fatalf("Esperava um erro de Timeout, chegou: %v", erro)
		}
	})
}

func TestDeletarCurso(t *testing.T) {
	t.Parallel()

	t.Run("OKAY", func(t *testing.T) {
		t.Parallel()
		id := adicionarCurso(t, criarCursoAleatória())
		erro := cursoBD.Deletar(id)
		if erro != nil {
			t.Fatalf("Não esperava um erro ao deletar curso: %v", erro)
		}

		_, erro = cursoBD.Pegar(id)
		if erro == nil || !erro.ÉPadrão(data.ErroCursoNãoEncontrado) {
			t.Fatalf("Esperava: %v, chegou: %v", data.ErroCursoNãoEncontrado, erro)
		}
	})

	t.Run("TimeOut", func(t *testing.T) {
		t.Parallel()
		erro := cursoBDInválido.Deletar(entidades.NovoID())
		if erro == nil || !mongo.IsTimeout(erro.ErroExterno) {
			t.Fatalf("Esperava um erro de Timeout, chegou: %v", erro)
		}
	})
}
