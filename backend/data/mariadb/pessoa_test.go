package mariadb

import (
	"reflect"
	"regexp"
	"testing"

	"thiagofelipe.com.br/sistema-faculdade-backend/aleatorio"
	. "thiagofelipe.com.br/sistema-faculdade-backend/data"
	"thiagofelipe.com.br/sistema-faculdade-backend/entidades"
)

func criarPessoaAleatória() *pessoa {
	dataAgora := entidades.DataAtual()

	pessoa := &pessoa{
		ID:               entidades.NovoID(),
		Nome:             aleatorio.Palavra(aleatorio.Número(tamanhoMáximoPalavra) + 1),
		CPF:              aleatorio.CPF(),
		DataDeNascimento: dataAgora,
		Senha:            aleatorio.Palavra(aleatorio.Número(tamanhoMáximoPalavra) + 1),
	}

	return pessoa
}

func adicionarPessoa(t *testing.T, pessoa *pessoa) {
	t.Helper()

	erro := pessoaBD.Inserir(pessoa)
	if erro != nil {
		t.Fatalf("Erro ao inserir a pessoa no banco de dados: %s", erro.Error())
	}

	pessoaSalva, erro := pessoaBD.Pegar(pessoa.ID)
	if erro != nil {
		t.Fatalf("Erro ao pegar a pessoa no banco de dados: %s", erro.Error())
	}

	if !reflect.DeepEqual(pessoa, pessoaSalva) {
		t.Fatalf(
			"Erro ao salvar a pessoa no banco de dados, queria %v, chegou %v",
			pessoa,
			pessoaSalva,
		)
	}

	t.Cleanup(func() {
		removerPessoa(t, pessoa.ID)
	})
}

func removerPessoa(t *testing.T, id id) {
	t.Helper()

	erro := pessoaBD.Deletar(id)
	if erro != nil {
		t.Fatalf("Erro ao tentar deletar a pessoa teste: %v", erro.Error())
	}
}

//nolint: funlen
func TestInserirPessoa(t *testing.T) {
	t.Parallel()

	t.Run("OKAY", func(t *testing.T) {
		t.Parallel()

		pessoaTeste := criarPessoaAleatória()

		adicionarPessoa(t, pessoaTeste)
	})

	t.Run("Duplicado/ID", func(t *testing.T) {
		t.Parallel()

		const texto = `Duplicate entry.*PRIMARY`
		padrão := regexp.MustCompile(texto)

		pessoaTeste := criarPessoaAleatória()

		adicionarPessoa(t, pessoaTeste)

		pesssoaDuplicada := pessoaTeste
		pesssoaDuplicada.CPF = aleatorio.CPF()

		erro := pessoaBD.Inserir(pessoaTeste)
		if erro == nil || erro.ErroExterno == nil {
			t.Fatalf("Não foi enviado erro do sistema")
		}

		if !padrão.MatchString(erro.ErroExterno.Error()) {
			t.Fatalf(
				"Erro ao inserir a pessoa queria: %s, chegou %s",
				texto,
				erro.ErroExterno.Error(),
			)
		}
	})

	t.Run("Duplicado/CPF", func(t *testing.T) {
		t.Parallel()

		const texto = `Duplicate entry.*CPF`
		padrão := regexp.MustCompile(texto)

		pessoaTeste := criarPessoaAleatória()

		adicionarPessoa(t, pessoaTeste)

		pesssoaDuplicada := pessoaTeste
		pesssoaDuplicada.ID = entidades.NovoID()

		erro := pessoaBD.Inserir(pessoaTeste)
		if erro == nil || erro.ErroExterno == nil {
			t.Fatalf("Não foi enviado erro do sistema")
		}

		if !padrão.MatchString(erro.ErroExterno.Error()) {
			t.Fatalf(
				"Erro ao inserir a pessoa queria: %s, chegou %s",
				texto,
				erro.ErroExterno.Error(),
			)
		}
	})
}

func TestAtualizarPessoa(t *testing.T) {
	t.Parallel()

	t.Run("OKAY", func(t *testing.T) {
		t.Parallel()

		pessoaTeste1 := criarPessoaAleatória()
		pessoaTeste2 := criarPessoaAleatória()

		adicionarPessoa(t, pessoaTeste1)
		idBuffer := pessoaTeste2.ID
		pessoaTeste2.ID = pessoaTeste1.ID
		defer func() { pessoaTeste2.ID = idBuffer }()

		erro := pessoaBD.Atualizar(pessoaTeste1.ID, pessoaTeste2)
		if erro != nil {
			t.Fatalf("Erro ao atualizar a pessoa teste: %s", erro.Error())
		}

		pessoaSalva, erro := pessoaBD.Pegar(pessoaTeste1.ID)
		if erro != nil {
			t.Fatalf("Erro ao pegar a pessoa no banco de dados: %s", erro.Error())
		}

		if !reflect.DeepEqual(pessoaTeste2, pessoaSalva) {
			t.Fatalf(
				"Erro ao salvar a pessoa no banco de dados, queria %v, chegou %v",
				pessoaTeste1,
				pessoaSalva,
			)
		}
	})

	t.Run("Duplicado/CPF", func(t *testing.T) {
		t.Parallel()

		const texto = `Duplicate entry.*CPF`
		padrão := regexp.MustCompile(texto)

		pessoaTeste1 := criarPessoaAleatória()
		pessoaTeste2 := criarPessoaAleatória()

		adicionarPessoa(t, pessoaTeste1)
		adicionarPessoa(t, pessoaTeste2)

		erro := pessoaBD.Atualizar(pessoaTeste1.ID, pessoaTeste2)
		if erro == nil || erro.ErroExterno == nil {
			t.Fatalf("Não foi enviado erro do sistema")
		}

		if !padrão.MatchString(erro.ErroExterno.Error()) {
			t.Fatalf(
				"Erro ao inserir a pessoa queria: %s, chegou %s",
				texto,
				erro.ErroExterno.Error(),
			)
		}
	})
}

//nolint: dupl
func TestPegarPessoa(t *testing.T) {
	t.Parallel()

	t.Run("OKAY", func(t *testing.T) {
		t.Parallel()

		pessoaTeste := criarPessoaAleatória()
		adicionarPessoa(t, pessoaTeste)

		pessoaSalva, erro := pessoaBD.Pegar(pessoaTeste.ID)
		if erro != nil {
			t.Fatalf("Erro ao pegar a pessoa no banco de dados: %s", erro.Error())
		}

		if !reflect.DeepEqual(pessoaTeste, pessoaSalva) {
			t.Fatalf(
				"Erro ao salvar a pessoa no banco de dados, queria %v, chegou %v",
				pessoaTeste,
				pessoaSalva,
			)
		}
	})

	t.Run("PessoaNãoEncontrada", func(t *testing.T) {
		t.Parallel()

		_, erro := pessoaBD.Pegar(entidades.NovoID())

		if erro == nil || erro.ErroExterno == nil {
			t.Fatalf("Não foi enviado erro do sistema")
		}

		if !erro.ÉPadrão(ErroPessoaNãoEncontrada) {
			t.Fatalf(
				"Erro ao pegar pessoa no banco de dados, queria %v, chegou %v",
				ErroPessoaNãoEncontrada.Error(),
				erro.Error(),
			)
		}
	})

	t.Run("TabelaInválida", func(t *testing.T) {
		t.Parallel()

		const texto = `Table .* doesn't exist`
		padrão := regexp.MustCompile(texto)

		_, erro := pessoaBDInválido.Pegar(entidades.NovoID())

		if erro == nil || erro.ErroExterno == nil {
			t.Fatalf("Não foi enviado erro do sistema")
		}

		if !padrão.MatchString(erro.ErroExterno.Error()) {
			t.Fatalf(
				"Erro ao pegar pessoa no banco de dados, queria %v, chegou %v",
				texto,
				erro.ErroExterno.Error(),
			)
		}
	})
}

//nolint: dupl
func TestPegarPessoaPorCPF(t *testing.T) {
	t.Parallel()

	t.Run("OKAY", func(t *testing.T) {
		t.Parallel()

		pessoaTeste := criarPessoaAleatória()
		adicionarPessoa(t, pessoaTeste)

		pessoaSalva, erro := pessoaBD.PegarPorCPF(pessoaTeste.CPF)
		if erro != nil {
			t.Fatalf("Erro ao pegar a pessoa no banco de dados: %s", erro.Error())
		}

		if !reflect.DeepEqual(pessoaTeste, pessoaSalva) {
			t.Fatalf(
				"Erro ao salvar a pessoa no banco de dados, queria %v, chegou %v",
				pessoaTeste,
				pessoaSalva,
			)
		}
	})

	t.Run("PessoaNãoEncontrada", func(t *testing.T) {
		t.Parallel()

		_, erro := pessoaBD.PegarPorCPF(aleatorio.CPF())

		if erro == nil || erro.ErroExterno == nil {
			t.Fatalf("Não foi enviado erro do sistema")
		}

		if !erro.ÉPadrão(ErroPessoaNãoEncontrada) {
			t.Fatalf(
				"Erro ao pegar pessoa no banco de dados, queria %v, chegou %v",
				ErroPessoaNãoEncontrada.Error(),
				erro.Error(),
			)
		}
	})

	t.Run("TabelaInválida", func(t *testing.T) {
		t.Parallel()

		const texto = `Table .* doesn't exist`
		padrão := regexp.MustCompile(texto)

		_, erro := pessoaBDInválido.PegarPorCPF(aleatorio.CPF())

		if erro == nil || erro.ErroExterno == nil {
			t.Fatalf("Não foi enviado erro do sistema")
		}

		if !padrão.MatchString(erro.ErroExterno.Error()) {
			t.Fatalf(
				"Erro ao pegar pessoa no banco de dados, queria %v, chegou %v",
				texto,
				erro.ErroExterno.Error(),
			)
		}
	})
}

func TestDeletarPessoa(t *testing.T) {
	t.Parallel()

	pessoaTeste := criarPessoaAleatória()

	t.Run("OKAY", func(t *testing.T) {
		t.Parallel()

		adicionarPessoa(t, pessoaTeste)

		removerPessoa(t, pessoaTeste.ID)

		_, erro := pessoaBD.Pegar(pessoaTeste.ID)
		if erro == nil || !erro.ÉPadrão(ErroPessoaNãoEncontrada) {
			t.Fatalf(
				"Deveria retonar um erro de pessoa não encontrada, retornou %s",
				erro,
			)
		}
	})

	t.Run("PessoaNãoEncontrada", func(t *testing.T) {
		t.Parallel()

		id := entidades.NovoID()

		removerPessoa(t, id)

		_, erro := pessoaBD.Pegar(id)
		if erro == nil || !erro.ÉPadrão(ErroPessoaNãoEncontrada) {
			t.Fatalf(
				"Deveria retonar um erro de pessoa não encontrada, retonou %s",
				erro,
			)
		}
	})

	t.Run("TabelaInválida", func(t *testing.T) {
		t.Parallel()

		const texto = `Table .* doesn't exist`
		padrão := regexp.MustCompile(texto)

		erro := pessoaBDInválido.Deletar(pessoaTeste.ID)

		if erro == nil || erro.ErroExterno == nil {
			t.Fatalf("Não foi enviado erro do sistema")
		}

		if !padrão.MatchString(erro.ErroExterno.Error()) {
			t.Fatalf(
				"Erro ao pegar pessoa no banco de dados, queria %v, chegou %v",
				texto,
				erro.ErroExterno.Error(),
			)
		}
	})
}
