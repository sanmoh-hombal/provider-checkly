package checks_test

import (
	"testing"

	config "github.com/crossplane-contrib/provider-checkly/config"
	checks "github.com/crossplane-contrib/provider-checkly/config/namespaced/checks"
)

func TestNamespacedCheckRegistered(t *testing.T) {
	p := config.GetProviderNamespaced()
	resources := []struct {
		name     string
		shortGrp string
		kind     string
	}{
		{"checkly_check", checks.ShortGroup, "Check"},
		{"checkly_check_group", checks.ShortGroup, "CheckGroup"},
		{"checkly_check_group_v2", checks.ShortGroup, "CheckGroupV2"},
		{"checkly_dns_monitor", checks.ShortGroup, "DNSMonitor"},
		{"checkly_heartbeat", checks.ShortGroup, "Heartbeat"},
		{"checkly_heartbeat_monitor", checks.ShortGroup, "HeartbeatMonitor"},
		{"checkly_tcp_check", checks.ShortGroup, "TCPCheck"},
		{"checkly_tcp_monitor", checks.ShortGroup, "TCPMonitor"},
		{"checkly_icmp_monitor", checks.ShortGroup, "ICMPMonitor"},
		{"checkly_url_monitor", checks.ShortGroup, "URLMonitor"},
		{"checkly_playwright_code_bundle", checks.ShortGroup, "PlaywrightCodeBundle"},
		{"checkly_playwright_check_suite", checks.ShortGroup, "PlaywrightCheckSuite"},
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
