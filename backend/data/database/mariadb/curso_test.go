package mariadb

import (
	"math/rand"
	"reflect"
	"regexp"
	"testing"
	"time"

	"github.com/google/uuid"
	"thiagofelipe.com.br/sistema-faculdade/aleatorio"
	"thiagofelipe.com.br/sistema-faculdade/entidades"
	"thiagofelipe.com.br/sistema-faculdade/errors"
)

const MATÉRIAS_MÁXIMAS = 20
const TAMANHO_MÁXIMO_PALAVRA = 25

func criarMatériasCursoAleatórios(idCurso id) *[]entidades.CursoMatéria {
	matérias := make([]entidades.CursoMatéria, rand.Intn(MATÉRIAS_MÁXIMAS)+1)

	for i := range matérias {
		matérias[i] = entidades.CursoMatéria{
			IDCurso:    idCurso,
			IDMatéria:  uuid.New(),
			Status:     aleatorio.Palavra(rand.Intn(TAMANHO_MÁXIMO_PALAVRA) + 1),
			Período:    aleatorio.Palavra(rand.Intn(TAMANHO_MÁXIMO_PALAVRA) + 1),
			Tipo:       aleatorio.Palavra(rand.Intn(TAMANHO_MÁXIMO_PALAVRA) + 1),
			Observação: aleatorio.Palavra(rand.Intn(TAMANHO_MÁXIMO_PALAVRA) + 1),
		}
	}

	return &matérias
}

func criarCursoAleatório() *entidades.Curso {
	id := uuid.New()

	dataAgora := time.Now().UTC()
	dataAgora = dataAgora.Truncate(24 * time.Hour)
	dataFutura := dataAgora.AddDate(rand.Intn(50), rand.Intn(12), rand.Intn(28))

	return &entidades.Curso{
		ID:                id,
		Nome:              aleatorio.Palavra(rand.Intn(TAMANHO_MÁXIMO_PALAVRA) + 1),
		DataDeInício:      dataAgora,
		DataDeDesativação: dataFutura,
		Matérias:          *criarMatériasCursoAleatórios(id),
	}
}

func adiconarCursoMatérias(matérias *[]entidades.CursoMatéria, t *testing.T) {
	erro := cursoBD.InserirMatérias(matérias)
	if erro != nil {
		t.Fatalf("Erro ao inserir as matérias do curso: %s", erro.Error())
	}

	chaves := make(map[id]bool)
	cursoIDs := []id{}

	// filtrando os IDS do curso
	for _, matéria := range *matérias {
		if _, valor := chaves[matéria.IDCurso]; !valor {
			chaves[matéria.IDCurso] = true
			cursoIDs = append(cursoIDs, matéria.IDCurso)
		}
	}

	var matériasSalvas []entidades.CursoMatéria
	for _, id := range cursoIDs {
		matériaSalva, erro := cursoBD.PegarMatérias(id)
		if erro != nil {
			t.Fatalf("Erro ao pegar as matérias do curso: %s", erro.Error())
		}
		matériasSalvas = append(matériasSalvas, *matériaSalva...)
	}

	if !reflect.DeepEqual(&matérias, &matériasSalvas) {
		t.Fatalf(
			"Erro ao salvar a pessoa no banco de dados, queria %v, chegou %v",
			matérias,
			matériasSalvas,
		)
	}

	t.Cleanup(func() {
		for _, id := range cursoIDs {
			removerCursoMatérias(id, t)
		}
	})
}

func removerCursoMatérias(idCurso id, t *testing.T) {
	erro := cursoBD.DeletarMatérias(idCurso)
	if erro != nil {
		t.Fatalf("Erro ao tentar deletar as matérias do: %v", erro.Error())
	}
}

func adiconarCurso(curso *entidades.Curso, t *testing.T) {

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
		removerCurso(curso.ID, t)
	})
}

func removerCurso(id id, t *testing.T) {
	erro := cursoBD.Deletar(id)
	if erro != nil {
		t.Fatalf("Erro ao tentar deletar o usuário teste: %v", erro.Error())
	}
}

func TestInserirCursoMatérias_semTamanhoMínimo(t *testing.T) {
	matériasVazia := &[]entidades.CursoMatéria{}

	erro := cursoBD.InserirMatérias(matériasVazia)
	if !erro.ÉPadrão(errors.InserirCursoMatériasTamanhoMínimo) {
		t.Fatalf(
			"Erro ao inserir as matérias do curso, esperava \"%s\", chegou \"%s\"",
			errors.InserirCursoMatériasTamanhoMínimo.Mensagem,
			erro.Mensagem,
		)
	}
}

func TestInserirCursoMatérias_semCurso(t *testing.T) {
	texto := `foreign key constraint fails`
	padrão, erroRegex := regexp.Compile(texto)
	if erroRegex != nil {
		t.Fatal("Erro ao compilar o regex")
	}

	var matérias []entidades.CursoMatéria

	índiceMáximo := 5
	for índice := 0; índiceMáximo > índice; índice++ {
		matérias = append(matérias, *criarMatériasCursoAleatórios(uuid.New())...)
	}

	erro := cursoBD.InserirMatérias(&matérias)
	if erro == nil || erro.ErroExterno == nil {
		t.Fatalf("Deveria ter um erro no sistema")
	}

	if !padrão.MatchString(erro.ErroExterno.Error()) {
		t.Fatalf(
			"Erro ao inserir as matérias do curso queria: \"%s\", chegou \"%s\"",
			texto,
			erro.ErroExterno.Error(),
		)
	}
}

func TestInserirCurso(t *testing.T) {
	cursoTeste := criarCursoAleatório()

	adiconarCurso(cursoTeste, t)
}

func TestInserirCurso_duplicadoID(t *testing.T) {
	texto := `Duplicate entry.*PRIMARY`
	padrão, erroRegex := regexp.Compile(texto)
	if erroRegex != nil {
		t.Fatal("Erro ao compilar o regex")
	}

	cursoTeste := criarCursoAleatório()

	adiconarCurso(cursoTeste, t)

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

func TestPegarCursoMatérias(t *testing.T) {
	cursoTeste := criarCursoAleatório()

	adiconarCurso(cursoTeste, t)

	matériasSalvas, erro := cursoBD.PegarMatérias(cursoTeste.ID)
	if erro != nil {
		t.Fatalf("Não deveria retorna um erro, retornou: %s", erro.Error())
	}

	if !reflect.DeepEqual(&cursoTeste.Matérias, matériasSalvas) {
		t.Fatalf(
			"Erro ao salvar as matérias do curso no banco de dados, queria %v, chegou %v",
			&cursoTeste.Matérias,
			matériasSalvas,
		)
	}

}

func TestPegarCursoMatérias_idInválido(t *testing.T) {
	matérias, erro := cursoBD.PegarMatérias(uuid.New())
	if erro != nil {
		t.Fatalf("Não deveria retornar um erro")
	}

	if len(*matérias) != 0 {
		t.Fatalf("Deveria retornar uma lista vazia de curso")
	}
}

func TestPegarCurso_bdInválido(t *testing.T) {
	texto := `Table .* doesn't exist`
	padrão, erroRegex := regexp.Compile(texto)
	if erroRegex != nil {
		t.Fatal("Erro ao compilar o regex")
	}

	_, erro := cursoBDInválido.PegarMatérias(uuid.New())
	if erro == nil || erro.ErroExterno == nil {
		t.Fatalf("Deveria ter um erro de tabela inválida")
	}

	if !padrão.MatchString(erro.ErroExterno.Error()) {
		t.Fatalf(
			"Erro ao pegar o curso queria: %s, chegou %s",
			texto,
			erro.ErroExterno.Error(),
		)
	}
}

func TestPegarCurso(t *testing.T) {
	cursoTeste := criarCursoAleatório()

	adiconarCurso(cursoTeste, t)
}

func TestPegarCurso_idInválido(t *testing.T) {
	_, erro := cursoBD.Pegar(uuid.New())
	if erro == nil {
		t.Fatalf("Deveria ter um erro de curso não encontrado")
	}

	if !erro.ÉPadrão(errors.CursoNãoEncontrado) {
		t.Fatalf(
			"Espera o erro de não encontrar o curso, queria \"%s\", chegou \"%s\"",
			errors.CursoNãoEncontrado,
			erro.Mensagem,
		)
	}
}

func TestCurso_tabelaInválida(t *testing.T) {
	texto := `Table .* doesn't exist`
	padrão, erroRegex := regexp.Compile(texto)
	if erroRegex != nil {
		t.Fatal("Erro ao compilar o regex")
	}

	_, erro := cursoBDInválido.Pegar(uuid.New())
	if erro == nil || erro.ErroInicial == nil || erro.ErroInicial.ErroExterno == nil {
		t.Fatalf("Deveria ter um erro de tabela inválida")
	}

	if !padrão.MatchString(erro.ErroInicial.ErroExterno.Error()) {
		t.Fatalf(
			"Erro ao pegar o curso queria: %s, chegou %s",
			texto,
			erro.ErroExterno.Error(),
		)
	}
}

func TestCurso_tabelaInválida2(t *testing.T) {
	texto := `Table .* doesn't exist`
	padrão, erroRegex := regexp.Compile(texto)
	if erroRegex != nil {
		t.Fatal("Erro ao compilar o regex")
	}

	_, erro := cursoBDInválido2.Pegar(uuid.New())
	if erro == nil || erro.ErroExterno == nil {
		t.Fatalf("Deveria ter um erro de tabela inválida")
	}

	if !padrão.MatchString(erro.ErroExterno.Error()) {
		t.Fatalf(
			"Erro ao pegar o curso queria: %s, chegou %s",
			texto,
			erro.ErroExterno.Error(),
		)
	}
}
