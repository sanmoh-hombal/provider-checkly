package triggers

import ujconfig "github.com/crossplane/upjet/v2/pkg/config"

// Configure adds per-resource overrides for the triggers short-group (namespaced scope).
func Configure(p *ujconfig.Provider) {
	configureTriggerCheck(p)
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
