# docs/architecture.md

# Arquitetura do Projeto

Este documento registra as principais decisões arquiteturais tomadas durante o desenvolvimento do projeto.

O objetivo é servir como um diário técnico da evolução da aplicação.

---

# Visão Geral

```text
CLI
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

---

# Organização em Camadas

## Main

Responsável por:

* iniciar a aplicação;
* criar as dependências;
* conectar as camadas.

Não possui regras de negócio.

---

## Service

Responsável por:

* implementar casos de uso;
* validar entradas;
* aplicar regras de negócio;
* traduzir erros da infraestrutura.

Não conhece SQL.

Não conhece detalhes de persistência.

---

## Repository

Responsável por:

* persistência;
* consultas SQL;
* comunicação com `database/sql`.

Não possui regras de negócio.

---

## Database

Responsável por:

* abrir o banco;
* configurar a conexão;
* criar as tabelas quando necessário.

---

# Decisões Arquiteturais

## Repository Pattern

O acesso ao banco é isolado no Repository.

Benefícios:

* separação de responsabilidades;
* facilidade de manutenção;
* reutilização;
* clareza.

---

## Service Layer

A camada de serviço implementa os casos de uso.

Ela impede que regras de negócio sejam espalhadas pela aplicação.

---

## Erros de Domínio

A camada de Service traduz erros da infraestrutura.

Exemplo:

```text
sql.ErrNoRows

↓

ErrTaskNotFound
```

As camadas superiores nunca precisam conhecer `database/sql`.

---

## Casos de Uso

O Service trabalha com linguagem do domínio.

Exemplos:

* CreateTask
* ListTasks
* CompleteTask
* DeleteTask

Evitamos métodos genéricos como:

* UpdateTask

porque a aplicação trabalha com intenções do usuário, não com operações do banco.

---

## Testes

Existem dois conjuntos de testes.

### Repository

Valida:

* SQL
* Persistência
* CRUD

Utiliza SQLite em memória.

---

### Service

Valida:

* regras de negócio;
* casos de uso;
* tradução de erros.

Também utiliza SQLite em memória.

---

# Filosofia de Desenvolvimento

Durante este projeto foram adotados alguns princípios.

## KISS

Keep It Simple, Stupid.

Sempre preferir a solução mais simples.

---

## YAGNI

You Aren't Gonna Need It.

Abstrações somente quando existe necessidade real.

---

## DRY

Don't Repeat Yourself.

Entretanto:

Nem toda repetição merece uma abstração.

A abstração deve surgir apenas quando a repetição representa um problema concreto.

---

## Responsabilidade Única

Cada camada possui apenas um papel.

* Main compõe.
* Service decide.
* Repository persiste.

---

## Testar comportamento

Os testes verificam comportamento observável.

Eles não verificam detalhes internos de implementação.

Isso permite refatorações sem quebrar a suíte de testes.

---

## Código Idiomático

Foram adotadas práticas comuns na comunidade Go:

* receivers pequenos;
* funções curtas;
* early return;
* uso de `context.Context`;
* zero value;
* erros explícitos;
* composição em vez de herança.

---

# Evolução do Projeto

## Commit 01

Estrutura inicial.

---

## Commit 02

Banco de dados.

Aprendizados:

* `database/sql`
* drivers
* pool de conexões

---

## Commit 03

Repository.

Aprendizados:

* CRUD
* SQL parametrizado
* Context
* Scan
* Rows

---

## Commit 04

Testes automatizados.

Aprendizados:

* pacote `testing`
* `t.Helper`
* `t.Cleanup`
* Table-Driven Tests
* testes de integração

---

## Commit 05

Service Layer.

Aprendizados:

* casos de uso;
* erros de domínio;
* tradução de erros;
* separação entre domínio e infraestrutura.

---

# Próximos Passos

* CLI com Cobra
* Logging
* Configuração via `.env`
* Docker
* GitHub Actions
* Distribuição da aplicação

---

# Objetivo Final

Construir um projeto pequeno, porém suficientemente completo para servir como base para novas aplicações Go.

O foco não é apenas implementar um gerenciador de tarefas, mas consolidar um conjunto de práticas idiomáticas da linguagem, produzindo um código simples, testável, organizado e alinhado com a filosofia da comunidade Go.
