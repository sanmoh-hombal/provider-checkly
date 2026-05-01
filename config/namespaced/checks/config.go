package checks

import (
	ujconfig "github.com/crossplane/upjet/v2/pkg/config"

	common "github.com/crossplane-contrib/provider-checkly/config/common/checks"
)

// ShortGroup is re-exported from the shared configurator.
const ShortGroup = common.ShortGroup

// Configure delegates to the shared checks configurator.
func Configure(p *ujconfig.Provider) { common.Configure(p) }
