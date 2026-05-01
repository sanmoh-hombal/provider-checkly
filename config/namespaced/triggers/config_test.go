package triggers_test

import (
	"testing"

	config "github.com/crossplane-contrib/provider-checkly/config"
)

func TestNamespacedTriggerResourcesRegistered(t *testing.T) {
	p := config.GetProviderNamespaced()
	for _, name := range []string{
		"checkly_trigger_check",
		"checkly_trigger_group",
	} {
		if _, ok := p.Resources[name]; !ok {
			t.Errorf("%s not registered in namespaced provider", name)
		}
	}
}

func TestNamespacedTriggerCheckConnectionDetails(t *testing.T) {
	p := config.GetProviderNamespaced()
	r := p.Resources["checkly_trigger_check"]
	fn := r.Sensitive.AdditionalConnectionDetailsFn
	if fn == nil {
		t.Fatal("expected AdditionalConnectionDetailsFn to be set")
	}

	out, err := fn(map[string]any{"url": "https://example.com/trigger"})
	if err != nil {
		t.Fatal(err)
	}
	if string(out["url"]) != "https://example.com/trigger" {
		t.Fatalf("unexpected url: %q", out["url"])
	}
}

func TestNamespacedTriggerGroupConnectionDetails(t *testing.T) {
	p := config.GetProviderNamespaced()
	r := p.Resources["checkly_trigger_group"]
	fn := r.Sensitive.AdditionalConnectionDetailsFn
	if fn == nil {
		t.Fatal("expected AdditionalConnectionDetailsFn to be set")
	}

	out, err := fn(map[string]any{"url": "https://example.com/group", "token": "tok"})
	if err != nil {
		t.Fatal(err)
	}
	if string(out["url"]) != "https://example.com/group" {
		t.Fatalf("unexpected url: %q", out["url"])
	}
	if string(out["token"]) != "tok" {
		t.Fatalf("unexpected token: %q", out["token"])
	}
}
