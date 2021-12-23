package mariadb

import (
	"math/rand"
	"reflect"
	"regexp"
	"testing"
	"time"

	"thiagofelipe.com.br/sistema-faculdade-backend/aleatorio"
	. "thiagofelipe.com.br/sistema-faculdade-backend/data/erros"
	"thiagofelipe.com.br/sistema-faculdade-backend/entidades"
)

func criarMatériasCursoAleatórios(idCurso entidades.ID) *[]entidades.CursoMatéria {
	matérias := make([]entidades.CursoMatéria, rand.Intn(MATÉRIAS_MÁXIMAS)+1)

	for i := range matérias {
		matérias[i] = entidades.CursoMatéria{
			IDCurso:    idCurso,
			IDMatéria:  entidades.NovoID(),
			Status:     aleatorio.Palavra(rand.Intn(TAMANHO_MÁXIMO_PALAVRA) + 1),
			Período:    aleatorio.Palavra(rand.Intn(TAMANHO_MÁXIMO_PALAVRA) + 1),
			Tipo:       aleatorio.Palavra(rand.Intn(TAMANHO_MÁXIMO_PALAVRA) + 1),
			Observação: aleatorio.Palavra(rand.Intn(TAMANHO_MÁXIMO_PALAVRA) + 1),
		}
	}

	return &matérias
}

func criarCursoAleatório() *entidades.Curso {
	id := entidades.NovoID()

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

	chaves := make(map[entidades.ID]bool)
	cursoIDs := []entidades.ID{}

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
			"Erro ao salvar a curso no banco de dados, queria %v, chegou %v",
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

func removerCursoMatérias(idCurso entidades.ID, t *testing.T) {
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

func removerCurso(id entidades.ID, t *testing.T) {
	erro := cursoBD.Deletar(id)
	if erro != nil {
		t.Fatalf("Erro ao tentar deletar o curso teste: %v", erro.Error())
	}
}

func TestInserirCursoMatérias_semTamanhoMínimo(t *testing.T) {
	matériasVazia := &[]entidades.CursoMatéria{}

	erro := cursoBD.InserirMatérias(matériasVazia)
	if !erro.ÉPadrão(ErroInserirCursoMatériasTamanhoMínimo) {
		t.Fatalf(
			"Erro ao inserir as matérias do curso, esperava \"%s\", chegou \"%s\"",
			ErroInserirCursoMatériasTamanhoMínimo.Mensagem,
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
		matérias = append(matérias, *criarMatériasCursoAleatórios(entidades.NovoID())...)
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

func TestAtualizarCursoMáterias(t *testing.T) {
	cursoTeste := criarCursoAleatório()

	adiconarCurso(cursoTeste, t)

	for índice := range cursoTeste.Matérias {
		cursoTeste.Matérias[índice].Observação =
			aleatorio.Palavra(rand.Intn(TAMANHO_MÁXIMO_PALAVRA) + 1)
		cursoTeste.Matérias[índice].Período =
			aleatorio.Palavra(rand.Intn(TAMANHO_MÁXIMO_PALAVRA) + 1)
		cursoTeste.Matérias[índice].Status =
			aleatorio.Palavra(rand.Intn(TAMANHO_MÁXIMO_PALAVRA) + 1)
		cursoTeste.Matérias[índice].Tipo =
			aleatorio.Palavra(rand.Intn(TAMANHO_MÁXIMO_PALAVRA) + 1)
	}

	erro := cursoBD.AtualizarMatérias(&cursoTeste.Matérias)
	if erro != nil {
		t.Fatalf("Erro ao atualizar as matérias do cursos: %s", erro.Error())
	}

	matériasSalvas, erro := cursoBD.PegarMatérias(cursoTeste.ID)
	if erro != nil {
		t.Fatalf("Erro ao pegar as matérias do curso no banco de dados: %s", erro.Error())
	}

	if !reflect.DeepEqual(&cursoTeste.Matérias, matériasSalvas) {
		t.Fatalf(
			"Erro ao salvar a pessoa no banco de dados, queria %v, chegou %v",
			&cursoTeste.Matérias,
			matériasSalvas,
		)
	}
}

func TestAtualizarCursoMáterias_tabelaInválida(t *testing.T) {
	texto := `Table .* doesn't exist`
	padrão, erroRegex := regexp.Compile(texto)
	if erroRegex != nil {
		t.Fatalf("Erro ao compilar o regex: %s", texto)
	}

	cursoTeste := criarCursoAleatório()

	adiconarCurso(cursoTeste, t)

	for índice := range cursoTeste.Matérias {
		cursoTeste.Matérias[índice].Observação =
			aleatorio.Palavra(rand.Intn(TAMANHO_MÁXIMO_PALAVRA) + 1)

		cursoTeste.Matérias[índice].Período =
			aleatorio.Palavra(rand.Intn(TAMANHO_MÁXIMO_PALAVRA) + 1)

		cursoTeste.Matérias[índice].Status =
			aleatorio.Palavra(rand.Intn(TAMANHO_MÁXIMO_PALAVRA) + 1)

		cursoTeste.Matérias[índice].Tipo =
			aleatorio.Palavra(rand.Intn(TAMANHO_MÁXIMO_PALAVRA) + 1)
	}

	erro := cursoBDInválido.AtualizarMatérias(&cursoTeste.Matérias)
	if erro == nil || erro.ErroExterno == nil {
		t.Fatalf("Não foi enviado erro do sistema")
	}

	if !padrão.MatchString(erro.ErroExterno.Error()) {
		t.Fatalf(
			"Erro ao atualizar as matérias do cursos no banco de dados, queria %v, chegou %v",
			texto,
			erro.ErroExterno.Error(),
		)
	}
}

func TestAtualizarCurso(t *testing.T) {
	cursoTeste := criarCursoAleatório()

	adiconarCurso(cursoTeste, t)

	dataAgora := time.Now().UTC()
	dataAgora = dataAgora.Truncate(24 * time.Hour)
	dataFutura1 := dataAgora.AddDate(rand.Intn(50), rand.Intn(12), rand.Intn(28))
	dataFutura2 := dataAgora.AddDate(rand.Intn(50), rand.Intn(12), rand.Intn(28))

	cursoTeste.Nome = "Novo Nome"
	cursoTeste.DataDeDesativação = dataFutura1
	cursoTeste.DataDeInício = dataFutura2

	erro := cursoBD.Atualizar(cursoTeste.ID, cursoTeste)
	if erro != nil {
		t.Fatalf("Erro ao atualizar o curso teste: %s", erro.Error())
	}

	cursoSalvo, erro := cursoBD.Pegar(cursoTeste.ID)
	if erro != nil {
		t.Fatalf("Erro ao pegar o curso no banco de dados: %s", erro.Error())
	}

	if !reflect.DeepEqual(cursoTeste, cursoSalvo) {
		t.Fatalf(
			"Erro ao salvar o curso no banco de dados, queria %v, chegou %v",
			cursoTeste,
			cursoSalvo,
		)
	}
}

func TestAtualizarCurso_tabelaInválida(t *testing.T) {
	texto := `Table .* doesn't exist`
	padrão, erroRegex := regexp.Compile(texto)
	if erroRegex != nil {
		t.Fatal("Erro ao compilar o regex")
	}

	cursoTeste := criarCursoAleatório()

	adiconarCurso(cursoTeste, t)

	erro := cursoBDInválido.Atualizar(cursoTeste.ID, cursoTeste)
	if erro == nil || erro.ErroInicial == nil || erro.ErroInicial.ErroExterno == nil {
		t.Fatalf("Não foi enviado erro do sistema")
	}

	if !padrão.MatchString(erro.ErroInicial.ErroExterno.Error()) {
		t.Fatalf(
			"Erro ao atualizar o curso queria: %s, chegou %s",
			texto,
			erro.ErroInicial.ErroExterno.Error(),
		)
	}
}

func TestAtualizarCurso_tabelaInválida2(t *testing.T) {
	texto := `Table .* doesn't exist`
	padrão, erroRegex := regexp.Compile(texto)
	if erroRegex != nil {
		t.Fatal("Erro ao compilar o regex")
	}

	cursoTeste := criarCursoAleatório()

	adiconarCurso(cursoTeste, t)

	erro := cursoBDInválido2.Atualizar(cursoTeste.ID, cursoTeste)
	if erro == nil || erro.ErroExterno == nil {
		t.Fatalf("Não foi enviado erro do sistema")
	}

	if !padrão.MatchString(erro.ErroExterno.Error()) {
		t.Fatalf(
			"Erro ao atualizar o curso queria: %s, chegou %s",
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
	matérias, erro := cursoBD.PegarMatérias(entidades.NovoID())
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

	_, erro := cursoBDInválido.PegarMatérias(entidades.NovoID())
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
	_, erro := cursoBD.Pegar(entidades.NovoID())
	if erro == nil {
		t.Fatalf("Deveria ter um erro de curso não encontrado")
	}

	if !erro.ÉPadrão(ErroCursoNãoEncontrado) {
		t.Fatalf(
			"Espera o erro de não encontrar o curso, queria \"%s\", chegou \"%s\"",
			ErroCursoNãoEncontrado,
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

	_, erro := cursoBDInválido.Pegar(entidades.NovoID())
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

	_, erro := cursoBDInválido2.Pegar(entidades.NovoID())
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

func TestDeletarCurso(t *testing.T) {
	cursoTeste := criarCursoAleatório()

	adiconarCurso(cursoTeste, t)

	removerCurso(cursoTeste.ID, t)

	_, erro := cursoBD.Pegar(cursoTeste.ID)
	if erro == nil || !erro.ÉPadrão(ErroCursoNãoEncontrado) {
		t.Fatalf(
			"Deveria retonar um erro de curso não encontrado, retonou %s",
			erro,
		)
	}
}

func TestDeletarCurso_invalídoID(t *testing.T) {
	id := entidades.NovoID()

	removerCurso(id, t)

	_, erro := cursoBD.Pegar(id)
	if erro == nil || !erro.ÉPadrão(ErroCursoNãoEncontrado) {
		t.Fatalf(
			"Deveria retonar um erro de curso não encontrado, retonou %s",
			erro,
		)
	}
}

func TestDeletarCurso_tabelaInválida(t *testing.T) {
	texto := `Table .* doesn't exist`
	padrão, erroRegex := regexp.Compile(texto)
	if erroRegex != nil {
		t.Fatalf("Erro ao compilar o regex: %s", texto)
	}

	erro := cursoBDInválido.Deletar(entidades.NovoID())

	if erro == nil || erro.ErroInicial == nil || erro.ErroInicial.ErroExterno == nil {
		t.Fatalf("Não foi enviado erro do sistema")
	}

	if !padrão.MatchString(erro.ErroInicial.ErroExterno.Error()) {
		t.Fatalf(
			"Erro ao pegar curso no banco de dados, queria %v, chegou %v",
			texto,
			erro.ErroInicial.ErroExterno.Error(),
		)
	}
}

func TestDeletarCurso_tabelaInválida2(t *testing.T) {
	texto := `Table .* doesn't exist`
	padrão, erroRegex := regexp.Compile(texto)
	if erroRegex != nil {
		t.Fatalf("Erro ao compilar o regex: %s", texto)
	}

	erro := cursoBDInválido2.Deletar(entidades.NovoID())

	if erro == nil || erro.ErroExterno == nil {
		t.Fatalf("Não foi enviado erro do sistema")
	}

	if !padrão.MatchString(erro.ErroExterno.Error()) {
		t.Fatalf(
			"Erro ao pegar curso no banco de dados, queria %v, chegou %v",
			texto,
			erro.ErroExterno.Error(),
		)
	}
}

func TestDeletarCursoMatérias(t *testing.T) {
	cursoTeste := criarCursoAleatório()

	adiconarCurso(cursoTeste, t)

	removerCursoMatérias(cursoTeste.ID, t)

	matérias, erro := cursoBD.PegarMatérias(cursoTeste.ID)
	if erro != nil || len(*matérias) != 0 {
		t.Fatalf(
			"Deveria retonar um erro de curso não encontrado, retonou %s",
			erro,
		)
	}
}

func TestDeletarCursoMatérias_invalídoID(t *testing.T) {
	id := entidades.NovoID()

	removerCursoMatérias(id, t)

	matérias, erro := cursoBD.PegarMatérias(id)
	if erro != nil || len(*matérias) != 0 {
		t.Fatalf(
			"Deveria retonar um erro de curso não encontrado, retonou %s",
			erro,
		)
	}
}

func TestDeletarCursoMatérias_tabelaInválida(t *testing.T) {
	texto := `Table .* doesn't exist`
	padrão, erroRegex := regexp.Compile(texto)
	if erroRegex != nil {
		t.Fatalf("Erro ao compilar o regex: %s", texto)
	}

	erro := cursoBDInválido.DeletarMatérias(entidades.NovoID())

	if erro == nil || erro.ErroExterno == nil {
		t.Fatalf("Não foi enviado erro do sistema")
	}

	if !padrão.MatchString(erro.ErroExterno.Error()) {
		t.Fatalf(
			"Erro ao pegar curso no banco de dados, queria %v, chegou %v",
			texto,
			erro.ErroExterno.Error(),
		)
	}
}
