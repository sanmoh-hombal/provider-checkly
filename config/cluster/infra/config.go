package infra

import (
	ujconfig "github.com/crossplane/upjet/v2/pkg/config"

	common "github.com/crossplane-contrib/provider-checkly/config/common/infra"
)

// Configure delegates to the shared infra configurator.
func Configure(p *ujconfig.Provider) { common.Configure(p) }
