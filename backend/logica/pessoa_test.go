package logica

import (
	"reflect"
	"testing"
	"time"

	"thiagofelipe.com.br/sistema-faculdade-backend/aleatorio"
	"thiagofelipe.com.br/sistema-faculdade-backend/entidades"
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
	t.Parallel()

	nome, cpf, dataDeNascimento, senha := criarPessoaAleatória()

	t.Run("Okay", func(t *testing.T) {
		adicionarPessoa(nome, cpf, dataDeNascimento, senha, t)
	})

	t.Run("CPFInválido", func(t *testing.T) {
		pessoa, erro := logicaTeste.Pessoa.Criar(nome, "00000000001", dataDeNascimento, senha)
		if erro == nil || !erro.ÉPadrão(ErroCPFInválido) {
			if pessoa != nil {
				removerPessoa(pessoa.ID, t)
			}
			t.Fatalf("Queria: %v\nChegou: %v", ErroCPFInválido, erro)
		}
	})

	t.Run("BDInválido", func(t *testing.T) {
		_, erro := pessoaBDInválido.Criar(nome, cpf, dataDeNascimento, senha)
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

	t.Run("DataInválida", func(t *testing.T) {
		_, erro := pessoaDataInvalida.Criar(nome, cpf, dataDeNascimento, senha)
		if erro == nil || !erro.ÉPadrão(ErroCriarPessoa) {
			t.Fatalf("Esperava: %v\nChegou: %v", ErroCriarPessoa, erro)
		}
	})
}

func TestPegarPessoa(t *testing.T) {
	t.Parallel()

	nome, cpf, dataDeNascimento, senha := criarPessoaAleatória()

	t.Run("OKAY", func(t *testing.T) {
		pessoaCriada, erro := logicaTeste.Pessoa.Criar(nome, cpf, dataDeNascimento, senha)

		defer func(pessoaCriada *entidades.Pessoa, t *testing.T) {
			if pessoaCriada != nil {
				removerPessoa(pessoaCriada.ID, t)
			}
		}(pessoaCriada, t)

		if erro != nil {
			t.Fatalf("Esperava: %v, chegou: %v", nil, erro)
		}

		pessoaSalva, erro := logicaTeste.Pessoa.Pegar(pessoaCriada.ID)
		if erro != nil {
			t.Fatalf("Esperava: %v, chegou: %v", nil, erro)
		}

		if !reflect.DeepEqual(pessoaCriada, pessoaSalva) {
			t.Fatalf("Esperava: %v\nChegou: %v", pessoaCriada, pessoaSalva)
		}
	})

	t.Run("PessoaNãoExiste", func(t *testing.T) {
		_, erro := logicaTeste.Pessoa.Pegar(entidades.NovoID())
		if erro == nil || !erro.ÉPadrão(ErroPessoaNãoEncontrada) {
			t.Fatalf("Esperava: %v\nChegou: %v", ErroPessoaNãoEncontrada, erro)
		}
	})

	t.Run("BDInválido", func(t *testing.T) {
		_, erro := pessoaBDInválido.Pegar(entidades.NovoID())
		if erro == nil || !erro.ÉPadrão(ErroPegarPessoa) {
			t.Fatalf("Esperava: %v\nChegou: %v", ErroPegarPessoa, erro)
		}
	})
}

func TestVerificarSenha(t *testing.T) {
	t.Parallel()

	nome, cpf, dataDeNascimento, senha := criarPessoaAleatória()

	t.Run("OKAY", func(t *testing.T) {
		id := adicionarPessoa(nome, cpf, dataDeNascimento, senha, t)

		igual, erro := logicaTeste.Pessoa.VerificarSenha(senha, id)
		if erro != nil {
			t.Fatalf("Esperava: %v, chegou: %v", nil, erro)
		}

		if !igual {
			t.Fatalf("Esperava: %t, chegou: %t", true, igual)
		}
	})

	t.Run("PessoaNãoEncontrada", func(t *testing.T) {
		_, erro := logicaTeste.Pessoa.VerificarSenha(senha, entidades.NovoID())
		if erro == nil || !erro.ÉPadrão(ErroPessoaNãoEncontrada) {
			t.Fatalf("Esperava: %v\nChegou: %v", ErroPessoaNãoEncontrada, erro)
		}
	})

	t.Run("BDInválido", func(t *testing.T) {
		_, erro := pessoaBDInválido.VerificarSenha(senha, entidades.NovoID())
		if erro == nil || !erro.ÉPadrão(ErroVerificarSenha) {
			t.Fatalf("Esperava: %v\nChegou: %v", ErroVerificarSenha, erro)
		}
	})

	t.Run("SenhasDiferentes", func(t *testing.T) {
		id := adicionarPessoa(nome, cpf, dataDeNascimento, senha, t)

		igual, erro := logicaTeste.Pessoa.VerificarSenha("senhaInválida", id)
		if erro != nil {
			t.Fatalf("Esperava: %v, chegou: %v", nil, erro)
		}

		if igual {
			t.Fatalf("Esperava: %t, chegou: %t", true, igual)
		}
	})
}

func TestAtualizar(t *testing.T) {
	t.Parallel()

	nome1, cpf1, dataDeNascimento1, senha1 := criarPessoaAleatória()
	nome2, cpf2, dataDeNascimento2, senha2 := criarPessoaAleatória()

	t.Run("OKAY", func(t *testing.T) {
		id := adicionarPessoa(nome1, cpf1, dataDeNascimento1, senha1, t)

		pessoaAtualizada, erro := logicaTeste.Pessoa.Atualizar(
			id,
			nome2,
			cpf2,
			dataDeNascimento2,
			senha2,
		)
		if erro != nil {
			t.Fatalf("Esperava: %v, chegou: %v", nil, erro)
		}

		pessoaSalva, erro := logicaTeste.Pessoa.Pegar(id)
		if erro != nil {
			t.Fatalf("Esperava: %v, chegou: %v", nil, erro)
		}

		if !reflect.DeepEqual(pessoaAtualizada, pessoaSalva) {
			t.Fatalf("Esperava: %v\nChegou: %v", pessoaAtualizada, pessoaSalva)
		}
	})

	t.Run("PessoaNãoEncontrada", func(t *testing.T) {
		_, erro := logicaTeste.Pessoa.Atualizar(entidades.NovoID(), nome1, cpf1, dataDeNascimento1, senha1)
		if erro == nil || !erro.ÉPadrão(ErroPessoaNãoEncontrada) {
			t.Fatalf("Esperava: %v\nChegou: %v", ErroPessoaNãoEncontrada, erro)
		}
	})

	t.Run("BDInválido", func(t *testing.T) {
		_, erro := pessoaBDInválido.Atualizar(entidades.NovoID(), nome1, cpf1, dataDeNascimento1, senha1)
		if erro == nil || !erro.ÉPadrão(ErroAtualizarPessoa) {
			t.Fatalf("Esperava: %v\nChegou: %v", ErroAtualizarPessoa, erro)
		}
	})

	t.Run("CPFInválido", func(t *testing.T) {
		id := adicionarPessoa(nome1, cpf1, dataDeNascimento1, senha1, t)

		_, erro := logicaTeste.Pessoa.Atualizar(id, nome1, "00000000001", dataDeNascimento1, senha1)
		if erro == nil || !erro.ÉPadrão(ErroCPFInválido) {
			t.Fatalf("Esperava: %v\nChegou: %v", ErroCPFInválido, erro)
		}
	})

	t.Run("CPFJáExiste", func(t *testing.T) {
		id := adicionarPessoa(nome1, cpf1, dataDeNascimento1, senha1, t)
		adicionarPessoa(nome2, cpf2, dataDeNascimento2, senha2, t)

		_, erro := logicaTeste.Pessoa.Atualizar(id, nome1, cpf2, dataDeNascimento1, senha1)
		if erro == nil || !erro.ÉPadrão(ErroCPFExiste) {
			t.Fatalf("Esperava: %v\nChegou: %v", ErroCPFExiste, erro)
		}
	})

	t.Run("DataDeNascimentoInválida", func(t *testing.T) {
		id := adicionarPessoa(nome1, cpf1, dataDeNascimento1, senha1, t)

		dataAtual := entidades.DataAtual().AddDate(1, 0, 0)

		_, erro := logicaTeste.Pessoa.Atualizar(id, nome1, cpf1, dataAtual, senha1)
		if erro == nil || !erro.ÉPadrão(ErroDataDeNascimentoInválido) {
			t.Fatalf("Esperava: %v\nChegou: %v", ErroDataDeNascimentoInválido, erro)
		}
	})

	t.Run("SenhaInválida", func(t *testing.T) {
		id := adicionarPessoa(nome1, cpf1, dataDeNascimento1, senha1, t)

		_, erro := logicaTeste.Pessoa.Atualizar(id, nome1, cpf1, dataDeNascimento1, "senha")
		if erro == nil || !erro.ÉPadrão(ErroSenhaInválida) {
			t.Fatalf("Esperava: %v\nChegou: %v", ErroSenhaInválida, erro)
		}
	})

	t.Run("DataInválida", func(t *testing.T) {
		id := adicionarPessoa(nome1, cpf1, dataDeNascimento1, senha1, t)

		_, erro := pessoaDataInvalida.Atualizar(id, nome1, cpf1, dataDeNascimento1, senha1)
		if erro == nil || !erro.ÉPadrão(ErroAtualizarPessoa) {
			t.Fatalf("Esperava: %v\nChegou: %v", ErroAtualizarPessoa, erro)
		}
	})

	t.Run("DataInválida2", func(t *testing.T) {
		id := adicionarPessoa(nome1, cpf1, dataDeNascimento1, senha1, t)

		_, erro := pessoaDataInvalida2.Atualizar(id, nome1, "00000000000", dataDeNascimento1, senha1)
		if erro == nil || !erro.ÉPadrão(ErroAtualizarPessoa) {
			t.Fatalf("Esperava: %v\nChegou: %v", ErroAtualizarPessoa, erro)
		}
	})
}

func TestDeletar(t *testing.T) {
	t.Parallel()
	nome, cpf, dataDeNascimento, senha := criarPessoaAleatória()

	t.Run("OKAY", func(t *testing.T) {
		pessoa, erro := logicaTeste.Pessoa.Criar(nome, cpf, dataDeNascimento, senha)
		if erro != nil {
			t.Fatalf("Esperava: %v, chegou: %v", nil, erro)
		}

		erro = logicaTeste.Pessoa.Deletar(pessoa.ID)
		if erro != nil {
			t.Fatalf("Esperava: %v, chegou: %v", nil, erro)
		}
	})

	t.Run("PessoaNãoEncontrada", func(t *testing.T) {
		erro := logicaTeste.Pessoa.Deletar(entidades.NovoID())
		if erro == nil || !erro.ÉPadrão(ErroPessoaNãoEncontrada) {
			t.Fatalf("Esperava: %v\nChegou: %v", ErroPessoaNãoEncontrada, erro)
		}
	})

	t.Run("BDInválido", func(t *testing.T) {
		erro := pessoaBDInválido.Deletar(entidades.NovoID())
		if erro == nil || !erro.ÉPadrão(ErroDeletarPessoa) {
			t.Fatalf("Esperava: %v\nChegou: %v", ErroDeletarPessoa, erro)
		}
	})

	t.Run("DataInválida", func(t *testing.T) {
		id := adicionarPessoa(nome, cpf, dataDeNascimento, senha, t)

		erro := pessoaDataInvalida.Deletar(id)
		if erro == nil || !erro.ÉPadrão(ErroDeletarPessoa) {
			t.Fatalf("Esperava: %v\nChegou: %v", ErroDeletarPessoa, erro)
		}
	})
}
