package clients

import "testing"

func TestCredentialKeysAreUnique(t *testing.T) {
	keys := []string{keyAPIKey, keyAccountID, keyAPIURL}
	seen := map[string]bool{}
	for _, k := range keys {
		if seen[k] {
			t.Fatalf("duplicate credential key: %s", k)
		}
		seen[k] = true
	}
}
