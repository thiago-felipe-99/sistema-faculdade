package logica

import (
	"reflect"
	"testing"
	"time"

	"thiagofelipe.com.br/sistema-faculdade-backend/aleatorio"
	"thiagofelipe.com.br/sistema-faculdade-backend/entidades"
)

func criarMatériaAleatória() (string, time.Duration, float32, string) {
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

		nome1, ch1, créditos1, tipo1 := criarMatériaAleatória()
		id1 := adicionarMatéria(t, nome1, ch1, créditos1, tipo1, []id{})

		nome2, ch2, créditos2, tipo2 := criarMatériaAleatória()
		id2 := adicionarMatéria(t, nome2, ch2, créditos2, tipo2, []id{})

		nome3, ch3, créditos3, tipo3 := criarMatériaAleatória()
		adicionarMatéria(t, nome3, ch3, créditos3, tipo3, []id{id1, id2})
	})

	t.Run("CargaHoráriaSemanalInválido", func(t *testing.T) {
		t.Parallel()

		nome, _, créditos, tipo := criarMatériaAleatória()

		_, erro := logicaTeste.Matéria.Criar(nome, time.Minute, créditos, tipo, []id{})
		if erro == nil || !erro.ÉPadrão(ErroCargaHoráriaMínima) {
			t.Fatalf("Esperava: %v\nChegou: %v", ErroCargaHoráriaMínima, erro)
		}
	})

	t.Run("CréditosInválido", func(t *testing.T) {
		t.Parallel()

		nome, ch, _, tipo := criarMatériaAleatória()

		_, erro := logicaTeste.Matéria.Criar(nome, ch, 0, tipo, []id{})
		if erro == nil || !erro.ÉPadrão(ErroCréditosInválido) {
			t.Fatalf("Esperava: %v\nChegou: %v", ErroCréditosInválido, erro)
		}
	})

	t.Run("PréRequisitosInválidos", func(t *testing.T) {
		t.Parallel()

		id1, id2 := entidades.NovoID(), entidades.NovoID()
		nome, ch, créditos, tipo := criarMatériaAleatória()

		_, erro := logicaTeste.Matéria.Criar(nome, ch, créditos, tipo, []id{id1, id2})
		if erro == nil || !erro.ÉPadrão(ErroPréRequisitosNãoExiste) {
			t.Fatalf("Esperava: %v\nChegou: %v", ErroPréRequisitosNãoExiste, erro)
		}
	})

	t.Run("TimeOut", func(t *testing.T) {
		t.Parallel()

		id1, id2 := entidades.NovoID(), entidades.NovoID()
		nome, ch, créditos, tipo := criarMatériaAleatória()

		_, erro := matériaBDTimeOut.Criar(nome, ch, créditos, tipo, []id{id1, id2})
		if erro == nil || erro.ErroInicial == nil {
			t.Fatalf("Esperava um erro de timeout: %v", erro)
		}
	})

	t.Run("TimeOut2", func(t *testing.T) {
		t.Parallel()

		nome, ch, créditos, tipo := criarMatériaAleatória()

		_, erro := matériaBDTimeOut.Criar(nome, ch, créditos, tipo, []id{})
		if erro == nil || erro.ErroInicial == nil {
			t.Fatalf("Esperava um erro de timeout: %v", erro)
		}
	})
}

//nolint: funlen, cyclop
func TestAtualizarMatéria(t *testing.T) {
	t.Parallel()

	nome1, ch1, créditos1, tipo1 := criarMatériaAleatória()
	id1 := adicionarMatéria(t, nome1, ch1, créditos1, tipo1, []id{})

	nome2, ch2, créditos2, tipo2 := criarMatériaAleatória()
	id2 := adicionarMatéria(t, nome2, ch2, créditos2, tipo2, []id{})

	id3, id4 := entidades.NovoID(), entidades.NovoID()

	t.Run("OKAY", func(t *testing.T) {
		t.Parallel()

		nome3, ch3, créditos3, tipo3 := criarMatériaAleatória()
		id3 := adicionarMatéria(t, nome3, ch3, créditos3, tipo3, []id{id1, id2})

		matéria, erro := logicaTeste.Matéria.Atualizar(
			id3,
			nome1,
			ch2,
			créditos3,
			tipo2,
			[]id{id1},
		)
		if erro != nil {
			t.Fatalf("Não esperava um erro ao atualizar o atulizar a matéria: %v", erro)
		}

		matériaSalva, erro := logicaTeste.Matéria.Pegar(id3)
		if erro != nil {
			t.Fatalf("Não esperava um erro ao pegar a matéria: %v", erro)
		}

		if !reflect.DeepEqual(matéria, matériaSalva) {
			t.Fatalf("Esperava: %v\nChegou: %v", matéria, matériaSalva)
		}
	})

	t.Run("IDNãoExiste", func(t *testing.T) {
		t.Parallel()

		_, erro := logicaTeste.Matéria.Atualizar(
			entidades.NovoID(),
			nome1,
			ch1,
			créditos1,
			tipo1,
			[]id{},
		)
		if erro == nil || !erro.ÉPadrão(ErroMatériaNãoEncontrada) {
			t.Fatalf("Esperava: %v\nChegou: %v", ErroMatériaNãoEncontrada, erro)
		}
	})

	t.Run("CargaHoráriaSemanalInválido", func(t *testing.T) {
		t.Parallel()

		_, erro := logicaTeste.Matéria.Atualizar(
			id1,
			nome1,
			time.Minute,
			créditos1,
			tipo1,
			[]id{},
		)
		if erro == nil || !erro.ÉPadrão(ErroCargaHoráriaMínima) {
			t.Fatalf("Esperava: %v\nChegou: %v", ErroCargaHoráriaMínima, erro)
		}
	})

	t.Run("CréditosInválido", func(t *testing.T) {
		t.Parallel()

		_, erro := logicaTeste.Matéria.Atualizar(id1, nome1, ch1, 0, tipo1, []id{})
		if erro == nil || !erro.ÉPadrão(ErroCréditosInválido) {
			t.Fatalf("Esperava: %v\nChegou: %v", ErroCréditosInválido, erro)
		}
	})

	t.Run("PréRequisitosInválidos", func(t *testing.T) {
		t.Parallel()

		_, erro := logicaTeste.Matéria.Atualizar(
			id1,
			nome1,
			ch1,
			créditos1,
			tipo1,
			[]id{id3, id4},
		)
		if erro == nil || !erro.ÉPadrão(ErroPréRequisitosNãoExiste) {
			t.Fatalf("Esperava: %v\nChegou: %v", ErroPréRequisitosNãoExiste, erro)
		}
	})

	t.Run("TimeOut", func(t *testing.T) {
		t.Parallel()

		_, erro := matériaBDTimeOut.Atualizar(id1, nome1, ch1, créditos1, tipo1, []id{id3})
		if erro == nil || !erro.ÉPadrão(ErroAtualizarMatéria) {
			t.Fatalf("Esperava um erro de timeout: %v", erro)
		}
	})

	t.Run("BDInválido", func(t *testing.T) {
		t.Parallel()

		_, erro := matériaBDInválido.Atualizar(id1, nome1, ch1, créditos1, tipo1, []id{})
		if erro == nil || !erro.ÉPadrão(ErroAtualizarMatéria) {
			t.Fatalf("Esperava um erro ao atualizar matéria: %v", erro)
		}
	})
}

//nolint: funlen, cyclop
func TestPréRequisitosCiclo(t *testing.T) {
	t.Parallel()

	nome1, ch1, créditos1, tipo1 := criarMatériaAleatória()
	id1 := adicionarMatéria(t, nome1, ch1, créditos1, tipo1, []id{})

	nome2, ch2, créditos2, tipo2 := criarMatériaAleatória()
	id2 := adicionarMatéria(t, nome2, ch2, créditos2, tipo2, []id{})

	nome3, ch3, créditos3, tipo3 := criarMatériaAleatória()
	id3 := adicionarMatéria(t, nome3, ch3, créditos3, tipo3, []id{id1, id2})

	nome4, ch4, créditos4, tipo4 := criarMatériaAleatória()
	id4 := adicionarMatéria(t, nome4, ch4, créditos4, tipo4, []id{id1, id2, id3})

	nome5, ch5, créditos5, tipo5 := criarMatériaAleatória()
	id5 := adicionarMatéria(t, nome5, ch5, créditos5, tipo5, []id{id1, id2, id4})

	t.Run("OKAY", func(t *testing.T) {
		t.Parallel()

		ciclo, erro := logicaTeste.Matéria.préRequisitoCiclicos(id3, []id{id1, id2})
		if erro != nil {
			t.Fatalf("Não esperava um erro ao verificar o ciclo: %v", erro)
		}

		if ciclo {
			t.Fatalf("Não esperava um ciclo")
		}
	})

	t.Run("Próprio", func(t *testing.T) {
		t.Parallel()

		_, erro := logicaTeste.Matéria.Atualizar(
			id3,
			nome3,
			ch3,
			créditos3,
			tipo3,
			[]id{id1, id2, id3},
		)
		if erro == nil || !erro.ÉPadrão(ErroPréRequisitosCiclo) {
			t.Fatalf("Esperava: %v\nChegou: %v", ErroPréRequisitosCiclo, erro)
		}
	})

	t.Run("Direto", func(t *testing.T) {
		t.Parallel()

		_, erro := logicaTeste.Matéria.Atualizar(
			id3,
			nome3,
			ch3,
			créditos3,
			tipo3,
			[]id{id1, id2, id4},
		)
		if erro == nil || !erro.ÉPadrão(ErroPréRequisitosCiclo) {
			t.Fatalf("Esperava: %v\nChegou: %v", ErroPréRequisitosCiclo, erro)
		}
	})

	t.Run("Indireto", func(t *testing.T) {
		t.Parallel()

		_, erro := logicaTeste.Matéria.Atualizar(
			id3,
			nome3,
			ch3,
			créditos3,
			tipo3,
			[]id{id1, id2, id5},
		)
		if erro == nil || !erro.ÉPadrão(ErroPréRequisitosCiclo) {
			t.Fatalf("Esperava: %v\nChegou: %v", ErroPréRequisitosCiclo, erro)
		}
	})

	t.Run("BDInválido", func(t *testing.T) {
		t.Parallel()

		_, erro := matériaBDInválido.Atualizar(
			id1,
			nome1,
			ch1,
			créditos1,
			tipo1,
			[]id{id2},
		)
		if erro == nil || !erro.ÉPadrão(ErroAtualizarMatéria) {
			t.Fatalf("Esperava um erro ao atualizar matéria: %v", erro)
		}
	})

	t.Run("BDInválido/2", func(t *testing.T) {
		t.Parallel()

		_, erro := matériaBDInválido2.Atualizar(
			id1,
			nome1,
			ch1,
			créditos1,
			tipo1,
			[]id{id2},
		)
		if erro == nil || !erro.ÉPadrão(ErroAtualizarMatéria) {
			t.Fatalf("Esperava um erro ao atualizar matéria: %v", erro)
		}
	})
}

func TestPegarMatéria(t *testing.T) {
	t.Parallel()

	t.Run("OKAY", func(t *testing.T) {
		t.Parallel()

		nome, ch, créditos, tipo := criarMatériaAleatória()
		adicionarMatéria(t, nome, ch, créditos, tipo, []id{})
	})

	t.Run("MatériaNãoEncontrada", func(t *testing.T) {
		t.Parallel()

		_, erro := logicaTeste.Matéria.Pegar(entidades.NovoID())
		if erro == nil || !erro.ÉPadrão(ErroMatériaNãoEncontrada) {
			t.Fatalf("Esperava: %v\nChegou: %v", ErroMatériaNãoEncontrada, erro)
		}
	})

	t.Run("TimeOut", func(t *testing.T) {
		t.Parallel()

		_, erro := matériaBDTimeOut.Pegar(entidades.NovoID())
		if erro == nil || !erro.ÉPadrão(ErroPegarMatéria) {
			t.Fatalf("Esperava um erro de timeout: %v", erro)
		}
	})
}

func TestDeletarMatéria(t *testing.T) {
	t.Parallel()

	t.Run("OKAY", func(t *testing.T) {
		t.Parallel()

		nome, ch, créditos, tipo := criarMatériaAleatória()
		adicionarMatéria(t, nome, ch, créditos, tipo, []id{})
	})

	t.Run("MatériaNãoEncontrada", func(t *testing.T) {
		t.Parallel()

		erro := logicaTeste.Matéria.Deletar(entidades.NovoID())
		if erro == nil || !erro.ÉPadrão(ErroMatériaNãoEncontrada) {
			t.Fatalf("Esperava: %v\nChegou: %v", ErroMatériaNãoEncontrada, erro)
		}
	})

	t.Run("TimeOut", func(t *testing.T) {
		t.Parallel()

		erro := matériaBDTimeOut.Deletar(entidades.NovoID())
		if erro == nil || !erro.ÉPadrão(ErroDeletarMatéria) {
			t.Fatalf("Esperava um erro de timeout: %v", erro)
		}
	})

	t.Run("BDInválido", func(t *testing.T) {
		t.Parallel()

		nome, ch, créditos, tipo := criarMatériaAleatória()
		id := adicionarMatéria(t, nome, ch, créditos, tipo, []id{})

		erro := matériaBDInválido.Deletar(id)
		if erro == nil || !erro.ÉPadrão(ErroDeletarMatéria) {
			t.Fatalf("Esperava um erro de deletar matéria: %v", erro)
		}
	})
}
