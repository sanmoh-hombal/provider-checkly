package checks_test

import (
	"testing"

	config "github.com/crossplane-contrib/provider-checkly/config"
)

func TestPlaywrightCheckSuiteRegistered(t *testing.T) {
	p := config.GetProvider()
	r, ok := p.Resources["checkly_playwright_check_suite"]
	if !ok {
		t.Fatal("checkly_playwright_check_suite not registered")
	}
	if r.ShortGroup != "checks" || r.Kind != "PlaywrightCheckSuite" {
		t.Fatalf("unexpected group/kind: %s/%s", r.ShortGroup, r.Kind)
	}

	// Verify references
	for _, ref := range []string{
		"bundle.id",
		"alert_channel_subscription.channel_id",
		"private_locations",
	} {
		if _, ok := r.References[ref]; !ok {
			t.Errorf("missing reference: %s", ref)
		}
	}
}
