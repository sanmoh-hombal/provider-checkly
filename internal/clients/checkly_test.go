package clients

import (
	"context"
	"strings"
	"testing"

	xpv1 "github.com/crossplane/crossplane-runtime/v2/apis/common/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
	"sigs.k8s.io/controller-runtime/pkg/client/interceptor"

	clusterChecks "github.com/sanmoh-hombal/provider-checkly/apis/cluster/checks/v1alpha1"
	clusterv1beta1 "github.com/sanmoh-hombal/provider-checkly/apis/cluster/v1beta1"
	namespacedChecks "github.com/sanmoh-hombal/provider-checkly/apis/namespaced/checks/v1alpha1"
	namespacedv1beta1 "github.com/sanmoh-hombal/provider-checkly/apis/namespaced/v1beta1"
)

// testScheme returns a runtime.Scheme with the types needed by these tests.
func testScheme(t *testing.T) *runtime.Scheme {
	t.Helper()
	s := runtime.NewScheme()
	for _, add := range []func(*runtime.Scheme) error{
		corev1.AddToScheme,
		clusterv1beta1.SchemeBuilder.AddToScheme,
		clusterChecks.AddToScheme,
		namespacedv1beta1.SchemeBuilder.AddToScheme,
		namespacedChecks.AddToScheme,
	} {
		if err := add(s); err != nil {
			t.Fatalf("add to scheme: %v", err)
		}
	}
	return s
}

// clusterScopeInterceptor strips the namespace from Get calls for
// cluster-scoped types. The fake client doesn't do this automatically
// like the real API server does.
func clusterScopeInterceptor() interceptor.Funcs {
	return interceptor.Funcs{
		Get: func(ctx context.Context, c client.WithWatch, key client.ObjectKey, obj client.Object, opts ...client.GetOption) error {
			if _, ok := obj.(*namespacedv1beta1.ClusterProviderConfig); ok {
				key.Namespace = ""
			}
			return c.Get(ctx, key, obj, opts...)
		},
	}
}

// newManagedCheck returns a cluster-scoped Check (LegacyManaged) pointing at
// the given ProviderConfig name.
func newManagedCheck(pcName string) *clusterChecks.Check {
	c := &clusterChecks.Check{
		ObjectMeta: metav1.ObjectMeta{
			Name: "test-check",
			UID:  types.UID("test-uid-1234"),
		},
	}
	c.SetProviderConfigReference(&xpv1.Reference{Name: pcName})
	return c
}

// newProviderConfig creates a cluster-scoped ProviderConfig that reads
// credentials from the named Secret/key.
func newProviderConfig(name, secretName, secretKey string) *clusterv1beta1.ProviderConfig {
	return &clusterv1beta1.ProviderConfig{
		ObjectMeta: metav1.ObjectMeta{Name: name},
		Spec: clusterv1beta1.ProviderConfigSpec{
			Credentials: clusterv1beta1.ProviderCredentials{
				Source: xpv1.CredentialsSourceSecret,
				CommonCredentialSelectors: xpv1.CommonCredentialSelectors{
					SecretRef: &xpv1.SecretKeySelector{
						SecretReference: xpv1.SecretReference{
							Name:      secretName,
							Namespace: "crossplane-system",
						},
						Key: secretKey,
					},
				},
			},
		},
	}
}

// newSecret creates a Secret with a single key holding the given JSON data.
func newSecret(name, ns, key, jsonData string) *corev1.Secret {
	return &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{Name: name, Namespace: ns},
		Data:       map[string][]byte{key: []byte(jsonData)},
	}
}

const (
	testVersion  = "1.0.0"
	testSource   = "checkly/checkly"
	testPVersion = "1.22.0"
)

func TestTerraformSetupBuilder_SecretSource(t *testing.T) {
	s := testScheme(t)
	fc := fake.NewClientBuilder().WithScheme(s).WithObjects(
		newProviderConfig("default", "checkly-creds", "credentials"),
		newSecret("checkly-creds", "crossplane-system", "credentials",
			`{"api_key":"cu_live_test","account_id":"123"}`),
	).Build()

	setup := TerraformSetupBuilder(testVersion, testSource, testPVersion)
	ps, err := setup(context.Background(), fc, newManagedCheck("default"))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if got := ps.Configuration[keyAPIKey]; got != "cu_live_test" {
		t.Errorf("api_key = %v, want cu_live_test", got)
	}
	if got := ps.Configuration[keyAPIURL]; got != defaultAPIURL {
		t.Errorf("api_url = %v, want %s", got, defaultAPIURL)
	}
	if got := ps.Configuration[keyAccountID]; got != "123" {
		t.Errorf("account_id = %v, want 123", got)
	}
}

func TestTerraformSetupBuilder_EnvFallback(t *testing.T) {
	t.Setenv(envAPIKey, "cu_env_key")

	s := testScheme(t)
	fc := fake.NewClientBuilder().WithScheme(s).WithObjects(
		newProviderConfig("default", "checkly-creds", "credentials"),
		newSecret("checkly-creds", "crossplane-system", "credentials", `{}`),
	).Build()

	setup := TerraformSetupBuilder(testVersion, testSource, testPVersion)
	ps, err := setup(context.Background(), fc, newManagedCheck("default"))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if got := ps.Configuration[keyAPIKey]; got != "cu_env_key" {
		t.Errorf("api_key = %v, want cu_env_key", got)
	}
}

func TestTerraformSetupBuilder_DefaultAPIURL(t *testing.T) {
	s := testScheme(t)
	fc := fake.NewClientBuilder().WithScheme(s).WithObjects(
		newProviderConfig("default", "checkly-creds", "credentials"),
		newSecret("checkly-creds", "crossplane-system", "credentials",
			`{"api_key":"cu_live_test"}`),
	).Build()

	setup := TerraformSetupBuilder(testVersion, testSource, testPVersion)
	ps, err := setup(context.Background(), fc, newManagedCheck("default"))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if got := ps.Configuration[keyAPIURL]; got != defaultAPIURL {
		t.Errorf("api_url = %v, want %s", got, defaultAPIURL)
	}
}

func TestTerraformSetupBuilder_AccountIDOptional(t *testing.T) {
	s := testScheme(t)
	fc := fake.NewClientBuilder().WithScheme(s).WithObjects(
		newProviderConfig("default", "checkly-creds", "credentials"),
		newSecret("checkly-creds", "crossplane-system", "credentials",
			`{"api_key":"cu_live_test"}`),
	).Build()

	setup := TerraformSetupBuilder(testVersion, testSource, testPVersion)
	ps, err := setup(context.Background(), fc, newManagedCheck("default"))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if _, ok := ps.Configuration[keyAccountID]; ok {
		t.Error("account_id should be absent when not provided, but was present")
	}
}

func TestTerraformSetupBuilder_MissingAPIKey(t *testing.T) {
	// Ensure no env fallback either.
	t.Setenv(envAPIKey, "")

	s := testScheme(t)
	fc := fake.NewClientBuilder().WithScheme(s).WithObjects(
		newProviderConfig("default", "checkly-creds", "credentials"),
		newSecret("checkly-creds", "crossplane-system", "credentials", `{}`),
	).Build()

	setup := TerraformSetupBuilder(testVersion, testSource, testPVersion)
	_, err := setup(context.Background(), fc, newManagedCheck("default"))
	if err == nil {
		t.Fatal("expected error for missing api_key, got nil")
	}
	if got := err.Error(); !strings.Contains(got, errMissingAPIKey) {
		t.Errorf("error = %q, want it to contain %q", got, errMissingAPIKey)
	}
}

func TestTerraformSetupBuilder_InvalidJSON(t *testing.T) {
	s := testScheme(t)
	fc := fake.NewClientBuilder().WithScheme(s).WithObjects(
		newProviderConfig("default", "checkly-creds", "credentials"),
		newSecret("checkly-creds", "crossplane-system", "credentials", `not-json`),
	).Build()

	setup := TerraformSetupBuilder(testVersion, testSource, testPVersion)
	_, err := setup(context.Background(), fc, newManagedCheck("default"))
	if err == nil {
		t.Fatal("expected error for invalid JSON, got nil")
	}
	if got := err.Error(); !strings.Contains(got, errUnmarshalCredentials) {
		t.Errorf("error = %q, want it to contain %q", got, errUnmarshalCredentials)
	}
}

// --- Modern (namespaced) path tests ---

// newNamespacedManagedCheck returns a namespaced Check (ModernManaged) pointing
// at a ClusterProviderConfig with the given name.
func newNamespacedManagedCheck(pcName, ns string) *namespacedChecks.Check {
	c := &namespacedChecks.Check{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-ns-check",
			Namespace: ns,
			UID:       types.UID("test-ns-uid-5678"),
		},
	}
	c.SetProviderConfigReference(&xpv1.ProviderConfigReference{
		Kind: "ClusterProviderConfig",
		Name: pcName,
	})
	return c
}

// newClusterProviderConfig creates a namespaced-package ClusterProviderConfig.
func newClusterProviderConfig(name, secretName, secretNS, secretKey string) *namespacedv1beta1.ClusterProviderConfig {
	return &namespacedv1beta1.ClusterProviderConfig{
		ObjectMeta: metav1.ObjectMeta{Name: name},
		Spec: namespacedv1beta1.ProviderConfigSpec{
			Credentials: namespacedv1beta1.ProviderCredentials{
				Source: xpv1.CredentialsSourceSecret,
				CommonCredentialSelectors: xpv1.CommonCredentialSelectors{
					SecretRef: &xpv1.SecretKeySelector{
						SecretReference: xpv1.SecretReference{
							Name:      secretName,
							Namespace: secretNS,
						},
						Key: secretKey,
					},
				},
			},
		},
	}
}

func TestTerraformSetupBuilder_ModernPath_ClusterPC(t *testing.T) {
	s := testScheme(t)
	fc := fake.NewClientBuilder().WithScheme(s).WithInterceptorFuncs(clusterScopeInterceptor()).WithObjects(
		newClusterProviderConfig("default", "checkly-creds", "crossplane-system", "credentials"),
		newSecret("checkly-creds", "crossplane-system", "credentials",
			`{"api_key":"cu_modern_test","account_id":"456"}`),
	).Build()

	setup := TerraformSetupBuilder(testVersion, testSource, testPVersion)
	ps, err := setup(context.Background(), fc, newNamespacedManagedCheck("default", "team-a"))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if got := ps.Configuration[keyAPIKey]; got != "cu_modern_test" {
		t.Errorf("api_key = %v, want cu_modern_test", got)
	}
	if got := ps.Configuration[keyAccountID]; got != "456" {
		t.Errorf("account_id = %v, want 456", got)
	}
	if got := ps.Configuration[keyAPIURL]; got != defaultAPIURL {
		t.Errorf("api_url = %v, want %s", got, defaultAPIURL)
	}
}

func TestTerraformSetupBuilder_ModernPath_NamespacedPC(t *testing.T) {
	s := testScheme(t)
	nsPC := &namespacedv1beta1.ProviderConfig{
		ObjectMeta: metav1.ObjectMeta{Name: "team-pc", Namespace: "team-a"},
		Spec: namespacedv1beta1.ProviderConfigSpec{
			Credentials: namespacedv1beta1.ProviderCredentials{
				Source: xpv1.CredentialsSourceSecret,
				CommonCredentialSelectors: xpv1.CommonCredentialSelectors{
					SecretRef: &xpv1.SecretKeySelector{
						SecretReference: xpv1.SecretReference{
							Name: "team-secret",
							// Namespace will be overridden to mg's namespace
						},
						Key: "credentials",
					},
				},
			},
		},
	}
	fc := fake.NewClientBuilder().WithScheme(s).WithInterceptorFuncs(clusterScopeInterceptor()).WithObjects(
		nsPC,
		newSecret("team-secret", "team-a", "credentials",
			`{"api_key":"cu_ns_test"}`),
	).Build()

	c := &namespacedChecks.Check{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "ns-check",
			Namespace: "team-a",
			UID:       types.UID("ns-uid-9999"),
		},
	}
	c.SetProviderConfigReference(&xpv1.ProviderConfigReference{
		Kind: "ProviderConfig",
		Name: "team-pc",
	})

	setup := TerraformSetupBuilder(testVersion, testSource, testPVersion)
	ps, err := setup(context.Background(), fc, c)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if got := ps.Configuration[keyAPIKey]; got != "cu_ns_test" {
		t.Errorf("api_key = %v, want cu_ns_test", got)
	}
}

// --- Error path tests ---

func TestTerraformSetupBuilder_NoProviderConfigRef(t *testing.T) {
	s := testScheme(t)
	fc := fake.NewClientBuilder().WithScheme(s).Build()

	// Legacy check with no providerConfigRef set.
	c := &clusterChecks.Check{
		ObjectMeta: metav1.ObjectMeta{Name: "no-ref", UID: types.UID("uid-1")},
	}

	setup := TerraformSetupBuilder(testVersion, testSource, testPVersion)
	_, err := setup(context.Background(), fc, c)
	if err == nil {
		t.Fatal("expected error for missing providerConfigRef, got nil")
	}
	if !strings.Contains(err.Error(), errNoProviderConfig) {
		t.Errorf("error = %q, want it to contain %q", err.Error(), errNoProviderConfig)
	}
}

func TestTerraformSetupBuilder_ProviderConfigNotFound(t *testing.T) {
	s := testScheme(t)
	fc := fake.NewClientBuilder().WithScheme(s).Build()

	setup := TerraformSetupBuilder(testVersion, testSource, testPVersion)
	_, err := setup(context.Background(), fc, newManagedCheck("nonexistent"))
	if err == nil {
		t.Fatal("expected error for missing ProviderConfig, got nil")
	}
	if !strings.Contains(err.Error(), errGetProviderConfig) {
		t.Errorf("error = %q, want it to contain %q", err.Error(), errGetProviderConfig)
	}
}

func TestTerraformSetupBuilder_ModernNoProviderConfigRef(t *testing.T) {
	s := testScheme(t)
	fc := fake.NewClientBuilder().WithScheme(s).Build()

	c := &namespacedChecks.Check{
		ObjectMeta: metav1.ObjectMeta{Name: "no-ref", Namespace: "ns", UID: types.UID("uid-2")},
	}

	setup := TerraformSetupBuilder(testVersion, testSource, testPVersion)
	_, err := setup(context.Background(), fc, c)
	if err == nil {
		t.Fatal("expected error for missing providerConfigRef, got nil")
	}
	if !strings.Contains(err.Error(), errNoProviderConfig) {
		t.Errorf("error = %q, want it to contain %q", err.Error(), errNoProviderConfig)
	}
}

func TestTerraformSetupBuilder_ModernPCNotFound(t *testing.T) {
	s := testScheme(t)
	fc := fake.NewClientBuilder().WithScheme(s).Build()

	c := &namespacedChecks.Check{
		ObjectMeta: metav1.ObjectMeta{Name: "check", Namespace: "ns", UID: types.UID("uid-3")},
	}
	c.SetProviderConfigReference(&xpv1.ProviderConfigReference{
		Kind: "ProviderConfig",
		Name: "nonexistent",
	})

	setup := TerraformSetupBuilder(testVersion, testSource, testPVersion)
	_, err := setup(context.Background(), fc, c)
	if err == nil {
		t.Fatal("expected error for missing ProviderConfig, got nil")
	}
	if !strings.Contains(err.Error(), errGetProviderConfig) {
		t.Errorf("error = %q, want it to contain %q", err.Error(), errGetProviderConfig)
	}
}

func TestTerraformSetupBuilder_ModernUnknownKind(t *testing.T) {
	s := testScheme(t)
	fc := fake.NewClientBuilder().WithScheme(s).Build()

	c := &namespacedChecks.Check{
		ObjectMeta: metav1.ObjectMeta{Name: "check", Namespace: "ns", UID: types.UID("uid-4")},
	}
	c.SetProviderConfigReference(&xpv1.ProviderConfigReference{
		Kind: "BogusKind",
		Name: "default",
	})

	setup := TerraformSetupBuilder(testVersion, testSource, testPVersion)
	_, err := setup(context.Background(), fc, c)
	if err == nil {
		t.Fatal("expected error for unknown kind")
	}
	if !strings.Contains(err.Error(), "unknown GVK") {
		t.Errorf("error = %q, want it to contain 'unknown GVK'", err.Error())
	}
}

func TestTerraformSetupBuilder_ExtractCredentialsError(t *testing.T) {
	s := testScheme(t)
	// ProviderConfig references a Secret that does not exist.
	fc := fake.NewClientBuilder().WithScheme(s).WithObjects(
		newProviderConfig("default", "missing-secret", "credentials"),
	).Build()

	setup := TerraformSetupBuilder(testVersion, testSource, testPVersion)
	_, err := setup(context.Background(), fc, newManagedCheck("default"))
	if err == nil {
		t.Fatal("expected error for credential extraction failure")
	}
	if !strings.Contains(err.Error(), errExtractCredentials) {
		t.Errorf("error = %q, want it to contain %q", err.Error(), errExtractCredentials)
	}
}

// --- applyEnvDefaults unit tests ---

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
