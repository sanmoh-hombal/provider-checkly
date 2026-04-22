// Package clients contains the credential parsing and Terraform setup for the Checkly provider.
package clients

// Credential keys used in the JSON blob stored in the ProviderConfig Secret.
const (
	keyAPIKey    = "api_key"
	keyAccountID = "account_id"
	keyAPIURL    = "api_url"
)

// Environment variable fallbacks matching the upstream Terraform provider's behaviour.
const (
	envAPIKey    = "CHECKLY_API_KEY"
	envAccountID = "CHECKLY_ACCOUNT_ID"
	envAPIURL    = "CHECKLY_API_URL"
)

// Default upstream API endpoint.
const defaultAPIURL = "https://api.checklyhq.com"
