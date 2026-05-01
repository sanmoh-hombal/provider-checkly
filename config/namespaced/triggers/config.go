package triggers

import (
	ujconfig "github.com/crossplane/upjet/v2/pkg/config"

	common "github.com/crossplane-contrib/provider-checkly/config/common/triggers"
)

// Configure delegates to the shared triggers configurator.
func Configure(p *ujconfig.Provider) { common.Configure(p) }
