package statuspage

import ujconfig "github.com/crossplane/upjet/v2/pkg/config"

// Configure adds per-resource overrides for the statuspage short-group (cluster scope).
func Configure(p *ujconfig.Provider) {
	configureDashboard(p)
	configureStatusPage(p)
}

func configureStatusPage(p *ujconfig.Provider) {
	p.AddResourceConfigurator("checkly_status_page", func(r *ujconfig.Resource) {
		r.ShortGroup = "statuspage"
		r.Kind = "StatusPage"
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
