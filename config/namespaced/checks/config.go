package checks

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	ujconfig "github.com/crossplane/upjet/v2/pkg/config"
)

// Configure adds per-resource overrides for checkly_check (namespaced scope).
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("checkly_check", func(r *ujconfig.Resource) {
		r.ShortGroup = "checks"
		r.Kind = "Check"

		// Cross-resource references
		r.References["group_id"] = ujconfig.Reference{
			TerraformName: "checkly_check_group",
		}
		r.References["alert_channel_subscription.channel_id"] = ujconfig.Reference{
			TerraformName: "checkly_alert_channel",
		}
		r.References["private_locations"] = ujconfig.Reference{
			TerraformName: "checkly_private_location",
			RefFieldName:  "PrivateLocationRefs",
		}

		// Sensitive fields — surfaces as SecretKeySelector in the XR spec.
		reqSchema := r.TerraformResource.Schema["request"].Elem.(*schema.Resource).Schema
		reqSchema["basic_auth"].Elem.(*schema.Resource).Schema["password"].Sensitive = true
		reqSchema["headers"].Sensitive = true // may contain bearer tokens

		// Status-only / computed fields that shouldn't late-init into spec.
		r.LateInitializer.IgnoredFields = append(r.LateInitializer.IgnoredFields,
			"request.follow_redirects", // has a default TF applies
		)
	})
}
