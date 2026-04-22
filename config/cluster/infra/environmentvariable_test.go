package infra_test

import (
	"testing"

	config "github.com/sanmoh-hombal/provider-checkly/config"
)

func TestEnvironmentVariableRegistered(t *testing.T) {
	p := config.GetProvider()
	r, ok := p.Resources["checkly_environment_variable"]
	if !ok {
		t.Fatal("checkly_environment_variable not registered")
	}
	if r.ShortGroup != "infra" || r.Kind != "EnvironmentVariable" {
		t.Fatalf("unexpected group/kind: %s/%s", r.ShortGroup, r.Kind)
	}
	if !r.TerraformResource.Schema["value"].Sensitive {
		t.Fatal("expected value to be marked sensitive")
	}
}
