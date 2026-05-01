// Package triggers provides shared resource configurators for the triggers
// short-group, used by both cluster-scoped and namespaced providers.
package triggers

import (
	ujconfig "github.com/crossplane/upjet/v2/pkg/config"

	"github.com/crossplane-contrib/provider-checkly/config/conndetails"
)

// Configure adds per-resource overrides for the triggers short-group.
func Configure(p *ujconfig.Provider) {
	configureTriggerCheck(p)
	configureTriggerGroup(p)
}

func configureTriggerGroup(p *ujconfig.Provider) {
	p.AddResourceConfigurator("checkly_trigger_group", func(r *ujconfig.Resource) {
		r.ShortGroup = "triggers"
		r.Kind = "TriggerGroup"

		r.References["group_id"] = ujconfig.Reference{
			TerraformName: "checkly_check_group",
		}

		r.Sensitive.AdditionalConnectionDetailsFn = conndetails.StringKeys(map[string]string{
			"url":   "url",
			"token": "token",
		})
	})
}

func configureTriggerCheck(p *ujconfig.Provider) {
	p.AddResourceConfigurator("checkly_trigger_check", func(r *ujconfig.Resource) {
		r.ShortGroup = "triggers"
		r.Kind = "TriggerCheck"

		r.References["check_id"] = ujconfig.Reference{
			TerraformName: "checkly_check",
		}

		r.Sensitive.AdditionalConnectionDetailsFn = conndetails.StringKeys(map[string]string{
			"url": "url",
		})
	})
}
