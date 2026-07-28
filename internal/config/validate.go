package config

// Validate ensures all configuration fields contain expected, safe values.
func (c Config) Validate() error {
	switch c.LogLevel {
	case "debug", "info", "warn", "error":
		// Valid log level
	default:
		return ErrInvalidLogLevel
	}

	if c.Namespace == "" {
		return ErrEmptyNamespace
	}

	return nil
}
