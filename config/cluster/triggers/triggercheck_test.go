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
