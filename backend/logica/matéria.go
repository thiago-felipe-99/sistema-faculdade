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

// existe verifica se a Matéria existe na aplicação.
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

// existeIDs verifica se as Matérias existe na aplicação.
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

// préRequisitoCiclico verifica se a um ciclo entre os pré requisitos.
// Para fazer essa verificação ele usa o algoritmo Depth First Traversal(DFS).
func (lógica Matéria) préRequisitoCiclicos( //nolint: cyclop
	idMatéria id, préRequisitos []id,
) (bool, erro) {
	matérias := préRequisitos
	matériasPréRequisitos := map[id][]id{idMatéria: préRequisitos}

	for len(matérias) != 0 {
		var matéria id

		matéria, matérias = matérias[0], matérias[1:]

		if _, existe := matériasPréRequisitos[matéria]; existe {
			continue
		}

		préRequisitos, erro := lógica.data.PegarPréRequisitos(matéria)
		if erro != nil {
			return false, erros.Novo(ErroVerificarPréRequisitosCiclos, erro, nil)
		}

		matérias = append(matérias, préRequisitos...)
		matériasPréRequisitos[matéria] = préRequisitos
	}

	visitados, pilha := map[id]bool{}, map[id]bool{}
	for matéria := range matériasPréRequisitos {
		visitados[matéria] = false
		pilha[matéria] = false
	}

	var éCiclico func(id) bool

	éCiclico = func(idMatéria id) bool {
		if pilha[idMatéria] {
			return true
		}

		if visitados[idMatéria] {
			return false
		}

		visitados[idMatéria], pilha[idMatéria] = true, true

		for _, matéria := range matériasPréRequisitos[idMatéria] {
			if éCiclico(matéria) {
				return true
			}
		}

		pilha[idMatéria] = false

		return false
	}

	for matéria := range matériasPréRequisitos {
		if éCiclico(matéria) {
			return true, nil
		}
	}

	return false, nil
}

// Criar adiciona uma Matéria na aplicação.
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

// Atualizar é um método que atualiza um Matéria na aplicação.
func (lógica Matéria) Atualizar( //nolint: cyclop
	idMatéria id,
	nome string,
	cargaHoráriaSemanal time.Duration,
	créditos float32,
	tipo string,
	préRequisitos []id,
) (*matéria, erro) {
	existe, erro := lógica.existe(idMatéria)
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

	ciclo, erro := lógica.préRequisitoCiclicos(idMatéria, préRequisitos)
	if erro != nil {
		return nil, erros.Novo(ErroAtualizarMatéria, erro, nil)
	}

	if ciclo {
		return nil, erros.Novo(ErroPréRequisitosCiclo, nil, nil)
	}

	matéria := &matéria{
		ID:                  idMatéria,
		Nome:                nome,
		CargaHoráriaSemanal: cargaHoráriaSemanal,
		Créditos:            créditos,
		PréRequisitos:       préRequisitos,
		Tipo:                tipo,
	}

	erro = lógica.data.Atualizar(idMatéria, matéria)
	if erro != nil {
		return nil, erros.Novo(ErroAtualizarMatéria, erro, nil)
	}

	return matéria, nil
}

// Pegar é um método que pega uma Matéria na aplicação.
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
func (lógica Matéria) Deletar(idMatéria id) erro {
	existe, erro := lógica.existe(idMatéria)
	if erro != nil {
		return erros.Novo(ErroDeletarMatéria, erro, nil)
	}

	if !existe {
		return erros.Novo(ErroMatériaNãoEncontrada, nil, nil)
	}

	erro = lógica.data.Deletar(idMatéria)
	if erro != nil {
		return erros.Novo(ErroDeletarMatéria, erro, nil)
	}

	return nil
}
