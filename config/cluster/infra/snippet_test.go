package infra_test

import (
	"testing"

	config "github.com/sanmoh-hombal/provider-checkly/config"
)

func TestSnippetRegistered(t *testing.T) {
	p := config.GetProvider()
	r, ok := p.Resources["checkly_snippet"]
	if !ok {
		t.Fatal("checkly_snippet not registered")
	}
	if r.ShortGroup != "infra" || r.Kind != "Snippet" {
		t.Fatalf("unexpected group/kind: %s/%s", r.ShortGroup, r.Kind)
	}
}
