package checks

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	ujconfig "github.com/crossplane/upjet/v2/pkg/config"
)

// Configure adds per-resource overrides for the checks short-group (namespaced scope).
func Configure(p *ujconfig.Provider) {
	configureCheck(p)
	configureCheckGroup(p)
	configureCheckGroupV2(p)
	configureDNSMonitor(p)
	configureHeartbeat(p)
	configureHeartbeatMonitor(p)
	configureTCPCheck(p)
	configureTCPMonitor(p)
	configureURLMonitor(p)
}

func configureCheck(p *ujconfig.Provider) {
	p.AddResourceConfigurator("checkly_check", func(r *ujconfig.Resource) {
		r.ShortGroup = "checks"
		r.Kind = "Check"

		// Cross-resource references
		r.References["group_id"] = ujconfig.Reference{
			TerraformName: "checkly_check_group",
		}
		r.References["alert_channel_subscription.channel_id"] = ujconfig.Reference{
			TerraformName: "checkly_alert_channel",
		}
		r.References["private_locations"] = ujconfig.Reference{
			TerraformName: "checkly_private_location",
			RefFieldName:  "PrivateLocationRefs",
		}

		// Sensitive fields — surfaces as SecretKeySelector in the XR spec.
		reqSchema := r.TerraformResource.Schema["request"].Elem.(*schema.Resource).Schema
		reqSchema["basic_auth"].Elem.(*schema.Resource).Schema["password"].Sensitive = true
		reqSchema["headers"].Sensitive = true // may contain bearer tokens

		// Status-only / computed fields that shouldn't late-init into spec.
		r.LateInitializer.IgnoredFields = append(r.LateInitializer.IgnoredFields,
			"request.follow_redirects", // has a default TF applies
		)
	})
}

func configureCheckGroup(p *ujconfig.Provider) {
	p.AddResourceConfigurator("checkly_check_group", func(r *ujconfig.Resource) {
		r.ShortGroup = "checks"
		r.Kind = "CheckGroup"

		// Cross-resource references
		r.References["alert_channel_subscription.channel_id"] = ujconfig.Reference{
			TerraformName: "checkly_alert_channel",
		}
		r.References["private_locations"] = ujconfig.Reference{
			TerraformName: "checkly_private_location",
			RefFieldName:  "PrivateLocationRefs",
		}

		// Sensitive fields — environment_variable values may hold secrets.
		envVarSchema := r.TerraformResource.Schema["environment_variable"].Elem.(*schema.Resource).Schema
		envVarSchema["value"].Sensitive = true
	})
}

func configureDNSMonitor(p *ujconfig.Provider) {
	p.AddResourceConfigurator("checkly_dns_monitor", func(r *ujconfig.Resource) {
		r.ShortGroup = "checks"
		r.Kind = "DNSMonitor"

		// Cross-resource references
		r.References["group_id"] = ujconfig.Reference{
			TerraformName: "checkly_check_group",
		}
		r.References["alert_channel_subscription.channel_id"] = ujconfig.Reference{
			TerraformName: "checkly_alert_channel",
		}
	})
}

func configureHeartbeat(p *ujconfig.Provider) {
	p.AddResourceConfigurator("checkly_heartbeat", func(r *ujconfig.Resource) {
		r.ShortGroup = "checks"
		r.Kind = "Heartbeat"

		// Cross-resource references
		r.References["alert_channel_subscription.channel_id"] = ujconfig.Reference{
			TerraformName: "checkly_alert_channel",
		}
	})
}

func configureHeartbeatMonitor(p *ujconfig.Provider) {
	p.AddResourceConfigurator("checkly_heartbeat_monitor", func(r *ujconfig.Resource) {
		r.ShortGroup = "checks"
		r.Kind = "HeartbeatMonitor"

		// Cross-resource references
		r.References["alert_channel_subscription.channel_id"] = ujconfig.Reference{
			TerraformName: "checkly_alert_channel",
		}
	})
}

func configureTCPCheck(p *ujconfig.Provider) {
	p.AddResourceConfigurator("checkly_tcp_check", func(r *ujconfig.Resource) {
		r.ShortGroup = "checks"
		r.Kind = "TCPCheck"

		// Cross-resource references
		r.References["group_id"] = ujconfig.Reference{
			TerraformName: "checkly_check_group",
		}
		r.References["alert_channel_subscription.channel_id"] = ujconfig.Reference{
			TerraformName: "checkly_alert_channel",
		}
		r.References["private_locations"] = ujconfig.Reference{
			TerraformName: "checkly_private_location",
			RefFieldName:  "PrivateLocationRefs",
		}
	})
}

func configureTCPMonitor(p *ujconfig.Provider) {
	p.AddResourceConfigurator("checkly_tcp_monitor", func(r *ujconfig.Resource) {
		r.ShortGroup = "checks"
		r.Kind = "TCPMonitor"

		// Cross-resource references
		r.References["group_id"] = ujconfig.Reference{
			TerraformName: "checkly_check_group",
		}
		r.References["alert_channel_subscription.channel_id"] = ujconfig.Reference{
			TerraformName: "checkly_alert_channel",
		}
		r.References["private_locations"] = ujconfig.Reference{
			TerraformName: "checkly_private_location",
			RefFieldName:  "PrivateLocationRefs",
		}
	})
}

func configureURLMonitor(p *ujconfig.Provider) {
	p.AddResourceConfigurator("checkly_url_monitor", func(r *ujconfig.Resource) {
		r.ShortGroup = "checks"
		r.Kind = "URLMonitor"

		// Cross-resource references
		r.References["group_id"] = ujconfig.Reference{
			TerraformName: "checkly_check_group",
		}
		r.References["alert_channel_subscription.channel_id"] = ujconfig.Reference{
			TerraformName: "checkly_alert_channel",
		}
		r.References["private_locations"] = ujconfig.Reference{
			TerraformName: "checkly_private_location",
			RefFieldName:  "PrivateLocationRefs",
		}
	})
}

func configureCheckGroupV2(p *ujconfig.Provider) {
	p.AddResourceConfigurator("checkly_check_group_v2", func(r *ujconfig.Resource) {
		r.ShortGroup = "checks"
		r.Kind = "CheckGroupV2"

		// Cross-resource references
		r.References["enforce_alert_settings.alert_channel_subscription.channel_id"] = ujconfig.Reference{
			TerraformName: "checkly_alert_channel",
		}
		r.References["enforce_locations.private_locations"] = ujconfig.Reference{
			TerraformName: "checkly_private_location",
			RefFieldName:  "PrivateLocationRefs",
		}
		r.References["setup_script.snippet_id"] = ujconfig.Reference{
			TerraformName: "checkly_snippet",
		}
		r.References["teardown_script.snippet_id"] = ujconfig.Reference{
			TerraformName: "checkly_snippet",
		}

		// Sensitive fields — environment_variable values may hold secrets.
		envVarSchema := r.TerraformResource.Schema["environment_variable"].Elem.(*schema.Resource).Schema
		envVarSchema["value"].Sensitive = true

		// API check default headers may contain bearer tokens.
		apiDefSchema := r.TerraformResource.Schema["api_check_defaults"].Elem.(*schema.Resource).Schema
		apiDefSchema["headers"].Sensitive = true
	})
}
