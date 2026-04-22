package config

import (
	_ "embed"

	ujconfig "github.com/crossplane/upjet/v2/pkg/config"

	// per-group configurators — cluster-scoped
	alertsCluster     "github.com/sanmoh-hombal/provider-checkly/config/cluster/alerts"
	checksCluster     "github.com/sanmoh-hombal/provider-checkly/config/cluster/checks"
	infraCluster      "github.com/sanmoh-hombal/provider-checkly/config/cluster/infra"
	statuspageCluster "github.com/sanmoh-hombal/provider-checkly/config/cluster/statuspage"
	triggersCluster   "github.com/sanmoh-hombal/provider-checkly/config/cluster/triggers"

	// per-group configurators — namespaced
	alertsNamespaced     "github.com/sanmoh-hombal/provider-checkly/config/namespaced/alerts"
	checksNamespaced     "github.com/sanmoh-hombal/provider-checkly/config/namespaced/checks"
	infraNamespaced      "github.com/sanmoh-hombal/provider-checkly/config/namespaced/infra"
	statuspageNamespaced "github.com/sanmoh-hombal/provider-checkly/config/namespaced/statuspage"
	triggersNamespaced   "github.com/sanmoh-hombal/provider-checkly/config/namespaced/triggers"
)

const (
	resourcePrefix = "checkly"
	modulePath     = "github.com/sanmoh-hombal/provider-checkly"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns the provider configuration for cluster-scoped CRDs.
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("checkly.crossplane.io"),
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

	for _, configure := range []func(provider *ujconfig.Provider){
		checksCluster.Configure,
		alertsCluster.Configure,
		infraCluster.Configure,
		statuspageCluster.Configure,
		triggersCluster.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}

// GetProviderNamespaced returns the provider configuration for namespaced CRDs.
func GetProviderNamespaced() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("checkly.m.crossplane.io"),
		ujconfig.WithShortName("checkly"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		),
		ujconfig.WithSkipList([]string{
			"checkly_static_ips$",
		}),
	)

	for _, configure := range []func(provider *ujconfig.Provider){
		checksNamespaced.Configure,
		alertsNamespaced.Configure,
		infraNamespaced.Configure,
		statuspageNamespaced.Configure,
		triggersNamespaced.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
