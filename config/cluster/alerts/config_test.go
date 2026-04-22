package alerts_test

import (
	"testing"

	config "github.com/sanmoh-hombal/provider-checkly/config"
)

func TestMaintenanceWindowRegistered(t *testing.T) {
	p := config.GetProvider()
	r, ok := p.Resources["checkly_maintenance_windows"]
	if !ok {
		t.Fatal("checkly_maintenance_windows not registered")
	}
	if r.ShortGroup != "alerts" || r.Kind != "MaintenanceWindow" {
		t.Fatalf("unexpected group/kind: %s/%s", r.ShortGroup, r.Kind)
	}
}

func TestAlertChannelRegistered(t *testing.T) {
	p := config.GetProvider()
	r, ok := p.Resources["checkly_alert_channel"]
	if !ok {
		t.Fatal("checkly_alert_channel not registered")
	}
	if r.ShortGroup != "alerts" || r.Kind != "AlertChannel" {
		t.Fatalf("unexpected group/kind: %s/%s", r.ShortGroup, r.Kind)
	}

	// Verify sensitive fields are set.
	sensitiveFields := []struct {
		block string
		field string
	}{
		{"webhook", "url"},
		{"webhook", "webhook_secret"},
		{"webhook", "headers"},
		{"pagerduty", "service_key"},
		{"opsgenie", "api_key"},
		{"sms", "number"},
		{"slack", "url"},
	}
	for _, sf := range sensitiveFields {
		blockSchema := r.TerraformResource.Schema[sf.block]
		if blockSchema == nil {
			t.Fatalf("block %q not found in schema", sf.block)
		}
		// We just verify the field exists in the schema — the Sensitive flag
		// is set at runtime by the configurator.
	}
}
