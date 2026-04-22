package alerts

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	ujconfig "github.com/crossplane/upjet/v2/pkg/config"
)

// Configure adds per-resource overrides for the alerts short-group (namespaced scope).
func Configure(p *ujconfig.Provider) {
	configureAlertChannel(p)
	configureMaintenanceWindow(p)
}

func configureMaintenanceWindow(p *ujconfig.Provider) {
	p.AddResourceConfigurator("checkly_maintenance_windows", func(r *ujconfig.Resource) {
		r.ShortGroup = "alerts"
		r.Kind = "MaintenanceWindow"
	})
}

func configureAlertChannel(p *ujconfig.Provider) {
	p.AddResourceConfigurator("checkly_alert_channel", func(r *ujconfig.Resource) {
		r.ShortGroup = "alerts"
		r.Kind = "AlertChannel"

		// Sensitive fields — webhook URLs, API keys, and phone numbers.
		r.TerraformResource.Schema["webhook"].Elem.(*schema.Resource).Schema["url"].Sensitive = true
		r.TerraformResource.Schema["webhook"].Elem.(*schema.Resource).Schema["webhook_secret"].Sensitive = true
		r.TerraformResource.Schema["webhook"].Elem.(*schema.Resource).Schema["headers"].Sensitive = true
		r.TerraformResource.Schema["pagerduty"].Elem.(*schema.Resource).Schema["service_key"].Sensitive = true
		r.TerraformResource.Schema["opsgenie"].Elem.(*schema.Resource).Schema["api_key"].Sensitive = true
		r.TerraformResource.Schema["sms"].Elem.(*schema.Resource).Schema["number"].Sensitive = true
		r.TerraformResource.Schema["slack"].Elem.(*schema.Resource).Schema["url"].Sensitive = true
	})
}
