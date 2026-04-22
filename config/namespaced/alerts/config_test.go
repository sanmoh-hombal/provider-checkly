package alerts_test

import (
	"testing"

	config "github.com/sanmoh-hombal/provider-checkly/config"
)

func TestNamespacedAlertChannelRegistered(t *testing.T) {
	p := config.GetProviderNamespaced()
	r, ok := p.Resources["checkly_alert_channel"]
	if !ok {
		t.Fatal("checkly_alert_channel not registered in namespaced provider")
	}
	if r.ShortGroup != "alerts" || r.Kind != "AlertChannel" {
		t.Fatalf("unexpected group/kind: %s/%s", r.ShortGroup, r.Kind)
	}
}

func TestNamespacedMaintenanceWindowRegistered(t *testing.T) {
	p := config.GetProviderNamespaced()
	if _, ok := p.Resources["checkly_maintenance_windows"]; !ok {
		t.Fatal("checkly_maintenance_windows not registered in namespaced provider")
	}
}
