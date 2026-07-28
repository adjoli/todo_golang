# Binários
BINARY_CLI=taskmanager
BINARY_API=taskmanager-api

# Caminho do banco
DB=data/tasks.db

CMD_CLI=./cmd/taskmanager
CMD_API=./cmd/taskmanager-api

.PHONY: help
help:
	@echo "Comandos disponíveis:"
	@echo "  make run        - Executa o CLI"
	@echo "  make run-api    - Executa o servidor HTTP"
	@echo "  make build      - Compila ambos os binários"
	@echo "  make build-cli  - Compila apenas o CLI"
	@echo "  make build-api  - Compila apenas o HTTP"
	@echo "  make clean      - Remove binários"
	@echo "  make fmt        - Formata o código"
	@echo "  make vet        - Executa go vet"
	@echo "  make test       - Executa testes"
	@echo "  make test-cover - Executa testes com cobertura"
	@echo "  make tidy       - Atualiza go.mod"
	@echo "  make deps       - Baixa dependências"
	@echo "  make reset-db   - Remove banco SQLite"
	@echo "  make dev        - fmt + vet + run"

.PHONY: run
run:
	go run $(CMD_CLI)

.PHONY: run-api
run-api:
	go run $(CMD_API)

.PHONY: build
build: build-cli build-api

.PHONY: build-cli
build-cli:
	go build -o $(BINARY_CLI) $(CMD_CLI)

.PHONY: build-api
build-api:
	go build -o $(BINARY_API) $(CMD_API)

.PHONY: clean
clean:
	rm -f $(BINARY_CLI) $(BINARY_API)

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: vet
vet:
	go vet ./...

.PHONY: test
test:
	go test ./...

.PHONY: test-cover
test-cover:
	go test -cover ./...

.PHONY: tidy
tidy:
	go mod tidy

.PHONY: deps
deps:
	go mod download

.PHONY: reset-db
reset-db:
	rm -f $(DB)

.PHONY: dev
dev: fmt vet run
