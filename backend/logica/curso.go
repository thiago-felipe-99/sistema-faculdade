package logica

import (
	"time"

	"thiagofelipe.com.br/sistema-faculdade-backend/data"
	"thiagofelipe.com.br/sistema-faculdade-backend/entidades"
	"thiagofelipe.com.br/sistema-faculdade-backend/erros"
)

// Curso representa operações que se possa fazer com a entidade Curso.
type Curso struct {
	data    data.Curso
	matéria Matéria
}

// existe verifica se a Curso existe na aplicação.
func (lógica Curso) existe(id id) (bool, erro) {
	_, erro := lógica.data.Pegar(id)
	if erro != nil {
		if erro.ÉPadrão(data.ErroCursoNãoEncontrado) {
			return false, nil
		}

		return false, erros.Novo(ErroVerificarID, erro, nil)
	}

	return true, nil
}

// Criar adiciona um Curso na aplicação.
func (lógica Curso) Criar(
	nome string,
	dataDeInício time.Time,
	dataDeDesativação time.Time,
	matérias []cursoMatéria,
) (*curso, erro) {
	dataDeInício = entidades.RemoverHorário(dataDeInício)
	dataDeDesativação = entidades.RemoverHorário(dataDeDesativação)

	if dataDeInício.Unix() >= dataDeDesativação.Unix() {
		return nil, erros.Novo(ErroDataDeInícioMaior, nil, nil)
	}

	matériasID := []id{}

	for _, matéria := range matérias {
		matériasID = append(matériasID, matéria.Matéria)
	}

	_, existe, erro := lógica.matéria.existeIDs(matériasID)
	if erro != nil && !erro.ÉPadrão(ErroIDsTamanho) {
		return nil, erros.Novo(ErroCriarCurso, erro, nil)
	}

	if !existe {
		return nil, erros.Novo(ErroMatériaNãoEncontrada, nil, nil)
	}

	curso := &curso{
		ID:                entidades.NovoID(),
		Nome:              nome,
		DataDeInício:      dataDeInício,
		DataDeDesativação: dataDeDesativação,
		Matérias:          matérias,
	}

	erro = lógica.data.Inserir(curso)
	if erro != nil {
		return nil, erros.Novo(ErroCriarCurso, erro, nil)
	}

	return curso, nil
}

// Pegar é um método que pega uma Curso na aplicação.
func (lógica Curso) Pegar(id id) (*curso, erro) {
	curso, erro := lógica.data.Pegar(id)
	if erro != nil {
		if erro.ÉPadrão(data.ErroCursoNãoEncontrado) {
			return nil, erros.Novo(ErroCursoNãoEncontrado, nil, nil)
		}

		return nil, erros.Novo(ErroPegarCurso, erro, nil)
	}

	return curso, nil
}

// Deletar é um método que deleta uma curso da aplicação.
func (lógica Curso) Deletar(id id) erro {
	existe, erro := lógica.existe(id)
	if erro != nil {
		return erros.Novo(ErroDeletarCurso, erro, nil)
	}

	if !existe {
		return erros.Novo(ErroCursoNãoEncontrado, nil, nil)
	}

	erro = lógica.data.Deletar(id)
	if erro != nil {
		return erros.Novo(ErroDeletarCurso, erro, nil)
	}

	return nil
}
