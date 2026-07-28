package config

import "os"

// Load constructs a Config by reading environment variables and running validation.
func Load() (Config, error) {
	cfg := Config{
		LogLevel:  getEnv("ANCHORSTATE_LOG_LEVEL", DefaultLogLevel),
		Namespace: getEnv("ANCHORSTATE_NAMESPACE", DefaultNamespace),
	}

	if err := cfg.Validate(); err != nil {
		return Config{}, err
	}

	return cfg, nil
}

// getEnv returns the value of an environment variable or fallback if unset.
func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists && value != "" {
		return value
	}
	return fallback
}
