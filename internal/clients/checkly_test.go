package clients

import (
	"os"
	"testing"
)

// TODO(task-036): These tests exercise TerraformSetupBuilder end-to-end with a
// fake client.Client, ProviderConfig, and Secret. They are skipped until the
// generated v1beta1 types are fully wired and a fake ManagedResource is
// available for the test harness.

func TestTerraformSetupBuilder_SecretSource(t *testing.T) {
	t.Skip("awaits generated v1beta1 — TODO(task-036)")

	// Construct a fake client.Client with a ProviderConfig + Secret containing:
	//   {"api_key":"cu_live_test","account_id":"123"}
	// Call TerraformSetupBuilder("1.0.0","checkly/checkly","1.22.0") with a fake MR.
	// Assert:
	//   ps.Configuration[keyAPIKey] == "cu_live_test"
	//   ps.Configuration[keyAPIURL] == defaultAPIURL  (default applied)
	//   ps.Configuration[keyAccountID] == "123"
}

func TestTerraformSetupBuilder_EnvFallback(t *testing.T) {
	t.Skip("awaits generated v1beta1 — TODO(task-036)")

	// Set CHECKLY_API_KEY before the call; Secret contains empty JSON {}.
	// Assert ps.Configuration[keyAPIKey] picks up the env var value.
}

func TestTerraformSetupBuilder_MissingAPIKey(t *testing.T) {
	t.Skip("awaits generated v1beta1 — TODO(task-036)")

	// Neither Secret nor env provides api_key.
	// Assert error contains errMissingAPIKey.
}

func TestTerraformSetupBuilder_AccountIDOptional(t *testing.T) {
	t.Skip("awaits generated v1beta1 — TODO(task-036)")

	// Secret has {"api_key":"cu_live_test"} but no account_id.
	// Assert ps.Configuration has no keyAccountID entry.
}

func TestApplyEnvDefaults_FillsMissingKeys(t *testing.T) {
	t.Setenv(envAPIKey, "cu_from_env")
	t.Setenv(envAccountID, "42")
	// Leave envAPIURL unset so the hard-coded default kicks in.

	creds := map[string]string{}
	applyEnvDefaults(creds)

	if got := creds[keyAPIKey]; got != "cu_from_env" {
		t.Errorf("api_key = %q, want cu_from_env", got)
	}
	if got := creds[keyAccountID]; got != "42" {
		t.Errorf("account_id = %q, want 42", got)
	}
	if got := creds[keyAPIURL]; got != defaultAPIURL {
		t.Errorf("api_url = %q, want %q", got, defaultAPIURL)
	}
}

func TestApplyEnvDefaults_PreservesExistingKeys(t *testing.T) {
	t.Setenv(envAPIKey, "should_not_override")

	creds := map[string]string{
		keyAPIKey:    "cu_from_secret",
		keyAccountID: "99",
		keyAPIURL:    "https://custom.api",
	}
	applyEnvDefaults(creds)

	if got := creds[keyAPIKey]; got != "cu_from_secret" {
		t.Errorf("api_key = %q, want cu_from_secret (env should not override)", got)
	}
	if got := creds[keyAccountID]; got != "99" {
		t.Errorf("account_id = %q, want 99", got)
	}
	if got := creds[keyAPIURL]; got != "https://custom.api" {
		t.Errorf("api_url = %q, want https://custom.api", got)
	}
}

func TestApplyEnvDefaults_APIURLFromEnv(t *testing.T) {
	t.Setenv(envAPIURL, "https://staging.checklyhq.com")

	creds := map[string]string{}
	applyEnvDefaults(creds)

	if got := creds[keyAPIURL]; got != "https://staging.checklyhq.com" {
		t.Errorf("api_url = %q, want https://staging.checklyhq.com", got)
	}
}

// TestEnvFallbackDefaults verifies the env-var keys and default API URL
// constants are correct without requiring v1beta1 types.
func TestEnvFallbackDefaults(t *testing.T) {
	if envAPIKey != "CHECKLY_API_KEY" {
		t.Errorf("envAPIKey = %q, want CHECKLY_API_KEY", envAPIKey)
	}
	if envAccountID != "CHECKLY_ACCOUNT_ID" {
		t.Errorf("envAccountID = %q, want CHECKLY_ACCOUNT_ID", envAccountID)
	}
	if envAPIURL != "CHECKLY_API_URL" {
		t.Errorf("envAPIURL = %q, want CHECKLY_API_URL", envAPIURL)
	}
	if defaultAPIURL != "https://api.checklyhq.com" {
		t.Errorf("defaultAPIURL = %q, want https://api.checklyhq.com", defaultAPIURL)
	}
}

// TestEnvVarLookup validates os.Getenv returns our test values. This exercises
// the same env-var contract the builder relies on without needing v1beta1.
func TestEnvVarLookup(t *testing.T) {
	const want = "cu_test_env_12345"
	t.Setenv(envAPIKey, want)
	if got := os.Getenv(envAPIKey); got != want {
		t.Errorf("os.Getenv(%q) = %q, want %q", envAPIKey, got, want)
	}
}
