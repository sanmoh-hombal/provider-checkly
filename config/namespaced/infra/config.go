package infra

import ujconfig "github.com/crossplane/upjet/v2/pkg/config"

// Configure adds per-resource overrides for the infra short-group (namespaced scope).
func Configure(p *ujconfig.Provider) {
	configureSnippet(p)
}

func configureSnippet(p *ujconfig.Provider) {
	p.AddResourceConfigurator("checkly_snippet", func(r *ujconfig.Resource) {
		r.ShortGroup = "infra"
		r.Kind = "Snippet"
	})
}
