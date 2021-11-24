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
  - IDs Das Matérias

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
  - IDs Das Turmas (Realizadas E Ativas)

Professor
- Herdar: Pessoa
- Atributos:
  - ID Professor
  - Número Do Professor
  - Data De Ingresso
  - Data De Saída
  - Status
  - Grau
  - Carga Horária Semanal
  - Horário De Aulas

Administrativo
- Herdar: Pessoa
- Atributos:
  - ID Professor
  - Número Do Administrador
  - Data De Ingresso
  - Data De Saída
  - Status
  - Grau
  - Horário De Trabalho

Matéria
- Herdar: Não
- Atributos:
  - ID matéria
  - Carga Horária Semanal
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
  - Notas
  - Data De Início
  - Data De Fim

## Variáveis De Ambiente
Para rodar o programa deve ter as seguintes variáveis de ambiente:
```Shell
DB_ADMINISTRATIVO_ROOT_PASSWORD= #Senha do root do banco de dados administrativo
DB_MATERIA_ROOT_PASSWORD= #Senha do root do banco de dados matéria
DB_ADMINISTRATIVO_PORT= #Porta que o banco de dados matéria vai ser diponibilizado
DB_MATERIA_PORT= #Porta que o banco de dados matéria vai ser diponibilizado
```

## Volumes
Os seguintes volumes deve ser criados para rodar o projeto localmente com docker
compose:
```Shell
./data/administrativo/mariadb
./data/matéria/mongodb
```

