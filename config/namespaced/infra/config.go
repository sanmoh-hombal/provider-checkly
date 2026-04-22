package infra

import ujconfig "github.com/crossplane/upjet/v2/pkg/config"

// Configure adds per-resource overrides for the infra short-group (namespaced scope).
func Configure(p *ujconfig.Provider) {
	configureSnippet(p)
	configureEnvironmentVariable(p)
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
