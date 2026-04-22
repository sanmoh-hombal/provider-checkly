package triggers_test

import (
	"testing"

	config "github.com/sanmoh-hombal/provider-checkly/config"
)

func TestTriggerCheckRegistered(t *testing.T) {
	p := config.GetProvider()
	r, ok := p.Resources["checkly_trigger_check"]
	if !ok {
		t.Fatal("checkly_trigger_check not registered")
	}
	if r.ShortGroup != "triggers" || r.Kind != "TriggerCheck" {
		t.Fatalf("unexpected group/kind: %s/%s", r.ShortGroup, r.Kind)
	}
	if _, ok := r.References["check_id"]; !ok {
		t.Fatal("missing reference for check_id")
	}
	if r.Sensitive.AdditionalConnectionDetailsFn == nil {
		t.Fatal("expected AdditionalConnectionDetailsFn to be set")
	}
}

func TestTriggerCheckConnectionDetails(t *testing.T) {
	p := config.GetProvider()
	r := p.Resources["checkly_trigger_check"]
	fn := r.Sensitive.AdditionalConnectionDetailsFn

	out, err := fn(map[string]any{"url": "https://api.checklyhq.com/trigger/abc"})
	if err != nil {
		t.Fatal(err)
	}
	if string(out["url"]) != "https://api.checklyhq.com/trigger/abc" {
		t.Fatalf("unexpected url: %q", out["url"])
	}

	out, err = fn(map[string]any{})
	if err != nil {
		t.Fatal(err)
	}
	if len(out) != 0 {
		t.Fatalf("expected empty map, got %v", out)
	}
}
