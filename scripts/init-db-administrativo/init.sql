CREATE DATABASE IF NOT EXISTS `Administrativo` DEFAULT CHARACTER SET utf8 COLLATE utf8_unicode_ci;

USE `Administrativo`;

CREATE USER IF NOT EXISTS 'Administrativo'@'%' IDENTIFIED BY 'Administrativo';

GRANT SELECT, INSERT, DELETE, UPDATE ON Administrativo.* TO 'Administrativo'@'%';

CREATE TABLE IF NOT EXISTS `Pessoa` (
  `ID` UUID NOT NULL,
  `Nome` VARCHAR(255) NOT NULL,
  `CPF` VARCHAR(11) NOT NULL UNIQUE,
  `Data_De_Nascimento` DATE NOT NULL,
  `Senha` VARCHAR(255) NOT NULL,
  PRIMARY KEY(ID)
);

CREATE TABLE IF NOT EXISTS `Curso` (
  `ID` UUID NOT NULL,
  `Nome` VARCHAR(255) NOT NULL,
  `Data_De_Início` DATE NOT NULL,
  `Data_De_Desativação` DATE NOT NULL,
  PRIMARY KEY(ID)
);

CREATE TABLE IF NOT EXISTS `CursoMatérias` (
  `ID_Curso` UUID NOT NULL,
  `ID_Matéria` UUID NOT NULL,
  `Período` VARCHAR(255) NOT NULL,
  `Tipo` VARCHAR(255) NOT NULL,
  `Status` VARCHAR(255) NOT NULL,
  `Observação` VARCHAR(255),
  FOREIGN KEY(ID_Curso) REFERENCES Curso(ID)
);

CREATE TABLE IF NOT EXISTS `Aluno` (
  `ID` UUID NOT NULL,
  `ID_Pessoa` UUID NOT NULL,
  `ID_Curso` UUID NOT NULL,
  `Matrícula` VARCHAR(11) NOT NULL UNIQUE,
  `Data_De_Ingresso` DATE NOT NULL,
  `Data_De_Saída` DATE NOT NULL,
  `Período` VARCHAR(255) NOT NULL,
  `Status` VARCHAR(255) NOT NULL,
  PRIMARY KEY(ID),
  FOREIGN KEY(ID_Pessoa) REFERENCES Pessoa(ID),
  FOREIGN KEY(ID_Curso) REFERENCES Curso(ID)
);

CREATE TABLE IF NOT EXISTS `AlunoTurma` (
  `ID_Aluno` UUID NOT NULL,
  `ID_Turma` UUID NOT NULL,
  `Status` VARCHAR(255) NOT NULL,
  FOREIGN KEY(ID_Aluno) REFERENCES Aluno(ID)
);

CREATE TABLE IF NOT EXISTS `Professor` (
  `ID` UUID NOT NULL,
  `ID_Pessoa` UUID NOT NULL,
  `Matrícula` VARCHAR(11) NOT NULL UNIQUE,
  `Data_De_Ingresso` DATE NOT NULL,
  `Data_De_Saída` DATE NOT NULL,
  `Status` VARCHAR(255) NOT NULL,
  `Grau` VARCHAR(255) NOT NULL,
  `Carga_Horária_Semanal` TIME NOT NULL,
  PRIMARY KEY(ID),
  FOREIGN KEY(ID_Pessoa) REFERENCES Pessoa(ID)
);

CREATE TABLE IF NOT EXISTS `ProfessorHorário` (
  `ID` UUID NOT NULL,
  `ID_Professor` UUID NOT NULL,
  `ID_Turma` UUID,
  `Nome` VARCHAR(255) NOT NULL,
  `Dia` VARCHAR(255) NOT NULL,
  `Horário_Inicial` TIME NOT NULL,
  `Horário_Final` TIME NOT NULL,
  `Observação` TEXT,
  PRIMARY KEY(ID), 
  FOREIGN KEY(ID_Professor) REFERENCES Professor(ID)
);

CREATE TABLE IF NOT EXISTS `Administrativo` (
  `ID` UUID NOT NULL,
  `ID_Pessoa` UUID NOT NULL,
  `Matrícula` VARCHAR(11) NOT NULL UNIQUE,
  `Data_De_Ingresso` DATE NOT NULL,
  `Data_De_Saída` DATE NOT NULL,
  `Status` VARCHAR(255) NOT NULL,
  `Grau` VARCHAR(255) NOT NULL,
  PRIMARY KEY(ID),
  FOREIGN KEY(ID_Pessoa) REFERENCES Pessoa(ID)
);

CREATE TABLE IF NOT EXISTS `AdministrativoHorário` (
  `ID` UUID NOT NULL,
  `ID_Administrativo` UUID NOT NULL,
  `ID_Turma` UUID,
  `Nome` VARCHAR(255) NOT NULL,
  `Dia` VARCHAR(255) NOT NULL,
  `Horário_Inicial` TIME NOT NULL,
  `Horário_Final` TIME NOT NULL,
  `Observação` TEXT,
  PRIMARY KEY(ID), 
  FOREIGN KEY(ID_Administrativo) REFERENCES Administrativo(ID)
);

CREATE DATABASE IF NOT EXISTS `Teste` DEFAULT CHARACTER SET utf8 COLLATE utf8_unicode_ci;

USE `Teste`;

CREATE USER IF NOT EXISTS 'Teste'@'%' IDENTIFIED BY 'Teste';

GRANT SELECT, INSERT, DELETE, UPDATE ON Teste.* TO 'Teste'@'%';

CREATE TABLE IF NOT EXISTS `Pessoa` (
  `ID` UUID NOT NULL,
  `Nome` VARCHAR(255) NOT NULL,
  `CPF` VARCHAR(11) NOT NULL UNIQUE,
  `Data_De_Nascimento` DATE NOT NULL,
  `Senha` VARCHAR(255) NOT NULL,
  PRIMARY KEY(ID)
);

CREATE TABLE IF NOT EXISTS `Curso` (
  `ID` UUID NOT NULL,
  `Nome` VARCHAR(255) NOT NULL,
  `Data_De_Início` DATE NOT NULL,
  `Data_De_Desativação` DATE NOT NULL,
  PRIMARY KEY(ID)
);

CREATE TABLE IF NOT EXISTS `CursoMatérias` (
  `ID_Curso` UUID NOT NULL,
  `ID_Matéria` UUID NOT NULL,
  `Período` VARCHAR(255) NOT NULL,
  `Tipo` VARCHAR(255) NOT NULL,
  `Status` VARCHAR(255) NOT NULL,
  `Observação` VARCHAR(255),
  FOREIGN KEY(ID_Curso) REFERENCES Curso(ID)
);

CREATE TABLE IF NOT EXISTS `Aluno` (
  `ID` UUID NOT NULL,
  `ID_Pessoa` UUID NOT NULL,
  `ID_Curso` UUID NOT NULL,
  `Matrícula` VARCHAR(11) NOT NULL UNIQUE,
  `Data_De_Ingresso` DATE NOT NULL,
  `Data_De_Saída` DATE NOT NULL,
  `Período` VARCHAR(255) NOT NULL,
  `Status` VARCHAR(255) NOT NULL,
  PRIMARY KEY(ID),
  FOREIGN KEY(ID_Pessoa) REFERENCES Pessoa(ID),
  FOREIGN KEY(ID_Curso) REFERENCES Curso(ID)
);

CREATE TABLE IF NOT EXISTS `AlunoTurma` (
  `ID_Aluno` UUID NOT NULL,
  `ID_Turma` UUID NOT NULL,
  `Status` VARCHAR(255) NOT NULL,
  FOREIGN KEY(ID_Aluno) REFERENCES Aluno(ID)
);

CREATE TABLE IF NOT EXISTS `Professor` (
  `ID` UUID NOT NULL,
  `ID_Pessoa` UUID NOT NULL,
  `Matrícula` VARCHAR(11) NOT NULL UNIQUE,
  `Data_De_Ingresso` DATE NOT NULL,
  `Data_De_Saída` DATE NOT NULL,
  `Status` VARCHAR(255) NOT NULL,
  `Grau` VARCHAR(255) NOT NULL,
  `Carga_Horária_Semanal` TIME NOT NULL,
  PRIMARY KEY(ID),
  FOREIGN KEY(ID_Pessoa) REFERENCES Pessoa(ID)
);

CREATE TABLE IF NOT EXISTS `ProfessorHorário` (
  `ID` UUID NOT NULL,
  `ID_Professor` UUID NOT NULL,
  `ID_Turma` UUID,
  `Nome` VARCHAR(255) NOT NULL,
  `Dia` VARCHAR(255) NOT NULL,
  `Horário_Inicial` TIME NOT NULL,
  `Horário_Final` TIME NOT NULL,
  `Observação` TEXT,
  PRIMARY KEY(ID), 
  FOREIGN KEY(ID_Professor) REFERENCES Professor(ID)
);

CREATE TABLE IF NOT EXISTS `Administrativo` (
  `ID` UUID NOT NULL,
  `ID_Pessoa` UUID NOT NULL,
  `Matrícula` VARCHAR(11) NOT NULL UNIQUE,
  `Data_De_Ingresso` DATE NOT NULL,
  `Data_De_Saída` DATE NOT NULL,
  `Status` VARCHAR(255) NOT NULL,
  `Grau` VARCHAR(255) NOT NULL,
  PRIMARY KEY(ID),
  FOREIGN KEY(ID_Pessoa) REFERENCES Pessoa(ID)
);

CREATE TABLE IF NOT EXISTS `AdministrativoHorário` (
  `ID` UUID NOT NULL,
  `ID_Administrativo` UUID NOT NULL,
  `ID_Turma` UUID,
  `Nome` VARCHAR(255) NOT NULL,
  `Dia` VARCHAR(255) NOT NULL,
  `Horário_Inicial` TIME NOT NULL,
  `Horário_Final` TIME NOT NULL,
  `Observação` TEXT,
  PRIMARY KEY(ID), 
  FOREIGN KEY(ID_Administrativo) REFERENCES Administrativo(ID)
);

