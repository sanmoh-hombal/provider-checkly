package triggers_test

import (
	"testing"

	config "github.com/sanmoh-hombal/provider-checkly/config"
)

func TestTriggerGroupRegistered(t *testing.T) {
	p := config.GetProvider()
	r, ok := p.Resources["checkly_trigger_group"]
	if !ok {
		t.Fatal("checkly_trigger_group not registered")
	}
	if r.ShortGroup != "triggers" || r.Kind != "TriggerGroup" {
		t.Fatalf("unexpected group/kind: %s/%s", r.ShortGroup, r.Kind)
	}
	if _, ok := r.References["group_id"]; !ok {
		t.Fatal("missing reference for group_id")
	}
	if r.Sensitive.AdditionalConnectionDetailsFn == nil {
		t.Fatal("expected AdditionalConnectionDetailsFn to be set")
	}
}

func TestTriggerGroupConnectionDetails(t *testing.T) {
	p := config.GetProvider()
	r := p.Resources["checkly_trigger_group"]
	fn := r.Sensitive.AdditionalConnectionDetailsFn

	out, err := fn(map[string]any{
		"url":   "https://api.checklyhq.com/trigger/group/abc",
		"token": "tok-secret",
	})
	if err != nil {
		t.Fatal(err)
	}
	if string(out["url"]) != "https://api.checklyhq.com/trigger/group/abc" {
		t.Fatalf("unexpected url: %q", out["url"])
	}
	if string(out["token"]) != "tok-secret" {
		t.Fatalf("unexpected token: %q", out["token"])
	}

	out, err = fn(map[string]any{})
	if err != nil {
		t.Fatal(err)
	}
	if len(out) != 0 {
		t.Fatalf("expected empty map, got %v", out)
	}
}
