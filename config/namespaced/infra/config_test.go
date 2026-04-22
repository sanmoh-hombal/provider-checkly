package infra_test

import (
	"testing"

	config "github.com/sanmoh-hombal/provider-checkly/config"
)

func TestNamespacedInfraResourcesRegistered(t *testing.T) {
	p := config.GetProviderNamespaced()
	for _, name := range []string{
		"checkly_snippet",
		"checkly_environment_variable",
		"checkly_private_location",
		"checkly_client_certificate",
	} {
		if _, ok := p.Resources[name]; !ok {
			t.Errorf("%s not registered in namespaced provider", name)
		}
	}
}

func TestNamespacedPrivateLocationConnectionDetails(t *testing.T) {
	p := config.GetProviderNamespaced()
	r := p.Resources["checkly_private_location"]
	fn := r.Sensitive.AdditionalConnectionDetailsFn
	if fn == nil {
		t.Fatal("expected AdditionalConnectionDetailsFn to be set")
	}

	out, err := fn(map[string]any{"keys": []any{"k0", "k1"}})
	if err != nil {
		t.Fatal(err)
	}
	if string(out["api_key"]) != "k0" {
		t.Fatalf("expected api_key=k0, got %q", out["api_key"])
	}
	if string(out["api_key_1"]) != "k1" {
		t.Fatalf("expected api_key_1=k1, got %q", out["api_key_1"])
	}
}
