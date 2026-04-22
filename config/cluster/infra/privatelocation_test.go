package infra_test

import (
	"testing"

	config "github.com/sanmoh-hombal/provider-checkly/config"
)

func TestPrivateLocationRegistered(t *testing.T) {
	p := config.GetProvider()
	r, ok := p.Resources["checkly_private_location"]
	if !ok {
		t.Fatal("checkly_private_location not registered")
	}
	if r.ShortGroup != "infra" || r.Kind != "PrivateLocation" {
		t.Fatalf("unexpected group/kind: %s/%s", r.ShortGroup, r.Kind)
	}
	if r.Sensitive.AdditionalConnectionDetailsFn == nil {
		t.Fatal("expected AdditionalConnectionDetailsFn to be set")
	}
}
