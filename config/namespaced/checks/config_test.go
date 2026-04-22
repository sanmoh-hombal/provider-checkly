package checks_test

import (
	"testing"

	config "github.com/sanmoh-hombal/provider-checkly/config"
)

func TestNamespacedCheckRegistered(t *testing.T) {
	p := config.GetProviderNamespaced()
	resources := []struct {
		name      string
		shortGrp  string
		kind      string
	}{
		{"checkly_check", "checks", "Check"},
		{"checkly_check_group", "checks", "CheckGroup"},
		{"checkly_check_group_v2", "checks", "CheckGroupV2"},
		{"checkly_dns_monitor", "checks", "DNSMonitor"},
		{"checkly_heartbeat", "checks", "Heartbeat"},
		{"checkly_heartbeat_monitor", "checks", "HeartbeatMonitor"},
		{"checkly_tcp_check", "checks", "TCPCheck"},
		{"checkly_tcp_monitor", "checks", "TCPMonitor"},
		{"checkly_icmp_monitor", "checks", "ICMPMonitor"},
		{"checkly_url_monitor", "checks", "URLMonitor"},
		{"checkly_playwright_code_bundle", "checks", "PlaywrightCodeBundle"},
		{"checkly_playwright_check_suite", "checks", "PlaywrightCheckSuite"},
	}
	for _, tc := range resources {
		t.Run(tc.kind, func(t *testing.T) {
			r, ok := p.Resources[tc.name]
			if !ok {
				t.Fatalf("%s not registered in namespaced provider", tc.name)
			}
			if r.ShortGroup != tc.shortGrp || r.Kind != tc.kind {
				t.Fatalf("unexpected group/kind: %s/%s", r.ShortGroup, r.Kind)
			}
		})
	}
}
