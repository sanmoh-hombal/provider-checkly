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
