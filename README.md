# README.md

# Task Manager

Um projeto didático desenvolvido em Go com o objetivo de aprender a linguagem de forma progressiva, aplicando boas práticas, arquitetura em camadas e utilizando prioritariamente a Standard Library.

O projeto é construído em pequenos commits, onde cada etapa introduz um novo conceito e consolida os conhecimentos adquiridos anteriormente.

## Objetivos

* Aprender Go de forma prática.
* Conhecer a Standard Library.
* Aplicar boas práticas da comunidade Go.
* Construir uma arquitetura simples, idiomática e de fácil manutenção.
* Utilizar Git de forma incremental, com commits representando pequenas evoluções do sistema.

---

## Funcionalidades atuais

* Criar tarefas
* Buscar tarefa por ID
* Listar tarefas
* Marcar tarefa como concluída
* Excluir tarefas

---

## Tecnologias

* Go
* SQLite
* `database/sql`
* Context API
* Standard Library

---

## Estrutura do projeto

```text
task-manager/
│
├── cmd/
│   └── taskmanager/
│       └── main.go
│
├── data/
│   └── taskmanager.db
│
├── docs/
│   └── architecture.md
│
├── internal/
│   ├── database/
│   │   ├── config.go
│   │   └── database.go
│   │
│   ├── models/
│   │   └── task.go
│   │
│   ├── repository/
│   │   ├── sql.go
│   │   ├── task_repository.go
│   │   ├── task_repository_test.go
│   │   └── test_helpers.go
│   │
│   └── service/
│       ├── errors.go
│       ├── task_service.go
│       ├── task_service_test.go
│       └── test_helpers.go
│
├── LICENSE
├── README.md
├── go.mod
└── go.sum
```

---

## Arquitetura

```text
                 CLI (em desenvolvimento)
                         │
                         ▼
                 TaskService
                         │
                         ▼
               TaskRepository
                         │
                         ▼
               database/sql
                         │
                         ▼
                     SQLite
```

Cada camada possui uma única responsabilidade.

### TaskService

Responsável pelas regras de negócio.

Exemplos:

* validar dados de entrada;
* criar entidades do domínio;
* traduzir erros da infraestrutura;
* implementar casos de uso.

### TaskRepository

Responsável exclusivamente pela persistência.

Exemplos:

* INSERT
* SELECT
* UPDATE
* DELETE

O Repository não conhece regras de negócio.

---

## Conceitos estudados

### Go

* Packages
* Modules
* Ponteiros
* Receivers
* Structs
* Métodos
* Zero Value
* Context

### Banco de Dados

* `database/sql`
* SQLite
* Connection Pool
* SQL parametrizado
* `ExecContext`
* `QueryContext`
* `QueryRowContext`
* `Rows`
* `Scan`
* `RowsAffected`
* `LastInsertId`

### Arquitetura

* Repository Pattern
* Service Layer
* Casos de Uso
* Erros de Domínio
* Separação de Responsabilidades
* YAGNI
* Código Idiomático

### Testes

* pacote `testing`
* Helpers
* `t.Helper`
* `t.Cleanup`
* Testes de Integração
* Table-Driven Tests
* Testar comportamento ao invés da implementação

---

## Executando

Compile e execute:

```bash
go run ./cmd/taskmanager
```

---

## Executando os testes

Todos os testes:

```bash
go test ./...
```

Cobertura:

```bash
go test -cover ./...
```

Executar o analisador estático:

```bash
go vet ./...
```

Formatar o projeto:

```bash
go fmt ./...
```

---

## Roadmap

### ✅ Commit 01

Estrutura inicial do projeto.

### ✅ Commit 02

Banco de dados.

### ✅ Commit 03

Repository.

### ✅ Commit 04

Testes automatizados.

### ✅ Commit 05

Service Layer.

### ⏳ Commit 06

CLI com Cobra.

### ⏳ Commit 07

Logging.

### ⏳ Commit 08

Configuração via `.env`.

### ⏳ Commit 09

Docker.

### ⏳ Commit 10

GitHub Actions.

---

## Filosofia do projeto

Este projeto procura seguir a filosofia da linguagem Go.

Alguns princípios adotados:

* simplicidade acima de abstração;
* clareza acima de inteligência;
* pequenas funções;
* responsabilidades bem definidas;
* abstrações apenas quando existe necessidade real;
* preferir código explícito;
* testes como parte do desenvolvimento.

---

## Licença

MIT License.
