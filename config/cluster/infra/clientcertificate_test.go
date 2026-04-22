package infra_test

import (
	"testing"

	config "github.com/sanmoh-hombal/provider-checkly/config"
)

func TestClientCertificateRegistered(t *testing.T) {
	p := config.GetProvider()
	r, ok := p.Resources["checkly_client_certificate"]
	if !ok {
		t.Fatal("checkly_client_certificate not registered")
	}
	if r.ShortGroup != "infra" || r.Kind != "ClientCertificate" {
		t.Fatalf("unexpected group/kind: %s/%s", r.ShortGroup, r.Kind)
	}
	if !r.TerraformResource.Schema["certificate"].Sensitive {
		t.Fatal("expected certificate to be marked sensitive")
	}
	if !r.TerraformResource.Schema["private_key"].Sensitive {
		t.Fatal("expected private_key to be marked sensitive")
	}
}
