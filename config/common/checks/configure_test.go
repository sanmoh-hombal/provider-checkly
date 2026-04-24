package checks_test

import (
	"testing"

	config "github.com/sanmoh-hombal/provider-checkly/config"
)

// resourceResponseTimeBounds maps each resource that exposes response-time
// fields to its documented upper bound (ms).
var resourceResponseTimeBounds = map[string]int{
	"checkly_check":       30000,
	"checkly_dns_monitor": 5000,
	"checkly_tcp_check":   5000,
	"checkly_tcp_monitor": 5000,
	"checkly_url_monitor": 30000,
}

func TestConfigureAddsResponseTimeValidation(t *testing.T) {
	t.Parallel()
	p := config.GetProvider()

	for name, max := range resourceResponseTimeBounds {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			r, ok := p.Resources[name]
			if !ok {
				t.Fatalf("%s not registered", name)
			}

			for _, field := range []string{"degraded_response_time", "max_response_time"} {
				s, ok := r.TerraformResource.Schema[field]
				if !ok {
					t.Fatalf("schema missing field %s", field)
				}
				if s.ValidateFunc == nil {
					t.Errorf("%s: ValidateFunc not set", field)
					continue
				}
				// boundary: max is accepted
				if _, errs := s.ValidateFunc(max, field); len(errs) > 0 {
					t.Errorf("%s: value %d should be valid: %v", field, max, errs)
				}
				// boundary: max+1 is rejected
				if _, errs := s.ValidateFunc(max+1, field); len(errs) == 0 {
					t.Errorf("%s: value %d should be invalid", field, max+1)
				}
			}
		})
	}
}
