package statuspage

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	ujconfig "github.com/crossplane/upjet/v2/pkg/config"
)

// Configure adds per-resource overrides for the statuspage short-group (namespaced scope).
func Configure(p *ujconfig.Provider) {
	configureDashboard(p)
	configureStatusPage(p)
	configureStatusPageService(p)
}

func configureStatusPage(p *ujconfig.Provider) {
	p.AddResourceConfigurator("checkly_status_page", func(r *ujconfig.Resource) {
		r.ShortGroup = "statuspage"
		r.Kind = "StatusPage"
	})
}

func configureStatusPageService(p *ujconfig.Provider) {
	p.AddResourceConfigurator("checkly_status_page_service", func(r *ujconfig.Resource) {
		r.ShortGroup = "statuspage"
		r.Kind = "StatusPageService"

		// The TF provider derives status_page_id from the composite import ID
		// (<status_page_id>/<service_id>) rather than exposing it as an
		// attribute. Add it so Upjet can populate the external-name template.
		r.TerraformResource.Schema["status_page_id"] = &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		}

		r.References["status_page_id"] = ujconfig.Reference{
			TerraformName: "checkly_status_page",
		}
	})
}

func configureDashboard(p *ujconfig.Provider) {
	p.AddResourceConfigurator("checkly_dashboard", func(r *ujconfig.Resource) {
		r.ShortGroup = "statuspage"
		r.Kind = "Dashboard"

		// "key" is already marked sensitive in the TF schema (computed access
		// key for private dashboards). Surface it as a connection detail so
		// consumers can mount it from a Secret.
		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]any) (map[string][]byte, error) {
			conn := map[string][]byte{}
			if v, ok := attr["key"].(string); ok && v != "" {
				conn["dashboard_key"] = []byte(v)
			}
			return conn, nil
		}
	})
}
