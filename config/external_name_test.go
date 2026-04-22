package config

import "testing"

func TestExternalNameConfigsCoverAllResources(t *testing.T) {
	expected := 23 // keep in sync with architecture doc §5
	if got := len(ExternalNameConfigs); got != expected {
		t.Fatalf("expected %d external-name entries, got %d", expected, got)
	}
}
