CREATE DATABASE IF NOT EXISTS `$DB_NAME`;

USE `$DB_NAME`;

CREATE USER IF NOT EXISTS '$NEW_USER'@'%' IDENTIFIED BY '$USER_PASSWORD';

GRANT SELECT, INSERT, DELETE, UPDATE ON $DB_NAME.* TO '$NEW_USER'@'%';

CREATE TABLE IF NOT EXISTS `Pessoa` (
  `ID` $UUID NOT NULL,
  `Nome` $STRING NOT NULL,
  `CPF` $CPF NOT NULL UNIQUE,
  `Data_De_Nascimento` $DATE NOT NULL,
  `Senha` $SENHA NOT NULL,
  PRIMARY KEY(ID)
);

CREATE TABLE IF NOT EXISTS `Curso` (
  `ID` $UUID NOT NULL,
  `Nome` $STRING NOT NULL,
  `Data_De_Início` $DATE NOT NULL,
  `Data_De_Desativação` $DATE NOT NULL,
  PRIMARY KEY(ID)
);

CREATE TABLE IF NOT EXISTS `CursoMaterias` (
  `ID_Curso` $UUID NOT NULL,
  `ID_Matéria` $UUID NOT NULL,
  `Periodo` $STRING NOT NULL,
  `Status` $STRING NOT NULL,
  `Observação` $STRING,
  FOREIGN KEY(ID_Curso) REFERENCES Curso(ID)
);

CREATE TABLE IF NOT EXISTS `Aluno` (
  `ID` $UUID NOT NULL,
  `ID_Pessoa` $UUID NOT NULL,
  `ID_Curso` $UUID NOT NULL,
  `Matrícula` $ID NOT NULL UNIQUE,
  `Data_De_Ingresso` $DATE NOT NULL,
  `Data_De_Saída` $DATE NOT NULL,
  `Período` $STRING NOT NULL,
  `Status` $STRING NOT NULL,
  PRIMARY KEY(ID),
  FOREIGN KEY(ID_Pessoa) REFERENCES Pessoa(ID),
  FOREIGN KEY(ID_Curso) REFERENCES Curso(ID)
);

CREATE TABLE IF NOT EXISTS `AlunoTurma` (
  `ID_Aluno` $UUID NOT NULL,
  `ID_Turma` $UUID NOT NULL,
  `Status` $UUID NOT NULL,
  FOREIGN KEY(ID_Aluno) REFERENCES Aluno(ID)
);

CREATE TABLE IF NOT EXISTS `Professor` (
  `ID` $UUID NOT NULL,
  `ID_Pessoa` $UUID NOT NULL,
  `Matrícula` $ID NOT NULL UNIQUE,
  `Data_De_Ingresso` $DATE NOT NULL,
  `Data_De_Saída` $DATE NOT NULL,
  `Status` $STRING NOT NULL,
  `Grau` $STRING NOT NULL,
  `Carga_Horária_Semanal` $TIME NOT NULL,
  PRIMARY KEY(ID),
  FOREIGN KEY(ID_Pessoa) REFERENCES Pessoa(ID)
);

CREATE TABLE IF NOT EXISTS `ProfessorHorario` (
  `ID` $UUID NOT NULL,
  `ID_Professor` $UUID NOT NULL,
  `ID_Turma` $UUID,
  `Nome` $STRING NOT NULL,
  `Dia` $STRING NOT NULL,
  `Horário_Inicial` $TIME NOT NULL,
  `Horário_Final` $TIME NOT NULL,
  `Observação` $TEXT,
  PRIMARY KEY(ID), 
  FOREIGN KEY(ID_Professor) REFERENCES Professor(ID)
);

CREATE TABLE IF NOT EXISTS `Administrativo` (
  `ID` $UUID NOT NULL,
  `ID_Pessoa` $UUID NOT NULL,
  `Matrícula` $ID NOT NULL UNIQUE,
  `Data_De_Ingresso` $DATE NOT NULL,
  `Data_De_Saída` $DATE NOT NULL,
  `Status` $STRING NOT NULL,
  `Grau` $STRING NOT NULL,
  PRIMARY KEY(ID),
  FOREIGN KEY(ID_Pessoa) REFERENCES Pessoa(ID)
);

CREATE TABLE IF NOT EXISTS `AdministrativoHorario` (
  `ID` $UUID NOT NULL,
  `ID_Administrativo` $UUID NOT NULL,
  `ID_Turma` $UUID,
  `Nome` $STRING NOT NULL,
  `Dia` $STRING NOT NULL,
  `Horário_Inicial` $TIME NOT NULL,
  `Horário_Final` $TIME NOT NULL,
  `Observação` $TEXT,
  PRIMARY KEY(ID), 
  FOREIGN KEY(ID_Administrativo) REFERENCES Administrativo(ID)
);

