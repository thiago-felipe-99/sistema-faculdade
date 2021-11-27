conn = new Mongo()

db = conn.getDB("Matéria")

db.createCollection("Matéria",{
  validator: {
    $jsonSchema: {
      bsonType: "object",
      required: [ 
        "Nome", 
        "Carga_Horária_Semanal", 
        "Créditos", 
        "Pré_requisitos",
        "Tipo"
      ],
      additionalProperties: false,
      properties: {
        _id: { bsonType: "binData" },
        Carga_Horária_Semanal: { bsonType: "date" },
        Créditos: { bsonType: "double" },
        Pré_requisitos: {
          bsonType: "array",
          items: { bsonType: "binData" },
          uniqueItems: true
        },
        Tipo: { bsonType: "string" }
      }
    }
  }
})

db.createCollection("Turma",{
  validator: {
    $jsonSchema: {
      bsonType: "object",
      required: [ 
        "Matéria",
        "Professores",
        "Alunos",
        "Cursos_Responsáveis",
        "Cursos_Ofertados",
        "Horário_Das_Aulas",
        "Notas",
        "Quantidade_De_Vagas",
        "Data_De_Início",
        "Data_De_Término"
      ],
      additionalProperties: false,
      properties: {
        _id: { bsonType: "binData" },
        Matéria: { bsonType: "binData" },
        Professores: {
          bsonType: "array",
          items: { bsonType: "binData" },
          uniqueItems: true
        },
        Alunos: {
          bsonType: "array",
          items: { bsonType: "binData" },
          uniqueItems: true
        },
        Cursos_Responsáveis: {
          bsonType: "array",
          items: { bsonType: "binData" },
          uniqueItems: true
        },
        Cursos_Ofertados: {
          bsonType: "array",
          items: { bsonType: "binData" },
          uniqueItems: true
        },
        Horário_Das_Aulas: {
          bsonType: "array",
          items: {
            bsonType: "object",
            properties: {
              Dia: { bsonType: "date" },
              Horária_Inicial: { bsonType: "date" },
              Horária_Final: { bsonType: "date" }
            }
          },
          minItems: 1
        },
        Notas: {
          bsonType: "array",
          items: {
            bsonType: "object",
            properties: {
              Aluno: { bsonType: "binData" },
              Nota: { bsonType: "double" },
              Status: { bsonType: "string" }
            }
          }
        },
        Quantidade_De_Vagas: { bsonType: "int" },
        Data_De_Início: { bsonType: "date" },
        Data_De_Término: { bsonType: "date" }
      }
    }
  }
})
