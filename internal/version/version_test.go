package version

import "testing"

func TestVersionMetadata(t *testing.T) {
	if Name != "AnchorState" {
		t.Errorf("expected Name to be AnchorState, got %s", Name)
	}

	if Version == "" {
		t.Error("expected Version to be non-empty")
	}
}
