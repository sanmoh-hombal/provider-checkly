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

func TestStatusPageServiceRegistered(t *testing.T) {
	p := config.GetProvider()
	r, ok := p.Resources["checkly_status_page_service"]
	if !ok {
		t.Fatal("checkly_status_page_service not registered")
	}
	if r.ShortGroup != "statuspage" || r.Kind != "StatusPageService" {
		t.Fatalf("unexpected group/kind: %s/%s", r.ShortGroup, r.Kind)
	}
	if _, ok := r.References["status_page_id"]; !ok {
		t.Fatal("missing reference for status_page_id")
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
