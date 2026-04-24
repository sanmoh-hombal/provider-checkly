package statuspage

import (
	ujconfig "github.com/crossplane/upjet/v2/pkg/config"

	common "github.com/sanmoh-hombal/provider-checkly/config/common/statuspage"
)

// Configure delegates to the shared statuspage configurator.
func Configure(p *ujconfig.Provider) { common.Configure(p) }
