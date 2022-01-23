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

// Criar adiciona uma matéria na aplicação.
func (lógica Matéria) Criar(
	nome string,
	cargaHoráriaSemanal time.Duration,
	créditos float32,
	préRequisitos []entidades.ID,
	tipo string,
) (*entidades.Matéria, *erros.Aplicação) {
	if cargaHoráriaSemanal < cargaHoráriaSemanalMímima {
		return nil, erros.Novo(ErroCargaHoráriaMínima, nil, nil)
	}

	if créditos <= 0 {
		return nil, erros.Novo(ErroCréditosInválido, nil, nil)
	}

	_, existe, erro := lógica.data.Existe(préRequisitos)
	if erro != nil {
		return nil, erros.Novo(ErroCriarMatéria, erro, nil)
	}
	if !existe {
		return nil, erros.Novo(ErroPréRequisitosNãoExiste, nil, nil)
	}

	matéria := &entidades.Matéria{
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
