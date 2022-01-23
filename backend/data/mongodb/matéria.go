package mongodb

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"thiagofelipe.com.br/sistema-faculdade-backend/data"
	"thiagofelipe.com.br/sistema-faculdade-backend/entidades"
	"thiagofelipe.com.br/sistema-faculdade-backend/erros"
)

type matériaParse struct {
	ID                  entidades.ID `bson:"_id"`
	Nome                string
	CargaHoráriaSemanal time.Duration `bson:"carga_horária_semanal"`
	Créditos            float32
	PréRequisitos       []entidades.ID `bson:"pré-requisitos"`
	Tipo                string
}

// MatériaBD representa a conexão com o banco de dados MongoDB para fazer
// alterações na entidade Matéria.
type MatériaBD struct {
	Conexão
	Collection *mongo.Collection
}

func (bd MatériaBD) inserirMúltiplas(matérias *[]entidades.Matéria) erro {
	ids := ""
	for _, matéria := range *matérias {
		ids += matéria.ID.String() + ","
	}

	bd.Log.Informação("Inserindo múltiplas matérias com os IDs:", ids)

	inserir := []interface{}{}

	for _, matéria := range *matérias {
		inserir = append(inserir, &matériaParse{
			ID:                  matéria.ID,
			Nome:                matéria.Nome,
			CargaHoráriaSemanal: matéria.CargaHoráriaSemanal,
			Créditos:            matéria.Créditos,
			PréRequisitos:       matéria.PréRequisitos,
			Tipo:                matéria.Tipo,
		})
	}

	ctx, cancel := context.WithTimeout(bd.ctx, bd.Timeout)
	defer cancel()

	result, err := bd.Collection.InsertMany(ctx, inserir)
	if err != nil {
		return erros.Novo(data.ErroInserirMatéria, nil, err)
	}

	bd.Log.Debug(result)

	return nil
}

// Inserir é uma método que adiciona uma entidade Matéria no banco de
// dados MongoDB.
func (bd MatériaBD) Inserir(matéria *entidades.Matéria) erro {
	bd.Log.Informação("Inserindo Matéria com ID:", matéria.ID.String())

	inserir := &matériaParse{
		ID:                  matéria.ID,
		Nome:                matéria.Nome,
		CargaHoráriaSemanal: matéria.CargaHoráriaSemanal,
		Créditos:            matéria.Créditos,
		PréRequisitos:       matéria.PréRequisitos,
		Tipo:                matéria.Tipo,
	}

	ctx, cancel := context.WithTimeout(bd.ctx, bd.Timeout)
	defer cancel()

	_, err := bd.Collection.InsertOne(ctx, inserir)
	if err != nil {
		return erros.Novo(data.ErroInserirMatéria, nil, err)
	}

	return nil
}

// Atualizar é uma método que faz a atualização de uma entidade Matéria no banco
// de dados MongoDB.
func (bd MatériaBD) Atualizar(id entidades.ID, matéria *entidades.Matéria) erro {
	bd.Log.Informação("Atualizando Matéria com ID:", matéria.ID.String())

	atualizar := &matériaParse{
		ID:                  id,
		Nome:                matéria.Nome,
		CargaHoráriaSemanal: matéria.CargaHoráriaSemanal,
		Créditos:            matéria.Créditos,
		PréRequisitos:       matéria.PréRequisitos,
		Tipo:                matéria.Tipo,
	}

	query := bson.D{{Key: "$set", Value: atualizar}}

	ctx, cancel := context.WithTimeout(bd.ctx, bd.Timeout)
	defer cancel()

	_, err := bd.Collection.UpdateByID(ctx, id, query)
	if err != nil {
		return erros.Novo(data.ErroAtualizarMatéria, nil, err)
	}

	return nil
}

// Pegar é uma método que retorna uma entidade Matéria no banco de dados MongoDB.
func (bd MatériaBD) Pegar(id entidades.ID) (*entidades.Matéria, erro) {
	bd.Log.Informação("Pegando Matéria com ID:", id)

	ctx, cancel := context.WithTimeout(bd.ctx, bd.Timeout)
	defer cancel()

	var resultado matériaParse

	err := bd.Collection.FindOne(ctx, bson.M{"_id": id}).Decode(&resultado)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, erros.Novo(data.ErroMatériaNãoEncontrada, nil, err)
		}

		return nil, erros.Novo(data.ErroPegarMatéria, nil, err)
	}

	return &entidades.Matéria{
		ID:                  resultado.ID,
		Nome:                resultado.Nome,
		CargaHoráriaSemanal: resultado.CargaHoráriaSemanal,
		Créditos:            resultado.Créditos,
		PréRequisitos:       resultado.PréRequisitos,
		Tipo:                resultado.Tipo,
	}, nil
}

// Existe é um método que retorna se as matérias existe no banco de dados
// MongoDB.
func (bd MatériaBD) Existe(ids []entidades.ID) ([]entidades.ID, bool, erro) {
	if len(ids) == 0 {
		return []entidades.ID{}, false, erros.Novo(data.ErroIDsTamanho, nil, nil)
	}

	idsÚnico := []entidades.ID{}

ids:
	for _, id := range ids {
		for _, idÚnico := range idsÚnico {
			if id == idÚnico {
				break ids
			}
		}
		idsÚnico = append(idsÚnico, id)
	}

	ctx, cancel := context.WithTimeout(bd.ctx, bd.Timeout)
	defer cancel()

	opts := options.Find().SetProjection(bson.M{"_id": 1})
	filtro := bson.M{"_id": bson.M{"$in": idsÚnico}}

	cursor, err := bd.Collection.Find(ctx, filtro, opts)
	if err != nil {
		return []entidades.ID{}, false, erros.Novo(data.ErroExisteMatérias, nil, err)
	}

	results := []struct {
		ID entidades.ID `bson:"_id"`
	}{}

	err = cursor.All(ctx, &results)
	if err != nil {
		return []entidades.ID{}, false, erros.Novo(data.ErroExisteMatérias, nil, err)
	}

	if len(results) != len(idsÚnico) {
		ids := []entidades.ID{}

		for _, result := range results {
			ids = append(ids, result.ID)
		}

		return ids, false, nil
	}

	return idsÚnico, true, nil
}

func (bd MatériaBD) deletarMúltiplas(ids []entidades.ID) erro {
	bd.Log.Informação("Deletando matérias com os IDs:", ids)

	ctx, cancel := context.WithTimeout(bd.ctx, bd.Timeout)
	defer cancel()

	_, err := bd.Collection.DeleteMany(ctx, bson.M{"_id": bson.M{"$in": ids}})
	if err != nil {
		return erros.Novo(data.ErroDeletarMatéria, nil, err)
	}

	return nil
}

// Deletar é uma método que remove uma entidade Matéria no banco de dados MongoDB.
func (bd MatériaBD) Deletar(id entidades.ID) erro {
	bd.Log.Informação("Deletando Matéria com ID:", id)

	ctx, cancel := context.WithTimeout(bd.ctx, bd.Timeout)
	defer cancel()

	_, err := bd.Collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return erros.Novo(data.ErroDeletarMatéria, nil, err)
	}

	return nil
}
