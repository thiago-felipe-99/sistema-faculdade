## Entidades

Pessoa
- Herdar: não
- Atributos:
  - ID
  - Nome
  - CPF
  - Data De Nascimento
  - Senha De Acesso

Curso
- Herdar: não
- Atributos:
  - ID
  - Nome
  - Data De Início
  - Data De Desativação
  - Matérias Do Curso

Aluno
- Herdar: Pessoa
- Atributos:
  - ID
  - Curso 
  - Matrícula
  - Data De Ingresso
  - Data De Saída
  - Período
  - Status
  - Turmas (Realizadas E Ativas)
  - Horário

Professor
- Herdar: Pessoa
- Atributos:
  - ID
  - Matrícula
  - Data De Ingresso
  - Data De Saída
  - Status
  - Grau
  - Carga Horária Semanal
  - Horário

Administrativo
- Herdar: Pessoa
- Atributos:
  - ID
  - Matrícula
  - Data De Ingresso
  - Data De Saída
  - Status
  - Grau
  - Carga Horária Semanal
  - Horário

Matéria
- Herdar: Não
- Atributos:
  - ID
  - Nome
  - Carga Horária Semanal
  - Créditos
  - Pré-requisitos
  - Tipo

Turma
- Herdar: Matéria
- Atributos:
  - ID
  - Matéria
  - Professores
  - Alunos
  - Curso Responsáveis
  - Cursos Ofertados
  - Horário Das Aulas
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
Os seguintes volumes serão criados ao rodar o projeto localmente com docker o
compose:
```Shell
./data/administrativo/mariadb
./data/matéria/mongodb
```

