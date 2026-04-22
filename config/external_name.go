package config

import "github.com/crossplane/upjet/v2/pkg/config"

// ExternalNameConfigs maps each Checkly terraform resource to an ExternalName
// strategy. Unless overridden, resources use IdentifierFromProvider because
// Checkly generates string IDs server-side.
var ExternalNameConfigs = map[string]config.ExternalName{
	"checkly_check":                  config.IdentifierFromProvider,
	"checkly_check_group":            config.IdentifierFromProvider,
	"checkly_check_group_v2":         config.IdentifierFromProvider,
	"checkly_heartbeat":              config.IdentifierFromProvider,
	"checkly_heartbeat_monitor":      config.IdentifierFromProvider,
	"checkly_tcp_check":              config.IdentifierFromProvider,
	"checkly_tcp_monitor":            config.IdentifierFromProvider,
	"checkly_url_monitor":            config.IdentifierFromProvider,
	"checkly_dns_monitor":            config.IdentifierFromProvider,
	"checkly_icmp_monitor":           config.IdentifierFromProvider,
	"checkly_playwright_check_suite": config.IdentifierFromProvider,
	"checkly_playwright_code_bundle": config.IdentifierFromProvider,

	"checkly_alert_channel":       config.IdentifierFromProvider,
	"checkly_maintenance_windows": config.IdentifierFromProvider,

	"checkly_snippet":            config.IdentifierFromProvider,
	"checkly_private_location":   config.IdentifierFromProvider,
	"checkly_client_certificate": config.IdentifierFromProvider,

	// env vars use the 'key' attribute as their ID.
	"checkly_environment_variable": config.TemplatedStringAsIdentifier("key", "{{ .external_name }}"),

	"checkly_dashboard":   config.IdentifierFromProvider,
	"checkly_status_page": config.IdentifierFromProvider,
	// composite ID: <status_page_id>/<service_id>
	"checkly_status_page_service": config.TemplatedStringAsIdentifier(
		"",
		"{{ .parameters.status_page_id }}/{{ .external_name }}",
	),

	"checkly_trigger_check": config.IdentifierFromProvider,
	"checkly_trigger_group": config.IdentifierFromProvider,
}

// ExternalNameConfigurations applies the ExternalNameConfigs to every matching resource.
func ExternalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		if e, ok := ExternalNameConfigs[r.Name]; ok {
			r.ExternalName = e
		}
	}
}

// ExternalNameConfigured returns the sorted list of resources with custom external-name strategies.
// Useful for include-lists in code generation.
func ExternalNameConfigured() []string {
	out := make([]string, 0, len(ExternalNameConfigs))
	for name := range ExternalNameConfigs {
		out = append(out, name+"$")
	}
	return out
}
