package checks

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TestAddResponseTimeValidation(t *testing.T) {
	t.Parallel()

	makeSchema := func() map[string]*schema.Schema {
		return map[string]*schema.Schema{
			"degraded_response_time": {Type: schema.TypeInt, Optional: true},
			"max_response_time":      {Type: schema.TypeInt, Optional: true},
		}
	}

	tests := []struct {
		name  string
		max   int
		field string
		value int
		ok    bool
	}{
		{"degraded at lower bound", 30000, "degraded_response_time", 0, true},
		{"degraded at upper bound", 30000, "degraded_response_time", 30000, true},
		{"degraded below lower bound", 30000, "degraded_response_time", -1, false},
		{"degraded above upper bound", 30000, "degraded_response_time", 30001, false},
		{"max at lower bound", 5000, "max_response_time", 0, true},
		{"max at upper bound", 5000, "max_response_time", 5000, true},
		{"max above upper bound", 5000, "max_response_time", 5001, false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			s := makeSchema()
			addResponseTimeValidation(s, tc.max)

			validateFn := s[tc.field].ValidateFunc
			if validateFn == nil {
				t.Fatal("ValidateFunc not set")
			}

			_, errs := validateFn(tc.value, tc.field)
			if tc.ok && len(errs) > 0 {
				t.Errorf("expected valid, got errors: %v", errs)
			}
			if !tc.ok && len(errs) == 0 {
				t.Error("expected validation error, got none")
			}
		})
	}
}

func TestAddResponseTimeValidation_MissingFields(t *testing.T) {
	t.Parallel()

	// Should not panic when fields are absent.
	s := map[string]*schema.Schema{
		"unrelated_field": {Type: schema.TypeString},
	}
	addResponseTimeValidation(s, 30000)

	if s["unrelated_field"].ValidateFunc != nil {
		t.Error("should not set ValidateFunc on unrelated fields")
	}
}
