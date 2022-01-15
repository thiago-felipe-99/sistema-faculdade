package logica

import (
	"reflect"
	"testing"
	"time"

	"thiagofelipe.com.br/sistema-faculdade-backend/aleatorio"
	"thiagofelipe.com.br/sistema-faculdade-backend/entidades"

	//nolint:revive, stylecheck
	. "thiagofelipe.com.br/sistema-faculdade-backend/logica/erros"
)

func criarPessoaAleatória() (string, string, time.Time, string) {
	nome := aleatorio.Palavra(aleatorio.Número(tamanhoMáximoDaPalavra) + 1)
	cpf := aleatorio.CPF()
	data := time.Now()
	senha := aleatorio.Senha()

	return nome, cpf, data, senha
}

func removerPessoa(id entidades.ID, t *testing.T) {
	erro := logicaTeste.Pessoa.Deletar(id)
	if erro != nil {
		t.Fatalf("Erro ao tentar deletar a pessoa teste: %v", erro.Traçado())
	}
}

func adicionarPessoa(
	nome string,
	cpf string,
	dataDeNascimento time.Time,
	senha string,
	t *testing.T,
) entidades.ID {
	pessoaCriada, erro := logicaTeste.Pessoa.Criar(nome, cpf, dataDeNascimento, senha)
	if erro != nil {
		t.Fatalf("Erro ao criar a pessoa: %s", erro.Traçado())
	}

	pessoaSalva, erro := logicaTeste.Pessoa.Pegar(pessoaCriada.ID)
	if erro != nil {
		t.Fatalf("Erro ao pegar a pessoa: %s", erro.Traçado())
	}

	if !reflect.DeepEqual(pessoaCriada, pessoaSalva) {
		t.Fatalf(
			"Erro ao criar a pessoa, queria %v, chegou %v",
			pessoaCriada,
			pessoaSalva,
		)
	}

	t.Cleanup(func() {
		removerPessoa(pessoaCriada.ID, t)
	})

	return pessoaCriada.ID
}

func TestCriarPessoa(t *testing.T) {

	nome, cpf, dataDeNascimento, senha := criarPessoaAleatória()
	t.Logf("Senha: %s", senha)

	t.Run("Okay", func(t *testing.T) {
		adicionarPessoa(nome, cpf, dataDeNascimento, senha, t)
	})

	t.Run("CPFInválido", func(t *testing.T) {
		_, erro := logicaTeste.Pessoa.Criar(nome, "00000000001", dataDeNascimento, senha)
		if erro == nil || !erro.ÉPadrão(ErroCPFInválido) {
			t.Fatalf("Queria: %v\nChegou: %v", ErroCPFInválido, erro)
		}
	})

	t.Run("BDInválido", func(t *testing.T) {
		_, erro := pessoaInválida.Criar(nome, cpf, dataDeNascimento, senha)
		if erro == nil || !erro.ÉPadrão(ErroCriarPessoa) {
			t.Fatalf("Queria: %v\nChegou: %v", ErroCriarPessoa, erro)
		}
	})

	t.Run("CPFJáExiste", func(t *testing.T) {
		adicionarPessoa(nome, cpf, dataDeNascimento, senha, t)

		_, erro := logicaTeste.Pessoa.Criar(nome, cpf, dataDeNascimento, senha)
		if erro == nil || !erro.ÉPadrão(ErroCPFExiste) {
			t.Fatalf("Queria: %v\nChegou: %v", ErroCPFExiste, erro)
		}
	})

	t.Run("DataDeNascimentoInválida", func(t *testing.T) {
		dataAtual := entidades.DataAtual().AddDate(1, 0, 0)

		_, erro := logicaTeste.Pessoa.Criar(nome, cpf, dataAtual, senha)
		if erro == nil || !erro.ÉPadrão(ErroDataDeNascimentoInválido) {
			t.Fatalf("Queria: %v\nChegou: %v", ErroDataDeNascimentoInválido, erro)
		}
	})

	t.Run("SenhaInválida", func(t *testing.T) {
		senhas := []string{
			"A",
			"AAAAAAAAAAAAAAAAA",
			"AAAAAAAAA",
			"aaaaaaaaa",
			"AaAaAaAaA",
			"AaAaAaAa0",
			"AaAaAaA =0",
		}
		for _, senha := range senhas {
			t.Run(senha, func(t *testing.T) {
				_, erro := logicaTeste.Pessoa.Criar(nome, cpf, dataDeNascimento, senha)
				if erro == nil || !erro.ÉPadrão(ErroSenhaInválida) {
					t.Fatalf("Queria: %v\nChegou: %v", ErroSenhaInválida, erro)
				}
			})
		}
	})
}

func TestPegarPessoa(t *testing.T) {
	nome, cpf, dataDeNascimento, senha := criarPessoaAleatória()

	t.Run("OKAY", func(t *testing.T) {
		pessoaCriada, err := logicaTeste.Pessoa.Criar(nome, cpf, dataDeNascimento, senha)
		if err != nil {
			t.Fatalf("Esperava: %v, chegou: %v", nil, err)
		}

		pessoaSalva, err := logicaTeste.Pessoa.Pegar(pessoaCriada.ID)
		if err != nil {
			t.Fatalf("Esperava: %v, chegou: %v", nil, err)
		}

		if !reflect.DeepEqual(pessoaCriada, pessoaSalva) {
			t.Fatalf("Esperava: %v\nChegou: %v", pessoaCriada, pessoaSalva)
		}
	})

	t.Run("PessoaNãoExiste", func(t *testing.T) {
		_, err := logicaTeste.Pessoa.Pegar(entidades.NovoID())
		if err == nil || !err.ÉPadrão(ErroPessoaNãoEncontrada) {
			t.Fatalf("Esperava: %v\nChegou: %v", ErroPessoaNãoEncontrada, err)
		}
	})

	t.Run("BDInválido", func(t *testing.T) {
		_, err := pessoaInválida.Pegar(entidades.NovoID())
		if err == nil || !err.ÉPadrão(ErroPegarPessoa) {
			t.Fatalf("Esperava: %v\nChegou: %v", ErroPegarPessoa, err)
		}
	})
}

func TestVerificarSenha(t *testing.T) {
	nome, cpf, dataDeNascimento, senha := criarPessoaAleatória()

	t.Run("OKAY", func(t *testing.T) {
		id := adicionarPessoa(nome, cpf, dataDeNascimento, senha, t)

		igual, err := logicaTeste.Pessoa.VerificarSenha(senha, id)
		if err != nil {
			t.Fatalf("Esperava: %v, chegou: %v", nil, err)
		}

		if !igual {
			t.Fatalf("Esperava: %t, chegou: %t", true, igual)
		}
	})

	t.Run("PessoaNãoEncontrada", func(t *testing.T) {
		_, err := logicaTeste.Pessoa.VerificarSenha(senha, entidades.NovoID())
		if err == nil || !err.ÉPadrão(ErroPessoaNãoEncontrada) {
			t.Fatalf("Esperava: %v\nChegou: %v", ErroPessoaNãoEncontrada, err)
		}
	})

	t.Run("BDInválido", func(t *testing.T) {
		_, err := pessoaInválida.VerificarSenha(senha, entidades.NovoID())
		if err == nil || !err.ÉPadrão(ErroVerificarSenha) {
			t.Fatalf("Esperava: %v\nChegou: %v", ErroVerificarSenha, err)
		}
	})

	t.Run("SenhasDiferentes", func(t *testing.T) {
		id := adicionarPessoa(nome, cpf, dataDeNascimento, senha, t)

		igual, err := logicaTeste.Pessoa.VerificarSenha("senhaInválida", id)
		if err != nil {
			t.Fatalf("Esperava: %v, chegou: %v", nil, err)
		}

		if igual {
			t.Fatalf("Esperava: %t, chegou: %t", true, igual)
		}
	})

}
