package config

import (
	"errors"
	"os"
	"testing"
)

func TestLoadDefaults(t *testing.T) {
	os.Unsetenv("ANCHORSTATE_LOG_LEVEL")
	os.Unsetenv("ANCHORSTATE_NAMESPACE")

	cfg, err := Load()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if cfg.LogLevel != DefaultLogLevel {
		t.Errorf("expected LogLevel %s, got %s", DefaultLogLevel, cfg.LogLevel)
	}

	if cfg.Namespace != DefaultNamespace {
		t.Errorf("expected Namespace %s, got %s", DefaultNamespace, cfg.Namespace)
	}
}

func TestLoadCustomEnv(t *testing.T) {
	t.Setenv("ANCHORSTATE_LOG_LEVEL", "debug")
	t.Setenv("ANCHORSTATE_NAMESPACE", "kube-system")

	cfg, err := Load()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if cfg.LogLevel != "debug" {
		t.Errorf("expected LogLevel debug, got %s", cfg.LogLevel)
	}

	if cfg.Namespace != "kube-system" {
		t.Errorf("expected Namespace kube-system, got %s", cfg.Namespace)
	}
}

func TestValidateInvalidLogLevel(t *testing.T) {
	cfg := Config{
		LogLevel:  "potato",
		Namespace: "default",
	}

	err := cfg.Validate()
	if !errors.Is(err, ErrInvalidLogLevel) {
		t.Errorf("expected ErrInvalidLogLevel, got %v", err)
	}
}

func TestValidateEmptyNamespace(t *testing.T) {
	cfg := Config{
		LogLevel:  "info",
		Namespace: "",
	}

	err := cfg.Validate()
	if !errors.Is(err, ErrEmptyNamespace) {
		t.Errorf("expected ErrEmptyNamespace, got %v", err)
	}
}
