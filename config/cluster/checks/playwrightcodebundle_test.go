package checks_test

import (
	"testing"

	config "github.com/sanmoh-hombal/provider-checkly/config"
	checks "github.com/sanmoh-hombal/provider-checkly/config/cluster/checks"
)

func TestPlaywrightCodeBundleRegistered(t *testing.T) {
	p := config.GetProvider()
	r, ok := p.Resources["checkly_playwright_code_bundle"]
	if !ok {
		t.Fatal("checkly_playwright_code_bundle not registered")
	}
	if r.ShortGroup != checks.ShortGroup || r.Kind != "PlaywrightCodeBundle" {
		t.Fatalf("unexpected group/kind: %s/%s", r.ShortGroup, r.Kind)
	}
}
