# Arquitetura do Projeto

Este documento registra a evolução da arquitetura da aplicação e, principalmente, os conceitos da linguagem Go e da Standard Library aprendidos durante o desenvolvimento.

O objetivo não é apenas documentar o código, mas também explicar as decisões de projeto, os padrões recorrentes encontrados na linguagem e o raciocínio por trás de cada escolha arquitetural.

---

# Commit 01 — Estrutura do Projeto

## Objetivos

* Criar a estrutura inicial do projeto.
* Organizar o código seguindo as convenções da comunidade Go.
* Preparar o projeto para crescimento sem complicar a estrutura.

## Conceitos aprendidos

* Organização de projetos Go.
* `cmd/` como ponto de entrada da aplicação.
* `internal/` para código privado da aplicação.
* Separação por responsabilidade.
* Organização em pacotes.

## Decisões arquiteturais

* Um único executável.
* Estrutura preparada para crescimento.
* Separação entre domínio, infraestrutura e ponto de entrada.
* Projeto organizado desde o início.

## Padrões da Standard Library observados

* Organização em pacotes pequenos.
* Responsabilidade única por pacote.
* Código simples e explícito.

## Filosofia Go descoberta neste commit

* Comece simples.
* Organize antes de crescer.
* A estrutura do projeto deve facilitar a navegação.

---

# Commit 02 — Banco de Dados

## Objetivos

* Configurar o SQLite.
* Inicializar automaticamente o banco.
* Criar a tabela `tasks`.

## Conceitos aprendidos

* `database/sql`
* Drivers SQL
* Importação anônima (`_`)
* Registro automático de drivers
* `*sql.DB`
* Pool de conexões
* `Ping()`
* `Exec()`
* `defer`
* Escape Analysis
* `os.MkdirAll()`
* Inicialização idempotente

## Decisões arquiteturais

* Um único `*sql.DB` compartilhado por toda a aplicação.
* Inicialização centralizada em `database.Open()`.
* Criação automática da estrutura do banco.
* SQL separado do restante do código.
* Um único pool de conexões durante toda a execução da aplicação.

## Padrões da Standard Library observados

* Recursos são fechados por quem os adquiriu (`defer`).
* Objetos de infraestrutura costumam ser compartilhados.
* Funções pequenas e lineares.
* Inicialização explícita.

## Filosofia Go descoberta neste commit

* Compartilhe recursos caros.
* O compilador decide quando um objeto deve ir para o heap (Escape Analysis).
* Prefira uma inicialização previsível e determinística.

---

# Commit 03 — Repository (em andamento)

## Objetivos

* Criar a entidade `Task`.
* Criar a camada Repository.
* Implementar o CRUD utilizando `database/sql`.

## Funcionalidades implementadas

* Model `Task`.
* `TaskRepository`.
* `Create()`.
* `FindByID()`.
* `List()`.

## Conceitos aprendidos

### Linguagem

* Structs
* Zero Values
* Ponteiros
* Passagem de parâmetros por valor
* Receivers
* Escape Analysis

### Standard Library

* `context.Context`
* `ExecContext()`
* `QueryRowContext()`
* `QueryContext()`
* `Scan()`
* `sql.Result`
* `LastInsertId()`
* `Rows`
* `rows.Next()`
* `rows.Close()`
* `rows.Err()`

### SQL

* SQL parametrizado
* Placeholders (`?`)
* SQL Injection
* `SELECT` explícito
* `ORDER BY`

## Decisões arquiteturais

* Repository concreto (interfaces serão introduzidas apenas quando houver necessidade).
* `context.Context` como primeiro parâmetro das operações de I/O.
* SQL armazenado em constantes.
* Entidades passadas por ponteiro quando podem ser modificadas.
* Não utilizar `SELECT *`.
* O `main.go` atua temporariamente como laboratório de testes.

## Padrões da Standard Library observados

* APIs preferem preencher estruturas fornecidas pelo chamador em vez de criar novos objetos.
* O controle da alocação normalmente pertence ao chamador.
* Ponteiros representam intenção de modificação.
* `Context` é propagado entre camadas.
* Recursos adquiridos devem ser liberados por quem os adquiriu.
* SQL e dados são mantidos separados através de placeholders.
* Ao iterar resultados, a API fornece um cursor (`Rows`) em vez de retornar todos os registros de uma vez.

## Princípios de Projeto Descobertos

* **YAGNI (You Aren't Gonna Need It)**: não implemente hoje uma solução para um problema que ainda não existe.
* Elimine redundâncias.
* Refatore quando houver um novo requisito, não antes.
* Prefira código que conte uma história.

## Filosofia Go descoberta neste commit

* O chamador controla a memória.
* A biblioteca apenas preenche estruturas fornecidas pelo usuário.
* Prefira funções pequenas.
* Prefira simplicidade à abstração prematura.
* Escreva código explícito.

---

# Glossário

## Receiver

Parâmetro especial que associa um método a um tipo.

Exemplo:

```go
func (r *TaskRepository) Create(...)
```

---

## Zero Value

Valor padrão de qualquer tipo em Go.

Exemplos:

* `0`
* `false`
* `""`
* `nil`
* `time.Time{}`

---

## Escape Analysis

Análise realizada pelo compilador para decidir se um objeto ficará na stack ou será movido para o heap.

---

## Context

Representa uma operação em andamento e permite propagar cancelamento, timeout e metadados por toda a cadeia de chamadas.

---

## Repository

Camada responsável por persistir e recuperar entidades do domínio, encapsulando os detalhes da tecnologia de armazenamento.

---

## Pool de Conexões

Conjunto de conexões gerenciado por `*sql.DB`, reutilizado durante toda a execução da aplicação.

---

## Placeholder

Marcador (`?`) utilizado em comandos SQL para separar o comando dos valores dos parâmetros.

Além de evitar SQL Injection, melhora a legibilidade e permite que o driver trate corretamente os dados.

---

## Cursor

Objeto que representa um conjunto de resultados retornados por uma consulta SQL.

No Go é representado por `*sql.Rows`.

---

## Scan

Método responsável por copiar os valores retornados pelo banco para variáveis fornecidas pelo chamador.

---

# Roadmap

## ✅ Concluído

* Estrutura do projeto.
* Inicialização do banco.
* Model `Task`.
* Repository.
* `Create()`.
* `FindByID()`.
* `List()`.

## 🚧 Em andamento

* `Update()`
* `Delete()`
* Testes completos do Repository.
* Tag `v0.3.0`

## Próximos passos

### Commit 04

* Service Layer.
* Regras de negócio.
* Tradução de erros de infraestrutura para erros de domínio.

### Commit 05

* CLI utilizando Cobra.
* Comandos:

  * `add`
  * `list`
  * `done`
  * `delete`

### Futuro

* Testes automatizados.
* Configuração via `.env`.
* Logging estruturado.
* Paginação.
* Filtros.
* Busca por texto.
* Migrações de banco.
* Docker.
* GitHub Actions.
* Releases automáticas.

## Princípios Arquiteturais

- Prefira duplicação simples a abstrações prematuras.
- Cada Repository representa uma entidade do domínio.
- Generics devem ser utilizados quando o algoritmo é realmente independente do tipo.
- Reflection deve ser evitada quando existe uma solução simples e explícita.
- Escreva código que seja fácil de ler, mesmo que repita algumas linhas.
