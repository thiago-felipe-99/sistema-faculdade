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
  - Data De Início: DATA
  - Data De Desativação: DATA

Tabela Aluno
- Entidade: Aluno
- Database: Banco De Dados Administrativo
- Atributos:
  - ID Aluno: UUID
  - Número Do Aluno: NÚMERO INTEIRO
  - ID Curso : UUID
  - Data De Ingresso: DATA
  - Data De Saída: DATA
  - Período: CARACTERES
  - Status: CARACTERES

Tabela Aluno~Turma
- Entidade: Aluno ~ IDS Das Turmas
- Database: Banco De Dados Administrativo
- Atributos:
  - ID Do Aluno: UUID
  - ID Da Turma: UUID
  
Tabela Professor
- Entidade: Professor
- Database: Banco De Dados Administrativo
- Atributos:
- ID Professor: UUID
- Número Do Professor: NÚMERO INTEIRO
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
  - Dia: CARACTERES
  - Horário Inicial: HORA
  - Horário Final: HORA
  - ID Da Turma: UUID

Tabela Administrativo
- Herdar: Pessoa
- Database: Banco De Dados Administrativo
- Atributos:
  - ID Professor: UUID
  - Número Do Administrador: NÚMERO INTEIRO
  - Data De Ingresso: DATA
  - Data De Saída: DATA
  - Status: CARACTERES
  - Grau: CARACTERES
  
Tabela Administrativo~Horário 
- Entidade: Administrativo ~ Horário De Trabalho
- Database: Banco De Dados Administrativo
- Atributos:
  - ID Do Horário: UUID
  - Dia: CARACTERES
  - Horário Inicial: HORA
  - Horário Final: HORA

Tabela Matéria
- Entidade: Matéria
- Database: Banco De Dados Matéria
- Atributos:
  - ID matéria: UUID
  - Carga Horária Semanal: HORA
  - Créditos: NÚMERO REAL
  - IDs Dos Pré-requisitos: ARRAY DE UUID
  - Tipo: CARACTERES

Tabela Turma
- Entidade: Turma
- Database: Banco De Dados Matéria
- Atributos:
  - ID Da Turma: UUID
  - IDs Dos Professores: ARRAY DE UUID
  - IDs Dos Alunos: ARRAY DE UUID
  - IDs Dos Curso Responsáveis: ARRAY DE UUID
  - IDs Dos Cursos Ofertados: ARRAY DE UUID
  - Horário Da Aulas: SEMANA
  - Quantidade De Vagas: NÚMERO INTEIRO
  - Notas: NOTA
  - Data De Início: DATA
  - Data De Fim: DATA 
