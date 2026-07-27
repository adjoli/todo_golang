package config

import (
	"os"
	"path/filepath"
	"testing"
)

// TestNew valida o comportamento de New em cenários de configuração
// padrão, sobrescrita via variáveis de ambiente e erro de DSN vazio.
func TestNew(t *testing.T) {
	tests := []struct {
		name      string
		envSet    bool
		envValue  string
		wantDSN   string
		wantError bool
	}{
		{
			name:    "uses default configuration",
			wantDSN: DefaultDSN,
		},
		{
			name:      "uses database DSN from environment",
			envSet:    true,
			envValue:  "custom.db",
			wantDSN:   "custom.db",
			wantError: false,
		},
		{
			name:      "returns error for empty database DSN",
			envSet:    true,
			envValue:  "",
			wantError: true,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			if tt.envSet {
				t.Setenv(EnvDatabaseDSN, tt.envValue)
			}

			cfg, err := New()

			if (err != nil) != tt.wantError {
				t.Fatalf("error = %v, wantError = %v", err, tt.wantError)
			}

			if err != nil {
				return
			}

			if got := cfg.Database.DSN; got != tt.wantDSN {
				t.Fatalf("Database.DSN = %q, want %q", got, tt.wantDSN)
			}
		})
	}
}

// TestNew_DefaultDriver verifica que o driver padrão é "sqlite".
func TestNew_DefaultDriver(t *testing.T) {
	cfg, err := New()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if cfg.Database.Driver != "sqlite" {
		t.Fatalf("Database.Driver = %q, want %q", cfg.Database.Driver, "sqlite")
	}
}

// TestNew_PostgresDriver verifica que TASKMANAGER_DB_DRIVER sobrescreve
// o driver para "postgres".
func TestNew_PostgresDriver(t *testing.T) {
	t.Setenv(EnvDatabaseDriver, "postgres")
	t.Setenv(EnvDatabaseDSN, "postgres://localhost:5432/testdb")

	cfg, err := New()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if cfg.Database.Driver != "postgres" {
		t.Fatalf("Database.Driver = %q, want %q", cfg.Database.Driver, "postgres")
	}

	if cfg.Database.DSN != "postgres://localhost:5432/testdb" {
		t.Fatalf("Database.DSN = %q, want %q", cfg.Database.DSN, "postgres://localhost:5432/testdb")
	}
}

// TestNew_InvalidDriver verifica que retorna erro quando o driver
// não é "sqlite" nem "postgres".
func TestNew_InvalidDriver(t *testing.T) {
	t.Setenv(EnvDatabaseDriver, "mysql")

	_, err := New()
	if err == nil {
		t.Fatal("expected error for invalid driver, got nil")
	}
}

// TestNew_EmptyDriver verifica que retorna erro quando o driver
// é definido mas vazio.
func TestNew_EmptyDriver(t *testing.T) {
	t.Setenv(EnvDatabaseDriver, "  ")

	_, err := New()
	if err == nil {
		t.Fatal("expected error for empty driver, got nil")
	}
}

// TestNew_LegacyEnvVar verifica que TASKMANAGER_DB funciona como
// alias para TASKMANAGER_DB_DSN (retrocompatibilidade).
func TestNew_LegacyEnvVar(t *testing.T) {
	t.Setenv(EnvDatabasePath, "legacy/path/tasks.db")

	cfg, err := New()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if cfg.Database.DSN != "legacy/path/tasks.db" {
		t.Fatalf("Database.DSN = %q, want %q", cfg.Database.DSN, "legacy/path/tasks.db")
	}
}

// TestNew_DSNTakesPrecedenceOverLegacy verifica que TASKMANAGER_DB_DSN
// tem prioridade sobre TASKMANAGER_DB quando ambas estão definidas.
func TestNew_DSNTakesPrecedenceOverLegacy(t *testing.T) {
	t.Setenv(EnvDatabaseDSN, "modern/path/tasks.db")
	t.Setenv(EnvDatabasePath, "legacy/path/tasks.db")

	cfg, err := New()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if cfg.Database.DSN != "modern/path/tasks.db" {
		t.Fatalf("Database.DSN = %q, want %q", cfg.Database.DSN, "modern/path/tasks.db")
	}
}

// TestNew_CreatesDatabaseDirectory verifica que New cria o diretório
// pai do caminho do banco de dados automaticamente para SQLite.
func TestNew_CreatesDatabaseDirectory(t *testing.T) {
	root := t.TempDir()

	dbPath := filepath.Join(
		root,
		"database",
		"tasks.db",
	)

	t.Setenv(EnvDatabaseDSN, dbPath)

	_, err := New()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	info, err := os.Stat(filepath.Dir(dbPath))
	if err != nil {
		t.Fatalf("directory was not created: %v", err)
	}

	if !info.IsDir() {
		t.Fatalf("expected a directory")
	}
}

// TestNew_NoDirectoryCreationForPostgres verifica que New não tenta
// criar diretórios quando o driver é Postgres.
func TestNew_NoDirectoryCreationForPostgres(t *testing.T) {
	t.Setenv(EnvDatabaseDriver, "postgres")
	t.Setenv(EnvDatabaseDSN, "postgres://localhost:5432/testdb")

	cfg, err := New()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if cfg.Database.Driver != "postgres" {
		t.Fatalf("Database.Driver = %q, want %q", cfg.Database.Driver, "postgres")
	}
}
