package mariadb

import (
	"reflect"
	"regexp"
	"testing"

	"thiagofelipe.com.br/sistema-faculdade-backend/aleatorio"
	. "thiagofelipe.com.br/sistema-faculdade-backend/data"
	"thiagofelipe.com.br/sistema-faculdade-backend/entidades"
)

func criarMatériasCursoAleatórios(idCurso entidades.ID) *[]entidades.CursoMatéria {
	matérias := make([]entidades.CursoMatéria, aleatorio.Número(matériasMáximas)+1)

	for i := range matérias {
		matérias[i] = entidades.CursoMatéria{
			IDCurso:    idCurso,
			IDMatéria:  entidades.NovoID(),
			Status:     aleatorio.Palavra(aleatorio.Número(tamanhoMáximoPalavra) + 1),
			Período:    aleatorio.Palavra(aleatorio.Número(tamanhoMáximoPalavra) + 1),
			Tipo:       aleatorio.Palavra(aleatorio.Número(tamanhoMáximoPalavra) + 1),
			Observação: aleatorio.Palavra(aleatorio.Número(tamanhoMáximoPalavra) + 1),
		}
	}

	return &matérias
}

func criarCursoAleatório() *entidades.Curso {
	id := entidades.NovoID()

	dataAgora := entidades.DataAtual()
	ano := int(aleatorio.Número(50))
	mês := int(aleatorio.Número(12))
	dia := int(aleatorio.Número(28))
	dataFutura := dataAgora.AddDate(ano, mês, dia)

	return &entidades.Curso{
		ID:                id,
		Nome:              aleatorio.Palavra(aleatorio.Número(tamanhoMáximoPalavra) + 1),
		DataDeInício:      dataAgora,
		DataDeDesativação: dataFutura,
		Matérias:          *criarMatériasCursoAleatórios(id),
	}
}

func removerCursoMatérias(t *testing.T, idCurso entidades.ID) {
	erro := cursoBD.DeletarMatérias(idCurso)
	if erro != nil {
		t.Fatalf("Erro ao tentar deletar as matérias do: %v", erro.Error())
	}
}

func adiconarCurso(t *testing.T, curso *entidades.Curso) {
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
		removerCurso(t, curso.ID)
	})
}

func removerCurso(t *testing.T, id entidades.ID) {
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
	const texto = `foreign key constraint fails`
	padrão := regexp.MustCompile(texto)

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

	adiconarCurso(t, cursoTeste)
}

func TestInserirCurso_duplicadoID(t *testing.T) {
	const texto = `Duplicate entry.*PRIMARY`
	padrão := regexp.MustCompile(texto)

	cursoTeste := criarCursoAleatório()

	adiconarCurso(t, cursoTeste)

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

	adiconarCurso(t, cursoTeste)

	for índice := range cursoTeste.Matérias {
		cursoTeste.Matérias[índice].Observação =
			aleatorio.Palavra(aleatorio.Número(tamanhoMáximoPalavra) + 1)
		cursoTeste.Matérias[índice].Período =
			aleatorio.Palavra(aleatorio.Número(tamanhoMáximoPalavra) + 1)
		cursoTeste.Matérias[índice].Status =
			aleatorio.Palavra(aleatorio.Número(tamanhoMáximoPalavra) + 1)
		cursoTeste.Matérias[índice].Tipo =
			aleatorio.Palavra(aleatorio.Número(tamanhoMáximoPalavra) + 1)
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
	const texto = `Table .* doesn't exist`
	padrão := regexp.MustCompile(texto)

	cursoTeste := criarCursoAleatório()

	adiconarCurso(t, cursoTeste)

	for índice := range cursoTeste.Matérias {
		cursoTeste.Matérias[índice].Observação =
			aleatorio.Palavra(aleatorio.Número(tamanhoMáximoPalavra) + 1)

		cursoTeste.Matérias[índice].Período =
			aleatorio.Palavra(aleatorio.Número(tamanhoMáximoPalavra) + 1)

		cursoTeste.Matérias[índice].Status =
			aleatorio.Palavra(aleatorio.Número(tamanhoMáximoPalavra) + 1)

		cursoTeste.Matérias[índice].Tipo =
			aleatorio.Palavra(aleatorio.Número(tamanhoMáximoPalavra) + 1)
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

	adiconarCurso(t, cursoTeste)

	dataAgora := entidades.DataAtual()
	ano1 := int(aleatorio.Número(50))
	mês1 := int(aleatorio.Número(12))
	dia1 := int(aleatorio.Número(28))
	dataFutura1 := dataAgora.AddDate(ano1, mês1, dia1)
	ano2 := int(aleatorio.Número(50))
	mês2 := int(aleatorio.Número(12))
	dia2 := int(aleatorio.Número(28))
	dataFutura2 := dataAgora.AddDate(ano2, mês2, dia2)

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
	const texto = `Table .* doesn't exist`
	padrão := regexp.MustCompile(texto)

	cursoTeste := criarCursoAleatório()

	adiconarCurso(t, cursoTeste)

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
	const texto = `Table .* doesn't exist`
	padrão := regexp.MustCompile(texto)

	cursoTeste := criarCursoAleatório()

	adiconarCurso(t, cursoTeste)

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

	adiconarCurso(t, cursoTeste)

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
	const texto = `Table .* doesn't exist`
	padrão := regexp.MustCompile(texto)

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

	adiconarCurso(t, cursoTeste)
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
	const texto = `Table .* doesn't exist`
	padrão := regexp.MustCompile(texto)

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
	const texto = `Table .* doesn't exist`
	padrão := regexp.MustCompile(texto)

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

	adiconarCurso(t, cursoTeste)

	removerCurso(t, cursoTeste.ID)

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

	removerCurso(t, id)

	_, erro := cursoBD.Pegar(id)
	if erro == nil || !erro.ÉPadrão(ErroCursoNãoEncontrado) {
		t.Fatalf(
			"Deveria retonar um erro de curso não encontrado, retonou %s",
			erro,
		)
	}
}

func TestDeletarCurso_tabelaInválida(t *testing.T) {
	const texto = `Table .* doesn't exist`
	padrão := regexp.MustCompile(texto)

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
	const texto = `Table .* doesn't exist`
	padrão := regexp.MustCompile(texto)

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

	adiconarCurso(t, cursoTeste)

	removerCursoMatérias(t, cursoTeste.ID)

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

	removerCursoMatérias(t, id)

	matérias, erro := cursoBD.PegarMatérias(id)
	if erro != nil || len(*matérias) != 0 {
		t.Fatalf(
			"Deveria retonar um erro de curso não encontrado, retonou %s",
			erro,
		)
	}
}

func TestDeletarCursoMatérias_tabelaInválida(t *testing.T) {
	const texto = `Table .* doesn't exist`
	padrão := regexp.MustCompile(texto)

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