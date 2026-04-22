package statuspage_test

import (
	"testing"

	config "github.com/sanmoh-hombal/provider-checkly/config"
)

func TestStatusPageRegistered(t *testing.T) {
	p := config.GetProvider()
	r, ok := p.Resources["checkly_status_page"]
	if !ok {
		t.Fatal("checkly_status_page not registered")
	}
	if r.ShortGroup != "statuspage" || r.Kind != "StatusPage" {
		t.Fatalf("unexpected group/kind: %s/%s", r.ShortGroup, r.Kind)
	}
}

func TestDashboardRegistered(t *testing.T) {
	p := config.GetProvider()
	r, ok := p.Resources["checkly_dashboard"]
	if !ok {
		t.Fatal("checkly_dashboard not registered")
	}
	if r.ShortGroup != "statuspage" || r.Kind != "Dashboard" {
		t.Fatalf("unexpected group/kind: %s/%s", r.ShortGroup, r.Kind)
	}
}
