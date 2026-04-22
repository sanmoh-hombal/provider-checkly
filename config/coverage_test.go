package config

import "testing"

func TestAll23ResourcesRegistered(t *testing.T) {
	p := GetProvider()
	if got := len(p.Resources); got != 23 {
		t.Fatalf("expected 23 resources, got %d", got)
	}
}

func TestGetProviderNamespaced(t *testing.T) {
	p := GetProviderNamespaced()
	if got := len(p.Resources); got != 23 {
		t.Fatalf("expected 23 resources, got %d", got)
	}

	// Verify root group differs from cluster-scoped provider.
	if p.RootGroup != "checkly.m.crossplane.io" {
		t.Fatalf("expected namespaced root group checkly.m.crossplane.io, got %s", p.RootGroup)
	}
}
