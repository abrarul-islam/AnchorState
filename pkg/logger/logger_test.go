package logger

import "testing"

func TestSetupLogger(t *testing.T) {
	log := SetupLogger("debug")
	if log == nil {
		t.Error("expected initialized logger, got nil")
	}
}
