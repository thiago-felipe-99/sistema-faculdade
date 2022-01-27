package logica

import (
	"reflect"
	"testing"
	"time"

	"thiagofelipe.com.br/sistema-faculdade-backend/aleatorio"
	"thiagofelipe.com.br/sistema-faculdade-backend/entidades"
)

func criarCursoAleatório(t *testing.T) (string, time.Time, time.Time, []cursoMatéria) {
	t.Helper()

	nome := aleatorio.Palavra(aleatorio.Número(tamanhoMáximoPalavra) + 1)
	início := entidades.DataAtual()
	final := início.AddDate(int(aleatorio.Número(tamanhoMáximoData)+1), 0, 0)
	matérias := []cursoMatéria{}

	for índice := uint(0); índice <= aleatorio.Número(tamanhoMáximoMatéria); índice++ {
		nome, ch, créditos, tipo := criarMatériaAleatória()
		id := adicionarMatéria(t, nome, ch, créditos, tipo, []id{})

		matérias = append(matérias, cursoMatéria{
			Matéria:    id,
			Período:    aleatorio.Palavra(aleatorio.Número(tamanhoMáximoPalavra) + 1),
			Tipo:       aleatorio.Palavra(aleatorio.Número(tamanhoMáximoPalavra) + 1),
			Status:     aleatorio.Palavra(aleatorio.Número(tamanhoMáximoPalavra) + 1),
			Observação: aleatorio.Palavra(aleatorio.Número(tamanhoMáximoPalavra) + 1),
		})
	}

	return nome, início, final, matérias
}

func adicionarCurso(
	t *testing.T,
	nome string,
	início time.Time,
	desativação time.Time,
	matérias []cursoMatéria,
) id {
	t.Helper()

	curso, erro := logicaTeste.Curso.Criar(nome, início, desativação, matérias)
	if erro != nil {
		t.Fatalf("Não esperava erro ao criar curso: %v", erro)
	}

	cursoSalva, erro := logicaTeste.Curso.Pegar(curso.ID)
	if erro != nil {
		t.Log("Ffoi4")
		t.Fatalf("Não esperava erro ao pegar o curso: %v", erro)
	}

	if !reflect.DeepEqual(curso, cursoSalva) {
		t.Fatalf("Esperava: %v\nChegou: %v", curso, cursoSalva)
	}

	t.Cleanup(func() {
		removerCurso(t, curso.ID)
	})

	return curso.ID
}

func removerCurso(t *testing.T, id id) {
	t.Helper()

	erro := logicaTeste.Curso.Deletar(id)
	if erro != nil {
		t.Errorf("Não esperava erro ao deletar o curso: %v", erro)
	}
}

func TestCriarCurso(t *testing.T) {
	t.Parallel()

	nome, dataDeInício, dataDeDesativação, matérias := criarCursoAleatório(t)

	t.Run("OKAY", func(t *testing.T) {
		t.Parallel()
		adicionarCurso(t, nome, dataDeInício, dataDeDesativação, matérias)
	})

	t.Run("DatasInválida", func(t *testing.T) {
		t.Parallel()

		_, erro := logicaTeste.Curso.
			Criar(nome, dataDeDesativação, dataDeInício, matérias)
		if erro == nil || !erro.ÉPadrão(ErroDataDeInícioMaior) {
			t.Fatalf("Esperava: %v\nChegou: %v", ErroDataDeInícioMaior, erro)
		}
	})

	t.Run("MatériaNãoEncontrada", func(t *testing.T) {
		t.Parallel()

		matéria := cursoMatéria{
			Matéria:    entidades.NovoID(),
			Período:    "erro",
			Tipo:       "teste",
			Status:     "teste",
			Observação: "teste",
		}

		_, erro := logicaTeste.Curso.
			Criar(nome, dataDeInício, dataDeDesativação, append(matérias, matéria))
		if erro == nil || !erro.ÉPadrão(ErroMatériaNãoEncontrada) {
			t.Fatalf("Esperava: %v\nChegou: %v", ErroMatériaNãoEncontrada, erro)
		}
	})

	t.Run("TimeOut", func(t *testing.T) {
		t.Parallel()

		_, erro := cursoBDTimeOut.
			Criar(nome, dataDeInício, dataDeDesativação, []cursoMatéria{})
		if erro == nil || !erro.ÉPadrão(ErroCriarCurso) {
			t.Fatalf("Esperava: %v\nChegou: %v", ErroCriarCurso, erro)
		}
	})

	t.Run("TimeOut/2", func(t *testing.T) {
		t.Parallel()

		_, erro := cursoBDTimeOut.
			Criar(nome, dataDeInício, dataDeDesativação, matérias)
		if erro == nil || !erro.ÉPadrão(ErroCriarCurso) {
			t.Fatalf("Esperava: %v\nChegou: %v", ErroCriarCurso, erro)
		}
	})
}

func TestPegarCurso(t *testing.T) {
	t.Parallel()

	t.Run("OKAY", func(t *testing.T) {
		t.Parallel()

		nome, dataDeInício, dataDeDesativação, tipo := criarCursoAleatório(t)
		adicionarCurso(t, nome, dataDeInício, dataDeDesativação, tipo)
	})

	t.Run("CursoNãoEncontrada", func(t *testing.T) {
		t.Parallel()

		_, erro := logicaTeste.Curso.Pegar(entidades.NovoID())
		if erro == nil || !erro.ÉPadrão(ErroCursoNãoEncontrado) {
			t.Fatalf("Esperava: %v\nChegou: %v", ErroCursoNãoEncontrado, erro)
		}
	})

	t.Run("TimeOut", func(t *testing.T) {
		t.Parallel()

		_, erro := cursoBDTimeOut.Pegar(entidades.NovoID())
		if erro == nil || !erro.ÉPadrão(ErroPegarCurso) {
			t.Fatalf("Esperava um erro de timeout: %v", erro)
		}
	})
}

func TestDeletarCurso(t *testing.T) {
	t.Parallel()

	t.Run("OKAY", func(t *testing.T) {
		t.Parallel()

		nome, dataDeInício, dataDeDesativação, tipo := criarCursoAleatório(t)
		adicionarCurso(t, nome, dataDeInício, dataDeDesativação, tipo)
	})

	t.Run("CursoNãoEncontrado", func(t *testing.T) {
		t.Parallel()

		erro := logicaTeste.Curso.Deletar(entidades.NovoID())
		if erro == nil || !erro.ÉPadrão(ErroCursoNãoEncontrado) {
			t.Fatalf("Esperava: %v\nChegou: %v", ErroCursoNãoEncontrado, erro)
		}
	})

	t.Run("TimeOut", func(t *testing.T) {
		t.Parallel()

		erro := cursoBDTimeOut.Deletar(entidades.NovoID())
		if erro == nil || !erro.ÉPadrão(ErroDeletarCurso) {
			t.Fatalf("Esperava um erro de timeout: %v", erro)
		}
	})

	t.Run("BDInválido", func(t *testing.T) {
		t.Parallel()

		nome, dataDeInício, dataDeDesativação, tipo := criarCursoAleatório(t)
		id := adicionarCurso(t, nome, dataDeInício, dataDeDesativação, tipo)

		erro := cursoBDInválido.Deletar(id)
		if erro == nil || !erro.ÉPadrão(ErroDeletarCurso) {
			t.Fatalf("Esperava um erro de deletar curso: %v", erro)
		}
	})
}
