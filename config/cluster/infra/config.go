package infra

import (
	"fmt"

	ujconfig "github.com/crossplane/upjet/v2/pkg/config"
)

// Configure adds per-resource overrides for the infra short-group (cluster scope).
func Configure(p *ujconfig.Provider) {
	configureSnippet(p)
	configureEnvironmentVariable(p)
	configurePrivateLocation(p)
	configureClientCertificate(p)
}

func configureSnippet(p *ujconfig.Provider) {
	p.AddResourceConfigurator("checkly_snippet", func(r *ujconfig.Resource) {
		r.ShortGroup = "infra"
		r.Kind = "Snippet"
	})
}

func configureEnvironmentVariable(p *ujconfig.Provider) {
	p.AddResourceConfigurator("checkly_environment_variable", func(r *ujconfig.Resource) {
		r.ShortGroup = "infra"
		r.Kind = "EnvironmentVariable"
		r.TerraformResource.Schema["value"].Sensitive = true
	})
}

func configureClientCertificate(p *ujconfig.Provider) {
	p.AddResourceConfigurator("checkly_client_certificate", func(r *ujconfig.Resource) {
		r.ShortGroup = "infra"
		r.Kind = "ClientCertificate"
		r.TerraformResource.Schema["certificate"].Sensitive = true
		r.TerraformResource.Schema["private_key"].Sensitive = true
	})
}

func configurePrivateLocation(p *ujconfig.Provider) {
	p.AddResourceConfigurator("checkly_private_location", func(r *ujconfig.Resource) {
		r.ShortGroup = "infra"
		r.Kind = "PrivateLocation"
		r.TerraformResource.Schema["keys"].Sensitive = true
		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]any) (map[string][]byte, error) {
			out := map[string][]byte{}
			if v, ok := attr["keys"]; ok {
				if keys, ok := v.([]any); ok {
					for i, k := range keys {
						if s, ok := k.(string); ok && s != "" {
							key := "api_key"
							if i > 0 {
								key = fmt.Sprintf("api_key_%d", i)
							}
							out[key] = []byte(s)
						}
					}
				}
			}
			return out, nil
		}
	})
}
