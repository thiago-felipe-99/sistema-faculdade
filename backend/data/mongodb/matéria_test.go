//nolint: dupl
package mongodb

import (
	"reflect"
	"sort"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"thiagofelipe.com.br/sistema-faculdade-backend/aleatorio"
	"thiagofelipe.com.br/sistema-faculdade-backend/data"
	"thiagofelipe.com.br/sistema-faculdade-backend/entidades"
)

func criarMatériaAleatória() *matéria {
	préRequisitos := []id{}
	índice := uint(0)

	for ; índice <= aleatorio.Número(tamanhoMáximoPréRequisito); índice++ {
		préRequisitos = append(préRequisitos, entidades.NovoID())
	}

	carga := aleatorio.Número(cargaHoráriaMáxima) + 1

	return &matéria{
		ID:                  entidades.NovoID(),
		Nome:                aleatorio.Palavra(aleatorio.Número(tamanhoMáximoPalavra) + 1),
		CargaHoráriaSemanal: time.Hour * time.Duration(carga),
		Créditos:            float32(carga),
		PréRequisitos:       préRequisitos,
		Tipo:                aleatorio.Palavra(aleatorio.Número(tamanhoMáximoPalavra) + 1),
	}
}

func criarMatériasAleatórias(quantidade uint) []matéria {
	matérias := []matéria{}

	for índice := uint(0); índice <= quantidade; índice++ {
		matérias = append(matérias, *criarMatériaAleatória())
	}

	return matérias
}

func adicionarMatéria(t *testing.T, matéria *matéria) id {
	t.Helper()

	erro := matériaBD.Inserir(matéria)
	if erro != nil {
		t.Fatalf("Erro ao inserir a matéria no banco de dados: %v", erro)
	}

	matériaSalva, erro := matériaBD.Pegar(matéria.ID)
	if erro != nil {
		t.Fatalf("Erro ao pegar a matéria no banco de dados: %v", erro)
	}

	if !reflect.DeepEqual(matéria, matériaSalva) {
		t.Fatalf(
			"Erro ao salvar a matéria no banco de dados, queria: %v, chegou: %v",
			matéria,
			matériaSalva,
		)
	}

	t.Cleanup(func() {
		removerMatéria(t, matéria.ID)
	})

	return matéria.ID
}

func removerMatéria(t *testing.T, id id) {
	t.Helper()

	if erro := matériaBD.Deletar(id); erro != nil {
		t.Fatalf("Erro ao tentar deletar a matéria teste: %v", erro.Error())
	}
}

func TestInserirMatéria(t *testing.T) {
	t.Parallel()

	t.Run("OKAY", func(t *testing.T) {
		t.Parallel()
		adicionarMatéria(t, criarMatériaAleatória())
	})

	t.Run("Duplicado/ID", func(t *testing.T) {
		t.Parallel()

		matériaTeste := criarMatériaAleatória()

		adicionarMatéria(t, matériaTeste)

		erro := matériaBD.Inserir(matériaTeste)
		if erro == nil || erro.ErroExterno == nil {
			t.Fatalf("Não foi enviado erro do sistema")
		}

		if !mongo.IsDuplicateKeyError(erro.ErroExterno) {
			t.Fatalf(
				"Erro ao inserir a matéria queria: %v, chegou: %v",
				"Erro de id duplicado",
				erro.ErroExterno,
			)
		}
	})

	t.Run("TimeOut", func(t *testing.T) {
		t.Parallel()
		erro := matériaBDInválido.Inserir(criarMatériaAleatória())
		if erro == nil || !mongo.IsTimeout(erro.ErroExterno) {
			t.Fatalf("Esperava um erro de Timeout, chegou: %v", erro)
		}
	})
}

func TestAtualizarMatéria(t *testing.T) {
	t.Parallel()

	id := adicionarMatéria(t, criarMatériaAleatória())
	novaMatéria := criarMatériaAleatória()
	novaMatéria.ID = id

	t.Run("OKAY", func(t *testing.T) {
		t.Parallel()

		erro := matériaBD.Atualizar(id, novaMatéria)
		if erro != nil {
			t.Fatalf("Não deveria ter erro ao atualizar matéria: %v", erro)
		}

		matériaSalva, erro := matériaBD.Pegar(id)
		if erro != nil {
			t.Fatalf("Erro ao pegar a matéria no banco de dados: %v", erro)
		}

		if !reflect.DeepEqual(novaMatéria, matériaSalva) {
			t.Fatalf(
				"Erro ao salvar a matéria no banco de dados, queria: %v, chegou: %v",
				novaMatéria,
				matériaSalva,
			)
		}
	})

	t.Run("TimeOut", func(t *testing.T) {
		t.Parallel()
		erro := matériaBDInválido.Atualizar(id, novaMatéria)
		if erro == nil || !mongo.IsTimeout(erro.ErroExterno) {
			t.Fatalf("Esperava um erro de Timeout, chegou: %v", erro)
		}
	})
}

func TestPegarMatéria(t *testing.T) {
	t.Parallel()

	t.Run("OKAY", func(t *testing.T) {
		t.Parallel()
		adicionarMatéria(t, criarMatériaAleatória())
	})

	t.Run("MatériaNãoEcontrada", func(t *testing.T) {
		t.Parallel()
		_, erro := matériaBD.Pegar(entidades.NovoID())
		if erro == nil || !erro.ÉPadrão(data.ErroMatériaNãoEncontrada) {
			t.Fatalf(
				"Erro ao pegar a matéria, queria: %v, chegou: %v",
				data.ErroMatériaNãoEncontrada,
				erro,
			)
		}
	})

	t.Run("TimeOut", func(t *testing.T) {
		t.Parallel()
		_, erro := matériaBDInválido.Pegar(entidades.NovoID())
		if erro == nil || !mongo.IsTimeout(erro.ErroExterno) {
			t.Fatalf("Esperava um erro de Timeout, chegou: %v", erro)
		}
	})
}

func TestPegarMatériaPréRequisitos(t *testing.T) {
	t.Parallel()

	t.Run("OKAY", func(t *testing.T) {
		t.Parallel()
		id := adicionarMatéria(t, criarMatériaAleatória())

		matéria, erro := matériaBD.Pegar(id)
		if erro != nil {
			t.Fatalf("Não esperava um erro ao pegar a matéria: %v", erro)
		}

		ids, erro := matériaBD.PegarPréRequisitos(id)
		if erro != nil {
			t.Fatalf("Não esperava um erro ao pegar as IDs: %v", erro)
		}

		if !reflect.DeepEqual(ids, matéria.PréRequisitos) {
			t.Fatalf("Esperava: %v\nChegou: %v", ids, matéria.ID)
		}
	})

	t.Run("MatériaNãoEcontrada", func(t *testing.T) {
		t.Parallel()
		_, erro := matériaBD.PegarPréRequisitos(entidades.NovoID())
		if erro == nil || !erro.ÉPadrão(data.ErroMatériaNãoEncontrada) {
			t.Fatalf(
				"Erro ao pegar a matéria, queria: %v, chegou: %v",
				data.ErroMatériaNãoEncontrada,
				erro,
			)
		}
	})

	t.Run("TimeOut", func(t *testing.T) {
		t.Parallel()
		_, erro := matériaBDInválido.PegarPréRequisitos(entidades.NovoID())
		if erro == nil || !mongo.IsTimeout(erro.ErroExterno) {
			t.Fatalf("Esperava um erro de Timeout, chegou: %v", erro)
		}
	})
}

//nolint: funlen, cyclop
func TestPegarMúltiplos(t *testing.T) {
	t.Parallel()

	t.Run("OKAY", func(t *testing.T) {
		t.Parallel()

		matérias := criarMatériasAleatórias(aleatorio.Número(tamanhoMáximoMatérias))

		erro := matériaBD.inserirMúltiplas(&matérias)
		if erro != nil {
			t.Fatalf("Erro ao inserir múltiplas matérias: %v", erro)
		}

		ids := []id{}
		for _, matéria := range matérias {
			ids = append(ids, matéria.ID)
		}

		defer func() {
			erro := matériaBD.deletarMúltiplas(ids)
			if erro != nil {
				t.Fatalf("Erro ao deletar múltiplas matérias: %v", erro)
			}
		}()

		matériasSalvas, erro := matériaBD.PegarMúltiplos(append(ids, ids[0]))
		if erro != nil {
			t.Fatalf("Erro ao verificar se existe matérias: %v", erro)
		}

		sort.Slice(matérias, func(i, j int) bool {
			return matérias[i].ID.ID() < matérias[j].ID.ID()
		})

		sort.Slice(matériasSalvas, func(i, j int) bool {
			return matériasSalvas[i].ID.ID() < matériasSalvas[j].ID.ID()
		})

		if !reflect.DeepEqual(matérias, matériasSalvas) {
			t.Fatalf("Esperava: %v, chegou: %v", matérias, matériasSalvas)
		}
	})

	t.Run("TamanhoInválido", func(t *testing.T) {
		t.Parallel()

		_, erro := matériaBD.PegarMúltiplos([]id{})
		if erro == nil || !erro.ÉPadrão(data.ErroIDsTamanho) {
			t.Fatalf("Esperava: %v\nChegou: %v", data.ErroIDsTamanho, erro)
		}
	})

	t.Run("TimeOut", func(t *testing.T) {
		t.Parallel()

		_, erro := matériaBDInválido.PegarMúltiplos([]id{entidades.NovoID()})
		if erro == nil || !mongo.IsTimeout(erro.ErroExterno) {
			t.Fatalf("Esperava um erro de Timeout, chegou: %v", erro)
		}
	})

	t.Run("IDsDiferentes", func(t *testing.T) {
		t.Parallel()

		matérias := criarMatériasAleatórias(aleatorio.Número(tamanhoMáximoMatérias))

		erro := matériaBD.inserirMúltiplas(&matérias)
		if erro != nil {
			t.Fatalf("Erro ao inserir múltiplas matérias: %v", erro)
		}

		ids := []id{}
		for _, matéria := range matérias {
			ids = append(ids, matéria.ID)
		}

		defer func() {
			erro := matériaBD.deletarMúltiplas(ids)
			if erro != nil {
				t.Fatalf("Erro ao deletar múltiplas matérias: %v", erro)
			}
		}()

		idsErrados := make([]id, len(ids))
		copy(idsErrados, ids)

		// nolint: makezero
		idsErrados = append(idsErrados, entidades.NovoID(), entidades.NovoID())

		matériasSalvas, erro := matériaBD.PegarMúltiplos(idsErrados)
		if erro != nil {
			t.Fatalf("Não esperava um erro ao verificar se matérias existe: %v", erro)
		}

		sort.Slice(matérias, func(i, j int) bool {
			return matérias[i].ID.ID() < matérias[j].ID.ID()
		})

		sort.Slice(matériasSalvas, func(i, j int) bool {
			return matériasSalvas[i].ID.ID() < matériasSalvas[j].ID.ID()
		})

		if !reflect.DeepEqual(matérias, matériasSalvas) {
			t.Fatalf("Esperava os ids: %v, chegou: %v", matérias, matériasSalvas)
		}
	})
}

func TestDeletarMatéria(t *testing.T) {
	t.Parallel()

	t.Run("OKAY", func(t *testing.T) {
		t.Parallel()
		id := adicionarMatéria(t, criarMatériaAleatória())
		erro := matériaBD.Deletar(id)
		if erro != nil {
			t.Fatalf("Não esperava um erro ao deletar matéria: %v", erro)
		}

		_, erro = matériaBD.Pegar(id)
		if erro == nil || !erro.ÉPadrão(data.ErroMatériaNãoEncontrada) {
			t.Fatalf("Esperava: %v, chegou: %v", data.ErroMatériaNãoEncontrada, erro)
		}
	})

	t.Run("TimeOut", func(t *testing.T) {
		t.Parallel()
		erro := matériaBDInválido.Deletar(entidades.NovoID())
		if erro == nil || !mongo.IsTimeout(erro.ErroExterno) {
			t.Fatalf("Esperava um erro de Timeout, chegou: %v", erro)
		}
	})
}

//nolint: funlen, cyclop
func TestMúltiplasMatérias(t *testing.T) {
	t.Parallel()

	t.Run("OKAY", func(t *testing.T) {
		t.Parallel()

		matérias := criarMatériasAleatórias(aleatorio.Número(tamanhoMáximoMatérias))

		erro := matériaBD.inserirMúltiplas(&matérias)
		if erro != nil {
			t.Errorf("Não esperava erro ao inserir múltiplas matérias: %v", erro)
		}

		ids := []id{}

		for _, matéria := range matérias {
			ids = append(ids, matéria.ID)
		}

		t.Run("Pegar", func(t *testing.T) {
			for _, matéria := range matérias {
				matéria := matéria
				t.Run("Pegar/"+matéria.ID.String(), func(t *testing.T) {
					t.Parallel()

					matériaSalva, erro := matériaBD.Pegar(matéria.ID)
					if erro != nil {
						t.Fatalf("Não esperava erro ao pegar a matéria: %v", erro)
					}

					if !reflect.DeepEqual(matériaSalva, &matéria) {
						t.Fatalf("Esperava matéria: %v\nChegou: %v", matéria, matériaSalva)
					}
				})
			}
		})

		erro = matériaBD.deletarMúltiplas(ids)
		if erro != nil {
			t.Errorf("Não esperava erro ao deletar múltiplas matérias: %v", erro)
		}
	})

	t.Run("TimeOut/Inserir", func(t *testing.T) {
		t.Parallel()

		matérias := criarMatériasAleatórias(aleatorio.Número(tamanhoMáximoMatérias))

		erro := matériaBDInválido.inserirMúltiplas(&matérias)
		if erro == nil || !mongo.IsTimeout(erro.ErroExterno) {
			t.Fatalf("Esperava um erro de Timeout, chegou: %v", erro)
		}
	})

	t.Run("TimeOut/Deletar", func(t *testing.T) {
		t.Parallel()

		erro := matériaBDInválido.deletarMúltiplas([]id{entidades.NovoID()})
		if erro == nil || !mongo.IsTimeout(erro.ErroExterno) {
			t.Fatalf("Esperava um erro de Timeout, chegou: %v", erro)
		}
	})
}
