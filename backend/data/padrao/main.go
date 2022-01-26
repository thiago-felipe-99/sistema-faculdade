package padrao

import (
	"context"
	"database/sql"

	"go.mongodb.org/mongo-driver/mongo"
	"thiagofelipe.com.br/sistema-faculdade-backend/data"
	"thiagofelipe.com.br/sistema-faculdade-backend/data/mariadb"
	"thiagofelipe.com.br/sistema-faculdade-backend/data/mongodb"
	"thiagofelipe.com.br/sistema-faculdade-backend/logs"
)

// DataPadrão cria um Data que pode ser utilizado na aplicação.
func DataPadrão(
	log *logs.Entidades,
	bdSQL *sql.DB,
	bdMongo *mongo.Database,
) *data.Data {
	pessoa := mariadb.PessoaBD{
		Conexão:      *mariadb.NovaConexão(log.Pessoa, bdSQL),
		NomeDaTabela: "Pessoa",
	}

	conexãoMongoDB := *mongodb.NovaConexão(context.Background(), log.Curso, bdMongo)
	curso := mongodb.CursoBD{
		Conexão:    conexãoMongoDB,
		Collection: conexãoMongoDB.BD.Collection("Curso"),
	}

	// conexãoMongoDB = *mongodb.NovaConexão(context.Background(), log.Aluno, bdMongo)
	// aluno := mongodb.AlunoBD{
	// 	Conexão: conexãoMongoDB,
	// }

	// conexãoMongoDB =
	// *mongodb.NovaConexão(context.Background(), log.Professor, bdMongo)
	// professor := mongodb.ProfessorBD{
	// 	Conexão: conexãoMongoDB,
	// }

	// conexãoMongoDB =
	// *mongodb.NovaConexão(context.Background(), log.Administrativo, bdMongo)
	// administrativo := mongodb.AdministrativoBD{
	// 	Conexão: conexãoMongoDB,
	// }

	conexãoMongoDB = *mongodb.NovaConexão(context.Background(), log.Matéria, bdMongo)
	matéria := mongodb.MatériaBD{
		Conexão:    conexãoMongoDB,
		Collection: conexãoMongoDB.BD.Collection("Matéria"),
	}

	conexãoMongoDB = *mongodb.NovaConexão(context.Background(), log.Turma, bdMongo)
	turma := mongodb.TurmaBD{
		Conexão: conexãoMongoDB,
	}

	return &data.Data{
		Pessoa:         pessoa,
		Curso:          curso,
		Aluno:          nil, // aluno,
		Professor:      nil, // professor,
		Administrativo: nil, // administrativo,
		Matéria:        matéria,
		Turma:          turma,
	}
}
