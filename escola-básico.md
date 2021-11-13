## Entidades

Pessoa
- Herdar: não
- Atributos:
  - ID Pessoa
  - Nome
  - CPF
  - Data De Nascimento
  - Senha De Acesso

Curso
- Herdar: não
- Atributos:
  - ID Curso
  - Data De Início
  - Data De Desativação

Aluno
- Herdar: Pessoa
- Atributos:
  - ID Aluno
  - Número Do Aluno
  - ID Curso 
  - Data De Ingresso
  - Data De Saída
  - Período
  - Status
  - ID Das Turmas Já Realizadas
  - ID Da Matérias Atuais

Professor
- Herdar: Pessoa
- Atributos:
  - ID Professor
  - Número Do Professor
  - Data De Ingresso
  - Data De Saída
  - Formação
  - Status
  - Grau
  - Horário De Aulas

Administrativo
- Herdar: Pessoa
- Atributos:
  - ID Professor
  - Número Do Professor
  - Data De Ingresso
  - Data De Saída
  - Status
  - Grau
  - Horário De Trabalho

Matéria
- Herdar: Não
- Atributos:
  - ID matéria
  - Quantidade De Horas Semanais
  - Créditos
  - IDs Dos Pré-requisitos
  - Tipo

Turma
- Herdar: Matéria
- Atributos:
  - ID Da Turma
  - IDs Dos Professores
  - IDs Dos Alunos
  - IDs Dos Curso Responsáveis
  - IDs Dos Cursos Ofertados
  - Horário Da Aulas
  - Quantidade De Vagas

