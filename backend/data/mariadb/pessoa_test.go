package mariadb

import (
	"reflect"
	"regexp"
	"testing"

	"thiagofelipe.com.br/sistema-faculdade-backend/aleatorio"
	. "thiagofelipe.com.br/sistema-faculdade-backend/data"
	"thiagofelipe.com.br/sistema-faculdade-backend/entidades"
)

func criarPessoaAleatória() *entidades.Pessoa {
	dataAgora := entidades.DataAtual()

	var pessoa = &entidades.Pessoa{
		ID:               entidades.NovoID(),
		Nome:             aleatorio.Palavra(aleatorio.Número(tamanhoMáximoPalavra) + 1),
		CPF:              aleatorio.CPF(),
		DataDeNascimento: dataAgora,
		Senha:            "Senha",
	}

	return pessoa
}

func adicionarPessoa(pessoa *entidades.Pessoa, t *testing.T) {
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

func removerPessoa(t *testing.T, id entidades.ID) {
	erro := pessoaBD.Deletar(id)
	if erro != nil {
		t.Fatalf("Erro ao tentar deletar a pessoa teste: %v", erro.Error())
	}
}

func TestInserirPessoa(t *testing.T) {
	pessoaTeste := criarPessoaAleatória()

	t.Run("OKAY", func(t *testing.T) {
		adicionarPessoa(pessoaTeste, t)
	})

	t.Run("Duplicado/ID", func(t *testing.T) {
		const texto = `Duplicate entry.*PRIMARY`
		padrão := regexp.MustCompile(texto)

		adicionarPessoa(pessoaTeste, t)

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
		const texto = `Duplicate entry.*CPF`
		padrão := regexp.MustCompile(texto)

		adicionarPessoa(pessoaTeste, t)

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
	pessoaTeste1 := criarPessoaAleatória()
	pessoaTeste2 := criarPessoaAleatória()

	t.Run("OKAY", func(t *testing.T) {
		adicionarPessoa(pessoaTeste1, t)
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
		const texto = `Duplicate entry.*CPF`
		padrão := regexp.MustCompile(texto)

		adicionarPessoa(pessoaTeste1, t)
		adicionarPessoa(pessoaTeste2, t)

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

func TestPegarPessoa(t *testing.T) {
	t.Run("OKAY", func(t *testing.T) {
		pessoaTeste := criarPessoaAleatória()

		adicionarPessoa(pessoaTeste, t)

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

func TestPegarPessoaPorCPF(t *testing.T) {
	t.Run("OKAY", func(t *testing.T) {
		pessoaTeste := criarPessoaAleatória()

		adicionarPessoa(pessoaTeste, t)

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
	pessoaTeste := criarPessoaAleatória()

	t.Run("OKAY", func(t *testing.T) {
		adicionarPessoa(pessoaTeste, t)

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
		removerPessoa(t, pessoaTeste.ID)

		_, erro := pessoaBD.Pegar(pessoaTeste.ID)
		if erro == nil || !erro.ÉPadrão(ErroPessoaNãoEncontrada) {
			t.Fatalf(
				"Deveria retonar um erro de pessoa não encontrada, retonou %s",
				erro,
			)
		}
	})

	t.Run("TabelaInválida", func(t *testing.T) {
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
