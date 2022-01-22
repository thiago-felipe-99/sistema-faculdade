package mongodb

import (
	"reflect"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"thiagofelipe.com.br/sistema-faculdade-backend/aleatorio"
	"thiagofelipe.com.br/sistema-faculdade-backend/data"
	"thiagofelipe.com.br/sistema-faculdade-backend/entidades"
)

func criarMatériaAleatória() *entidades.Matéria {
	var préRequisitos []entidades.ID

	var indice uint

	for indice = 0; indice <= aleatorio.Número(tamanhoMáximoPréRequisito); indice++ {
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

func TestDeletar(t *testing.T) {
	t.Parallel()
	t.Run("OKAY", func(t *testing.T) {
		t.Parallel()
		id := adicionarMatéria(t, criarMatériaAleatória())
		erro := matériaBD.Deletar(id)
		if erro != nil {
			t.Fatalf("Não esperava um erro ao deletar usuário: %v", erro)
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
