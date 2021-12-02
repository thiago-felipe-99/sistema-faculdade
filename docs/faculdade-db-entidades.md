Tipo De Atributos Padrões
  - UUID
  - CARACTERES
  - NÚMERO INTEIRO
  - NÚMERO REAL
  - HORA
  - DATA

Tipo De Atributos Criadas
  - SEMANA: ARRAY DO TIPO {
    dia: CARACTERES
    Horária Inicial: HORA
    Horária Final: HORA

  }

  - NOTA: {
    ID Do Aluno: UUID
    Nota: NÚMERO INTEIRO
    Status: CARACTERES
  }

  - CURSO: {
    ID Do Curso: UUID
    Vagas: NÚMERO REAL
    Período: CARACTERES
  }

Tabela Pessoa
- Entidade: Pessoa
- Database: Banco De Dados Administrativo
- Atributos:
  - ID Pessoa: UUID
  - Nome: CARACTERES
  - CPF: CARACTERES
  - Data De Nascimento: DATA
  - Senha De Acesso: CARACTERES

Tabela Curso
- Entidade: Curso
- Database: Banco De Dados Administrativo
- Atributos:
  - ID Curso: UUID
  - Nome: CARACTERES
  - Data De Início: DATA
  - Data De Desativação: DATA

Tabela Curso~Matérias
- Entidade: Curso ~ Matérias
- Database: Banco De Dados Administrativo
- Atributos:
  - ID Do Curso: UUID
  - ID Da Matéria: UUID
  - Período: NÚMERO INTEIRO
  - Status: CARACTERES
  - Observação: CARACTERES

Tabela Aluno
- Entidade: Aluno
- Database: Banco De Dados Administrativo
- Atributos:
  - ID Aluno: UUID
  - ID Pessoa: UUID
  - ID Curso : UUID
  - Matrícula: NÚMERO INTEIRO
  - Data De Ingresso: DATA
  - Data De Saída: DATA
  - Período: CARACTERES
  - Status: CARACTERES

Tabela Aluno~Turma
- Entidade: Aluno ~ Turmas
- Database: Banco De Dados Administrativo
- Atributos:
  - ID Do Aluno: UUID
  - ID Da Turma: UUID
  - Status
  
Tabela Professor
- Entidade: Professor
- Database: Banco De Dados Administrativo
- Atributos:
  - ID Professor: UUID
  - ID Pessoa: UUID
  - Matrícula: NÚMERO INTEIRO
  - Data De Ingresso: DATA
  - Data De Saída: DATA
  - Status: CARACTERES
  - Grau: CARACTERES
  - Carga Horária Semanal: HORA
  
Tabela Professor~Horário 
- Entidade: Professor ~ Horário De Aulas
- Database: Banco De Dados Administrativo
- Atributos:
  - ID Do Horário: UUID
  - ID Do Professor: UUID
  - ID Da Turma: UUID
  - Nome: CARACTERES
  - Dia: CARACTERES
  - Horário Inicial: HORA
  - Horário Final: HORA
  - Observação: CARACTERES

Tabela Administrativo
- Herdar: Pessoa
- Database: Banco De Dados Administrativo
- Atributos:
  - ID Administrativo: UUID
  - ID Pessoa: UUID
  - Matrícula: NÚMERO INTEIRO
  - Data De Ingresso: DATA
  - Data De Saída: DATA
  - Status: CARACTERES
  - Grau: CARACTERES
  - Carga Horário Semanal: HORA
  
Tabela Administrativo~Horário 
- Entidade: Administrativo ~ Horário De Trabalho
- Database: Banco De Dados Administrativo
- Atributos:
  - ID Do Horário: UUID
  - ID Do Administrativo: UUID
  - Nome: CARACTERES
  - Dia: CARACTERES
  - Horário Inicial: HORA
  - Horário Final: HORA
  - Observação: CARACTERES

Tabela Matéria
- Entidade: Matéria
- Database: Banco De Dados Matéria
- Atributos:
  - ID matéria: UUID
  - Nome: CARACTERES
  - Carga Horária Semanal: HORA
  - Créditos: NÚMERO REAL
  - IDs Dos Pré-requisitos: ARRAY DE UUID
  - Tipo: CARACTERES

Tabela Turma
- Entidade: Turma
- Database: Banco De Dados Matéria
- Atributos:
  - ID Da Turma: UUID
  - ID Da Matéria: UUID
  - IDs Dos Professores: ARRAY DE UUID
  - IDs Dos Alunos: ARRAY DE UUID
  - IDs Dos Cursos Responsáveis: ARRAY DE UUID
  - IDs Dos Cursos Ofertados: ARRAY DE UUID
  - Horário Das Aulas: SEMANA
  - Notas: NOTA
  - Quantidade De Vagas: NÚMERO INTEIRO
  - Data De Início: DATA
  - Data De Fim: DATA 
