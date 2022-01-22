package mariadb

import (
	"reflect"
	"testing"

	"thiagofelipe.com.br/sistema-faculdade-backend/aleatorio"
	"thiagofelipe.com.br/sistema-faculdade-backend/entidades"
)

func criarTurmasAlunoAleatório(idAluno entidades.ID) *[]entidades.TurmaAluno {
	turmas := make([]entidades.TurmaAluno, aleatorio.Número(matériasMáximas)+1)

	for i := range turmas {
		turmas[i] = entidades.TurmaAluno{
			IDAluno: idAluno,
			IDTurma: entidades.NovoID(),
			Status:  aleatorio.Palavra(aleatorio.Número(tamanhoMáximoPalavra) + 1),
		}
	}

	return &turmas
}

func criarAlunoAleatório(t *testing.T) *entidades.Aluno {
	pessoa := criarPessoaAleatória()
	adicionarPessoa(pessoa, t)

	curso := criarCursoAleatório()
	adiconarCurso(t, curso)

	id := entidades.NovoID()

	dataAgora := entidades.DataAtual()
	ano := int(aleatorio.Número(50))
	mês := int(aleatorio.Número(12))
	dia := int(aleatorio.Número(28))
	dataFutura := dataAgora.AddDate(ano, mês, dia)

	return &entidades.Aluno{
		ID:             id,
		Pessoa:         pessoa.ID,
		Curso:          curso.ID,
		Matrícula:      aleatorio.Palavra(aleatorio.Número(tamanhoMáximoMatrícula) + 1),
		DataDeIngresso: dataAgora,
		DataDeSaída:    dataFutura,
		Período:        aleatorio.Palavra(aleatorio.Número(tamanhoMáximoPalavra) + 1),
		Status:         aleatorio.Palavra(aleatorio.Número(tamanhoMáximoPalavra) + 1),
		Turmas:         *criarTurmasAlunoAleatório(id),
	}
}

func adicionarAluno(t *testing.T, aluno *entidades.Aluno) {
	erro := alunoBD.Inserir(aluno)
	if erro != nil {
		t.Fatalf("Erro ao inserir o aluno no banco de dados: %s", erro.Error())
	}

	alunoSalvo, erro := alunoBD.Pegar(aluno.ID)
	if erro != nil {
		t.Fatalf("Erro ao pegar o aluno no banco de dados: %s", erro.Error())
	}

	if !reflect.DeepEqual(aluno, alunoSalvo) {
		t.Fatalf(
			"Erro ao salvar o aluno no banco de dados, queria %v, chegou %v",
			aluno,
			alunoSalvo,
		)
	}

	t.Cleanup(func() {
		removerAluno(aluno.ID, t)
	})
}

func removerAluno(id entidades.ID, t *testing.T) {
	_, erro := alunoBD.BD.Exec("DELETE FROM AlunoTurma;")
	if erro != nil {
		t.Fatalf("Erro ao tentar deletar o aluno teste: %v", erro.Error())
	}

	_, erro = alunoBD.BD.Exec("DELETE FROM Aluno;")
	if erro != nil {
		t.Fatalf("Erro ao tentar deletar o aluno teste: %v", erro.Error())
	}
}

func TestInserirAluno(t *testing.T) {
	alunoTeste := criarAlunoAleatório(t)

	adicionarAluno(t, alunoTeste)
}