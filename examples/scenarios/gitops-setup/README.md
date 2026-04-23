# Scenario: GitOps Setup

A single `kustomization.yaml` that composes a ProviderConfig, a shared snippet,
and 10 API checks from a base template — demonstrating how to manage monitoring
at scale with Kustomize overlays.

## What Gets Created

| Resource | Kind | Purpose |
|----------|------|---------|
| `default` | ProviderConfig | Checkly API credentials (from a Secret) |
| `auth-helper` | Snippet | Shared setup snippet injected into all checks |
| `svc-*` (x10) | Check (API) | Health checks for 10 microservices |

## Before

Each microservice team creates checks manually in the Checkly UI. No
consistency in naming, tags, or alert configuration. No audit trail.

## After

A single Kustomize base defines the standard check template. Adding a new service
means adding one resource entry. All checks share the same tags, locations, and
naming convention. Changes are reviewed via pull requests.

## Usage

```bash
# Create the credentials secret first
kubectl create secret generic checkly-creds \
  --namespace crossplane-system \
  --from-literal=credentials='{"api_key":"YOUR_KEY","account_id":"YOUR_ACCOUNT"}'

# Apply the full stack
kubectl apply -k examples/scenarios/gitops-setup/
```
