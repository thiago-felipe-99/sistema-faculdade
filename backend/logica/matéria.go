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

// ExisteIDs verifica se as matérias existe na aplicação.
func (lógica Matéria) ExisteIDs(ids []id) ([]id, bool, erro) {
	if len(ids) == 0 {
		return []id{}, true, erros.Novo(ErroIDsTamanho, nil, nil)
	}

	idsÚnico := entidades.IDsÚnicos(ids)

	matérias, erro := lógica.data.PegarMúltiplos(idsÚnico)
	if erro != nil {
		return []id{}, false, erros.Novo(ErroIDsExiste, erro, nil)
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

	préRequisitos, existe, erro := lógica.ExisteIDs(préRequisitos)
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

// Pegar é um método que pega uma matéria da aplicação.
func (lógica Matéria) Pegar(id id) (*matéria, erro) {
	matéria, erro := lógica.data.Pegar(id)
	if erro != nil {
		if !erro.ÉPadrão(data.ErroMatériaNãoEncontrada) {
			return nil, erros.Novo(ErroMatériaNãoEncontrada, nil, nil)
		}

		return nil, erros.Novo(ErroPegarMatéria, erro, nil)
	}

	return matéria, nil
}

// Deletar é um método que deleta uma matéria da aplicação.
func (lógica Matéria) Deletar(id id) erro {
	_, erro := lógica.data.Pegar(id)
	if erro != nil {
		if !erro.ÉPadrão(ErroMatériaNãoEncontrada) {
			return erros.Novo(ErroMatériaNãoEncontrada, nil, nil)
		}

		return erros.Novo(ErroDeletarMatéria, erro, nil)
	}

	erro = lógica.data.Deletar(id)
	if erro != nil {
		return erros.Novo(ErroDeletarMatéria, erro, nil)
	}

	return nil
}
