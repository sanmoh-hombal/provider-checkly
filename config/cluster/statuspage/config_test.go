package statuspage_test

import (
	"testing"

	config "github.com/crossplane-contrib/provider-checkly/config"
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
	if r.Sensitive.AdditionalConnectionDetailsFn == nil {
		t.Fatal("expected AdditionalConnectionDetailsFn to be set")
	}
}

func TestDashboardConnectionDetails(t *testing.T) {
	p := config.GetProvider()
	r := p.Resources["checkly_dashboard"]
	fn := r.Sensitive.AdditionalConnectionDetailsFn

	t.Run("with key", func(t *testing.T) {
		out, err := fn(map[string]any{"key": "dash-secret"})
		if err != nil {
			t.Fatal(err)
		}
		if string(out["dashboard_key"]) != "dash-secret" {
			t.Fatalf("expected dashboard_key=dash-secret, got %q", out["dashboard_key"])
		}
	})

	t.Run("empty key", func(t *testing.T) {
		out, err := fn(map[string]any{"key": ""})
		if err != nil {
			t.Fatal(err)
		}
		if len(out) != 0 {
			t.Fatalf("expected empty map, got %v", out)
		}
	})

	t.Run("missing key", func(t *testing.T) {
		out, err := fn(map[string]any{})
		if err != nil {
			t.Fatal(err)
		}
		if len(out) != 0 {
			t.Fatalf("expected empty map, got %v", out)
		}
	})
}
