package data

// Data representa quais são as operações para modificar as entidades de uma
// forma definitiva
type Data struct {
	Pessoa         PessoaData
	Curso          CursoData
	Aluno          AlunoData
	Professor      ProfessorData
	Administrativo AdministrativoData
	Matéria        MatériaData
	Turma          TurmaData
}
