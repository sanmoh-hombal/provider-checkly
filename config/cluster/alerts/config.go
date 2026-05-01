package alerts

import (
	ujconfig "github.com/crossplane/upjet/v2/pkg/config"

	common "github.com/crossplane-contrib/provider-checkly/config/common/alerts"
)

// Configure delegates to the shared alerts configurator.
func Configure(p *ujconfig.Provider) { common.Configure(p) }
