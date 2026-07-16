package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name      string
		envSet    bool
		envValue  string
		wantPath  string
		wantError bool
	}{
		{
			name:     "uses default configuration",
			wantPath: DefaultDatabasePath,
		},
		{
			name:      "uses database path from environment",
			envSet:    true,
			envValue:  "custom.db",
			wantPath:  "custom.db",
			wantError: false,
		},
		{
			name:      "returns erro fro empty database path",
			envSet:    true,
			envValue:  "",
			wantError: true,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			if tt.envSet {
				t.Setenv(EnvDatabasePath, tt.envValue)
			}

			cfg, err := New()

			if (err != nil) != tt.wantError {
				t.Fatalf("error = %v, wantError = %v", err, tt.wantError)
			}

			if err != nil {
				return
			}

			if got := cfg.Database.Path; got != tt.wantPath {
				t.Fatalf("Database.Path = %q, want %q", got, tt.wantPath)
			}
		})
	}
}

func TestNew_CreatesDatabaseDirectory(t *testing.T) {
	root := t.TempDir()

	dbPath := filepath.Join(
		root,
		"database",
		"tasks.db",
	)

	t.Setenv(EnvDatabasePath, dbPath)

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
