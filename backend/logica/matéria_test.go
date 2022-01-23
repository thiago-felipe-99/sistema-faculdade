package logica

import (
	"reflect"
	"testing"
	"time"

	"thiagofelipe.com.br/sistema-faculdade-backend/aleatorio"
	"thiagofelipe.com.br/sistema-faculdade-backend/entidades"
)

func criarMatériaAleatórira() (string, time.Duration, float32, string) {
	nome := aleatorio.Palavra(aleatorio.Número(tamanhoMáximoPalavra) + 1)
	ch := time.Duration(aleatorio.Número(tamanhoMáximoCargaHorária)+1) * time.Hour
	créditos := float32(aleatorio.Número(tamanhoMáximoCréditos) + 1)
	tipo := aleatorio.Palavra(aleatorio.Número(tamanhoMáximoPalavra) + 1)

	return nome, ch, créditos, tipo
}

func adicionarMatéria(
	t *testing.T,
	nome string,
	ch time.Duration,
	créditos float32,
	tipo string,
	préRequisitos []id,
) id {
	t.Helper()

	matéria, erro := logicaTeste.Matéria.Criar(nome, ch, créditos, tipo, préRequisitos)
	if erro != nil {
		t.Fatalf("Não esperava erro ao criar pessoa: %v", erro)
	}

	matériaSalva, erro := logicaTeste.Matéria.Pegar(matéria.ID)
	if erro != nil {
		t.Fatalf("Não esperava erro ao pegar a pessoa: %v", erro)
	}

	if !reflect.DeepEqual(matéria, matériaSalva) {
		t.Fatalf("Esperava: %v\nChegou: %v", matéria, matériaSalva)
	}

	t.Cleanup(func() {
		removerMatéria(t, matéria.ID)
	})

	return matéria.ID
}

func removerMatéria(t *testing.T, id id) {
	t.Helper()

	erro := logicaTeste.Matéria.Deletar(id)
	if erro != nil {
		t.Errorf("Não esperava erro ao deletar a matéria: %v", erro)
	}
}

//nolint: funlen, cyclop
func TestCriarMatéria(t *testing.T) {
	t.Parallel()

	t.Run("OKAY", func(t *testing.T) {
		t.Parallel()

		nome1, ch1, créditos1, tipo1 := criarMatériaAleatórira()
		id1 := adicionarMatéria(t, nome1, ch1, créditos1, tipo1, []id{})

		nome2, ch2, créditos2, tipo2 := criarMatériaAleatórira()
		id2 := adicionarMatéria(t, nome2, ch2, créditos2, tipo2, []id{})

		nome3, ch3, créditos3, tipo3 := criarMatériaAleatórira()
		adicionarMatéria(t, nome3, ch3, créditos3, tipo3, []id{id1, id2})
	})

	t.Run("CargaHoráriaSemanalInválido", func(t *testing.T) {
		t.Parallel()

		nome, _, créditos, tipo := criarMatériaAleatórira()

		_, erro := logicaTeste.Matéria.Criar(nome, time.Minute, créditos, tipo, []id{})
		if erro == nil || !erro.ÉPadrão(ErroCargaHoráriaMínima) {
			t.Fatalf("Esperava: %v\nChegou: %v", ErroCargaHoráriaMínima, erro)
		}
	})

	t.Run("CréditosInválido", func(t *testing.T) {
		t.Parallel()

		nome, ch, _, tipo := criarMatériaAleatórira()

		_, erro := logicaTeste.Matéria.Criar(nome, ch, 0, tipo, []id{})
		if erro == nil || !erro.ÉPadrão(ErroCréditosInválido) {
			t.Fatalf("Esperava: %v\nChegou: %v", ErroCréditosInválido, erro)
		}
	})

	t.Run("PréRequisitosInválidos", func(t *testing.T) {
		t.Parallel()

		id1, id2 := entidades.NovoID(), entidades.NovoID()
		nome, ch, créditos, tipo := criarMatériaAleatórira()

		_, erro := logicaTeste.Matéria.Criar(nome, ch, créditos, tipo, []id{id1, id2})
		if erro == nil || !erro.ÉPadrão(ErroPréRequisitosNãoExiste) {
			t.Fatalf("Esperava: %v\nChegou: %v", ErroPréRequisitosNãoExiste, erro)
		}
	})

	t.Run("TimeOut", func(t *testing.T) {
		t.Parallel()

		id1, id2 := entidades.NovoID(), entidades.NovoID()
		nome, ch, créditos, tipo := criarMatériaAleatórira()

		_, erro := matériaBDTimeOut.Criar(nome, ch, créditos, tipo, []id{id1, id2})
		if erro == nil || erro.ErroInicial == nil {
			t.Fatalf("Esperava um erro de timeout: %v", erro)
		}
	})

	t.Run("TimeOut2", func(t *testing.T) {
		t.Parallel()

		nome, ch, créditos, tipo := criarMatériaAleatórira()

		_, erro := matériaBDTimeOut.Criar(nome, ch, créditos, tipo, []id{})
		if erro == nil || erro.ErroInicial == nil {
			t.Fatalf("Esperava um erro de timeout: %v", erro)
		}
	})
}
