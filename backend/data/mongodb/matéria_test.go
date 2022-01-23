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

func criarMatériaAleatória() *entidades.Matéria {
	préRequisitos := []entidades.ID{}
	índice := uint(0)

	for ; índice <= aleatorio.Número(tamanhoMáximoPréRequisito); índice++ {
		préRequisitos = append(préRequisitos, entidades.NovoID())
	}

	ch := aleatorio.Número(cargaHoráriaMáxima) + 1

	matéria := &entidades.Matéria{
		ID:                  entidades.NovoID(),
		Nome:                aleatorio.Palavra(aleatorio.Número(tamanhoMáximoPalavra) + 1),
		CargaHoráriaSemanal: time.Hour * time.Duration(ch),
		Créditos:            float32(ch),
		PréRequisitos:       préRequisitos,
		Tipo:                aleatorio.Palavra(aleatorio.Número(tamanhoMáximoPalavra) + 1),
	}

	return matéria
}

func criarMatériasAleatórias(quantidade uint) *[]entidades.Matéria {
	matérias := []entidades.Matéria{}

	for índice := uint(0); índice <= quantidade; índice++ {
		matérias = append(matérias, *criarMatériaAleatória())
	}

	return &matérias
}

func adicionarMatéria(t *testing.T, matéria *entidades.Matéria) entidades.ID {
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

func removerMatéria(t *testing.T, id entidades.ID) {
	t.Helper()

	erro := matériaBD.Deletar(id)
	if erro != nil {
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
				"Erro de id fuplicado",
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

//nolint: funlen, gocognit, cyclop
func TestExisteMatérias(t *testing.T) {
	t.Parallel()

	t.Run("OKAY", func(t *testing.T) {
		t.Parallel()

		matérias := criarMatériasAleatórias(aleatorio.Número(tamanhoMáximoMatérias))

		erro := matériaBD.inserirMúltiplas(matérias)
		if erro != nil {
			t.Fatalf("Erro ao inserir múltiplas matérias: %v", erro)
		}

		ids := []entidades.ID{}
		for _, matéria := range *matérias {
			ids = append(ids, matéria.ID)
		}

		defer func() {
			erro := matériaBD.deletarMúltiplas(ids)
			if erro != nil {
				t.Fatalf("Erro ao deletar múltiplas matérias: %v", erro)
			}
		}()

		idsExiste, existeTodas, erro := matériaBD.Existe(append(ids, ids[0]))
		if erro != nil {
			t.Fatalf("Erro ao verificar se existe matérias: %v", erro)
		}

		if !existeTodas {
			t.Fatalf("Esperava existir todas as matérias: %v", erro)
		}

		if !reflect.DeepEqual(ids, idsExiste) {
			t.Fatalf("Esperava os ids: %v, chegou: %v", ids, idsExiste)
		}
	})

	t.Run("TamanhoInválido", func(t *testing.T) {
		t.Parallel()

		_, _, erro := matériaBD.Existe([]entidades.ID{})
		if erro == nil || !erro.ÉPadrão(data.ErroIDsTamanho) {
			t.Fatalf("Esperava: %v\nChegou: %v", data.ErroIDsTamanho, erro)
		}
	})

	t.Run("TimeOut", func(t *testing.T) {
		t.Parallel()

		_, _, erro := matériaBDInválido.Existe([]entidades.ID{entidades.NovoID()})
		if erro == nil || !mongo.IsTimeout(erro.ErroExterno) {
			t.Fatalf("Esperava um erro de Timeout, chegou: %v", erro)
		}
	})

	t.Run("IDsDiferentes", func(t *testing.T) {
		t.Parallel()

		matérias := criarMatériasAleatórias(aleatorio.Número(tamanhoMáximoMatérias))

		erro := matériaBD.inserirMúltiplas(matérias)
		if erro != nil {
			t.Fatalf("Erro ao inserir múltiplas matérias: %v", erro)
		}

		ids := []entidades.ID{}
		for _, matéria := range *matérias {
			ids = append(ids, matéria.ID)
		}

		defer func() {
			erro := matériaBD.deletarMúltiplas(ids)
			if erro != nil {
				t.Fatalf("Erro ao deletar múltiplas matérias: %v", erro)
			}
		}()

		idsErrados := append(ids, entidades.NovoID(), entidades.NovoID())

		idsExiste, existeTudo, erro := matériaBD.Existe(idsErrados)
		if erro != nil {
			t.Fatalf("Não esperava um erro ao verificar se matérias existe: %v", erro)
		}

		if existeTudo {
			t.Fatalf("Esperava não existe todos IDs")
		}

		sort.Slice(ids, func(i, j int) bool {
			return ids[i].ID() < ids[j].ID()
		})

		sort.Slice(idsExiste, func(i, j int) bool {
			return idsExiste[i].ID() < idsExiste[j].ID()
		})

		if !reflect.DeepEqual(ids, idsExiste) {
			t.Fatalf("Esperava os ids: %v, chegou: %v", ids, idsExiste)
		}
	})
}

func TestDeletar(t *testing.T) {
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
func TestMúltiplos(t *testing.T) {
	t.Parallel()

	t.Run("OKAY", func(t *testing.T) {
		t.Parallel()

		matérias := criarMatériasAleatórias(aleatorio.Número(tamanhoMáximoMatérias))

		erro := matériaBD.inserirMúltiplas(matérias)
		if erro != nil {
			t.Errorf("Não esperava erro ao inserir múltiplas matérias: %v", erro)
		}

		ids := []entidades.ID{}

		for _, matéria := range *matérias {
			ids = append(ids, matéria.ID)
		}

		t.Run("Pegar", func(t *testing.T) {
			for _, matéria := range *matérias {
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

		erro := matériaBDInválido.inserirMúltiplas(matérias)
		if erro == nil || !mongo.IsTimeout(erro.ErroExterno) {
			t.Fatalf("Esperava um erro de Timeout, chegou: %v", erro)
		}
	})

	t.Run("TimeOut/Deletar", func(t *testing.T) {
		t.Parallel()

		erro := matériaBDInválido.deletarMúltiplas([]entidades.ID{entidades.NovoID()})
		if erro == nil || !mongo.IsTimeout(erro.ErroExterno) {
			t.Fatalf("Esperava um erro de Timeout, chegou: %v", erro)
		}
	})
}
