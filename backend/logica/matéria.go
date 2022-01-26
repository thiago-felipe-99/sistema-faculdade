package logica

import (
	"time"

	"thiagofelipe.com.br/sistema-faculdade-backend/data"
	"thiagofelipe.com.br/sistema-faculdade-backend/entidades"
	"thiagofelipe.com.br/sistema-faculdade-backend/erros"
)

// Matéria representa operações que se possa fazer com a entidade Matéria.
type Matéria struct {
	data data.Matéria
}

const cargaHoráriaSemanalMímima = time.Hour

// existe verifica se a matéria existe na aplicação.
func (lógica Matéria) existe(id id) (bool, erro) {
	_, erro := lógica.data.Pegar(id)
	if erro != nil {
		if erro.ÉPadrão(data.ErroMatériaNãoEncontrada) {
			return false, nil
		}

		return false, erros.Novo(ErroVerificarID, erro, nil)
	}

	return true, nil
}

// existeIDs verifica se as matérias existe na aplicação.
func (lógica Matéria) existeIDs(ids []id) ([]id, bool, erro) {
	if len(ids) == 0 {
		return []id{}, true, erros.Novo(ErroIDsTamanho, nil, nil)
	}

	idsÚnico := entidades.IDsÚnicos(ids)

	matérias, erro := lógica.data.PegarMúltiplos(idsÚnico)
	if erro != nil {
		return []id{}, false, erros.Novo(ErroVerificarIDs, erro, nil)
	}

	idsSalvos := []id{}

	for _, matéria := range matérias {
		idsSalvos = append(idsSalvos, matéria.ID)
	}

	return idsSalvos, len(idsSalvos) == len(idsÚnico), nil
}

// Criar adiciona uma matéria na aplicação.
func (lógica Matéria) Criar(
	nome string,
	cargaHoráriaSemanal time.Duration,
	créditos float32,
	tipo string,
	préRequisitos []id,
) (*matéria, erro) {
	if cargaHoráriaSemanal < cargaHoráriaSemanalMímima {
		return nil, erros.Novo(ErroCargaHoráriaMínima, nil, nil)
	}

	if créditos <= 0 {
		return nil, erros.Novo(ErroCréditosInválido, nil, nil)
	}

	préRequisitos, existe, erro := lógica.existeIDs(préRequisitos)
	if erro != nil && !erro.ÉPadrão(ErroIDsTamanho) {
		return nil, erros.Novo(ErroCriarMatéria, erro, nil)
	}

	if !existe {
		return nil, erros.Novo(ErroPréRequisitosNãoExiste, nil, nil)
	}

	matéria := &matéria{
		ID:                  entidades.NovoID(),
		Nome:                nome,
		CargaHoráriaSemanal: cargaHoráriaSemanal,
		Créditos:            créditos,
		PréRequisitos:       préRequisitos,
		Tipo:                tipo,
	}

	erro = lógica.data.Inserir(matéria)
	if erro != nil {
		return nil, erros.Novo(ErroCriarMatéria, erro, nil)
	}

	return matéria, nil
}

// Atualizar é um método que atualiza um matéria na aplicação.
func (lógica Matéria) Atualizar(
	id id,
	nome string,
	cargaHoráriaSemanal time.Duration,
	créditos float32,
	tipo string,
	préRequisitos []id,
) (*matéria, erro) {
	existe, erro := lógica.existe(id)
	if erro != nil {
		return nil, erros.Novo(ErroAtualizarMatéria, erro, nil)
	}

	if !existe {
		return nil, erros.Novo(ErroMatériaNãoEncontrada, nil, nil)
	}

	if cargaHoráriaSemanal < cargaHoráriaSemanalMímima {
		return nil, erros.Novo(ErroCargaHoráriaMínima, nil, nil)
	}

	if créditos <= 0 {
		return nil, erros.Novo(ErroCréditosInválido, nil, nil)
	}

	préRequisitos, existe, erro = lógica.existeIDs(préRequisitos)
	if erro != nil && !erro.ÉPadrão(ErroIDsTamanho) {
		return nil, erros.Novo(ErroAtualizarMatéria, erro, nil)
	}

	if !existe {
		return nil, erros.Novo(ErroPréRequisitosNãoExiste, nil, nil)
	}

	matéria := &matéria{
		ID:                  id,
		Nome:                nome,
		CargaHoráriaSemanal: cargaHoráriaSemanal,
		Créditos:            créditos,
		PréRequisitos:       préRequisitos,
		Tipo:                tipo,
	}

	erro = lógica.data.Atualizar(id, matéria)
	if erro != nil {
		return nil, erros.Novo(ErroAtualizarMatéria, erro, nil)
	}

	return matéria, nil
}

// Pegar é um método que pega uma matéria na aplicação.
func (lógica Matéria) Pegar(id id) (*matéria, erro) {
	matéria, erro := lógica.data.Pegar(id)
	if erro != nil {
		if erro.ÉPadrão(data.ErroMatériaNãoEncontrada) {
			return nil, erros.Novo(ErroMatériaNãoEncontrada, nil, nil)
		}

		return nil, erros.Novo(ErroPegarMatéria, erro, nil)
	}

	return matéria, nil
}

// Deletar é um método que deleta uma matéria da aplicação.
func (lógica Matéria) Deletar(id id) erro {
	existe, erro := lógica.existe(id)
	if erro != nil {
		return erros.Novo(ErroDeletarMatéria, erro, nil)
	}

	if !existe {
		return erros.Novo(ErroMatériaNãoEncontrada, nil, nil)
	}

	erro = lógica.data.Deletar(id)
	if erro != nil {
		return erros.Novo(ErroDeletarMatéria, erro, nil)
	}

	return nil
}
