# Task Manager em Go

Uma aplicação didática desenvolvida em Go com o objetivo de estudar boas práticas da linguagem, arquitetura em camadas e utilização da Standard Library.

O projeto é construído de forma incremental, onde cada commit representa um novo conceito aprendido e uma evolução da aplicação.

> **Objetivo:** servir como um modelo para projetos Go simples e como material de estudo da linguagem.

---

## Funcionalidades

* Cadastro de tarefas
* Busca de tarefa por ID
* Listagem de tarefas
* Atualização de tarefas
* Exclusão de tarefas

---

## Tecnologias

* Go
* SQLite
* Standard Library (`database/sql`)
* Context API

---

## Estrutura do Projeto

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
│   ├── models/
│   └── repository/
│
├── go.mod
└── README.md
```

---

## Conceitos estudados até o momento

* Organização de projetos Go
* `database/sql`
* SQLite
* Pool de conexões
* `context.Context`
* Repository Pattern
* CRUD
* SQL parametrizado
* Ponteiros
* Receivers
* `ExecContext`
* `QueryRowContext`
* `QueryContext`
* `Scan`
* `Rows`
* `LastInsertId`
* `RowsAffected`

---

## Como executar

Clone o repositório:

```bash
git clone <URL_DO_REPOSITORIO>
```

Entre na pasta:

```bash
cd task-manager
```

Execute:

```bash
go run ./cmd/taskmanager
```

---

## Organização dos Commits

Cada commit representa um conjunto de conceitos aprendidos.

Exemplo:

* Commit 01 — Estrutura do Projeto
* Commit 02 — Banco de Dados
* Commit 03 — Repository (CRUD)

Essa abordagem permite acompanhar a evolução da aplicação passo a passo.

---

## Documentação

A documentação da arquitetura encontra-se em:

```text
docs/architecture.md
```

Ela registra:

* decisões arquiteturais;
* conceitos estudados;
* padrões da Standard Library;
* filosofia da linguagem Go;
* roadmap do projeto.

---

## Roadmap

* ✅ Estrutura do projeto
* ✅ Banco de dados
* ✅ Repository (CRUD)
* ⏳ Testes automatizados
* ⏳ Service Layer
* ⏳ CLI
* ⏳ Logging
* ⏳ Configuração via `.env`
* ⏳ Docker
* ⏳ GitHub Actions

---

## Licença

Este projeto está licenciado sob a licença MIT.

