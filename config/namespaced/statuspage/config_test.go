package statuspage_test

import (
	"testing"

	config "github.com/crossplane-contrib/provider-checkly/config"
)

func TestNamespacedStatusPageResourcesRegistered(t *testing.T) {
	p := config.GetProviderNamespaced()
	for _, name := range []string{
		"checkly_dashboard",
		"checkly_status_page",
		"checkly_status_page_service",
	} {
		if _, ok := p.Resources[name]; !ok {
			t.Errorf("%s not registered in namespaced provider", name)
		}
	}
}

func TestNamespacedDashboardConnectionDetails(t *testing.T) {
	p := config.GetProviderNamespaced()
	r := p.Resources["checkly_dashboard"]
	fn := r.Sensitive.AdditionalConnectionDetailsFn
	if fn == nil {
		t.Fatal("expected AdditionalConnectionDetailsFn to be set")
	}

	out, err := fn(map[string]any{"key": "dash-key"})
	if err != nil {
		t.Fatal(err)
	}
	if string(out["dashboard_key"]) != "dash-key" {
		t.Fatalf("expected dashboard_key=dash-key, got %q", out["dashboard_key"])
	}
}
