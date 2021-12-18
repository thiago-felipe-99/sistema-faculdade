package mariadb

import (
	"math/rand"
	"reflect"
	"testing"
	"time"

	"thiagofelipe.com.br/sistema-faculdade/aleatorio"
	"thiagofelipe.com.br/sistema-faculdade/entidades"
)

func criarTurmasAlunoAleatório(idAluno entidades.ID) *[]entidades.TurmaAluno {
	turmas := make([]entidades.TurmaAluno, rand.Intn(MATÉRIAS_MÁXIMAS)+1)

	for i := range turmas {
		turmas[i] = entidades.TurmaAluno{
			IDAluno: idAluno,
			IDTurma: entidades.NovoID(),
			Status:  aleatorio.Palavra(rand.Intn(TAMANHO_MÁXIMO_PALAVRA) + 1),
		}
	}

	return &turmas
}

func criarAlunoAleatório(t *testing.T) *entidades.Aluno {
	pessoa := criarPessoaAleatória()
	adiconarPessoa(pessoa, t)

	curso := criarCursoAleatório()
	adiconarCurso(curso, t)

	id := entidades.NovoID()

	dataAgora := time.Now().UTC()
	dataAgora = dataAgora.Truncate(24 * time.Hour)
	dataFutura := dataAgora.AddDate(rand.Intn(50), rand.Intn(12), rand.Intn(28))

	return &entidades.Aluno{
		ID:             id,
		Pessoa:         pessoa.ID,
		Curso:          curso.ID,
		Matrícula:      aleatorio.Palavra(rand.Intn(TAMANHO_MÁXIMO_MATRÍCULA) + 1),
		DataDeIngresso: dataAgora,
		DataDeSaída:    dataFutura,
		Período:        aleatorio.Palavra(rand.Intn(TAMANHO_MÁXIMO_PALAVRA) + 1),
		Status:         aleatorio.Palavra(rand.Intn(TAMANHO_MÁXIMO_PALAVRA) + 1),
		Turmas:         *criarTurmasAlunoAleatório(id),
	}
}

func adicionarAluno(aluno *entidades.Aluno, t *testing.T) {
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
	alunoTest := criarAlunoAleatório(t)

	adicionarAluno(alunoTest, t)
}
