// Package infra provides shared resource configurators for the infra
// short-group, used by both cluster-scoped and namespaced providers.
package infra

import (
	ujconfig "github.com/crossplane/upjet/v2/pkg/config"

	"github.com/crossplane-contrib/provider-checkly/config/conndetails"
)

// Configure adds per-resource overrides for the infra short-group.
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
		r.Sensitive.AdditionalConnectionDetailsFn = conndetails.IndexedSlice("keys", "api_key")
	})
}
