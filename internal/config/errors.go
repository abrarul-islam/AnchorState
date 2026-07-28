package config

import "errors"

var (
	ErrInvalidLogLevel = errors.New("invalid log level: must be debug, info, warn, or error")
	ErrEmptyNamespace  = errors.New("namespace cannot be empty")
)
