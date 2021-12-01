package mariadb

import (
	"math/rand"
	"reflect"
	"regexp"
	"testing"
	"time"

	"github.com/google/uuid"
	"thiagofelipe.com.br/sistema-faculdade/aleatorio"
	"thiagofelipe.com.br/sistema-faculdade/data"
)

const MATÉRIAS_MÁXIMAS = 20
const TAMANHO_MÁXIMO_PALAVRA = 25

func criarMatériasCursoAleatórios(idCurso id) *[]data.CursoMatéria {
	matérias := make([]data.CursoMatéria, rand.Intn(MATÉRIAS_MÁXIMAS))

	for i := range matérias {
		matérias[i] = data.CursoMatéria{
			ID_Curso:   idCurso,
			ID_Matéria: uuid.New(),
			Status:     aleatorio.Palavra(rand.Intn(TAMANHO_MÁXIMO_PALAVRA)),
			Período:    aleatorio.Palavra(rand.Intn(TAMANHO_MÁXIMO_PALAVRA)),
			Tipo:       aleatorio.Palavra(rand.Intn(TAMANHO_MÁXIMO_PALAVRA)),
			Observação: aleatorio.Palavra(rand.Intn(TAMANHO_MÁXIMO_PALAVRA)),
		}
	}

	return &matérias
}

func criarCursoAleatório() *data.Curso {
	id := uuid.New()

	dataAgora := time.Now().UTC()
	dataAgora = dataAgora.Truncate(24 * time.Hour)
	dataFutura := dataAgora.AddDate(rand.Intn(50), rand.Intn(12), rand.Intn(28))

	return &data.Curso{
		ID:                id,
		Nome:              aleatorio.Palavra(rand.Intn(TAMANHO_MÁXIMO_PALAVRA)),
		DataDeInício:      dataAgora,
		DataDeDesativação: dataFutura,
		Matérias:          *criarMatériasCursoAleatórios(id),
	}
}

func adiconarCursoTeste(curso *data.Curso, t *testing.T) {

	erro := cursoBD.Inserir(curso)
	if erro != nil {
		t.Fatalf("Erro ao inserir o curso no banco de dados: %s", erro.Error())
	}

	cursoSalvo, erro := cursoBD.Pegar(curso.ID)
	if erro != nil {
		t.Fatalf("Erro ao pegar o curso no banco de dados: %s", erro.Error())
	}

	if !reflect.DeepEqual(curso, cursoSalvo) {
		t.Fatalf(
			"Erro ao salvar o curso no banco de dados, queria %v, chegou %v",
			curso,
			cursoSalvo,
		)
	}

	t.Cleanup(func() {
		removerCursoTeste(curso.ID, t)
	})
}

func removerCursoTeste(id id, t *testing.T) {
	erro := cursoBD.Deletar(id)
	if erro != nil {
		t.Fatalf("Erro ao tentar deletar o usuário teste: %v", erro.Error())
	}
}

func TestInserirCursoMatérias(t *testing.T) {
	var matérias []data.CursoMatéria

	índiceMáximo := 5
	for índice := 0; índiceMáximo > índice; índice++ {
		matérias = append(matérias, *criarMatériasCursoAleatórios(uuid.New())...)
	}

}

func TestInserirCurso(t *testing.T) {
	cursoTeste := criarCursoAleatório()

	adiconarCursoTeste(cursoTeste, t)
}

func TestInserirCurso_duplicadoID(t *testing.T) {
	texto := `Duplicate entry.*PRIMARY`
	padrão, erroRegex := regexp.Compile(texto)
	if erroRegex != nil {
		t.Fatal("Erro ao compilar o regex")
	}

	cursoTeste := criarCursoAleatório()

	adiconarCursoTeste(cursoTeste, t)

	erro := cursoBD.Inserir(cursoTeste)
	if erro == nil || erro.ErroExterno == nil {
		t.Fatalf("Não foi enviado erro do sistema")
	}

	if !padrão.MatchString(erro.ErroExterno.Error()) {
		t.Fatalf(
			"Erro ao inserir o curso queria: %s, chegou %s",
			texto,
			erro.ErroExterno.Error(),
		)
	}
}
