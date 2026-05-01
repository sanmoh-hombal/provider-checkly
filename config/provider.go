package config

import (
	_ "embed"

	ujconfig "github.com/crossplane/upjet/v2/pkg/config"

	"github.com/crossplane-contrib/provider-checkly/config/common/alerts"
	"github.com/crossplane-contrib/provider-checkly/config/common/checks"
	"github.com/crossplane-contrib/provider-checkly/config/common/infra"
	"github.com/crossplane-contrib/provider-checkly/config/common/statuspage"
	"github.com/crossplane-contrib/provider-checkly/config/common/triggers"
)

const (
	resourcePrefix = "checkly"
	modulePath     = "github.com/crossplane-contrib/provider-checkly"

	rootGroupCluster    = "checkly.crossplane.io"
	rootGroupNamespaced = "checkly.m.crossplane.io"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// configurators is the ordered list of per-group Configure functions
// shared by both cluster-scoped and namespaced providers.
var configurators = []func(*ujconfig.Provider){
	checks.Configure,
	alerts.Configure,
	infra.Configure,
	statuspage.Configure,
	triggers.Configure,
}

// GetProvider returns the provider configuration for cluster-scoped CRDs.
func GetProvider() *ujconfig.Provider {
	return buildProvider(rootGroupCluster)
}

// GetProviderNamespaced returns the provider configuration for namespaced CRDs.
func GetProviderNamespaced() *ujconfig.Provider {
	return buildProvider(rootGroupNamespaced)
}

// buildProvider constructs a fully-configured ujconfig.Provider. The only
// axis of variation is the root API group, which determines whether the
// generated CRDs are cluster-scoped or namespaced.
func buildProvider(rootGroup string) *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup(rootGroup),
		ujconfig.WithShortName("checkly"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		),
		ujconfig.WithSkipList([]string{
			// Data sources — out of scope for v0.1.
			"checkly_static_ips$",
		}),
	)

	for _, configure := range configurators {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
