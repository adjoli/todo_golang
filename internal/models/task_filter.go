package models

// TaskFilter contém os critérios utilizados para consultar tarefas.
// Campos ponteiro indicam filtros opcionais:
//   - nil   => não aplicar o filtro
//   - valor => aplicar o filtro com o valor informado
type TaskFilter struct {
	Completed *bool
}
