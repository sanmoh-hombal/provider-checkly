package triggers

import ujconfig "github.com/crossplane/upjet/v2/pkg/config"

// Configure adds per-resource overrides for the triggers short-group (cluster scope).
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

		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]any) (map[string][]byte, error) {
			out := map[string][]byte{}
			if v, ok := attr["url"].(string); ok && v != "" {
				out["url"] = []byte(v)
			}
			if v, ok := attr["token"].(string); ok && v != "" {
				out["token"] = []byte(v)
			}
			return out, nil
		}
	})
}

func configureTriggerCheck(p *ujconfig.Provider) {
	p.AddResourceConfigurator("checkly_trigger_check", func(r *ujconfig.Resource) {
		r.ShortGroup = "triggers"
		r.Kind = "TriggerCheck"

		r.References["check_id"] = ujconfig.Reference{
			TerraformName: "checkly_check",
		}

		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]any) (map[string][]byte, error) {
			out := map[string][]byte{}
			if v, ok := attr["url"].(string); ok && v != "" {
				out["url"] = []byte(v)
			}
			return out, nil
		}
	})
}
