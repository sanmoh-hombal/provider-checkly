package infra_test

import (
	"testing"

	config "github.com/crossplane-contrib/provider-checkly/config"
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
	if !r.TerraformResource.Schema["keys"].Sensitive {
		t.Fatal("expected keys to be marked sensitive")
	}
	if r.Sensitive.AdditionalConnectionDetailsFn == nil {
		t.Fatal("expected AdditionalConnectionDetailsFn to be set")
	}
}

func TestPrivateLocationConnectionDetails(t *testing.T) {
	p := config.GetProvider()
	r := p.Resources["checkly_private_location"]
	fn := r.Sensitive.AdditionalConnectionDetailsFn

	t.Run("single key", func(t *testing.T) {
		out, err := fn(map[string]any{"keys": []any{"abc123"}})
		if err != nil {
			t.Fatal(err)
		}
		if string(out["api_key"]) != "abc123" {
			t.Fatalf("expected api_key=abc123, got %q", out["api_key"])
		}
	})

	t.Run("multiple keys", func(t *testing.T) {
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
	})

	t.Run("empty attr", func(t *testing.T) {
		out, err := fn(map[string]any{})
		if err != nil {
			t.Fatal(err)
		}
		if len(out) != 0 {
			t.Fatalf("expected empty map, got %v", out)
		}
	})
}
