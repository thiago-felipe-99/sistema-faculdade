package mongodb

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"thiagofelipe.com.br/sistema-faculdade-backend/data"
	"thiagofelipe.com.br/sistema-faculdade-backend/erros"
)

type cursoMatériaParse struct {
	Matéria    id     `bson:"matéria"`
	Período    string `bson:"período"`
	Tipo       string `bson:"tipo"`
	Status     string `bson:"status"`
	Observação string `bson:"observação"`
}

type cursoParse struct {
	ID                id                  `bson:"_id"`
	Nome              string              `bson:"nome"`
	DataDeInício      time.Time           `bson:"data_de_ínicio"`
	DataDeDesativação time.Time           `bson:"data_de_desativação"`
	Matérias          []cursoMatériaParse `bson:"matérias"`
}

func paraCursoParse(curso *curso) cursoParse {
	matérias := []cursoMatériaParse{}

	for _, matéria := range curso.Matérias {
		matérias = append(matérias, cursoMatériaParse{
			Matéria:    matéria.Matéria,
			Período:    matéria.Período,
			Tipo:       matéria.Tipo,
			Status:     matéria.Status,
			Observação: matéria.Observação,
		})
	}

	return cursoParse{
		ID:                curso.ID,
		Nome:              curso.Nome,
		DataDeInício:      curso.DataDeInício,
		DataDeDesativação: curso.DataDeDesativação,
		Matérias:          matérias,
	}
}

func paraCurso(cursoparse *cursoParse) curso {
	matérias := []cursoMatéria{}

	for _, matéria := range cursoparse.Matérias {
		matérias = append(matérias, cursoMatéria{
			Matéria:    matéria.Matéria,
			Período:    matéria.Período,
			Tipo:       matéria.Tipo,
			Status:     matéria.Status,
			Observação: matéria.Observação,
		})
	}

	return curso{
		ID:                cursoparse.ID,
		Nome:              cursoparse.Nome,
		DataDeInício:      cursoparse.DataDeInício,
		DataDeDesativação: cursoparse.DataDeDesativação,
		Matérias:          matérias,
	}
}

// CursoBD representa a conexão com o banco de dados MongoDB para fazer
// alterações na entidade Curso.
type CursoBD struct {
	Conexão
	Collection *mongo.Collection
}

// Inserir é uma método que adiciona uma entidade Curso no banco de
// dados MongoDB.
func (bd CursoBD) Inserir(curso *curso) erro {
	bd.Log.Informação("Inserindo Curso com ID:", curso.ID.String())

	inserir := paraCursoParse(curso)

	ctx, cancel := context.WithTimeout(bd.ctx, bd.Timeout)
	defer cancel()

	_, err := bd.Collection.InsertOne(ctx, inserir)
	if err != nil {
		return erros.Novo(data.ErroInserirCurso, nil, err)
	}

	return nil
}

// Atualizar é uma método que faz a atualização de uma entidade Curso no banco
// de dados MongoDB.
func (bd CursoBD) Atualizar(id id, curso *curso) erro {
	bd.Log.Informação("Atualizando Curso com ID:", curso.ID.String())

	curso.ID = id
	atualizar := paraCursoParse(curso)
	query := bson.D{{Key: "$set", Value: atualizar}}

	ctx, cancel := context.WithTimeout(bd.ctx, bd.Timeout)
	defer cancel()

	_, err := bd.Collection.UpdateByID(ctx, id, query)
	if err != nil {
		return erros.Novo(data.ErroAtualizarCurso, nil, err)
	}

	return nil
}

// Pegar é uma método que retorna uma entidade Curso no banco de dados MongoDB.
func (bd CursoBD) Pegar(id id) (*curso, erro) {
	bd.Log.Informação("Pegando Curso com ID:", id)

	ctx, cancel := context.WithTimeout(bd.ctx, bd.Timeout)
	defer cancel()

	var cursoparse cursoParse

	err := bd.Collection.FindOne(ctx, bson.M{"_id": id}).Decode(&cursoparse)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, erros.Novo(data.ErroCursoNãoEncontrado, nil, err)
		}

		return nil, erros.Novo(data.ErroPegarCurso, nil, err)
	}

	curso := paraCurso(&cursoparse)

	return &curso, nil
}

// Deletar é uma método que remove uma entidade Curso no banco de dados MongoDB.
func (bd CursoBD) Deletar(id id) erro {
	bd.Log.Informação("Deletando Curso com ID:", id)

	ctx, cancel := context.WithTimeout(bd.ctx, bd.Timeout)
	defer cancel()

	_, err := bd.Collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return erros.Novo(data.ErroDeletarCurso, nil, err)
	}

	return nil
}
