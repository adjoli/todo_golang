## Commit 03 (Parte 1)

### Conceitos estudados

- Structs
- Zero values
- Ponteiros
- Passagem de parâmetros por valor
- Escape Analysis
- Repository Pattern
- `context.Context`
- `ExecContext`
- SQL parametrizado
- `sql.Result`
- `LastInsertId`

### Decisões arquiteturais

- Um único `*sql.DB` compartilhado por toda a aplicação.
- Repository concreto (sem interface por enquanto).
- SQL separado em constantes.
- `context.Context` como primeiro parâmetro de todas as operações de banco.
- Entidades passadas por ponteiro quando podem ser modificadas.

### Estado atual

- Banco inicializado automaticamente.
- Model `Task` criado.
- `TaskRepository` implementado.
- Método `Create()` funcional.

### Próximo passo

- Implementar `FindByID()`.
- Validar leitura de uma tarefa.