package checks_test

import (
	"testing"

	config "github.com/sanmoh-hombal/provider-checkly/config"
	checks "github.com/sanmoh-hombal/provider-checkly/config/cluster/checks"
)

// ---------- Check ----------

func TestCheckRegistered(t *testing.T) {
	p := config.GetProvider()
	r, ok := p.Resources["checkly_check"]
	if !ok {
		t.Fatal("checkly_check not registered")
	}
	if r.ShortGroup != checks.ShortGroup || r.Kind != "Check" {
		t.Fatalf("unexpected group/kind: %s/%s", r.ShortGroup, r.Kind)
	}

	for _, ref := range []string{
		"group_id",
		"alert_channel_subscription.channel_id",
		"private_locations",
	} {
		if _, ok := r.References[ref]; !ok {
			t.Errorf("missing reference: %s", ref)
		}
	}
}

// ---------- CheckGroup ----------

func TestCheckGroupRegistered(t *testing.T) {
	p := config.GetProvider()
	r, ok := p.Resources["checkly_check_group"]
	if !ok {
		t.Fatal("checkly_check_group not registered")
	}
	if r.ShortGroup != checks.ShortGroup || r.Kind != "CheckGroup" {
		t.Fatalf("unexpected group/kind: %s/%s", r.ShortGroup, r.Kind)
	}

	for _, ref := range []string{
		"alert_channel_subscription.channel_id",
		"private_locations",
	} {
		if _, ok := r.References[ref]; !ok {
			t.Errorf("missing reference: %s", ref)
		}
	}
}

// ---------- CheckGroupV2 ----------

func TestCheckGroupV2Registered(t *testing.T) {
	p := config.GetProvider()
	r, ok := p.Resources["checkly_check_group_v2"]
	if !ok {
		t.Fatal("checkly_check_group_v2 not registered")
	}
	if r.ShortGroup != checks.ShortGroup || r.Kind != "CheckGroupV2" {
		t.Fatalf("unexpected group/kind: %s/%s", r.ShortGroup, r.Kind)
	}

	for _, ref := range []string{
		"enforce_alert_settings.alert_channel_subscription.channel_id",
		"enforce_locations.private_locations",
		"setup_script.snippet_id",
		"teardown_script.snippet_id",
	} {
		if _, ok := r.References[ref]; !ok {
			t.Errorf("missing reference: %s", ref)
		}
	}
}

// ---------- DNSMonitor ----------

func TestDNSMonitorRegistered(t *testing.T) {
	p := config.GetProvider()
	r, ok := p.Resources["checkly_dns_monitor"]
	if !ok {
		t.Fatal("checkly_dns_monitor not registered")
	}
	if r.ShortGroup != checks.ShortGroup || r.Kind != "DNSMonitor" {
		t.Fatalf("unexpected group/kind: %s/%s", r.ShortGroup, r.Kind)
	}

	for _, ref := range []string{
		"group_id",
		"alert_channel_subscription.channel_id",
	} {
		if _, ok := r.References[ref]; !ok {
			t.Errorf("missing reference: %s", ref)
		}
	}
}

// ---------- Heartbeat ----------

func TestHeartbeatRegistered(t *testing.T) {
	p := config.GetProvider()
	r, ok := p.Resources["checkly_heartbeat"]
	if !ok {
		t.Fatal("checkly_heartbeat not registered")
	}
	if r.ShortGroup != checks.ShortGroup || r.Kind != "Heartbeat" {
		t.Fatalf("unexpected group/kind: %s/%s", r.ShortGroup, r.Kind)
	}
	if _, ok := r.References["alert_channel_subscription.channel_id"]; !ok {
		t.Fatal("missing reference: alert_channel_subscription.channel_id")
	}
}

// ---------- HeartbeatMonitor ----------

func TestHeartbeatMonitorRegistered(t *testing.T) {
	p := config.GetProvider()
	r, ok := p.Resources["checkly_heartbeat_monitor"]
	if !ok {
		t.Fatal("checkly_heartbeat_monitor not registered")
	}
	if r.ShortGroup != checks.ShortGroup || r.Kind != "HeartbeatMonitor" {
		t.Fatalf("unexpected group/kind: %s/%s", r.ShortGroup, r.Kind)
	}
	if _, ok := r.References["alert_channel_subscription.channel_id"]; !ok {
		t.Fatal("missing reference: alert_channel_subscription.channel_id")
	}
}

// ---------- TCPCheck ----------

func TestTCPCheckRegistered(t *testing.T) {
	p := config.GetProvider()
	r, ok := p.Resources["checkly_tcp_check"]
	if !ok {
		t.Fatal("checkly_tcp_check not registered")
	}
	if r.ShortGroup != checks.ShortGroup || r.Kind != "TCPCheck" {
		t.Fatalf("unexpected group/kind: %s/%s", r.ShortGroup, r.Kind)
	}

	for _, ref := range []string{
		"group_id",
		"alert_channel_subscription.channel_id",
		"private_locations",
	} {
		if _, ok := r.References[ref]; !ok {
			t.Errorf("missing reference: %s", ref)
		}
	}
}

// ---------- TCPMonitor ----------

func TestTCPMonitorRegistered(t *testing.T) {
	p := config.GetProvider()
	r, ok := p.Resources["checkly_tcp_monitor"]
	if !ok {
		t.Fatal("checkly_tcp_monitor not registered")
	}
	if r.ShortGroup != checks.ShortGroup || r.Kind != "TCPMonitor" {
		t.Fatalf("unexpected group/kind: %s/%s", r.ShortGroup, r.Kind)
	}

	for _, ref := range []string{
		"group_id",
		"alert_channel_subscription.channel_id",
		"private_locations",
	} {
		if _, ok := r.References[ref]; !ok {
			t.Errorf("missing reference: %s", ref)
		}
	}
}

// ---------- ICMPMonitor ----------

func TestICMPMonitorRegistered(t *testing.T) {
	p := config.GetProvider()
	r, ok := p.Resources["checkly_icmp_monitor"]
	if !ok {
		t.Fatal("checkly_icmp_monitor not registered")
	}
	if r.ShortGroup != checks.ShortGroup || r.Kind != "ICMPMonitor" {
		t.Fatalf("unexpected group/kind: %s/%s", r.ShortGroup, r.Kind)
	}

	for _, ref := range []string{
		"group_id",
		"alert_channel_subscription.channel_id",
	} {
		if _, ok := r.References[ref]; !ok {
			t.Errorf("missing reference: %s", ref)
		}
	}
}

// ---------- URLMonitor ----------

func TestURLMonitorRegistered(t *testing.T) {
	p := config.GetProvider()
	r, ok := p.Resources["checkly_url_monitor"]
	if !ok {
		t.Fatal("checkly_url_monitor not registered")
	}
	if r.ShortGroup != checks.ShortGroup || r.Kind != "URLMonitor" {
		t.Fatalf("unexpected group/kind: %s/%s", r.ShortGroup, r.Kind)
	}

	for _, ref := range []string{
		"group_id",
		"alert_channel_subscription.channel_id",
		"private_locations",
	} {
		if _, ok := r.References[ref]; !ok {
			t.Errorf("missing reference: %s", ref)
		}
	}
}
